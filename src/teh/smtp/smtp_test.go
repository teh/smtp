package smtp

import (
	"bytes"
	"testing"
)

func TestValidHelo(t *testing.T) {
	parser := NewProtocolParser()
	entries := []string{
		"HELO some.domain\r\n",
		"EHLO some.domain\r\n",
		"EHLO [192.167.1.1]\r\n",
		"EHLO [0x7f.1]\r\n",
		"EHLO [0177.0.1]\r\n",
		"EHLO [0177.0.1]\n", // though not valid just \n should still work
		"EHLO somedomain\r\n",
	}
	for _, entry := range entries {
		remaining, err := parser.Feed([]byte(entry))
		if err != nil {
			t.Errorf("Valid HELO/EHLO doesn't parse: %s: %s", entry, err)
		}
		if len(remaining) != 0 {
			t.Errorf("Parse not clean for %#v: remaining `%#v`", entry, string(remaining))
		}
	}
}

func TestInvalidHelo(t *testing.T) {
	parser := NewProtocolParser()
	entries := []string{
		"HELO  some.domain\r\n",
		"EHLO \r\n",
		"EHLO []\r\n",
	}
	for _, entry := range entries {
		_, err := parser.Feed([]byte(entry))
		if err == nil {
			t.Errorf("Invalid helo parses: %s", entry)
		}
	}
}

func TestValidMail(t *testing.T) {
	parser := NewProtocolParser()
	entries := []string{
		"MAIL FROM:<me@example.com>\r\n",
		"MAIL fROM:<root>\r\n",
		"MAIL FrOM:<\"ro\\\"o\"t@some>\r\n",
		"MAIL FRoM:<\\m\\o\\d@some>\r\n",
		"MAIL FROm:<a+b+c@some>\r\n",
		"MAIL FroM:<@q@@10.com>\r\n",
		"MAIL fROm:<@route.1.com,@route2.com:peter@example.net>\r\n",
		"MAIL FROM:<root> BODY=8BITMIME\r\n",
		"MAIL FROM:<root> BODY=7BIT\r\n",
		"MAIL FROM:<root> SIZE=1024\r\n",
		"MAIL FROM:<root> SIZE=1024 BODY=7BIT\r\n",
		"MAIL FROM:<root> BODY=7BIT SIZE=2048\r\n",
		"MAIL FROM:<root> \r\n",
		"MAIL FROM:<root_underscore> \r\n",
		"MAIL FROM:<root-hyphen> \r\n",
		"MAIL FROM:<user@[IPv6:2001:db8:1ff::a0b:dbd0]>\r\n",
		"MAIL FROM:<\"()<>[]:,;@\\\"!#$%&'*+-/=?^_`{}| ~.a\"@example.org>\r\n",
	}
	for _, entry := range entries {
		_, err := parser.Feed([]byte(entry))
		if err != nil {
			t.Errorf("Valid email doesn't parse %s: %s", entry, err)
		}
	}
}

func TestInvalidMail(t *testing.T) {
	parser := NewProtocolParser()
	entries := []string{
		"MAIL FROM: <me@example.com>\r\n",
		"MAIL FROM:<\"me@example.com>\r\n",
		"MAIL FROM:<\"m\"e\"x\"x\"@example.com>\r\n",
		"MAIL FROM:<a\"b(c)d,e:f;g<h>i[j\\k]l@example.com>\r\n",
	}
	for _, entry := range entries {
		_, err := parser.Feed([]byte(entry))
		if err == nil {
			t.Errorf("Invalid email parses %s", entry)
		}
	}
}

func TestValidVerbs(t *testing.T) {
	parser := NewProtocolParser()
	entries := []string{
		"QUIT\r\n",
		"VRFY root\r\n",
		"VRFY <sue@example.com>\r\n",
		"RCPT TO:<sue@example.com>\r\n",
		"DATA\r\n",
		"AUTH PLAIN text\r\n",
		"AUTH PLAIN\r\n",
		"STARTTLS\r\n",
	}
	for _, entry := range entries {
		_, err := parser.Feed([]byte(entry))
		if err != nil {
			t.Errorf("Valid line doesn't parse %s: %s", entry, err)
		}
	}
}

func TestSplitParsing(t *testing.T) {
	parser := NewProtocolParser()
	entry := "EHLO [192.167.1.1]\r\nQUIT\r\n"

	for i := range entry[:19] {
		remaining, err := parser.Feed([]byte(entry[i : i+1]))
		if err != Dangling {
			t.Errorf("Expected Dangling return code, got: %s (remaining: %#v)", err, remaining)
		}
	}
	remaining, err := parser.Feed([]byte(entry[19:20]))
	if err != nil {
		t.Errorf("Expected nil return code, got: %s (remaining: %#v)", err, remaining)
	}
	if parser.current.Verb != VerbEHLO {
		t.Errorf("Expected EHLO verb (%d), got: %d", VerbEHLO, parser.current.Verb)
	}
	if bytes.Compare(parser.current.Data, []byte("[192.167.1.1]")) != 0 {
		t.Errorf("Misparsed Host: %s instead of %s", string(parser.current.Data), "[192.167.1.1]")
	}

	remaining, err = parser.Feed([]byte(entry[20:]))
	if err != nil {
		t.Errorf("Expected nil return code, got: %s (remaining: %#v)", err, remaining)
	}
	if parser.current.Verb != VerbQUIT {
		t.Errorf("Expected QUIT verb (%d), got: %d", VerbQUIT, parser.current.Verb)
	}
}

func TestCleanMessage(t *testing.T) {
	parser := NewMessageParser()
	messages := []string{
		".\r\n",
		"..\r\n.\r\n",
		"......\r\n.\r\n",
		"xx\r\n.\r\n",
		"\n.\r\n",
	}
	for _, m := range messages {
		remaining, err := parser.Feed([]byte(m))
		if err != nil {
			t.Errorf("Valid message doesn't parse %#v: %s", m, err)
		}
		if len(remaining) != 0 {
			t.Errorf("Parse not clean for %#v: remaining `%#v`", m, string(remaining))
		}
	}
}

func TestRemainderMessage(t *testing.T) {
	parser := NewMessageParser()
	messages := []string{
		".\r\nxxx",
		"..\r\n.\r\nxxx",
		"......\r\n.\r\nxxx",
		"\nx\r\nx\r\n.\r\nxxx",
	}
	for _, m := range messages {
		remaining, err := parser.Feed([]byte(m))
		if err != nil {
			t.Errorf("Valid message doesn't parse %#v: %s", m, err)
		}
		if bytes.Compare(remaining, []byte("xxx")) != 0 {
			t.Errorf("Missing remaining for %#v: actual remaining `%#v`", m, string(remaining))
		}
	}
}

func TestChunkedMessage(t *testing.T) {
	parser := NewMessageParser()
	remaining, err := parser.Feed([]byte("..some message\n"))
	if err != Dangling {
		t.Errorf("Expected Dangling return code, got: %s (remaining: %#v)", err, remaining)
	}
	remaining, err = parser.Feed([]byte("more\r\n"))
	if err != Dangling {
		t.Errorf("Expected Dangling return code, got: %s (remaining: %#v)", err, remaining)
	}
	remaining, err = parser.Feed([]byte(""))

	remaining, err = parser.Feed([]byte(".\r\nxxx"))
	if err != nil {
		t.Errorf("Expected nil return code, got: %s (remaining: %#v)", err, remaining)
	}
	if bytes.Compare(remaining, []byte("xxx")) != 0 {
		t.Errorf("Missing remaining for %#v: actual remaining `%#v`", "xxx", string(remaining))
	}

	expected := "..some message\nmore\r\n"
	if parser.buffer.String() != expected {
		t.Errorf("Wrong parse. Expected: %#v, got %#v", expected, parser.buffer.String())
	}
}
