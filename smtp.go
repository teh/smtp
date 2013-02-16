
// line 1 "smtp.ragel"
package smtp

import (
	"fmt"
	"errors"
	"bytes"
)

// Good introduction: http://cr.yp.to/smtp/request.html

// line 64 "smtp.ragel"



// line 18 "smtp.go"
const smtp_start = 1;
const smtp_first_final = 144;
const smtp_error = 0;

const smtp_en_main = 1


// line 67 "smtp.ragel"

var Dangling = errors.New("DANGLING")

func NewParser() *Parser {
	return &Parser{cs: smtp_en_main}
}

// Receive data, keep track of state in parser.current. If error is
// Dangling more data is needed for a successfull parse. If error is
// nil then parser.current contains the new action item.
func (parser *Parser) Feed(data []byte) (remaining []byte, err error) {
	var p, pb int
	cs := parser.cs
	pe := len(data)

	
// line 43 "smtp.go"
	{
	if p == pe {
		goto _test_eof
	}
	switch cs {
	case 1:
		goto st_case_1
	case 0:
		goto st_case_0
	case 2:
		goto st_case_2
	case 3:
		goto st_case_3
	case 4:
		goto st_case_4
	case 5:
		goto st_case_5
	case 144:
		goto st_case_144
	case 6:
		goto st_case_6
	case 7:
		goto st_case_7
	case 8:
		goto st_case_8
	case 9:
		goto st_case_9
	case 10:
		goto st_case_10
	case 11:
		goto st_case_11
	case 12:
		goto st_case_12
	case 13:
		goto st_case_13
	case 14:
		goto st_case_14
	case 15:
		goto st_case_15
	case 16:
		goto st_case_16
	case 17:
		goto st_case_17
	case 18:
		goto st_case_18
	case 19:
		goto st_case_19
	case 20:
		goto st_case_20
	case 21:
		goto st_case_21
	case 22:
		goto st_case_22
	case 23:
		goto st_case_23
	case 24:
		goto st_case_24
	case 25:
		goto st_case_25
	case 26:
		goto st_case_26
	case 27:
		goto st_case_27
	case 28:
		goto st_case_28
	case 29:
		goto st_case_29
	case 30:
		goto st_case_30
	case 31:
		goto st_case_31
	case 32:
		goto st_case_32
	case 33:
		goto st_case_33
	case 34:
		goto st_case_34
	case 35:
		goto st_case_35
	case 36:
		goto st_case_36
	case 37:
		goto st_case_37
	case 38:
		goto st_case_38
	case 39:
		goto st_case_39
	case 40:
		goto st_case_40
	case 41:
		goto st_case_41
	case 42:
		goto st_case_42
	case 43:
		goto st_case_43
	case 44:
		goto st_case_44
	case 45:
		goto st_case_45
	case 46:
		goto st_case_46
	case 47:
		goto st_case_47
	case 48:
		goto st_case_48
	case 49:
		goto st_case_49
	case 50:
		goto st_case_50
	case 51:
		goto st_case_51
	case 52:
		goto st_case_52
	case 53:
		goto st_case_53
	case 54:
		goto st_case_54
	case 55:
		goto st_case_55
	case 56:
		goto st_case_56
	case 57:
		goto st_case_57
	case 58:
		goto st_case_58
	case 59:
		goto st_case_59
	case 60:
		goto st_case_60
	case 61:
		goto st_case_61
	case 62:
		goto st_case_62
	case 63:
		goto st_case_63
	case 64:
		goto st_case_64
	case 65:
		goto st_case_65
	case 66:
		goto st_case_66
	case 67:
		goto st_case_67
	case 68:
		goto st_case_68
	case 69:
		goto st_case_69
	case 70:
		goto st_case_70
	case 71:
		goto st_case_71
	case 72:
		goto st_case_72
	case 73:
		goto st_case_73
	case 74:
		goto st_case_74
	case 75:
		goto st_case_75
	case 76:
		goto st_case_76
	case 77:
		goto st_case_77
	case 78:
		goto st_case_78
	case 79:
		goto st_case_79
	case 80:
		goto st_case_80
	case 81:
		goto st_case_81
	case 82:
		goto st_case_82
	case 83:
		goto st_case_83
	case 84:
		goto st_case_84
	case 85:
		goto st_case_85
	case 86:
		goto st_case_86
	case 87:
		goto st_case_87
	case 88:
		goto st_case_88
	case 89:
		goto st_case_89
	case 90:
		goto st_case_90
	case 91:
		goto st_case_91
	case 92:
		goto st_case_92
	case 93:
		goto st_case_93
	case 94:
		goto st_case_94
	case 95:
		goto st_case_95
	case 96:
		goto st_case_96
	case 97:
		goto st_case_97
	case 98:
		goto st_case_98
	case 99:
		goto st_case_99
	case 100:
		goto st_case_100
	case 101:
		goto st_case_101
	case 102:
		goto st_case_102
	case 145:
		goto st_case_145
	case 103:
		goto st_case_103
	case 104:
		goto st_case_104
	case 105:
		goto st_case_105
	case 106:
		goto st_case_106
	case 107:
		goto st_case_107
	case 108:
		goto st_case_108
	case 109:
		goto st_case_109
	case 110:
		goto st_case_110
	case 111:
		goto st_case_111
	case 112:
		goto st_case_112
	case 113:
		goto st_case_113
	case 114:
		goto st_case_114
	case 115:
		goto st_case_115
	case 116:
		goto st_case_116
	case 117:
		goto st_case_117
	case 118:
		goto st_case_118
	case 119:
		goto st_case_119
	case 120:
		goto st_case_120
	case 121:
		goto st_case_121
	case 122:
		goto st_case_122
	case 123:
		goto st_case_123
	case 124:
		goto st_case_124
	case 125:
		goto st_case_125
	case 126:
		goto st_case_126
	case 127:
		goto st_case_127
	case 128:
		goto st_case_128
	case 129:
		goto st_case_129
	case 130:
		goto st_case_130
	case 131:
		goto st_case_131
	case 146:
		goto st_case_146
	case 132:
		goto st_case_132
	case 133:
		goto st_case_133
	case 134:
		goto st_case_134
	case 135:
		goto st_case_135
	case 136:
		goto st_case_136
	case 137:
		goto st_case_137
	case 138:
		goto st_case_138
	case 139:
		goto st_case_139
	case 140:
		goto st_case_140
	case 141:
		goto st_case_141
	case 142:
		goto st_case_142
	case 143:
		goto st_case_143
	}
goto st_end
st_case_1:
	switch data[p] {
	case 68:
goto st2
	case 69:
goto st7
	case 72:
goto st43
	case 77:
goto st79
	case 81:
goto st106
	case 82:
goto st110
	case 86:
goto st138
	}
{
	goto st0

}
st_case_0:
st0:
cs = 0
	goto _out
st2:
	if p++; p == pe {
		goto _test_eof2
	}
st_case_2:
	if data[p] == 65 {
		goto st3
	}
{
	goto st0

}
st3:
	if p++; p == pe {
		goto _test_eof3
	}
st_case_3:
	if data[p] == 84 {
		goto st4
	}
{
	goto st0

}
st4:
	if p++; p == pe {
		goto _test_eof4
	}
st_case_4:
	if data[p] == 65 {
		goto st5
	}
{
	goto st0

}
st5:
	if p++; p == pe {
		goto _test_eof5
	}
st_case_5:
	switch data[p] {
	case 10:
goto ctr11
	case 13:
goto ctr12
	}
{
	goto st0

}
ctr11:
// line 56 "smtp.ragel"
	{
parser.current.Verb = VerbDATA}
// line 63 "smtp.ragel"
	{
{ p++; cs = 144; goto _out }
}
	goto st144
ctr13:
// line 63 "smtp.ragel"
	{
{ p++; cs = 144; goto _out }
}
	goto st144
ctr20:
// line 20 "smtp.ragel"
	{

	parser.buffer = append(parser.buffer, data[pb:p])
	parser.current.Data = bytes.Join(parser.buffer, nil)
	parser.buffer = nil
	parser.recording = false
}
// line 52 "smtp.ragel"
	{
parser.current.Verb = VerbEHLO}
// line 63 "smtp.ragel"
	{
{ p++; cs = 144; goto _out }
}
	goto st144
ctr58:
// line 20 "smtp.ragel"
	{

	parser.buffer = append(parser.buffer, data[pb:p])
	parser.current.Data = bytes.Join(parser.buffer, nil)
	parser.buffer = nil
	parser.recording = false
}
// line 51 "smtp.ragel"
	{
parser.current.Verb = VerbHELO}
// line 63 "smtp.ragel"
	{
{ p++; cs = 144; goto _out }
}
	goto st144
ctr110:
// line 54 "smtp.ragel"
	{
parser.current.Verb = VerbMAIL}
// line 63 "smtp.ragel"
	{
{ p++; cs = 144; goto _out }
}
	goto st144
ctr126:
// line 57 "smtp.ragel"
	{
parser.current.Verb = VerbQUIT}
// line 63 "smtp.ragel"
	{
{ p++; cs = 144; goto _out }
}
	goto st144
ctr147:
// line 55 "smtp.ragel"
	{
parser.current.Verb = VerbRCPT}
// line 63 "smtp.ragel"
	{
{ p++; cs = 144; goto _out }
}
	goto st144
ctr162:
// line 53 "smtp.ragel"
	{
parser.current.Verb = VerbRSET}
// line 63 "smtp.ragel"
	{
{ p++; cs = 144; goto _out }
}
	goto st144
ctr169:
// line 58 "smtp.ragel"
	{
parser.current.Verb = VerbVRFY}
// line 63 "smtp.ragel"
	{
{ p++; cs = 144; goto _out }
}
	goto st144
st144:
	if p++; p == pe {
		goto _test_eof144
	}
st_case_144:
// line 520 "smtp.go"
{
	goto st0

}
ctr12:
// line 56 "smtp.ragel"
	{
parser.current.Verb = VerbDATA}
	goto st6
ctr21:
// line 20 "smtp.ragel"
	{

	parser.buffer = append(parser.buffer, data[pb:p])
	parser.current.Data = bytes.Join(parser.buffer, nil)
	parser.buffer = nil
	parser.recording = false
}
// line 52 "smtp.ragel"
	{
parser.current.Verb = VerbEHLO}
	goto st6
ctr59:
// line 20 "smtp.ragel"
	{

	parser.buffer = append(parser.buffer, data[pb:p])
	parser.current.Data = bytes.Join(parser.buffer, nil)
	parser.buffer = nil
	parser.recording = false
}
// line 51 "smtp.ragel"
	{
parser.current.Verb = VerbHELO}
	goto st6
ctr111:
// line 54 "smtp.ragel"
	{
parser.current.Verb = VerbMAIL}
	goto st6
ctr127:
// line 57 "smtp.ragel"
	{
parser.current.Verb = VerbQUIT}
	goto st6
ctr148:
// line 55 "smtp.ragel"
	{
parser.current.Verb = VerbRCPT}
	goto st6
ctr163:
// line 53 "smtp.ragel"
	{
parser.current.Verb = VerbRSET}
	goto st6
ctr170:
// line 58 "smtp.ragel"
	{
parser.current.Verb = VerbVRFY}
	goto st6
st6:
	if p++; p == pe {
		goto _test_eof6
	}
st_case_6:
// line 586 "smtp.go"
	if data[p] == 10 {
		goto ctr13
	}
{
	goto st0

}
st7:
	if p++; p == pe {
		goto _test_eof7
	}
st_case_7:
	if data[p] == 72 {
		goto st8
	}
{
	goto st0

}
st8:
	if p++; p == pe {
		goto _test_eof8
	}
st_case_8:
	if data[p] == 76 {
		goto st9
	}
{
	goto st0

}
st9:
	if p++; p == pe {
		goto _test_eof9
	}
st_case_9:
	if data[p] == 79 {
		goto st10
	}
{
	goto st0

}
st10:
	if p++; p == pe {
		goto _test_eof10
	}
st_case_10:
	if data[p] == 32 {
		goto st11
	}
{
	goto st0

}
st11:
	if p++; p == pe {
		goto _test_eof11
	}
st_case_11:
	if data[p] == 91 {
		goto ctr19
	}
	switch {
	case data[p] < 65:
		if 48 <= data[p] && data[p] <= 57 {
			goto ctr18
		}
	case data[p] > 90:
		if 97 <= data[p] && data[p] <= 122 {
			goto ctr18
		}
	default:
		goto ctr18
	}
{
	goto st0

}
ctr18:
// line 27 "smtp.ragel"
	{

	pb = p
	parser.recording = true
}
	goto st12
st12:
	if p++; p == pe {
		goto _test_eof12
	}
st_case_12:
// line 679 "smtp.go"
	switch data[p] {
	case 10:
goto ctr20
	case 13:
goto ctr21
	case 46:
goto st13
	}
	switch {
	case data[p] < 65:
		if 48 <= data[p] && data[p] <= 57 {
			goto st12
		}
	case data[p] > 90:
		if 97 <= data[p] && data[p] <= 122 {
			goto st12
		}
	default:
		goto st12
	}
{
	goto st0

}
st13:
	if p++; p == pe {
		goto _test_eof13
	}
st_case_13:
	switch {
	case data[p] < 65:
		if 48 <= data[p] && data[p] <= 57 {
			goto st12
		}
	case data[p] > 90:
		if 97 <= data[p] && data[p] <= 122 {
			goto st12
		}
	default:
		goto st12
	}
{
	goto st0

}
ctr19:
// line 27 "smtp.ragel"
	{

	pb = p
	parser.recording = true
}
	goto st14
st14:
	if p++; p == pe {
		goto _test_eof14
	}
st_case_14:
// line 738 "smtp.go"
	if data[p] == 48 {
		goto st15
	}
	if 49 <= data[p] && data[p] <= 57 {
		goto st38
	}
{
	goto st0

}
st15:
	if p++; p == pe {
		goto _test_eof15
	}
st_case_15:
	switch data[p] {
	case 46:
goto st16
	case 120:
goto st41
	}
	if 48 <= data[p] && data[p] <= 57 {
		goto st38
	}
{
	goto st0

}
st16:
	if p++; p == pe {
		goto _test_eof16
	}
st_case_16:
	if data[p] == 48 {
		goto st17
	}
	if 49 <= data[p] && data[p] <= 57 {
		goto st33
	}
{
	goto st0

}
st17:
	if p++; p == pe {
		goto _test_eof17
	}
st_case_17:
	switch data[p] {
	case 46:
goto st18
	case 93:
goto st25
	case 120:
goto st36
	}
	if 48 <= data[p] && data[p] <= 57 {
		goto st33
	}
{
	goto st0

}
st18:
	if p++; p == pe {
		goto _test_eof18
	}
st_case_18:
	if data[p] == 48 {
		goto st19
	}
	if 49 <= data[p] && data[p] <= 57 {
		goto st28
	}
{
	goto st0

}
st19:
	if p++; p == pe {
		goto _test_eof19
	}
st_case_19:
	switch data[p] {
	case 46:
goto st20
	case 93:
goto st25
	case 120:
goto st31
	}
	if 48 <= data[p] && data[p] <= 57 {
		goto st28
	}
{
	goto st0

}
st20:
	if p++; p == pe {
		goto _test_eof20
	}
st_case_20:
	if data[p] == 48 {
		goto st21
	}
	if 49 <= data[p] && data[p] <= 57 {
		goto st22
	}
{
	goto st0

}
st21:
	if p++; p == pe {
		goto _test_eof21
	}
st_case_21:
	switch data[p] {
	case 93:
goto st25
	case 120:
goto st26
	}
	if 48 <= data[p] && data[p] <= 57 {
		goto st22
	}
{
	goto st0

}
st22:
	if p++; p == pe {
		goto _test_eof22
	}
st_case_22:
	if data[p] == 93 {
		goto st25
	}
	if 48 <= data[p] && data[p] <= 57 {
		goto st23
	}
{
	goto st0

}
st23:
	if p++; p == pe {
		goto _test_eof23
	}
st_case_23:
	if data[p] == 93 {
		goto st25
	}
	if 48 <= data[p] && data[p] <= 57 {
		goto st24
	}
{
	goto st0

}
st24:
	if p++; p == pe {
		goto _test_eof24
	}
st_case_24:
	if data[p] == 93 {
		goto st25
	}
{
	goto st0

}
st25:
	if p++; p == pe {
		goto _test_eof25
	}
st_case_25:
	switch data[p] {
	case 10:
goto ctr20
	case 13:
goto ctr21
	}
{
	goto st0

}
st26:
	if p++; p == pe {
		goto _test_eof26
	}
st_case_26:
	switch {
	case data[p] < 65:
		if 48 <= data[p] && data[p] <= 57 {
			goto st27
		}
	case data[p] > 70:
		if 97 <= data[p] && data[p] <= 102 {
			goto st27
		}
	default:
		goto st27
	}
{
	goto st0

}
st27:
	if p++; p == pe {
		goto _test_eof27
	}
st_case_27:
	if data[p] == 93 {
		goto st25
	}
	switch {
	case data[p] < 65:
		if 48 <= data[p] && data[p] <= 57 {
			goto st24
		}
	case data[p] > 70:
		if 97 <= data[p] && data[p] <= 102 {
			goto st24
		}
	default:
		goto st24
	}
{
	goto st0

}
st28:
	if p++; p == pe {
		goto _test_eof28
	}
st_case_28:
	switch data[p] {
	case 46:
goto st20
	case 93:
goto st25
	}
	if 48 <= data[p] && data[p] <= 57 {
		goto st29
	}
{
	goto st0

}
st29:
	if p++; p == pe {
		goto _test_eof29
	}
st_case_29:
	switch data[p] {
	case 46:
goto st20
	case 93:
goto st25
	}
	if 48 <= data[p] && data[p] <= 57 {
		goto st30
	}
{
	goto st0

}
st30:
	if p++; p == pe {
		goto _test_eof30
	}
st_case_30:
	switch data[p] {
	case 46:
goto st20
	case 93:
goto st25
	}
{
	goto st0

}
st31:
	if p++; p == pe {
		goto _test_eof31
	}
st_case_31:
	switch {
	case data[p] < 65:
		if 48 <= data[p] && data[p] <= 57 {
			goto st32
		}
	case data[p] > 70:
		if 97 <= data[p] && data[p] <= 102 {
			goto st32
		}
	default:
		goto st32
	}
{
	goto st0

}
st32:
	if p++; p == pe {
		goto _test_eof32
	}
st_case_32:
	switch data[p] {
	case 46:
goto st20
	case 93:
goto st25
	}
	switch {
	case data[p] < 65:
		if 48 <= data[p] && data[p] <= 57 {
			goto st30
		}
	case data[p] > 70:
		if 97 <= data[p] && data[p] <= 102 {
			goto st30
		}
	default:
		goto st30
	}
{
	goto st0

}
st33:
	if p++; p == pe {
		goto _test_eof33
	}
st_case_33:
	switch data[p] {
	case 46:
goto st18
	case 93:
goto st25
	}
	if 48 <= data[p] && data[p] <= 57 {
		goto st34
	}
{
	goto st0

}
st34:
	if p++; p == pe {
		goto _test_eof34
	}
st_case_34:
	switch data[p] {
	case 46:
goto st18
	case 93:
goto st25
	}
	if 48 <= data[p] && data[p] <= 57 {
		goto st35
	}
{
	goto st0

}
st35:
	if p++; p == pe {
		goto _test_eof35
	}
st_case_35:
	switch data[p] {
	case 46:
goto st18
	case 93:
goto st25
	}
{
	goto st0

}
st36:
	if p++; p == pe {
		goto _test_eof36
	}
st_case_36:
	switch {
	case data[p] < 65:
		if 48 <= data[p] && data[p] <= 57 {
			goto st37
		}
	case data[p] > 70:
		if 97 <= data[p] && data[p] <= 102 {
			goto st37
		}
	default:
		goto st37
	}
{
	goto st0

}
st37:
	if p++; p == pe {
		goto _test_eof37
	}
st_case_37:
	switch data[p] {
	case 46:
goto st18
	case 93:
goto st25
	}
	switch {
	case data[p] < 65:
		if 48 <= data[p] && data[p] <= 57 {
			goto st35
		}
	case data[p] > 70:
		if 97 <= data[p] && data[p] <= 102 {
			goto st35
		}
	default:
		goto st35
	}
{
	goto st0

}
st38:
	if p++; p == pe {
		goto _test_eof38
	}
st_case_38:
	if data[p] == 46 {
		goto st16
	}
	if 48 <= data[p] && data[p] <= 57 {
		goto st39
	}
{
	goto st0

}
st39:
	if p++; p == pe {
		goto _test_eof39
	}
st_case_39:
	if data[p] == 46 {
		goto st16
	}
	if 48 <= data[p] && data[p] <= 57 {
		goto st40
	}
{
	goto st0

}
st40:
	if p++; p == pe {
		goto _test_eof40
	}
st_case_40:
	if data[p] == 46 {
		goto st16
	}
{
	goto st0

}
st41:
	if p++; p == pe {
		goto _test_eof41
	}
st_case_41:
	switch {
	case data[p] < 65:
		if 48 <= data[p] && data[p] <= 57 {
			goto st42
		}
	case data[p] > 70:
		if 97 <= data[p] && data[p] <= 102 {
			goto st42
		}
	default:
		goto st42
	}
{
	goto st0

}
st42:
	if p++; p == pe {
		goto _test_eof42
	}
st_case_42:
	if data[p] == 46 {
		goto st16
	}
	switch {
	case data[p] < 65:
		if 48 <= data[p] && data[p] <= 57 {
			goto st40
		}
	case data[p] > 70:
		if 97 <= data[p] && data[p] <= 102 {
			goto st40
		}
	default:
		goto st40
	}
{
	goto st0

}
st43:
	if p++; p == pe {
		goto _test_eof43
	}
st_case_43:
	if data[p] == 69 {
		goto st44
	}
{
	goto st0

}
st44:
	if p++; p == pe {
		goto _test_eof44
	}
st_case_44:
	if data[p] == 76 {
		goto st45
	}
{
	goto st0

}
st45:
	if p++; p == pe {
		goto _test_eof45
	}
st_case_45:
	if data[p] == 79 {
		goto st46
	}
{
	goto st0

}
st46:
	if p++; p == pe {
		goto _test_eof46
	}
st_case_46:
	if data[p] == 32 {
		goto st47
	}
{
	goto st0

}
st47:
	if p++; p == pe {
		goto _test_eof47
	}
st_case_47:
	if data[p] == 91 {
		goto ctr57
	}
	switch {
	case data[p] < 65:
		if 48 <= data[p] && data[p] <= 57 {
			goto ctr56
		}
	case data[p] > 90:
		if 97 <= data[p] && data[p] <= 122 {
			goto ctr56
		}
	default:
		goto ctr56
	}
{
	goto st0

}
ctr56:
// line 27 "smtp.ragel"
	{

	pb = p
	parser.recording = true
}
	goto st48
st48:
	if p++; p == pe {
		goto _test_eof48
	}
st_case_48:
// line 1342 "smtp.go"
	switch data[p] {
	case 10:
goto ctr58
	case 13:
goto ctr59
	case 46:
goto st49
	}
	switch {
	case data[p] < 65:
		if 48 <= data[p] && data[p] <= 57 {
			goto st48
		}
	case data[p] > 90:
		if 97 <= data[p] && data[p] <= 122 {
			goto st48
		}
	default:
		goto st48
	}
{
	goto st0

}
st49:
	if p++; p == pe {
		goto _test_eof49
	}
st_case_49:
	switch {
	case data[p] < 65:
		if 48 <= data[p] && data[p] <= 57 {
			goto st48
		}
	case data[p] > 90:
		if 97 <= data[p] && data[p] <= 122 {
			goto st48
		}
	default:
		goto st48
	}
{
	goto st0

}
ctr57:
// line 27 "smtp.ragel"
	{

	pb = p
	parser.recording = true
}
	goto st50
st50:
	if p++; p == pe {
		goto _test_eof50
	}
st_case_50:
// line 1401 "smtp.go"
	if data[p] == 48 {
		goto st51
	}
	if 49 <= data[p] && data[p] <= 57 {
		goto st74
	}
{
	goto st0

}
st51:
	if p++; p == pe {
		goto _test_eof51
	}
st_case_51:
	switch data[p] {
	case 46:
goto st52
	case 120:
goto st77
	}
	if 48 <= data[p] && data[p] <= 57 {
		goto st74
	}
{
	goto st0

}
st52:
	if p++; p == pe {
		goto _test_eof52
	}
st_case_52:
	if data[p] == 48 {
		goto st53
	}
	if 49 <= data[p] && data[p] <= 57 {
		goto st69
	}
{
	goto st0

}
st53:
	if p++; p == pe {
		goto _test_eof53
	}
st_case_53:
	switch data[p] {
	case 46:
goto st54
	case 93:
goto st61
	case 120:
goto st72
	}
	if 48 <= data[p] && data[p] <= 57 {
		goto st69
	}
{
	goto st0

}
st54:
	if p++; p == pe {
		goto _test_eof54
	}
st_case_54:
	if data[p] == 48 {
		goto st55
	}
	if 49 <= data[p] && data[p] <= 57 {
		goto st64
	}
{
	goto st0

}
st55:
	if p++; p == pe {
		goto _test_eof55
	}
st_case_55:
	switch data[p] {
	case 46:
goto st56
	case 93:
goto st61
	case 120:
goto st67
	}
	if 48 <= data[p] && data[p] <= 57 {
		goto st64
	}
{
	goto st0

}
st56:
	if p++; p == pe {
		goto _test_eof56
	}
st_case_56:
	if data[p] == 48 {
		goto st57
	}
	if 49 <= data[p] && data[p] <= 57 {
		goto st58
	}
{
	goto st0

}
st57:
	if p++; p == pe {
		goto _test_eof57
	}
st_case_57:
	switch data[p] {
	case 93:
goto st61
	case 120:
goto st62
	}
	if 48 <= data[p] && data[p] <= 57 {
		goto st58
	}
{
	goto st0

}
st58:
	if p++; p == pe {
		goto _test_eof58
	}
st_case_58:
	if data[p] == 93 {
		goto st61
	}
	if 48 <= data[p] && data[p] <= 57 {
		goto st59
	}
{
	goto st0

}
st59:
	if p++; p == pe {
		goto _test_eof59
	}
st_case_59:
	if data[p] == 93 {
		goto st61
	}
	if 48 <= data[p] && data[p] <= 57 {
		goto st60
	}
{
	goto st0

}
st60:
	if p++; p == pe {
		goto _test_eof60
	}
st_case_60:
	if data[p] == 93 {
		goto st61
	}
{
	goto st0

}
st61:
	if p++; p == pe {
		goto _test_eof61
	}
st_case_61:
	switch data[p] {
	case 10:
goto ctr58
	case 13:
goto ctr59
	}
{
	goto st0

}
st62:
	if p++; p == pe {
		goto _test_eof62
	}
st_case_62:
	switch {
	case data[p] < 65:
		if 48 <= data[p] && data[p] <= 57 {
			goto st63
		}
	case data[p] > 70:
		if 97 <= data[p] && data[p] <= 102 {
			goto st63
		}
	default:
		goto st63
	}
{
	goto st0

}
st63:
	if p++; p == pe {
		goto _test_eof63
	}
st_case_63:
	if data[p] == 93 {
		goto st61
	}
	switch {
	case data[p] < 65:
		if 48 <= data[p] && data[p] <= 57 {
			goto st60
		}
	case data[p] > 70:
		if 97 <= data[p] && data[p] <= 102 {
			goto st60
		}
	default:
		goto st60
	}
{
	goto st0

}
st64:
	if p++; p == pe {
		goto _test_eof64
	}
st_case_64:
	switch data[p] {
	case 46:
goto st56
	case 93:
goto st61
	}
	if 48 <= data[p] && data[p] <= 57 {
		goto st65
	}
{
	goto st0

}
st65:
	if p++; p == pe {
		goto _test_eof65
	}
st_case_65:
	switch data[p] {
	case 46:
goto st56
	case 93:
goto st61
	}
	if 48 <= data[p] && data[p] <= 57 {
		goto st66
	}
{
	goto st0

}
st66:
	if p++; p == pe {
		goto _test_eof66
	}
st_case_66:
	switch data[p] {
	case 46:
goto st56
	case 93:
goto st61
	}
{
	goto st0

}
st67:
	if p++; p == pe {
		goto _test_eof67
	}
st_case_67:
	switch {
	case data[p] < 65:
		if 48 <= data[p] && data[p] <= 57 {
			goto st68
		}
	case data[p] > 70:
		if 97 <= data[p] && data[p] <= 102 {
			goto st68
		}
	default:
		goto st68
	}
{
	goto st0

}
st68:
	if p++; p == pe {
		goto _test_eof68
	}
st_case_68:
	switch data[p] {
	case 46:
goto st56
	case 93:
goto st61
	}
	switch {
	case data[p] < 65:
		if 48 <= data[p] && data[p] <= 57 {
			goto st66
		}
	case data[p] > 70:
		if 97 <= data[p] && data[p] <= 102 {
			goto st66
		}
	default:
		goto st66
	}
{
	goto st0

}
st69:
	if p++; p == pe {
		goto _test_eof69
	}
st_case_69:
	switch data[p] {
	case 46:
goto st54
	case 93:
goto st61
	}
	if 48 <= data[p] && data[p] <= 57 {
		goto st70
	}
{
	goto st0

}
st70:
	if p++; p == pe {
		goto _test_eof70
	}
st_case_70:
	switch data[p] {
	case 46:
goto st54
	case 93:
goto st61
	}
	if 48 <= data[p] && data[p] <= 57 {
		goto st71
	}
{
	goto st0

}
st71:
	if p++; p == pe {
		goto _test_eof71
	}
st_case_71:
	switch data[p] {
	case 46:
goto st54
	case 93:
goto st61
	}
{
	goto st0

}
st72:
	if p++; p == pe {
		goto _test_eof72
	}
st_case_72:
	switch {
	case data[p] < 65:
		if 48 <= data[p] && data[p] <= 57 {
			goto st73
		}
	case data[p] > 70:
		if 97 <= data[p] && data[p] <= 102 {
			goto st73
		}
	default:
		goto st73
	}
{
	goto st0

}
st73:
	if p++; p == pe {
		goto _test_eof73
	}
st_case_73:
	switch data[p] {
	case 46:
goto st54
	case 93:
goto st61
	}
	switch {
	case data[p] < 65:
		if 48 <= data[p] && data[p] <= 57 {
			goto st71
		}
	case data[p] > 70:
		if 97 <= data[p] && data[p] <= 102 {
			goto st71
		}
	default:
		goto st71
	}
{
	goto st0

}
st74:
	if p++; p == pe {
		goto _test_eof74
	}
st_case_74:
	if data[p] == 46 {
		goto st52
	}
	if 48 <= data[p] && data[p] <= 57 {
		goto st75
	}
{
	goto st0

}
st75:
	if p++; p == pe {
		goto _test_eof75
	}
st_case_75:
	if data[p] == 46 {
		goto st52
	}
	if 48 <= data[p] && data[p] <= 57 {
		goto st76
	}
{
	goto st0

}
st76:
	if p++; p == pe {
		goto _test_eof76
	}
st_case_76:
	if data[p] == 46 {
		goto st52
	}
{
	goto st0

}
st77:
	if p++; p == pe {
		goto _test_eof77
	}
st_case_77:
	switch {
	case data[p] < 65:
		if 48 <= data[p] && data[p] <= 57 {
			goto st78
		}
	case data[p] > 70:
		if 97 <= data[p] && data[p] <= 102 {
			goto st78
		}
	default:
		goto st78
	}
{
	goto st0

}
st78:
	if p++; p == pe {
		goto _test_eof78
	}
st_case_78:
	if data[p] == 46 {
		goto st52
	}
	switch {
	case data[p] < 65:
		if 48 <= data[p] && data[p] <= 57 {
			goto st76
		}
	case data[p] > 70:
		if 97 <= data[p] && data[p] <= 102 {
			goto st76
		}
	default:
		goto st76
	}
{
	goto st0

}
st79:
	if p++; p == pe {
		goto _test_eof79
	}
st_case_79:
	if data[p] == 65 {
		goto st80
	}
{
	goto st0

}
st80:
	if p++; p == pe {
		goto _test_eof80
	}
st_case_80:
	if data[p] == 73 {
		goto st81
	}
{
	goto st0

}
st81:
	if p++; p == pe {
		goto _test_eof81
	}
st_case_81:
	if data[p] == 76 {
		goto st82
	}
{
	goto st0

}
st82:
	if p++; p == pe {
		goto _test_eof82
	}
st_case_82:
	if data[p] == 32 {
		goto st83
	}
{
	goto st0

}
st83:
	if p++; p == pe {
		goto _test_eof83
	}
st_case_83:
	switch data[p] {
	case 70:
goto st84
	case 102:
goto st84
	}
{
	goto st0

}
st84:
	if p++; p == pe {
		goto _test_eof84
	}
st_case_84:
	switch data[p] {
	case 82:
goto st85
	case 114:
goto st85
	}
{
	goto st0

}
st85:
	if p++; p == pe {
		goto _test_eof85
	}
st_case_85:
	switch data[p] {
	case 79:
goto st86
	case 111:
goto st86
	}
{
	goto st0

}
st86:
	if p++; p == pe {
		goto _test_eof86
	}
st_case_86:
	switch data[p] {
	case 77:
goto st87
	case 109:
goto st87
	}
{
	goto st0

}
st87:
	if p++; p == pe {
		goto _test_eof87
	}
st_case_87:
	if data[p] == 58 {
		goto st88
	}
{
	goto st0

}
st88:
	if p++; p == pe {
		goto _test_eof88
	}
st_case_88:
	if data[p] == 60 {
		goto st89
	}
{
	goto st0

}
st89:
	if p++; p == pe {
		goto _test_eof89
	}
st_case_89:
	switch data[p] {
	case 34:
goto ctr100
	case 43:
goto ctr101
	case 46:
goto ctr101
	case 64:
goto ctr102
	case 92:
goto ctr103
	}
	switch {
	case data[p] < 65:
		if 48 <= data[p] && data[p] <= 57 {
			goto ctr101
		}
	case data[p] > 90:
		if 97 <= data[p] && data[p] <= 122 {
			goto ctr101
		}
	default:
		goto ctr101
	}
{
	goto st0

}
ctr100:
// line 27 "smtp.ragel"
	{

	pb = p
	parser.recording = true
}
	goto st90
st90:
	if p++; p == pe {
		goto _test_eof90
	}
st_case_90:
// line 2098 "smtp.go"
	switch data[p] {
	case 43:
goto st91
	case 46:
goto st91
	case 92:
goto st95
	}
	switch {
	case data[p] < 64:
		if 48 <= data[p] && data[p] <= 57 {
			goto st91
		}
	case data[p] > 90:
		if 97 <= data[p] && data[p] <= 122 {
			goto st91
		}
	default:
		goto st91
	}
{
	goto st0

}
st91:
	if p++; p == pe {
		goto _test_eof91
	}
st_case_91:
	switch data[p] {
	case 34:
goto st92
	case 43:
goto st91
	case 46:
goto st91
	case 92:
goto st95
	}
	switch {
	case data[p] < 64:
		if 48 <= data[p] && data[p] <= 57 {
			goto st91
		}
	case data[p] > 90:
		if 97 <= data[p] && data[p] <= 122 {
			goto st91
		}
	default:
		goto st91
	}
{
	goto st0

}
ctr101:
// line 27 "smtp.ragel"
	{

	pb = p
	parser.recording = true
}
	goto st92
st92:
	if p++; p == pe {
		goto _test_eof92
	}
st_case_92:
// line 2167 "smtp.go"
	switch data[p] {
	case 34:
goto st90
	case 43:
goto st92
	case 46:
goto st92
	case 62:
goto ctr108
	case 92:
goto st94
	}
	switch {
	case data[p] < 64:
		if 48 <= data[p] && data[p] <= 57 {
			goto st92
		}
	case data[p] > 90:
		if 97 <= data[p] && data[p] <= 122 {
			goto st92
		}
	default:
		goto st92
	}
{
	goto st0

}
ctr108:
// line 20 "smtp.ragel"
	{

	parser.buffer = append(parser.buffer, data[pb:p])
	parser.current.Data = bytes.Join(parser.buffer, nil)
	parser.buffer = nil
	parser.recording = false
}
	goto st93
st93:
	if p++; p == pe {
		goto _test_eof93
	}
st_case_93:
// line 2211 "smtp.go"
	switch data[p] {
	case 10:
goto ctr110
	case 13:
goto ctr111
	}
{
	goto st0

}
ctr103:
// line 27 "smtp.ragel"
	{

	pb = p
	parser.recording = true
}
	goto st94
st94:
	if p++; p == pe {
		goto _test_eof94
	}
st_case_94:
// line 2235 "smtp.go"
	switch data[p] {
	case 32:
goto st92
	case 34:
goto st92
	case 43:
goto st92
	case 46:
goto st92
	}
	switch {
	case data[p] < 64:
		if 48 <= data[p] && data[p] <= 57 {
			goto st92
		}
	case data[p] > 90:
		if 97 <= data[p] && data[p] <= 122 {
			goto st92
		}
	default:
		goto st92
	}
{
	goto st0

}
st95:
	if p++; p == pe {
		goto _test_eof95
	}
st_case_95:
	switch data[p] {
	case 32:
goto st91
	case 34:
goto st91
	case 43:
goto st91
	case 46:
goto st91
	}
	switch {
	case data[p] < 64:
		if 48 <= data[p] && data[p] <= 57 {
			goto st91
		}
	case data[p] > 90:
		if 97 <= data[p] && data[p] <= 122 {
			goto st91
		}
	default:
		goto st91
	}
{
	goto st0

}
ctr102:
// line 27 "smtp.ragel"
	{

	pb = p
	parser.recording = true
}
	goto st96
st96:
	if p++; p == pe {
		goto _test_eof96
	}
st_case_96:
// line 2306 "smtp.go"
	switch data[p] {
	case 34:
goto st99
	case 43:
goto st101
	case 46:
goto st101
	case 47:
goto st97
	case 58:
goto st0
	case 62:
goto ctr115
	case 63:
goto st97
	case 91:
goto st97
	case 92:
goto st104
	}
	switch {
	case data[p] < 48:
		switch {
		case data[p] < 35:
			if data[p] <= 33 {
				goto st97
			}
		case data[p] > 42:
			if 44 <= data[p] && data[p] <= 45 {
				goto st97
			}
		default:
			goto st97
		}
	case data[p] > 57:
		switch {
		case data[p] < 64:
			if 59 <= data[p] && data[p] <= 61 {
				goto st97
			}
		case data[p] > 90:
			switch {
			case data[p] > 96:
				if data[p] <= 122 {
					goto st101
				}
			case data[p] >= 93:
				goto st97
			}
		default:
			goto st101
		}
	default:
		goto st101
	}
{
	goto st97

}
st97:
	if p++; p == pe {
		goto _test_eof97
	}
st_case_97:
	if data[p] == 58 {
		goto st98
	}
	if data[p] <= 57 {
		goto st97
	}
{
	goto st97

}
st98:
	if p++; p == pe {
		goto _test_eof98
	}
st_case_98:
	switch data[p] {
	case 34:
goto ctr100
	case 43:
goto ctr101
	case 46:
goto ctr101
	case 92:
goto ctr103
	}
	switch {
	case data[p] < 64:
		if 48 <= data[p] && data[p] <= 57 {
			goto ctr101
		}
	case data[p] > 90:
		if 97 <= data[p] && data[p] <= 122 {
			goto ctr101
		}
	default:
		goto ctr101
	}
{
	goto st0

}
st99:
	if p++; p == pe {
		goto _test_eof99
	}
st_case_99:
	switch data[p] {
	case 43:
goto st100
	case 46:
goto st100
	case 47:
goto st97
	case 58:
goto st98
	case 91:
goto st97
	case 92:
goto st105
	}
	switch {
	case data[p] < 59:
		switch {
		case data[p] < 44:
			if data[p] <= 42 {
				goto st97
			}
		case data[p] > 45:
			if 48 <= data[p] && data[p] <= 57 {
				goto st100
			}
		default:
			goto st97
		}
	case data[p] > 63:
		switch {
		case data[p] < 93:
			if data[p] <= 90 {
				goto st100
			}
		case data[p] > 96:
			if data[p] <= 122 {
				goto st100
			}
		default:
			goto st97
		}
	default:
		goto st97
	}
{
	goto st97

}
st100:
	if p++; p == pe {
		goto _test_eof100
	}
st_case_100:
	switch data[p] {
	case 34:
goto st101
	case 43:
goto st100
	case 46:
goto st100
	case 47:
goto st97
	case 58:
goto st98
	case 91:
goto st97
	case 92:
goto st105
	}
	switch {
	case data[p] < 48:
		switch {
		case data[p] < 35:
			if data[p] <= 33 {
				goto st97
			}
		case data[p] > 42:
			if 44 <= data[p] && data[p] <= 45 {
				goto st97
			}
		default:
			goto st97
		}
	case data[p] > 57:
		switch {
		case data[p] < 64:
			if 59 <= data[p] {
				goto st97
			}
		case data[p] > 90:
			switch {
			case data[p] > 96:
				if data[p] <= 122 {
					goto st100
				}
			case data[p] >= 93:
				goto st97
			}
		default:
			goto st100
		}
	default:
		goto st100
	}
{
	goto st97

}
st101:
	if p++; p == pe {
		goto _test_eof101
	}
st_case_101:
	switch data[p] {
	case 34:
goto st99
	case 43:
goto st101
	case 46:
goto st101
	case 47:
goto st97
	case 58:
goto st98
	case 62:
goto ctr115
	case 63:
goto st97
	case 91:
goto st97
	case 92:
goto st104
	}
	switch {
	case data[p] < 48:
		switch {
		case data[p] < 35:
			if data[p] <= 33 {
				goto st97
			}
		case data[p] > 42:
			if 44 <= data[p] && data[p] <= 45 {
				goto st97
			}
		default:
			goto st97
		}
	case data[p] > 57:
		switch {
		case data[p] < 64:
			if 59 <= data[p] && data[p] <= 61 {
				goto st97
			}
		case data[p] > 90:
			switch {
			case data[p] > 96:
				if data[p] <= 122 {
					goto st101
				}
			case data[p] >= 93:
				goto st97
			}
		default:
			goto st101
		}
	default:
		goto st101
	}
{
	goto st97

}
ctr115:
// line 20 "smtp.ragel"
	{

	parser.buffer = append(parser.buffer, data[pb:p])
	parser.current.Data = bytes.Join(parser.buffer, nil)
	parser.buffer = nil
	parser.recording = false
}
	goto st102
st102:
	if p++; p == pe {
		goto _test_eof102
	}
st_case_102:
// line 2604 "smtp.go"
	switch data[p] {
	case 10:
goto ctr120
	case 13:
goto ctr121
	case 58:
goto st98
	}
	switch {
	case data[p] < 11:
		if data[p] <= 9 {
			goto st97
		}
	case data[p] > 12:
		if 14 <= data[p] && data[p] <= 57 {
			goto st97
		}
	default:
		goto st97
	}
{
	goto st97

}
ctr122:
// line 63 "smtp.ragel"
	{
{ p++; cs = 145; goto _out }
}
	goto st145
ctr120:
// line 54 "smtp.ragel"
	{
parser.current.Verb = VerbMAIL}
// line 63 "smtp.ragel"
	{
{ p++; cs = 145; goto _out }
}
	goto st145
st145:
	if p++; p == pe {
		goto _test_eof145
	}
st_case_145:
// line 2649 "smtp.go"
	if data[p] == 58 {
		goto st98
	}
	if data[p] <= 57 {
		goto st97
	}
{
	goto st97

}
ctr121:
// line 54 "smtp.ragel"
	{
parser.current.Verb = VerbMAIL}
	goto st103
st103:
	if p++; p == pe {
		goto _test_eof103
	}
st_case_103:
// line 2670 "smtp.go"
	switch data[p] {
	case 10:
goto ctr122
	case 58:
goto st98
	}
	switch {
	case data[p] > 9:
		if 11 <= data[p] && data[p] <= 57 {
			goto st97
		}
	default:
		goto st97
	}
{
	goto st97

}
st104:
	if p++; p == pe {
		goto _test_eof104
	}
st_case_104:
	switch data[p] {
	case 32:
goto st101
	case 33:
goto st97
	case 34:
goto st101
	case 43:
goto st101
	case 46:
goto st101
	case 47:
goto st97
	case 58:
goto st98
	}
	switch {
	case data[p] < 48:
		switch {
		case data[p] < 35:
			if data[p] <= 31 {
				goto st97
			}
		case data[p] > 42:
			if 44 <= data[p] && data[p] <= 45 {
				goto st97
			}
		default:
			goto st97
		}
	case data[p] > 57:
		switch {
		case data[p] < 64:
			if 59 <= data[p] {
				goto st97
			}
		case data[p] > 90:
			switch {
			case data[p] > 96:
				if data[p] <= 122 {
					goto st101
				}
			default:
				goto st97
			}
		default:
			goto st101
		}
	default:
		goto st101
	}
{
	goto st97

}
st105:
	if p++; p == pe {
		goto _test_eof105
	}
st_case_105:
	switch data[p] {
	case 32:
goto st100
	case 33:
goto st97
	case 34:
goto st100
	case 43:
goto st100
	case 46:
goto st100
	case 47:
goto st97
	case 58:
goto st98
	}
	switch {
	case data[p] < 48:
		switch {
		case data[p] < 35:
			if data[p] <= 31 {
				goto st97
			}
		case data[p] > 42:
			if 44 <= data[p] && data[p] <= 45 {
				goto st97
			}
		default:
			goto st97
		}
	case data[p] > 57:
		switch {
		case data[p] < 64:
			if 59 <= data[p] {
				goto st97
			}
		case data[p] > 90:
			switch {
			case data[p] > 96:
				if data[p] <= 122 {
					goto st100
				}
			default:
				goto st97
			}
		default:
			goto st100
		}
	default:
		goto st100
	}
{
	goto st97

}
st106:
	if p++; p == pe {
		goto _test_eof106
	}
st_case_106:
	if data[p] == 85 {
		goto st107
	}
{
	goto st0

}
st107:
	if p++; p == pe {
		goto _test_eof107
	}
st_case_107:
	if data[p] == 73 {
		goto st108
	}
{
	goto st0

}
st108:
	if p++; p == pe {
		goto _test_eof108
	}
st_case_108:
	if data[p] == 84 {
		goto st109
	}
{
	goto st0

}
st109:
	if p++; p == pe {
		goto _test_eof109
	}
st_case_109:
	switch data[p] {
	case 10:
goto ctr126
	case 13:
goto ctr127
	}
{
	goto st0

}
st110:
	if p++; p == pe {
		goto _test_eof110
	}
st_case_110:
	switch data[p] {
	case 67:
goto st111
	case 83:
goto st135
	}
{
	goto st0

}
st111:
	if p++; p == pe {
		goto _test_eof111
	}
st_case_111:
	if data[p] == 80 {
		goto st112
	}
{
	goto st0

}
st112:
	if p++; p == pe {
		goto _test_eof112
	}
st_case_112:
	if data[p] == 84 {
		goto st113
	}
{
	goto st0

}
st113:
	if p++; p == pe {
		goto _test_eof113
	}
st_case_113:
	if data[p] == 32 {
		goto st114
	}
{
	goto st0

}
st114:
	if p++; p == pe {
		goto _test_eof114
	}
st_case_114:
	switch data[p] {
	case 84:
goto st115
	case 116:
goto st115
	}
{
	goto st0

}
st115:
	if p++; p == pe {
		goto _test_eof115
	}
st_case_115:
	switch data[p] {
	case 79:
goto st116
	case 111:
goto st116
	}
{
	goto st0

}
st116:
	if p++; p == pe {
		goto _test_eof116
	}
st_case_116:
	if data[p] == 58 {
		goto st117
	}
{
	goto st0

}
st117:
	if p++; p == pe {
		goto _test_eof117
	}
st_case_117:
	if data[p] == 60 {
		goto st118
	}
{
	goto st0

}
st118:
	if p++; p == pe {
		goto _test_eof118
	}
st_case_118:
	switch data[p] {
	case 34:
goto ctr137
	case 43:
goto ctr138
	case 46:
goto ctr138
	case 64:
goto ctr139
	case 92:
goto ctr140
	}
	switch {
	case data[p] < 65:
		if 48 <= data[p] && data[p] <= 57 {
			goto ctr138
		}
	case data[p] > 90:
		if 97 <= data[p] && data[p] <= 122 {
			goto ctr138
		}
	default:
		goto ctr138
	}
{
	goto st0

}
ctr137:
// line 27 "smtp.ragel"
	{

	pb = p
	parser.recording = true
}
	goto st119
st119:
	if p++; p == pe {
		goto _test_eof119
	}
st_case_119:
// line 3011 "smtp.go"
	switch data[p] {
	case 43:
goto st120
	case 46:
goto st120
	case 92:
goto st124
	}
	switch {
	case data[p] < 64:
		if 48 <= data[p] && data[p] <= 57 {
			goto st120
		}
	case data[p] > 90:
		if 97 <= data[p] && data[p] <= 122 {
			goto st120
		}
	default:
		goto st120
	}
{
	goto st0

}
st120:
	if p++; p == pe {
		goto _test_eof120
	}
st_case_120:
	switch data[p] {
	case 34:
goto st121
	case 43:
goto st120
	case 46:
goto st120
	case 92:
goto st124
	}
	switch {
	case data[p] < 64:
		if 48 <= data[p] && data[p] <= 57 {
			goto st120
		}
	case data[p] > 90:
		if 97 <= data[p] && data[p] <= 122 {
			goto st120
		}
	default:
		goto st120
	}
{
	goto st0

}
ctr138:
// line 27 "smtp.ragel"
	{

	pb = p
	parser.recording = true
}
	goto st121
st121:
	if p++; p == pe {
		goto _test_eof121
	}
st_case_121:
// line 3080 "smtp.go"
	switch data[p] {
	case 34:
goto st119
	case 43:
goto st121
	case 46:
goto st121
	case 62:
goto ctr145
	case 92:
goto st123
	}
	switch {
	case data[p] < 64:
		if 48 <= data[p] && data[p] <= 57 {
			goto st121
		}
	case data[p] > 90:
		if 97 <= data[p] && data[p] <= 122 {
			goto st121
		}
	default:
		goto st121
	}
{
	goto st0

}
ctr145:
// line 20 "smtp.ragel"
	{

	parser.buffer = append(parser.buffer, data[pb:p])
	parser.current.Data = bytes.Join(parser.buffer, nil)
	parser.buffer = nil
	parser.recording = false
}
	goto st122
st122:
	if p++; p == pe {
		goto _test_eof122
	}
st_case_122:
// line 3124 "smtp.go"
	switch data[p] {
	case 10:
goto ctr147
	case 13:
goto ctr148
	}
{
	goto st0

}
ctr140:
// line 27 "smtp.ragel"
	{

	pb = p
	parser.recording = true
}
	goto st123
st123:
	if p++; p == pe {
		goto _test_eof123
	}
st_case_123:
// line 3148 "smtp.go"
	switch data[p] {
	case 32:
goto st121
	case 34:
goto st121
	case 43:
goto st121
	case 46:
goto st121
	}
	switch {
	case data[p] < 64:
		if 48 <= data[p] && data[p] <= 57 {
			goto st121
		}
	case data[p] > 90:
		if 97 <= data[p] && data[p] <= 122 {
			goto st121
		}
	default:
		goto st121
	}
{
	goto st0

}
st124:
	if p++; p == pe {
		goto _test_eof124
	}
st_case_124:
	switch data[p] {
	case 32:
goto st120
	case 34:
goto st120
	case 43:
goto st120
	case 46:
goto st120
	}
	switch {
	case data[p] < 64:
		if 48 <= data[p] && data[p] <= 57 {
			goto st120
		}
	case data[p] > 90:
		if 97 <= data[p] && data[p] <= 122 {
			goto st120
		}
	default:
		goto st120
	}
{
	goto st0

}
ctr139:
// line 27 "smtp.ragel"
	{

	pb = p
	parser.recording = true
}
	goto st125
st125:
	if p++; p == pe {
		goto _test_eof125
	}
st_case_125:
// line 3219 "smtp.go"
	switch data[p] {
	case 34:
goto st128
	case 43:
goto st130
	case 46:
goto st130
	case 47:
goto st126
	case 58:
goto st0
	case 62:
goto ctr152
	case 63:
goto st126
	case 91:
goto st126
	case 92:
goto st133
	}
	switch {
	case data[p] < 48:
		switch {
		case data[p] < 35:
			if data[p] <= 33 {
				goto st126
			}
		case data[p] > 42:
			if 44 <= data[p] && data[p] <= 45 {
				goto st126
			}
		default:
			goto st126
		}
	case data[p] > 57:
		switch {
		case data[p] < 64:
			if 59 <= data[p] && data[p] <= 61 {
				goto st126
			}
		case data[p] > 90:
			switch {
			case data[p] > 96:
				if data[p] <= 122 {
					goto st130
				}
			case data[p] >= 93:
				goto st126
			}
		default:
			goto st130
		}
	default:
		goto st130
	}
{
	goto st126

}
st126:
	if p++; p == pe {
		goto _test_eof126
	}
st_case_126:
	if data[p] == 58 {
		goto st127
	}
	if data[p] <= 57 {
		goto st126
	}
{
	goto st126

}
st127:
	if p++; p == pe {
		goto _test_eof127
	}
st_case_127:
	switch data[p] {
	case 34:
goto ctr137
	case 43:
goto ctr138
	case 46:
goto ctr138
	case 92:
goto ctr140
	}
	switch {
	case data[p] < 64:
		if 48 <= data[p] && data[p] <= 57 {
			goto ctr138
		}
	case data[p] > 90:
		if 97 <= data[p] && data[p] <= 122 {
			goto ctr138
		}
	default:
		goto ctr138
	}
{
	goto st0

}
st128:
	if p++; p == pe {
		goto _test_eof128
	}
st_case_128:
	switch data[p] {
	case 43:
goto st129
	case 46:
goto st129
	case 47:
goto st126
	case 58:
goto st127
	case 91:
goto st126
	case 92:
goto st134
	}
	switch {
	case data[p] < 59:
		switch {
		case data[p] < 44:
			if data[p] <= 42 {
				goto st126
			}
		case data[p] > 45:
			if 48 <= data[p] && data[p] <= 57 {
				goto st129
			}
		default:
			goto st126
		}
	case data[p] > 63:
		switch {
		case data[p] < 93:
			if data[p] <= 90 {
				goto st129
			}
		case data[p] > 96:
			if data[p] <= 122 {
				goto st129
			}
		default:
			goto st126
		}
	default:
		goto st126
	}
{
	goto st126

}
st129:
	if p++; p == pe {
		goto _test_eof129
	}
st_case_129:
	switch data[p] {
	case 34:
goto st130
	case 43:
goto st129
	case 46:
goto st129
	case 47:
goto st126
	case 58:
goto st127
	case 91:
goto st126
	case 92:
goto st134
	}
	switch {
	case data[p] < 48:
		switch {
		case data[p] < 35:
			if data[p] <= 33 {
				goto st126
			}
		case data[p] > 42:
			if 44 <= data[p] && data[p] <= 45 {
				goto st126
			}
		default:
			goto st126
		}
	case data[p] > 57:
		switch {
		case data[p] < 64:
			if 59 <= data[p] {
				goto st126
			}
		case data[p] > 90:
			switch {
			case data[p] > 96:
				if data[p] <= 122 {
					goto st129
				}
			case data[p] >= 93:
				goto st126
			}
		default:
			goto st129
		}
	default:
		goto st129
	}
{
	goto st126

}
st130:
	if p++; p == pe {
		goto _test_eof130
	}
st_case_130:
	switch data[p] {
	case 34:
goto st128
	case 43:
goto st130
	case 46:
goto st130
	case 47:
goto st126
	case 58:
goto st127
	case 62:
goto ctr152
	case 63:
goto st126
	case 91:
goto st126
	case 92:
goto st133
	}
	switch {
	case data[p] < 48:
		switch {
		case data[p] < 35:
			if data[p] <= 33 {
				goto st126
			}
		case data[p] > 42:
			if 44 <= data[p] && data[p] <= 45 {
				goto st126
			}
		default:
			goto st126
		}
	case data[p] > 57:
		switch {
		case data[p] < 64:
			if 59 <= data[p] && data[p] <= 61 {
				goto st126
			}
		case data[p] > 90:
			switch {
			case data[p] > 96:
				if data[p] <= 122 {
					goto st130
				}
			case data[p] >= 93:
				goto st126
			}
		default:
			goto st130
		}
	default:
		goto st130
	}
{
	goto st126

}
ctr152:
// line 20 "smtp.ragel"
	{

	parser.buffer = append(parser.buffer, data[pb:p])
	parser.current.Data = bytes.Join(parser.buffer, nil)
	parser.buffer = nil
	parser.recording = false
}
	goto st131
st131:
	if p++; p == pe {
		goto _test_eof131
	}
st_case_131:
// line 3517 "smtp.go"
	switch data[p] {
	case 10:
goto ctr157
	case 13:
goto ctr158
	case 58:
goto st127
	}
	switch {
	case data[p] < 11:
		if data[p] <= 9 {
			goto st126
		}
	case data[p] > 12:
		if 14 <= data[p] && data[p] <= 57 {
			goto st126
		}
	default:
		goto st126
	}
{
	goto st126

}
ctr159:
// line 63 "smtp.ragel"
	{
{ p++; cs = 146; goto _out }
}
	goto st146
ctr157:
// line 55 "smtp.ragel"
	{
parser.current.Verb = VerbRCPT}
// line 63 "smtp.ragel"
	{
{ p++; cs = 146; goto _out }
}
	goto st146
st146:
	if p++; p == pe {
		goto _test_eof146
	}
st_case_146:
// line 3562 "smtp.go"
	if data[p] == 58 {
		goto st127
	}
	if data[p] <= 57 {
		goto st126
	}
{
	goto st126

}
ctr158:
// line 55 "smtp.ragel"
	{
parser.current.Verb = VerbRCPT}
	goto st132
st132:
	if p++; p == pe {
		goto _test_eof132
	}
st_case_132:
// line 3583 "smtp.go"
	switch data[p] {
	case 10:
goto ctr159
	case 58:
goto st127
	}
	switch {
	case data[p] > 9:
		if 11 <= data[p] && data[p] <= 57 {
			goto st126
		}
	default:
		goto st126
	}
{
	goto st126

}
st133:
	if p++; p == pe {
		goto _test_eof133
	}
st_case_133:
	switch data[p] {
	case 32:
goto st130
	case 33:
goto st126
	case 34:
goto st130
	case 43:
goto st130
	case 46:
goto st130
	case 47:
goto st126
	case 58:
goto st127
	}
	switch {
	case data[p] < 48:
		switch {
		case data[p] < 35:
			if data[p] <= 31 {
				goto st126
			}
		case data[p] > 42:
			if 44 <= data[p] && data[p] <= 45 {
				goto st126
			}
		default:
			goto st126
		}
	case data[p] > 57:
		switch {
		case data[p] < 64:
			if 59 <= data[p] {
				goto st126
			}
		case data[p] > 90:
			switch {
			case data[p] > 96:
				if data[p] <= 122 {
					goto st130
				}
			default:
				goto st126
			}
		default:
			goto st130
		}
	default:
		goto st130
	}
{
	goto st126

}
st134:
	if p++; p == pe {
		goto _test_eof134
	}
st_case_134:
	switch data[p] {
	case 32:
goto st129
	case 33:
goto st126
	case 34:
goto st129
	case 43:
goto st129
	case 46:
goto st129
	case 47:
goto st126
	case 58:
goto st127
	}
	switch {
	case data[p] < 48:
		switch {
		case data[p] < 35:
			if data[p] <= 31 {
				goto st126
			}
		case data[p] > 42:
			if 44 <= data[p] && data[p] <= 45 {
				goto st126
			}
		default:
			goto st126
		}
	case data[p] > 57:
		switch {
		case data[p] < 64:
			if 59 <= data[p] {
				goto st126
			}
		case data[p] > 90:
			switch {
			case data[p] > 96:
				if data[p] <= 122 {
					goto st129
				}
			default:
				goto st126
			}
		default:
			goto st129
		}
	default:
		goto st129
	}
{
	goto st126

}
st135:
	if p++; p == pe {
		goto _test_eof135
	}
st_case_135:
	if data[p] == 69 {
		goto st136
	}
{
	goto st0

}
st136:
	if p++; p == pe {
		goto _test_eof136
	}
st_case_136:
	if data[p] == 84 {
		goto st137
	}
{
	goto st0

}
st137:
	if p++; p == pe {
		goto _test_eof137
	}
st_case_137:
	switch data[p] {
	case 10:
goto ctr162
	case 13:
goto ctr163
	}
{
	goto st0

}
st138:
	if p++; p == pe {
		goto _test_eof138
	}
st_case_138:
	if data[p] == 82 {
		goto st139
	}
{
	goto st0

}
st139:
	if p++; p == pe {
		goto _test_eof139
	}
st_case_139:
	if data[p] == 70 {
		goto st140
	}
{
	goto st0

}
st140:
	if p++; p == pe {
		goto _test_eof140
	}
st_case_140:
	if data[p] == 89 {
		goto st141
	}
{
	goto st0

}
st141:
	if p++; p == pe {
		goto _test_eof141
	}
st_case_141:
	if data[p] == 32 {
		goto st142
	}
{
	goto st0

}
st142:
	if p++; p == pe {
		goto _test_eof142
	}
st_case_142:
	switch data[p] {
	case 10:
goto st0
	case 13:
goto st0
	}
	switch {
	case data[p] > 9:
		if 11 <= data[p] && data[p] <= 12 {
			goto st143
		}
	default:
		goto st143
	}
{
	goto st143

}
st143:
	if p++; p == pe {
		goto _test_eof143
	}
st_case_143:
	switch data[p] {
	case 10:
goto ctr169
	case 13:
goto ctr170
	}
	switch {
	case data[p] > 9:
		if 11 <= data[p] && data[p] <= 12 {
			goto st143
		}
	default:
		goto st143
	}
{
	goto st143

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
	_test_eof144: cs = 144
	goto _test_eof
	_test_eof6: cs = 6
	goto _test_eof
	_test_eof7: cs = 7
	goto _test_eof
	_test_eof8: cs = 8
	goto _test_eof
	_test_eof9: cs = 9
	goto _test_eof
	_test_eof10: cs = 10
	goto _test_eof
	_test_eof11: cs = 11
	goto _test_eof
	_test_eof12: cs = 12
	goto _test_eof
	_test_eof13: cs = 13
	goto _test_eof
	_test_eof14: cs = 14
	goto _test_eof
	_test_eof15: cs = 15
	goto _test_eof
	_test_eof16: cs = 16
	goto _test_eof
	_test_eof17: cs = 17
	goto _test_eof
	_test_eof18: cs = 18
	goto _test_eof
	_test_eof19: cs = 19
	goto _test_eof
	_test_eof20: cs = 20
	goto _test_eof
	_test_eof21: cs = 21
	goto _test_eof
	_test_eof22: cs = 22
	goto _test_eof
	_test_eof23: cs = 23
	goto _test_eof
	_test_eof24: cs = 24
	goto _test_eof
	_test_eof25: cs = 25
	goto _test_eof
	_test_eof26: cs = 26
	goto _test_eof
	_test_eof27: cs = 27
	goto _test_eof
	_test_eof28: cs = 28
	goto _test_eof
	_test_eof29: cs = 29
	goto _test_eof
	_test_eof30: cs = 30
	goto _test_eof
	_test_eof31: cs = 31
	goto _test_eof
	_test_eof32: cs = 32
	goto _test_eof
	_test_eof33: cs = 33
	goto _test_eof
	_test_eof34: cs = 34
	goto _test_eof
	_test_eof35: cs = 35
	goto _test_eof
	_test_eof36: cs = 36
	goto _test_eof
	_test_eof37: cs = 37
	goto _test_eof
	_test_eof38: cs = 38
	goto _test_eof
	_test_eof39: cs = 39
	goto _test_eof
	_test_eof40: cs = 40
	goto _test_eof
	_test_eof41: cs = 41
	goto _test_eof
	_test_eof42: cs = 42
	goto _test_eof
	_test_eof43: cs = 43
	goto _test_eof
	_test_eof44: cs = 44
	goto _test_eof
	_test_eof45: cs = 45
	goto _test_eof
	_test_eof46: cs = 46
	goto _test_eof
	_test_eof47: cs = 47
	goto _test_eof
	_test_eof48: cs = 48
	goto _test_eof
	_test_eof49: cs = 49
	goto _test_eof
	_test_eof50: cs = 50
	goto _test_eof
	_test_eof51: cs = 51
	goto _test_eof
	_test_eof52: cs = 52
	goto _test_eof
	_test_eof53: cs = 53
	goto _test_eof
	_test_eof54: cs = 54
	goto _test_eof
	_test_eof55: cs = 55
	goto _test_eof
	_test_eof56: cs = 56
	goto _test_eof
	_test_eof57: cs = 57
	goto _test_eof
	_test_eof58: cs = 58
	goto _test_eof
	_test_eof59: cs = 59
	goto _test_eof
	_test_eof60: cs = 60
	goto _test_eof
	_test_eof61: cs = 61
	goto _test_eof
	_test_eof62: cs = 62
	goto _test_eof
	_test_eof63: cs = 63
	goto _test_eof
	_test_eof64: cs = 64
	goto _test_eof
	_test_eof65: cs = 65
	goto _test_eof
	_test_eof66: cs = 66
	goto _test_eof
	_test_eof67: cs = 67
	goto _test_eof
	_test_eof68: cs = 68
	goto _test_eof
	_test_eof69: cs = 69
	goto _test_eof
	_test_eof70: cs = 70
	goto _test_eof
	_test_eof71: cs = 71
	goto _test_eof
	_test_eof72: cs = 72
	goto _test_eof
	_test_eof73: cs = 73
	goto _test_eof
	_test_eof74: cs = 74
	goto _test_eof
	_test_eof75: cs = 75
	goto _test_eof
	_test_eof76: cs = 76
	goto _test_eof
	_test_eof77: cs = 77
	goto _test_eof
	_test_eof78: cs = 78
	goto _test_eof
	_test_eof79: cs = 79
	goto _test_eof
	_test_eof80: cs = 80
	goto _test_eof
	_test_eof81: cs = 81
	goto _test_eof
	_test_eof82: cs = 82
	goto _test_eof
	_test_eof83: cs = 83
	goto _test_eof
	_test_eof84: cs = 84
	goto _test_eof
	_test_eof85: cs = 85
	goto _test_eof
	_test_eof86: cs = 86
	goto _test_eof
	_test_eof87: cs = 87
	goto _test_eof
	_test_eof88: cs = 88
	goto _test_eof
	_test_eof89: cs = 89
	goto _test_eof
	_test_eof90: cs = 90
	goto _test_eof
	_test_eof91: cs = 91
	goto _test_eof
	_test_eof92: cs = 92
	goto _test_eof
	_test_eof93: cs = 93
	goto _test_eof
	_test_eof94: cs = 94
	goto _test_eof
	_test_eof95: cs = 95
	goto _test_eof
	_test_eof96: cs = 96
	goto _test_eof
	_test_eof97: cs = 97
	goto _test_eof
	_test_eof98: cs = 98
	goto _test_eof
	_test_eof99: cs = 99
	goto _test_eof
	_test_eof100: cs = 100
	goto _test_eof
	_test_eof101: cs = 101
	goto _test_eof
	_test_eof102: cs = 102
	goto _test_eof
	_test_eof145: cs = 145
	goto _test_eof
	_test_eof103: cs = 103
	goto _test_eof
	_test_eof104: cs = 104
	goto _test_eof
	_test_eof105: cs = 105
	goto _test_eof
	_test_eof106: cs = 106
	goto _test_eof
	_test_eof107: cs = 107
	goto _test_eof
	_test_eof108: cs = 108
	goto _test_eof
	_test_eof109: cs = 109
	goto _test_eof
	_test_eof110: cs = 110
	goto _test_eof
	_test_eof111: cs = 111
	goto _test_eof
	_test_eof112: cs = 112
	goto _test_eof
	_test_eof113: cs = 113
	goto _test_eof
	_test_eof114: cs = 114
	goto _test_eof
	_test_eof115: cs = 115
	goto _test_eof
	_test_eof116: cs = 116
	goto _test_eof
	_test_eof117: cs = 117
	goto _test_eof
	_test_eof118: cs = 118
	goto _test_eof
	_test_eof119: cs = 119
	goto _test_eof
	_test_eof120: cs = 120
	goto _test_eof
	_test_eof121: cs = 121
	goto _test_eof
	_test_eof122: cs = 122
	goto _test_eof
	_test_eof123: cs = 123
	goto _test_eof
	_test_eof124: cs = 124
	goto _test_eof
	_test_eof125: cs = 125
	goto _test_eof
	_test_eof126: cs = 126
	goto _test_eof
	_test_eof127: cs = 127
	goto _test_eof
	_test_eof128: cs = 128
	goto _test_eof
	_test_eof129: cs = 129
	goto _test_eof
	_test_eof130: cs = 130
	goto _test_eof
	_test_eof131: cs = 131
	goto _test_eof
	_test_eof146: cs = 146
	goto _test_eof
	_test_eof132: cs = 132
	goto _test_eof
	_test_eof133: cs = 133
	goto _test_eof
	_test_eof134: cs = 134
	goto _test_eof
	_test_eof135: cs = 135
	goto _test_eof
	_test_eof136: cs = 136
	goto _test_eof
	_test_eof137: cs = 137
	goto _test_eof
	_test_eof138: cs = 138
	goto _test_eof
	_test_eof139: cs = 139
	goto _test_eof
	_test_eof140: cs = 140
	goto _test_eof
	_test_eof141: cs = 141
	goto _test_eof
	_test_eof142: cs = 142
	goto _test_eof
	_test_eof143: cs = 143
	goto _test_eof

	_test_eof: {}
	_out: {}
	}

// line 83 "smtp.ragel"

	if cs == smtp_error {
		return data[p:], fmt.Errorf("Invalid character in pos %d: `%c`.", p, data[p])
	}

	// Not yet a full parse, remeber everything from pb to p if we are
	// recording.
	if cs < smtp_first_final {
		if parser.recording {
			parser.buffer = append(parser.buffer, data[pb:p])
		}
		parser.cs = cs
		return data[p:], Dangling
	}

	// Full parse:
	parser.cs = smtp_en_main
	return data[p:], nil
}
