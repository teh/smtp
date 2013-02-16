package smtp

import "testing"

func TestValidHelo(t *testing.T) {
	entries := []string{
		"HELO some.domain\r\n",
		"EHLO some.domain\r\n",
		"EHLO [192.167.1.1]\r\n",
		"EHLO [0x7f.1]\r\n",
		"EHLO [0177.0.1]\r\n",
		"EHLO [0177.0.1]\n", // though not valid just \n should still work
	}
	for _, entry := range entries {
		err := ParseEntry([]byte(entry))
		if err != nil {
			t.Errorf("Valid HELO/EHLO doesn't parse: %s: %s", entry, err)
		}
	}
}

func TestInvalidHelo(t *testing.T) {
	entries := []string{
		"HELO  some.domain\r\n",
		"EHLO \r\n",
		"EHLO []\r\n",
	}
	for _, entry := range entries {
		err := ParseEntry([]byte(entry))
		if err == nil {
			t.Errorf("Invalid helo parses: %s", entry)
		}
	}
}

func TestValidMail(t *testing.T) {
	entries := []string{
		"MAIL FROM:<me@example.com>\r\n",
		"MAIL FROM:<root>\r\n",
		"MAIL FROM:<\"ro\"o\\\"t@some>\r\n",
		"MAIL FROM:<\\m\\o\\d@some>\r\n",
		"MAIL FROM:<a+b+c@some>\r\n",
		"MAIL FROM:<@q@@10.com>\r\n",
	}
	for _, entry := range entries {
		err := ParseEntry([]byte(entry))
		if err != nil {
			t.Errorf("Valid email doesn't parse %s: %s", entry, err)
		}
	}
}

func TestInvalidMail(t *testing.T) {
	entries := []string{
		"MAIL FROM: <me@example.com>\r\n",
		"MAIL FROM:<\"me@example.com>\r\n",
		"MAIL FROM:<\"m\"e\"x\"x\"@example.com>\r\n",
	}
	for _, entry := range entries {
		err := ParseEntry([]byte(entry))
		if err == nil {
			t.Errorf("Invalid email parses %s", entry)
		}
	}
}
