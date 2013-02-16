
// line 1 "smtp_message.ragel"
package smtp

import (
	"fmt"
	"bytes"
)

// Good introduction: http://cr.yp.to/smtp/request.html

// line 20 "smtp_message.ragel"



// line 17 "smtp_message.go"
const smtp_message_start = 1;
const smtp_message_first_final = 8;
const smtp_message_error = 0;

const smtp_message_en_main = 1


// line 23 "smtp_message.ragel"

func NewMessageParser() *MessageParser {
	return &MessageParser{
		cs: smtp_message_en_main,
		buffer: bytes.NewBuffer(nil),
	}
}

// Feed new bytes while keepting track of state in parser.current. If
// error is `Dangling` then more data is needed for a successfull
// parse. When error is nil then parser.buffer contains the full
// message without the terminating .\r\n
//
// See Parser.Feed for gotchas.
func (parser *MessageParser) Feed(data []byte) ([]byte, error) {
	var p int
	cs := parser.cs
	pe := len(data)

	
// line 46 "smtp_message.go"
	{
	if p == pe {
		goto _test_eof
	}
	switch cs {
	case 1:
		goto st_case_1
	case 2:
		goto st_case_2
	case 3:
		goto st_case_3
	case 4:
		goto st_case_4
	case 8:
		goto st_case_8
	case 5:
		goto st_case_5
	case 6:
		goto st_case_6
	case 0:
		goto st_case_0
	case 9:
		goto st_case_9
	case 7:
		goto st_case_7
	}
goto st_end
st_case_1:
	switch data[p] {
	case 10:
goto st3
	case 46:
goto st6
	}
	switch {
	case data[p] > 9:
		if 11 <= data[p] && data[p] <= 45 {
			goto st2
		}
	default:
		goto st2
	}
{
	goto st2

}
st2:
	if p++; p == pe {
		goto _test_eof2
	}
st_case_2:
	if data[p] == 10 {
		goto st3
	}
	if data[p] <= 9 {
		goto st2
	}
{
	goto st2

}
st3:
	if p++; p == pe {
		goto _test_eof3
	}
st_case_3:
	switch data[p] {
	case 10:
goto st3
	case 46:
goto st4
	}
	switch {
	case data[p] > 9:
		if 11 <= data[p] && data[p] <= 45 {
			goto st2
		}
	default:
		goto st2
	}
{
	goto st2

}
st4:
	if p++; p == pe {
		goto _test_eof4
	}
st_case_4:
	switch data[p] {
	case 10:
goto ctr4
	case 13:
goto st5
	}
	switch {
	case data[p] > 9:
		if 11 <= data[p] && data[p] <= 12 {
			goto st2
		}
	default:
		goto st2
	}
{
	goto st2

}
ctr4:
// line 19 "smtp_message.ragel"
	{
{ p++; cs = 8; goto _out }
}
	goto st8
st8:
	if p++; p == pe {
		goto _test_eof8
	}
st_case_8:
// line 165 "smtp_message.go"
	switch data[p] {
	case 10:
goto st3
	case 46:
goto st4
	}
	switch {
	case data[p] > 9:
		if 11 <= data[p] && data[p] <= 45 {
			goto st2
		}
	default:
		goto st2
	}
{
	goto st2

}
st5:
	if p++; p == pe {
		goto _test_eof5
	}
st_case_5:
	if data[p] == 10 {
		goto ctr4
	}
	if data[p] <= 9 {
		goto st2
	}
{
	goto st2

}
st6:
	if p++; p == pe {
		goto _test_eof6
	}
st_case_6:
	switch data[p] {
	case 10:
goto ctr6
	case 13:
goto st7
	case 46:
goto st2
	}
{
	goto st0

}
st_case_0:
st0:
cs = 0
	goto _out
ctr6:
// line 19 "smtp_message.ragel"
	{
{ p++; cs = 9; goto _out }
}
	goto st9
st9:
	if p++; p == pe {
		goto _test_eof9
	}
st_case_9:
// line 231 "smtp_message.go"
{
	goto st0

}
st7:
	if p++; p == pe {
		goto _test_eof7
	}
st_case_7:
	if data[p] == 10 {
		goto ctr6
	}
{
	goto st0

}
st_end:
	_test_eof2: cs = 2
	goto _test_eof
	_test_eof3: cs = 3
	goto _test_eof
	_test_eof4: cs = 4
	goto _test_eof
	_test_eof8: cs = 8
	goto _test_eof
	_test_eof5: cs = 5
	goto _test_eof
	_test_eof6: cs = 6
	goto _test_eof
	_test_eof9: cs = 9
	goto _test_eof
	_test_eof7: cs = 7
	goto _test_eof

	_test_eof: {}
	_out: {}
	}

// line 43 "smtp_message.ragel"

	if cs == smtp_message_error {
		return data[p:], fmt.Errorf("Invalid character in pos %d: `%#v`.", p, data[p])
	}
	if cs < smtp_message_first_final {
		_, err := parser.buffer.Write(data[:p])
		if err != nil {
			return data[p:], fmt.Errorf("Could not buffer data: %s", err)
		}
		parser.cs = cs
		return data[p:], Dangling
	}
	parser.cs = smtp_message_en_main
	return data[p:], nil
}
