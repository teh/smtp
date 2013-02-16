
// line 1 "smtp_message.ragel"
package smtp

import (
	"fmt"
)

// Good introduction: http://cr.yp.to/smtp/request.html

// line 19 "smtp_message.ragel"



// line 16 "smtp_message.go"
const smtp_message_start = 1;
const smtp_message_first_final = 9;
const smtp_message_error = 0;

const smtp_message_en_main = 1


// line 22 "smtp_message.ragel"

// Return remaining etc.
func ParseMessage(data []byte) ([]byte, []byte, error) {
	var cs, p int
	pe := len(data)

	
// line 32 "smtp_message.go"
	{
	cs = smtp_message_start
	}

// line 29 "smtp_message.ragel"
	
// line 39 "smtp_message.go"
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
	case 5:
		goto st_case_5
	case 6:
		goto st_case_6
	case 9:
		goto st_case_9
	case 0:
		goto st_case_0
	case 7:
		goto st_case_7
	case 8:
		goto st_case_8
	case 10:
		goto st_case_10
	}
goto st_end
st_case_1:
	switch data[p] {
	case 10:
goto st0
	case 13:
goto st0
	case 46:
goto st7
	}
	switch {
	case data[p] < 11:
		if data[p] <= 9 {
			goto st2
		}
	case data[p] > 12:
		if 14 <= data[p] && data[p] <= 45 {
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
	if data[p] == 13 {
		goto st3
	}
	if data[p] <= 12 {
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
goto st4
	case 13:
goto st3
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
st4:
	if p++; p == pe {
		goto _test_eof4
	}
st_case_4:
	switch data[p] {
	case 13:
goto st3
	case 46:
goto st5
	}
	switch {
	case data[p] > 12:
		if 14 <= data[p] && data[p] <= 45 {
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
	if data[p] == 13 {
		goto st6
	}
	if data[p] <= 12 {
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
goto ctr7
	case 13:
goto st3
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
ctr7:
// line 18 "smtp_message.ragel"
	{
{ p++; cs = 9; goto _out }
}
	goto st9
st9:
	if p++; p == pe {
		goto _test_eof9
	}
st_case_9:
// line 204 "smtp_message.go"
	switch data[p] {
	case 13:
goto st3
	case 46:
goto st5
	}
	switch {
	case data[p] > 12:
		if 14 <= data[p] && data[p] <= 45 {
			goto st2
		}
	default:
		goto st2
	}
{
	goto st2

}
st_case_0:
st0:
cs = 0
	goto _out
st7:
	if p++; p == pe {
		goto _test_eof7
	}
st_case_7:
	switch data[p] {
	case 13:
goto st8
	case 46:
goto st2
	}
{
	goto st0

}
st8:
	if p++; p == pe {
		goto _test_eof8
	}
st_case_8:
	if data[p] == 10 {
		goto ctr9
	}
{
	goto st0

}
ctr9:
// line 18 "smtp_message.ragel"
	{
{ p++; cs = 10; goto _out }
}
	goto st10
st10:
	if p++; p == pe {
		goto _test_eof10
	}
st_case_10:
// line 265 "smtp_message.go"
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
	_test_eof5: cs = 5
	goto _test_eof
	_test_eof6: cs = 6
	goto _test_eof
	_test_eof9: cs = 9
	goto _test_eof
	_test_eof7: cs = 7
	goto _test_eof
	_test_eof8: cs = 8
	goto _test_eof
	_test_eof10: cs = 10
	goto _test_eof

	_test_eof: {}
	_out: {}
	}

// line 30 "smtp_message.ragel"

	if cs == smtp_message_error {
		return []byte(""), []byte(""), fmt.Errorf("Invalid character in pos %d: `%c`.", p, data[p-1])
	}
	if cs < smtp_message_first_final {
		return []byte(""),  []byte(""), fmt.Errorf("Message doesn't terminate properly.")
	}
	return data[:p], data[p:], nil
}
