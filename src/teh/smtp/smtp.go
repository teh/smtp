
// line 1 "smtp.ragel"
package smtp

import (
	"fmt"
	"errors"
	"bytes"
)

// Good introduction on smtp: http://cr.yp.to/smtp/request.html

// line 73 "smtp.ragel"



// line 18 "smtp.go"
const smtp_start = 1;
const smtp_first_final = 200;
const smtp_error = 0;

const smtp_en_main = 1


// line 76 "smtp.ragel"

var Dangling = errors.New("DANGLING")

func NewParser() *Parser {
	return &Parser{cs: smtp_en_main}
}

// Feed new bytes while keepting track of state in parser.current. If
// error is `Dangling` then more data is needed for a successfull
// parse. If error is nil then parser.current contains the new action
// item. Bytes from `data` are copied, enabling the called to recycle
// data slice.
//
// Gotchas: * current is not transactional and can contain
// intermediate values. * Parsing is not bounded: the internal
// parser.buffer grows as much as the input, restrictions need to be
// imposed from the outside (using e.g. a LimitReader).
func (parser *Parser) Feed(data []byte) (remaining []byte, err error) {
	var p, pb int
	cs := parser.cs
	pe := len(data)

	
// line 50 "smtp.go"
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
	case 200:
		goto st_case_200
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
	case 201:
		goto st_case_201
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
	case 144:
		goto st_case_144
	case 145:
		goto st_case_145
	case 146:
		goto st_case_146
	case 147:
		goto st_case_147
	case 148:
		goto st_case_148
	case 149:
		goto st_case_149
	case 150:
		goto st_case_150
	case 151:
		goto st_case_151
	case 152:
		goto st_case_152
	case 153:
		goto st_case_153
	case 154:
		goto st_case_154
	case 155:
		goto st_case_155
	case 156:
		goto st_case_156
	case 157:
		goto st_case_157
	case 158:
		goto st_case_158
	case 159:
		goto st_case_159
	case 160:
		goto st_case_160
	case 161:
		goto st_case_161
	case 162:
		goto st_case_162
	case 163:
		goto st_case_163
	case 164:
		goto st_case_164
	case 165:
		goto st_case_165
	case 166:
		goto st_case_166
	case 167:
		goto st_case_167
	case 168:
		goto st_case_168
	case 169:
		goto st_case_169
	case 170:
		goto st_case_170
	case 171:
		goto st_case_171
	case 172:
		goto st_case_172
	case 173:
		goto st_case_173
	case 174:
		goto st_case_174
	case 175:
		goto st_case_175
	case 176:
		goto st_case_176
	case 177:
		goto st_case_177
	case 178:
		goto st_case_178
	case 179:
		goto st_case_179
	case 202:
		goto st_case_202
	case 180:
		goto st_case_180
	case 181:
		goto st_case_181
	case 182:
		goto st_case_182
	case 183:
		goto st_case_183
	case 184:
		goto st_case_184
	case 185:
		goto st_case_185
	case 186:
		goto st_case_186
	case 187:
		goto st_case_187
	case 188:
		goto st_case_188
	case 189:
		goto st_case_189
	case 190:
		goto st_case_190
	case 191:
		goto st_case_191
	case 192:
		goto st_case_192
	case 193:
		goto st_case_193
	case 194:
		goto st_case_194
	case 195:
		goto st_case_195
	case 196:
		goto st_case_196
	case 197:
		goto st_case_197
	case 198:
		goto st_case_198
	case 199:
		goto st_case_199
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
goto st154
	case 82:
goto st158
	case 83:
goto st186
	case 86:
goto st194
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
goto ctr12
	case 13:
goto ctr13
	}
{
	goto st0

}
ctr12:
// line 64 "smtp.ragel"
	{
parser.current.Verb = VerbDATA}
// line 72 "smtp.ragel"
	{
{ p++; cs = 200; goto _out }
}
	goto st200
ctr14:
// line 72 "smtp.ragel"
	{
{ p++; cs = 200; goto _out }
}
	goto st200
ctr21:
// line 20 "smtp.ragel"
	{

	parser.buffer.Write(data[pb:p])
	parser.current.Data = parser.buffer.Bytes()
	parser.buffer = nil
}
// line 60 "smtp.ragel"
	{
parser.current.Verb = VerbEHLO}
// line 72 "smtp.ragel"
	{
{ p++; cs = 200; goto _out }
}
	goto st200
ctr59:
// line 20 "smtp.ragel"
	{

	parser.buffer.Write(data[pb:p])
	parser.current.Data = parser.buffer.Bytes()
	parser.buffer = nil
}
// line 59 "smtp.ragel"
	{
parser.current.Verb = VerbHELO}
// line 72 "smtp.ragel"
	{
{ p++; cs = 200; goto _out }
}
	goto st200
ctr111:
// line 62 "smtp.ragel"
	{
parser.current.Verb = VerbMAIL}
// line 72 "smtp.ragel"
	{
{ p++; cs = 200; goto _out }
}
	goto st200
ctr125:
// line 51 "smtp.ragel"
	{
parser.current.BodyType = BodyType7BIT}
// line 62 "smtp.ragel"
	{
parser.current.Verb = VerbMAIL}
// line 72 "smtp.ragel"
	{
{ p++; cs = 200; goto _out }
}
	goto st200
ctr135:
// line 49 "smtp.ragel"
	{
parser.current.BodyType = BodyType8BITMIME}
// line 62 "smtp.ragel"
	{
parser.current.Verb = VerbMAIL}
// line 72 "smtp.ragel"
	{
{ p++; cs = 200; goto _out }
}
	goto st200
ctr187:
// line 65 "smtp.ragel"
	{
parser.current.Verb = VerbQUIT}
// line 72 "smtp.ragel"
	{
{ p++; cs = 200; goto _out }
}
	goto st200
ctr208:
// line 63 "smtp.ragel"
	{
parser.current.Verb = VerbRCPT}
// line 72 "smtp.ragel"
	{
{ p++; cs = 200; goto _out }
}
	goto st200
ctr223:
// line 61 "smtp.ragel"
	{
parser.current.Verb = VerbRSET}
// line 72 "smtp.ragel"
	{
{ p++; cs = 200; goto _out }
}
	goto st200
ctr232:
// line 66 "smtp.ragel"
	{
parser.current.Verb = VerbSTARTTLS}
// line 72 "smtp.ragel"
	{
{ p++; cs = 200; goto _out }
}
	goto st200
ctr239:
// line 67 "smtp.ragel"
	{
parser.current.Verb = VerbVRFY}
// line 72 "smtp.ragel"
	{
{ p++; cs = 200; goto _out }
}
	goto st200
st200:
	if p++; p == pe {
		goto _test_eof200
	}
st_case_200:
// line 672 "smtp.go"
{
	goto st0

}
ctr13:
// line 64 "smtp.ragel"
	{
parser.current.Verb = VerbDATA}
	goto st6
ctr22:
// line 20 "smtp.ragel"
	{

	parser.buffer.Write(data[pb:p])
	parser.current.Data = parser.buffer.Bytes()
	parser.buffer = nil
}
// line 60 "smtp.ragel"
	{
parser.current.Verb = VerbEHLO}
	goto st6
ctr60:
// line 20 "smtp.ragel"
	{

	parser.buffer.Write(data[pb:p])
	parser.current.Data = parser.buffer.Bytes()
	parser.buffer = nil
}
// line 59 "smtp.ragel"
	{
parser.current.Verb = VerbHELO}
	goto st6
ctr112:
// line 62 "smtp.ragel"
	{
parser.current.Verb = VerbMAIL}
	goto st6
ctr126:
// line 51 "smtp.ragel"
	{
parser.current.BodyType = BodyType7BIT}
// line 62 "smtp.ragel"
	{
parser.current.Verb = VerbMAIL}
	goto st6
ctr136:
// line 49 "smtp.ragel"
	{
parser.current.BodyType = BodyType8BITMIME}
// line 62 "smtp.ragel"
	{
parser.current.Verb = VerbMAIL}
	goto st6
ctr188:
// line 65 "smtp.ragel"
	{
parser.current.Verb = VerbQUIT}
	goto st6
ctr209:
// line 63 "smtp.ragel"
	{
parser.current.Verb = VerbRCPT}
	goto st6
ctr224:
// line 61 "smtp.ragel"
	{
parser.current.Verb = VerbRSET}
	goto st6
ctr233:
// line 66 "smtp.ragel"
	{
parser.current.Verb = VerbSTARTTLS}
	goto st6
ctr240:
// line 67 "smtp.ragel"
	{
parser.current.Verb = VerbVRFY}
	goto st6
st6:
	if p++; p == pe {
		goto _test_eof6
	}
st_case_6:
// line 757 "smtp.go"
	if data[p] == 10 {
		goto ctr14
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
		goto ctr20
	}
	switch {
	case data[p] < 65:
		if 48 <= data[p] && data[p] <= 57 {
			goto ctr19
		}
	case data[p] > 90:
		if 97 <= data[p] && data[p] <= 122 {
			goto ctr19
		}
	default:
		goto ctr19
	}
{
	goto st0

}
ctr19:
// line 26 "smtp.ragel"
	{

	pb = p
	parser.buffer = bytes.NewBuffer(nil)
}
	goto st12
st12:
	if p++; p == pe {
		goto _test_eof12
	}
st_case_12:
// line 850 "smtp.go"
	switch data[p] {
	case 10:
goto ctr21
	case 13:
goto ctr22
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
ctr20:
// line 26 "smtp.ragel"
	{

	pb = p
	parser.buffer = bytes.NewBuffer(nil)
}
	goto st14
st14:
	if p++; p == pe {
		goto _test_eof14
	}
st_case_14:
// line 909 "smtp.go"
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
goto ctr21
	case 13:
goto ctr22
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
		goto ctr58
	}
	switch {
	case data[p] < 65:
		if 48 <= data[p] && data[p] <= 57 {
			goto ctr57
		}
	case data[p] > 90:
		if 97 <= data[p] && data[p] <= 122 {
			goto ctr57
		}
	default:
		goto ctr57
	}
{
	goto st0

}
ctr57:
// line 26 "smtp.ragel"
	{

	pb = p
	parser.buffer = bytes.NewBuffer(nil)
}
	goto st48
st48:
	if p++; p == pe {
		goto _test_eof48
	}
st_case_48:
// line 1513 "smtp.go"
	switch data[p] {
	case 10:
goto ctr59
	case 13:
goto ctr60
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
ctr58:
// line 26 "smtp.ragel"
	{

	pb = p
	parser.buffer = bytes.NewBuffer(nil)
}
	goto st50
st50:
	if p++; p == pe {
		goto _test_eof50
	}
st_case_50:
// line 1572 "smtp.go"
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
goto ctr59
	case 13:
goto ctr60
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
goto ctr101
	case 43:
goto ctr102
	case 64:
goto ctr103
	case 92:
goto ctr104
	case 95:
goto ctr102
	}
	switch {
	case data[p] < 48:
		if 45 <= data[p] && data[p] <= 46 {
			goto ctr102
		}
	case data[p] > 57:
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto ctr102
			}
		case data[p] >= 65:
			goto ctr102
		}
	default:
		goto ctr102
	}
{
	goto st0

}
ctr101:
// line 26 "smtp.ragel"
	{

	pb = p
	parser.buffer = bytes.NewBuffer(nil)
}
	goto st90
st90:
	if p++; p == pe {
		goto _test_eof90
	}
st_case_90:
// line 2274 "smtp.go"
	switch data[p] {
	case 43:
goto st91
	case 92:
goto st119
	case 95:
goto st91
	}
	switch {
	case data[p] < 48:
		if 45 <= data[p] && data[p] <= 46 {
			goto st91
		}
	case data[p] > 57:
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st91
			}
		case data[p] >= 64:
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
	case 92:
goto st119
	case 95:
goto st91
	}
	switch {
	case data[p] < 48:
		if 45 <= data[p] && data[p] <= 46 {
			goto st91
		}
	case data[p] > 57:
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st91
			}
		case data[p] >= 64:
			goto st91
		}
	default:
		goto st91
	}
{
	goto st0

}
ctr102:
// line 26 "smtp.ragel"
	{

	pb = p
	parser.buffer = bytes.NewBuffer(nil)
}
	goto st92
st92:
	if p++; p == pe {
		goto _test_eof92
	}
st_case_92:
// line 2353 "smtp.go"
	switch data[p] {
	case 34:
goto st90
	case 43:
goto st92
	case 62:
goto ctr109
	case 92:
goto st118
	case 95:
goto st92
	}
	switch {
	case data[p] < 48:
		if 45 <= data[p] && data[p] <= 46 {
			goto st92
		}
	case data[p] > 57:
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st92
			}
		case data[p] >= 64:
			goto st92
		}
	default:
		goto st92
	}
{
	goto st0

}
ctr109:
// line 20 "smtp.ragel"
	{

	parser.buffer.Write(data[pb:p])
	parser.current.Data = parser.buffer.Bytes()
	parser.buffer = nil
}
	goto st93
st93:
	if p++; p == pe {
		goto _test_eof93
	}
st_case_93:
// line 2401 "smtp.go"
	switch data[p] {
	case 10:
goto ctr111
	case 13:
goto ctr112
	case 32:
goto st94
	}
{
	goto st0

}
ctr127:
// line 51 "smtp.ragel"
	{
parser.current.BodyType = BodyType7BIT}
	goto st94
ctr137:
// line 49 "smtp.ragel"
	{
parser.current.BodyType = BodyType8BITMIME}
	goto st94
st94:
	if p++; p == pe {
		goto _test_eof94
	}
st_case_94:
// line 2429 "smtp.go"
	switch data[p] {
	case 10:
goto ctr111
	case 13:
goto ctr112
	case 32:
goto st94
	case 66:
goto st95
	case 83:
goto st112
	}
{
	goto st0

}
st95:
	if p++; p == pe {
		goto _test_eof95
	}
st_case_95:
	if data[p] == 79 {
		goto st96
	}
{
	goto st0

}
st96:
	if p++; p == pe {
		goto _test_eof96
	}
st_case_96:
	if data[p] == 68 {
		goto st97
	}
{
	goto st0

}
st97:
	if p++; p == pe {
		goto _test_eof97
	}
st_case_97:
	if data[p] == 89 {
		goto st98
	}
{
	goto st0

}
st98:
	if p++; p == pe {
		goto _test_eof98
	}
st_case_98:
	if data[p] == 61 {
		goto st99
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
	case 55:
goto st100
	case 56:
goto st104
	}
{
	goto st0

}
st100:
	if p++; p == pe {
		goto _test_eof100
	}
st_case_100:
	if data[p] == 66 {
		goto st101
	}
{
	goto st0

}
st101:
	if p++; p == pe {
		goto _test_eof101
	}
st_case_101:
	if data[p] == 73 {
		goto st102
	}
{
	goto st0

}
st102:
	if p++; p == pe {
		goto _test_eof102
	}
st_case_102:
	if data[p] == 84 {
		goto st103
	}
{
	goto st0

}
st103:
	if p++; p == pe {
		goto _test_eof103
	}
st_case_103:
	switch data[p] {
	case 10:
goto ctr125
	case 13:
goto ctr126
	case 32:
goto ctr127
	}
{
	goto st0

}
st104:
	if p++; p == pe {
		goto _test_eof104
	}
st_case_104:
	if data[p] == 66 {
		goto st105
	}
{
	goto st0

}
st105:
	if p++; p == pe {
		goto _test_eof105
	}
st_case_105:
	if data[p] == 73 {
		goto st106
	}
{
	goto st0

}
st106:
	if p++; p == pe {
		goto _test_eof106
	}
st_case_106:
	if data[p] == 84 {
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
	if data[p] == 77 {
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
	if data[p] == 73 {
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
	if data[p] == 77 {
		goto st110
	}
{
	goto st0

}
st110:
	if p++; p == pe {
		goto _test_eof110
	}
st_case_110:
	if data[p] == 69 {
		goto st111
	}
{
	goto st0

}
st111:
	if p++; p == pe {
		goto _test_eof111
	}
st_case_111:
	switch data[p] {
	case 10:
goto ctr135
	case 13:
goto ctr136
	case 32:
goto ctr137
	}
{
	goto st0

}
st112:
	if p++; p == pe {
		goto _test_eof112
	}
st_case_112:
	if data[p] == 73 {
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
	if data[p] == 90 {
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
	if data[p] == 69 {
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
	if data[p] == 61 {
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
	if 48 <= data[p] && data[p] <= 57 {
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
	switch data[p] {
	case 10:
goto ctr111
	case 13:
goto ctr112
	case 32:
goto st94
	}
	if 48 <= data[p] && data[p] <= 57 {
		goto st117
	}
{
	goto st0

}
ctr104:
// line 26 "smtp.ragel"
	{

	pb = p
	parser.buffer = bytes.NewBuffer(nil)
}
	goto st118
st118:
	if p++; p == pe {
		goto _test_eof118
	}
st_case_118:
// line 2756 "smtp.go"
	switch data[p] {
	case 32:
goto st92
	case 34:
goto st92
	case 43:
goto st92
	case 95:
goto st92
	}
	switch {
	case data[p] < 48:
		if 45 <= data[p] && data[p] <= 46 {
			goto st92
		}
	case data[p] > 57:
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st92
			}
		case data[p] >= 64:
			goto st92
		}
	default:
		goto st92
	}
{
	goto st0

}
st119:
	if p++; p == pe {
		goto _test_eof119
	}
st_case_119:
	switch data[p] {
	case 32:
goto st91
	case 34:
goto st91
	case 43:
goto st91
	case 95:
goto st91
	}
	switch {
	case data[p] < 48:
		if 45 <= data[p] && data[p] <= 46 {
			goto st91
		}
	case data[p] > 57:
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st91
			}
		case data[p] >= 64:
			goto st91
		}
	default:
		goto st91
	}
{
	goto st0

}
ctr103:
// line 26 "smtp.ragel"
	{

	pb = p
	parser.buffer = bytes.NewBuffer(nil)
}
	goto st120
st120:
	if p++; p == pe {
		goto _test_eof120
	}
st_case_120:
// line 2837 "smtp.go"
	switch data[p] {
	case 34:
goto st123
	case 43:
goto st125
	case 44:
goto st121
	case 47:
goto st121
	case 58:
goto st0
	case 62:
goto ctr146
	case 63:
goto st121
	case 91:
goto st121
	case 92:
goto st152
	case 95:
goto st125
	case 96:
goto st121
	}
	switch {
	case data[p] < 48:
		switch {
		case data[p] < 35:
			if data[p] <= 33 {
				goto st121
			}
		case data[p] > 42:
			if 45 <= data[p] && data[p] <= 46 {
				goto st125
			}
		default:
			goto st121
		}
	case data[p] > 57:
		switch {
		case data[p] < 64:
			if 59 <= data[p] && data[p] <= 61 {
				goto st121
			}
		case data[p] > 90:
			switch {
			case data[p] > 94:
				if 97 <= data[p] && data[p] <= 122 {
					goto st125
				}
			case data[p] >= 93:
				goto st121
			}
		default:
			goto st125
		}
	default:
		goto st125
	}
{
	goto st121

}
st121:
	if p++; p == pe {
		goto _test_eof121
	}
st_case_121:
	if data[p] == 58 {
		goto st122
	}
	if data[p] <= 57 {
		goto st121
	}
{
	goto st121

}
st122:
	if p++; p == pe {
		goto _test_eof122
	}
st_case_122:
	switch data[p] {
	case 34:
goto ctr101
	case 43:
goto ctr102
	case 92:
goto ctr104
	case 95:
goto ctr102
	}
	switch {
	case data[p] < 48:
		if 45 <= data[p] && data[p] <= 46 {
			goto ctr102
		}
	case data[p] > 57:
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto ctr102
			}
		case data[p] >= 64:
			goto ctr102
		}
	default:
		goto ctr102
	}
{
	goto st0

}
st123:
	if p++; p == pe {
		goto _test_eof123
	}
st_case_123:
	switch data[p] {
	case 43:
goto st124
	case 44:
goto st121
	case 47:
goto st121
	case 58:
goto st122
	case 91:
goto st121
	case 92:
goto st153
	case 95:
goto st124
	case 96:
goto st121
	}
	switch {
	case data[p] < 59:
		switch {
		case data[p] < 45:
			if data[p] <= 42 {
				goto st121
			}
		case data[p] > 46:
			if 48 <= data[p] && data[p] <= 57 {
				goto st124
			}
		default:
			goto st124
		}
	case data[p] > 63:
		switch {
		case data[p] < 93:
			if data[p] <= 90 {
				goto st124
			}
		case data[p] > 94:
			if 97 <= data[p] && data[p] <= 122 {
				goto st124
			}
		default:
			goto st121
		}
	default:
		goto st121
	}
{
	goto st121

}
st124:
	if p++; p == pe {
		goto _test_eof124
	}
st_case_124:
	switch data[p] {
	case 34:
goto st125
	case 43:
goto st124
	case 44:
goto st121
	case 47:
goto st121
	case 58:
goto st122
	case 91:
goto st121
	case 92:
goto st153
	case 95:
goto st124
	case 96:
goto st121
	}
	switch {
	case data[p] < 48:
		switch {
		case data[p] < 35:
			if data[p] <= 33 {
				goto st121
			}
		case data[p] > 42:
			if 45 <= data[p] && data[p] <= 46 {
				goto st124
			}
		default:
			goto st121
		}
	case data[p] > 57:
		switch {
		case data[p] < 64:
			if 59 <= data[p] {
				goto st121
			}
		case data[p] > 90:
			switch {
			case data[p] > 94:
				if 97 <= data[p] && data[p] <= 122 {
					goto st124
				}
			case data[p] >= 93:
				goto st121
			}
		default:
			goto st124
		}
	default:
		goto st124
	}
{
	goto st121

}
st125:
	if p++; p == pe {
		goto _test_eof125
	}
st_case_125:
	switch data[p] {
	case 34:
goto st123
	case 43:
goto st125
	case 44:
goto st121
	case 47:
goto st121
	case 58:
goto st122
	case 62:
goto ctr146
	case 63:
goto st121
	case 91:
goto st121
	case 92:
goto st152
	case 95:
goto st125
	case 96:
goto st121
	}
	switch {
	case data[p] < 48:
		switch {
		case data[p] < 35:
			if data[p] <= 33 {
				goto st121
			}
		case data[p] > 42:
			if 45 <= data[p] && data[p] <= 46 {
				goto st125
			}
		default:
			goto st121
		}
	case data[p] > 57:
		switch {
		case data[p] < 64:
			if 59 <= data[p] && data[p] <= 61 {
				goto st121
			}
		case data[p] > 90:
			switch {
			case data[p] > 94:
				if 97 <= data[p] && data[p] <= 122 {
					goto st125
				}
			case data[p] >= 93:
				goto st121
			}
		default:
			goto st125
		}
	default:
		goto st125
	}
{
	goto st121

}
ctr146:
// line 20 "smtp.ragel"
	{

	parser.buffer.Write(data[pb:p])
	parser.current.Data = parser.buffer.Bytes()
	parser.buffer = nil
}
	goto st126
st126:
	if p++; p == pe {
		goto _test_eof126
	}
st_case_126:
// line 3155 "smtp.go"
	switch data[p] {
	case 10:
goto ctr151
	case 13:
goto ctr152
	case 32:
goto st128
	case 58:
goto st122
	}
	switch {
	case data[p] < 11:
		if data[p] <= 9 {
			goto st121
		}
	case data[p] > 12:
		switch {
		case data[p] > 31:
			if 33 <= data[p] && data[p] <= 57 {
				goto st121
			}
		case data[p] >= 14:
			goto st121
		}
	default:
		goto st121
	}
{
	goto st121

}
ctr154:
// line 72 "smtp.ragel"
	{
{ p++; cs = 201; goto _out }
}
	goto st201
ctr151:
// line 62 "smtp.ragel"
	{
parser.current.Verb = VerbMAIL}
// line 72 "smtp.ragel"
	{
{ p++; cs = 201; goto _out }
}
	goto st201
ctr166:
// line 51 "smtp.ragel"
	{
parser.current.BodyType = BodyType7BIT}
// line 62 "smtp.ragel"
	{
parser.current.Verb = VerbMAIL}
// line 72 "smtp.ragel"
	{
{ p++; cs = 201; goto _out }
}
	goto st201
ctr176:
// line 49 "smtp.ragel"
	{
parser.current.BodyType = BodyType8BITMIME}
// line 62 "smtp.ragel"
	{
parser.current.Verb = VerbMAIL}
// line 72 "smtp.ragel"
	{
{ p++; cs = 201; goto _out }
}
	goto st201
st201:
	if p++; p == pe {
		goto _test_eof201
	}
st_case_201:
// line 3231 "smtp.go"
	if data[p] == 58 {
		goto st122
	}
	if data[p] <= 57 {
		goto st121
	}
{
	goto st121

}
ctr152:
// line 62 "smtp.ragel"
	{
parser.current.Verb = VerbMAIL}
	goto st127
ctr167:
// line 51 "smtp.ragel"
	{
parser.current.BodyType = BodyType7BIT}
// line 62 "smtp.ragel"
	{
parser.current.Verb = VerbMAIL}
	goto st127
ctr177:
// line 49 "smtp.ragel"
	{
parser.current.BodyType = BodyType8BITMIME}
// line 62 "smtp.ragel"
	{
parser.current.Verb = VerbMAIL}
	goto st127
st127:
	if p++; p == pe {
		goto _test_eof127
	}
st_case_127:
// line 3268 "smtp.go"
	switch data[p] {
	case 10:
goto ctr154
	case 58:
goto st122
	}
	switch {
	case data[p] > 9:
		if 11 <= data[p] && data[p] <= 57 {
			goto st121
		}
	default:
		goto st121
	}
{
	goto st121

}
ctr168:
// line 51 "smtp.ragel"
	{
parser.current.BodyType = BodyType7BIT}
	goto st128
ctr178:
// line 49 "smtp.ragel"
	{
parser.current.BodyType = BodyType8BITMIME}
	goto st128
st128:
	if p++; p == pe {
		goto _test_eof128
	}
st_case_128:
// line 3302 "smtp.go"
	switch data[p] {
	case 10:
goto ctr151
	case 13:
goto ctr152
	case 32:
goto st128
	case 58:
goto st122
	case 66:
goto st129
	case 83:
goto st146
	}
	switch {
	case data[p] < 14:
		switch {
		case data[p] > 9:
			if 11 <= data[p] && data[p] <= 12 {
				goto st121
			}
		default:
			goto st121
		}
	case data[p] > 31:
		switch {
		case data[p] < 59:
			if 33 <= data[p] && data[p] <= 57 {
				goto st121
			}
		case data[p] > 65:
			if 67 <= data[p] && data[p] <= 82 {
				goto st121
			}
		default:
			goto st121
		}
	default:
		goto st121
	}
{
	goto st121

}
st129:
	if p++; p == pe {
		goto _test_eof129
	}
st_case_129:
	switch data[p] {
	case 58:
goto st122
	case 79:
goto st130
	}
	switch {
	case data[p] > 57:
		if 59 <= data[p] && data[p] <= 78 {
			goto st121
		}
	default:
		goto st121
	}
{
	goto st121

}
st130:
	if p++; p == pe {
		goto _test_eof130
	}
st_case_130:
	switch data[p] {
	case 58:
goto st122
	case 68:
goto st131
	}
	switch {
	case data[p] > 57:
		if 59 <= data[p] && data[p] <= 67 {
			goto st121
		}
	default:
		goto st121
	}
{
	goto st121

}
st131:
	if p++; p == pe {
		goto _test_eof131
	}
st_case_131:
	switch data[p] {
	case 58:
goto st122
	case 89:
goto st132
	}
	switch {
	case data[p] > 57:
		if 59 <= data[p] && data[p] <= 88 {
			goto st121
		}
	default:
		goto st121
	}
{
	goto st121

}
st132:
	if p++; p == pe {
		goto _test_eof132
	}
st_case_132:
	switch data[p] {
	case 58:
goto st122
	case 61:
goto st133
	}
	switch {
	case data[p] > 57:
		if 59 <= data[p] && data[p] <= 60 {
			goto st121
		}
	default:
		goto st121
	}
{
	goto st121

}
st133:
	if p++; p == pe {
		goto _test_eof133
	}
st_case_133:
	switch data[p] {
	case 55:
goto st134
	case 56:
goto st138
	case 57:
goto st121
	case 58:
goto st122
	}
	if data[p] <= 54 {
		goto st121
	}
{
	goto st121

}
st134:
	if p++; p == pe {
		goto _test_eof134
	}
st_case_134:
	switch data[p] {
	case 58:
goto st122
	case 66:
goto st135
	}
	switch {
	case data[p] > 57:
		if 59 <= data[p] && data[p] <= 65 {
			goto st121
		}
	default:
		goto st121
	}
{
	goto st121

}
st135:
	if p++; p == pe {
		goto _test_eof135
	}
st_case_135:
	switch data[p] {
	case 58:
goto st122
	case 73:
goto st136
	}
	switch {
	case data[p] > 57:
		if 59 <= data[p] && data[p] <= 72 {
			goto st121
		}
	default:
		goto st121
	}
{
	goto st121

}
st136:
	if p++; p == pe {
		goto _test_eof136
	}
st_case_136:
	switch data[p] {
	case 58:
goto st122
	case 84:
goto st137
	}
	switch {
	case data[p] > 57:
		if 59 <= data[p] && data[p] <= 83 {
			goto st121
		}
	default:
		goto st121
	}
{
	goto st121

}
st137:
	if p++; p == pe {
		goto _test_eof137
	}
st_case_137:
	switch data[p] {
	case 10:
goto ctr166
	case 13:
goto ctr167
	case 32:
goto ctr168
	case 58:
goto st122
	}
	switch {
	case data[p] < 11:
		if data[p] <= 9 {
			goto st121
		}
	case data[p] > 12:
		switch {
		case data[p] > 31:
			if 33 <= data[p] && data[p] <= 57 {
				goto st121
			}
		case data[p] >= 14:
			goto st121
		}
	default:
		goto st121
	}
{
	goto st121

}
st138:
	if p++; p == pe {
		goto _test_eof138
	}
st_case_138:
	switch data[p] {
	case 58:
goto st122
	case 66:
goto st139
	}
	switch {
	case data[p] > 57:
		if 59 <= data[p] && data[p] <= 65 {
			goto st121
		}
	default:
		goto st121
	}
{
	goto st121

}
st139:
	if p++; p == pe {
		goto _test_eof139
	}
st_case_139:
	switch data[p] {
	case 58:
goto st122
	case 73:
goto st140
	}
	switch {
	case data[p] > 57:
		if 59 <= data[p] && data[p] <= 72 {
			goto st121
		}
	default:
		goto st121
	}
{
	goto st121

}
st140:
	if p++; p == pe {
		goto _test_eof140
	}
st_case_140:
	switch data[p] {
	case 58:
goto st122
	case 84:
goto st141
	}
	switch {
	case data[p] > 57:
		if 59 <= data[p] && data[p] <= 83 {
			goto st121
		}
	default:
		goto st121
	}
{
	goto st121

}
st141:
	if p++; p == pe {
		goto _test_eof141
	}
st_case_141:
	switch data[p] {
	case 58:
goto st122
	case 77:
goto st142
	}
	switch {
	case data[p] > 57:
		if 59 <= data[p] && data[p] <= 76 {
			goto st121
		}
	default:
		goto st121
	}
{
	goto st121

}
st142:
	if p++; p == pe {
		goto _test_eof142
	}
st_case_142:
	switch data[p] {
	case 58:
goto st122
	case 73:
goto st143
	}
	switch {
	case data[p] > 57:
		if 59 <= data[p] && data[p] <= 72 {
			goto st121
		}
	default:
		goto st121
	}
{
	goto st121

}
st143:
	if p++; p == pe {
		goto _test_eof143
	}
st_case_143:
	switch data[p] {
	case 58:
goto st122
	case 77:
goto st144
	}
	switch {
	case data[p] > 57:
		if 59 <= data[p] && data[p] <= 76 {
			goto st121
		}
	default:
		goto st121
	}
{
	goto st121

}
st144:
	if p++; p == pe {
		goto _test_eof144
	}
st_case_144:
	switch data[p] {
	case 58:
goto st122
	case 69:
goto st145
	}
	switch {
	case data[p] > 57:
		if 59 <= data[p] && data[p] <= 68 {
			goto st121
		}
	default:
		goto st121
	}
{
	goto st121

}
st145:
	if p++; p == pe {
		goto _test_eof145
	}
st_case_145:
	switch data[p] {
	case 10:
goto ctr176
	case 13:
goto ctr177
	case 32:
goto ctr178
	case 58:
goto st122
	}
	switch {
	case data[p] < 11:
		if data[p] <= 9 {
			goto st121
		}
	case data[p] > 12:
		switch {
		case data[p] > 31:
			if 33 <= data[p] && data[p] <= 57 {
				goto st121
			}
		case data[p] >= 14:
			goto st121
		}
	default:
		goto st121
	}
{
	goto st121

}
st146:
	if p++; p == pe {
		goto _test_eof146
	}
st_case_146:
	switch data[p] {
	case 58:
goto st122
	case 73:
goto st147
	}
	switch {
	case data[p] > 57:
		if 59 <= data[p] && data[p] <= 72 {
			goto st121
		}
	default:
		goto st121
	}
{
	goto st121

}
st147:
	if p++; p == pe {
		goto _test_eof147
	}
st_case_147:
	switch data[p] {
	case 58:
goto st122
	case 90:
goto st148
	}
	switch {
	case data[p] > 57:
		if 59 <= data[p] && data[p] <= 89 {
			goto st121
		}
	default:
		goto st121
	}
{
	goto st121

}
st148:
	if p++; p == pe {
		goto _test_eof148
	}
st_case_148:
	switch data[p] {
	case 58:
goto st122
	case 69:
goto st149
	}
	switch {
	case data[p] > 57:
		if 59 <= data[p] && data[p] <= 68 {
			goto st121
		}
	default:
		goto st121
	}
{
	goto st121

}
st149:
	if p++; p == pe {
		goto _test_eof149
	}
st_case_149:
	switch data[p] {
	case 58:
goto st122
	case 61:
goto st150
	}
	switch {
	case data[p] > 57:
		if 59 <= data[p] && data[p] <= 60 {
			goto st121
		}
	default:
		goto st121
	}
{
	goto st121

}
st150:
	if p++; p == pe {
		goto _test_eof150
	}
st_case_150:
	if data[p] == 58 {
		goto st122
	}
	switch {
	case data[p] > 47:
		if data[p] <= 57 {
			goto st151
		}
	default:
		goto st121
	}
{
	goto st121

}
st151:
	if p++; p == pe {
		goto _test_eof151
	}
st_case_151:
	switch data[p] {
	case 10:
goto ctr151
	case 13:
goto ctr152
	case 32:
goto st128
	case 58:
goto st122
	}
	switch {
	case data[p] < 14:
		switch {
		case data[p] > 9:
			if 11 <= data[p] && data[p] <= 12 {
				goto st121
			}
		default:
			goto st121
		}
	case data[p] > 31:
		switch {
		case data[p] > 47:
			if data[p] <= 57 {
				goto st151
			}
		case data[p] >= 33:
			goto st121
		}
	default:
		goto st121
	}
{
	goto st121

}
st152:
	if p++; p == pe {
		goto _test_eof152
	}
st_case_152:
	switch data[p] {
	case 32:
goto st125
	case 33:
goto st121
	case 34:
goto st125
	case 43:
goto st125
	case 44:
goto st121
	case 47:
goto st121
	case 58:
goto st122
	case 95:
goto st125
	case 96:
goto st121
	}
	switch {
	case data[p] < 48:
		switch {
		case data[p] < 35:
			if data[p] <= 31 {
				goto st121
			}
		case data[p] > 42:
			if 45 <= data[p] && data[p] <= 46 {
				goto st125
			}
		default:
			goto st121
		}
	case data[p] > 57:
		switch {
		case data[p] < 64:
			if 59 <= data[p] {
				goto st121
			}
		case data[p] > 90:
			switch {
			case data[p] > 94:
				if 97 <= data[p] && data[p] <= 122 {
					goto st125
				}
			default:
				goto st121
			}
		default:
			goto st125
		}
	default:
		goto st125
	}
{
	goto st121

}
st153:
	if p++; p == pe {
		goto _test_eof153
	}
st_case_153:
	switch data[p] {
	case 32:
goto st124
	case 33:
goto st121
	case 34:
goto st124
	case 43:
goto st124
	case 44:
goto st121
	case 47:
goto st121
	case 58:
goto st122
	case 95:
goto st124
	case 96:
goto st121
	}
	switch {
	case data[p] < 48:
		switch {
		case data[p] < 35:
			if data[p] <= 31 {
				goto st121
			}
		case data[p] > 42:
			if 45 <= data[p] && data[p] <= 46 {
				goto st124
			}
		default:
			goto st121
		}
	case data[p] > 57:
		switch {
		case data[p] < 64:
			if 59 <= data[p] {
				goto st121
			}
		case data[p] > 90:
			switch {
			case data[p] > 94:
				if 97 <= data[p] && data[p] <= 122 {
					goto st124
				}
			default:
				goto st121
			}
		default:
			goto st124
		}
	default:
		goto st124
	}
{
	goto st121

}
st154:
	if p++; p == pe {
		goto _test_eof154
	}
st_case_154:
	if data[p] == 85 {
		goto st155
	}
{
	goto st0

}
st155:
	if p++; p == pe {
		goto _test_eof155
	}
st_case_155:
	if data[p] == 73 {
		goto st156
	}
{
	goto st0

}
st156:
	if p++; p == pe {
		goto _test_eof156
	}
st_case_156:
	if data[p] == 84 {
		goto st157
	}
{
	goto st0

}
st157:
	if p++; p == pe {
		goto _test_eof157
	}
st_case_157:
	switch data[p] {
	case 10:
goto ctr187
	case 13:
goto ctr188
	}
{
	goto st0

}
st158:
	if p++; p == pe {
		goto _test_eof158
	}
st_case_158:
	switch data[p] {
	case 67:
goto st159
	case 83:
goto st183
	}
{
	goto st0

}
st159:
	if p++; p == pe {
		goto _test_eof159
	}
st_case_159:
	if data[p] == 80 {
		goto st160
	}
{
	goto st0

}
st160:
	if p++; p == pe {
		goto _test_eof160
	}
st_case_160:
	if data[p] == 84 {
		goto st161
	}
{
	goto st0

}
st161:
	if p++; p == pe {
		goto _test_eof161
	}
st_case_161:
	if data[p] == 32 {
		goto st162
	}
{
	goto st0

}
st162:
	if p++; p == pe {
		goto _test_eof162
	}
st_case_162:
	switch data[p] {
	case 84:
goto st163
	case 116:
goto st163
	}
{
	goto st0

}
st163:
	if p++; p == pe {
		goto _test_eof163
	}
st_case_163:
	switch data[p] {
	case 79:
goto st164
	case 111:
goto st164
	}
{
	goto st0

}
st164:
	if p++; p == pe {
		goto _test_eof164
	}
st_case_164:
	if data[p] == 58 {
		goto st165
	}
{
	goto st0

}
st165:
	if p++; p == pe {
		goto _test_eof165
	}
st_case_165:
	if data[p] == 60 {
		goto st166
	}
{
	goto st0

}
st166:
	if p++; p == pe {
		goto _test_eof166
	}
st_case_166:
	switch data[p] {
	case 34:
goto ctr198
	case 43:
goto ctr199
	case 64:
goto ctr200
	case 92:
goto ctr201
	case 95:
goto ctr199
	}
	switch {
	case data[p] < 48:
		if 45 <= data[p] && data[p] <= 46 {
			goto ctr199
		}
	case data[p] > 57:
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto ctr199
			}
		case data[p] >= 65:
			goto ctr199
		}
	default:
		goto ctr199
	}
{
	goto st0

}
ctr198:
// line 26 "smtp.ragel"
	{

	pb = p
	parser.buffer = bytes.NewBuffer(nil)
}
	goto st167
st167:
	if p++; p == pe {
		goto _test_eof167
	}
st_case_167:
// line 4251 "smtp.go"
	switch data[p] {
	case 43:
goto st168
	case 92:
goto st172
	case 95:
goto st168
	}
	switch {
	case data[p] < 48:
		if 45 <= data[p] && data[p] <= 46 {
			goto st168
		}
	case data[p] > 57:
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st168
			}
		case data[p] >= 64:
			goto st168
		}
	default:
		goto st168
	}
{
	goto st0

}
st168:
	if p++; p == pe {
		goto _test_eof168
	}
st_case_168:
	switch data[p] {
	case 34:
goto st169
	case 43:
goto st168
	case 92:
goto st172
	case 95:
goto st168
	}
	switch {
	case data[p] < 48:
		if 45 <= data[p] && data[p] <= 46 {
			goto st168
		}
	case data[p] > 57:
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st168
			}
		case data[p] >= 64:
			goto st168
		}
	default:
		goto st168
	}
{
	goto st0

}
ctr199:
// line 26 "smtp.ragel"
	{

	pb = p
	parser.buffer = bytes.NewBuffer(nil)
}
	goto st169
st169:
	if p++; p == pe {
		goto _test_eof169
	}
st_case_169:
// line 4330 "smtp.go"
	switch data[p] {
	case 34:
goto st167
	case 43:
goto st169
	case 62:
goto ctr206
	case 92:
goto st171
	case 95:
goto st169
	}
	switch {
	case data[p] < 48:
		if 45 <= data[p] && data[p] <= 46 {
			goto st169
		}
	case data[p] > 57:
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st169
			}
		case data[p] >= 64:
			goto st169
		}
	default:
		goto st169
	}
{
	goto st0

}
ctr206:
// line 20 "smtp.ragel"
	{

	parser.buffer.Write(data[pb:p])
	parser.current.Data = parser.buffer.Bytes()
	parser.buffer = nil
}
	goto st170
st170:
	if p++; p == pe {
		goto _test_eof170
	}
st_case_170:
// line 4378 "smtp.go"
	switch data[p] {
	case 10:
goto ctr208
	case 13:
goto ctr209
	}
{
	goto st0

}
ctr201:
// line 26 "smtp.ragel"
	{

	pb = p
	parser.buffer = bytes.NewBuffer(nil)
}
	goto st171
st171:
	if p++; p == pe {
		goto _test_eof171
	}
st_case_171:
// line 4402 "smtp.go"
	switch data[p] {
	case 32:
goto st169
	case 34:
goto st169
	case 43:
goto st169
	case 95:
goto st169
	}
	switch {
	case data[p] < 48:
		if 45 <= data[p] && data[p] <= 46 {
			goto st169
		}
	case data[p] > 57:
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st169
			}
		case data[p] >= 64:
			goto st169
		}
	default:
		goto st169
	}
{
	goto st0

}
st172:
	if p++; p == pe {
		goto _test_eof172
	}
st_case_172:
	switch data[p] {
	case 32:
goto st168
	case 34:
goto st168
	case 43:
goto st168
	case 95:
goto st168
	}
	switch {
	case data[p] < 48:
		if 45 <= data[p] && data[p] <= 46 {
			goto st168
		}
	case data[p] > 57:
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st168
			}
		case data[p] >= 64:
			goto st168
		}
	default:
		goto st168
	}
{
	goto st0

}
ctr200:
// line 26 "smtp.ragel"
	{

	pb = p
	parser.buffer = bytes.NewBuffer(nil)
}
	goto st173
st173:
	if p++; p == pe {
		goto _test_eof173
	}
st_case_173:
// line 4483 "smtp.go"
	switch data[p] {
	case 34:
goto st176
	case 43:
goto st178
	case 44:
goto st174
	case 47:
goto st174
	case 58:
goto st0
	case 62:
goto ctr213
	case 63:
goto st174
	case 91:
goto st174
	case 92:
goto st181
	case 95:
goto st178
	case 96:
goto st174
	}
	switch {
	case data[p] < 48:
		switch {
		case data[p] < 35:
			if data[p] <= 33 {
				goto st174
			}
		case data[p] > 42:
			if 45 <= data[p] && data[p] <= 46 {
				goto st178
			}
		default:
			goto st174
		}
	case data[p] > 57:
		switch {
		case data[p] < 64:
			if 59 <= data[p] && data[p] <= 61 {
				goto st174
			}
		case data[p] > 90:
			switch {
			case data[p] > 94:
				if 97 <= data[p] && data[p] <= 122 {
					goto st178
				}
			case data[p] >= 93:
				goto st174
			}
		default:
			goto st178
		}
	default:
		goto st178
	}
{
	goto st174

}
st174:
	if p++; p == pe {
		goto _test_eof174
	}
st_case_174:
	if data[p] == 58 {
		goto st175
	}
	if data[p] <= 57 {
		goto st174
	}
{
	goto st174

}
st175:
	if p++; p == pe {
		goto _test_eof175
	}
st_case_175:
	switch data[p] {
	case 34:
goto ctr198
	case 43:
goto ctr199
	case 92:
goto ctr201
	case 95:
goto ctr199
	}
	switch {
	case data[p] < 48:
		if 45 <= data[p] && data[p] <= 46 {
			goto ctr199
		}
	case data[p] > 57:
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto ctr199
			}
		case data[p] >= 64:
			goto ctr199
		}
	default:
		goto ctr199
	}
{
	goto st0

}
st176:
	if p++; p == pe {
		goto _test_eof176
	}
st_case_176:
	switch data[p] {
	case 43:
goto st177
	case 44:
goto st174
	case 47:
goto st174
	case 58:
goto st175
	case 91:
goto st174
	case 92:
goto st182
	case 95:
goto st177
	case 96:
goto st174
	}
	switch {
	case data[p] < 59:
		switch {
		case data[p] < 45:
			if data[p] <= 42 {
				goto st174
			}
		case data[p] > 46:
			if 48 <= data[p] && data[p] <= 57 {
				goto st177
			}
		default:
			goto st177
		}
	case data[p] > 63:
		switch {
		case data[p] < 93:
			if data[p] <= 90 {
				goto st177
			}
		case data[p] > 94:
			if 97 <= data[p] && data[p] <= 122 {
				goto st177
			}
		default:
			goto st174
		}
	default:
		goto st174
	}
{
	goto st174

}
st177:
	if p++; p == pe {
		goto _test_eof177
	}
st_case_177:
	switch data[p] {
	case 34:
goto st178
	case 43:
goto st177
	case 44:
goto st174
	case 47:
goto st174
	case 58:
goto st175
	case 91:
goto st174
	case 92:
goto st182
	case 95:
goto st177
	case 96:
goto st174
	}
	switch {
	case data[p] < 48:
		switch {
		case data[p] < 35:
			if data[p] <= 33 {
				goto st174
			}
		case data[p] > 42:
			if 45 <= data[p] && data[p] <= 46 {
				goto st177
			}
		default:
			goto st174
		}
	case data[p] > 57:
		switch {
		case data[p] < 64:
			if 59 <= data[p] {
				goto st174
			}
		case data[p] > 90:
			switch {
			case data[p] > 94:
				if 97 <= data[p] && data[p] <= 122 {
					goto st177
				}
			case data[p] >= 93:
				goto st174
			}
		default:
			goto st177
		}
	default:
		goto st177
	}
{
	goto st174

}
st178:
	if p++; p == pe {
		goto _test_eof178
	}
st_case_178:
	switch data[p] {
	case 34:
goto st176
	case 43:
goto st178
	case 44:
goto st174
	case 47:
goto st174
	case 58:
goto st175
	case 62:
goto ctr213
	case 63:
goto st174
	case 91:
goto st174
	case 92:
goto st181
	case 95:
goto st178
	case 96:
goto st174
	}
	switch {
	case data[p] < 48:
		switch {
		case data[p] < 35:
			if data[p] <= 33 {
				goto st174
			}
		case data[p] > 42:
			if 45 <= data[p] && data[p] <= 46 {
				goto st178
			}
		default:
			goto st174
		}
	case data[p] > 57:
		switch {
		case data[p] < 64:
			if 59 <= data[p] && data[p] <= 61 {
				goto st174
			}
		case data[p] > 90:
			switch {
			case data[p] > 94:
				if 97 <= data[p] && data[p] <= 122 {
					goto st178
				}
			case data[p] >= 93:
				goto st174
			}
		default:
			goto st178
		}
	default:
		goto st178
	}
{
	goto st174

}
ctr213:
// line 20 "smtp.ragel"
	{

	parser.buffer.Write(data[pb:p])
	parser.current.Data = parser.buffer.Bytes()
	parser.buffer = nil
}
	goto st179
st179:
	if p++; p == pe {
		goto _test_eof179
	}
st_case_179:
// line 4801 "smtp.go"
	switch data[p] {
	case 10:
goto ctr218
	case 13:
goto ctr219
	case 58:
goto st175
	}
	switch {
	case data[p] < 11:
		if data[p] <= 9 {
			goto st174
		}
	case data[p] > 12:
		if 14 <= data[p] && data[p] <= 57 {
			goto st174
		}
	default:
		goto st174
	}
{
	goto st174

}
ctr220:
// line 72 "smtp.ragel"
	{
{ p++; cs = 202; goto _out }
}
	goto st202
ctr218:
// line 63 "smtp.ragel"
	{
parser.current.Verb = VerbRCPT}
// line 72 "smtp.ragel"
	{
{ p++; cs = 202; goto _out }
}
	goto st202
st202:
	if p++; p == pe {
		goto _test_eof202
	}
st_case_202:
// line 4846 "smtp.go"
	if data[p] == 58 {
		goto st175
	}
	if data[p] <= 57 {
		goto st174
	}
{
	goto st174

}
ctr219:
// line 63 "smtp.ragel"
	{
parser.current.Verb = VerbRCPT}
	goto st180
st180:
	if p++; p == pe {
		goto _test_eof180
	}
st_case_180:
// line 4867 "smtp.go"
	switch data[p] {
	case 10:
goto ctr220
	case 58:
goto st175
	}
	switch {
	case data[p] > 9:
		if 11 <= data[p] && data[p] <= 57 {
			goto st174
		}
	default:
		goto st174
	}
{
	goto st174

}
st181:
	if p++; p == pe {
		goto _test_eof181
	}
st_case_181:
	switch data[p] {
	case 32:
goto st178
	case 33:
goto st174
	case 34:
goto st178
	case 43:
goto st178
	case 44:
goto st174
	case 47:
goto st174
	case 58:
goto st175
	case 95:
goto st178
	case 96:
goto st174
	}
	switch {
	case data[p] < 48:
		switch {
		case data[p] < 35:
			if data[p] <= 31 {
				goto st174
			}
		case data[p] > 42:
			if 45 <= data[p] && data[p] <= 46 {
				goto st178
			}
		default:
			goto st174
		}
	case data[p] > 57:
		switch {
		case data[p] < 64:
			if 59 <= data[p] {
				goto st174
			}
		case data[p] > 90:
			switch {
			case data[p] > 94:
				if 97 <= data[p] && data[p] <= 122 {
					goto st178
				}
			default:
				goto st174
			}
		default:
			goto st178
		}
	default:
		goto st178
	}
{
	goto st174

}
st182:
	if p++; p == pe {
		goto _test_eof182
	}
st_case_182:
	switch data[p] {
	case 32:
goto st177
	case 33:
goto st174
	case 34:
goto st177
	case 43:
goto st177
	case 44:
goto st174
	case 47:
goto st174
	case 58:
goto st175
	case 95:
goto st177
	case 96:
goto st174
	}
	switch {
	case data[p] < 48:
		switch {
		case data[p] < 35:
			if data[p] <= 31 {
				goto st174
			}
		case data[p] > 42:
			if 45 <= data[p] && data[p] <= 46 {
				goto st177
			}
		default:
			goto st174
		}
	case data[p] > 57:
		switch {
		case data[p] < 64:
			if 59 <= data[p] {
				goto st174
			}
		case data[p] > 90:
			switch {
			case data[p] > 94:
				if 97 <= data[p] && data[p] <= 122 {
					goto st177
				}
			default:
				goto st174
			}
		default:
			goto st177
		}
	default:
		goto st177
	}
{
	goto st174

}
st183:
	if p++; p == pe {
		goto _test_eof183
	}
st_case_183:
	if data[p] == 69 {
		goto st184
	}
{
	goto st0

}
st184:
	if p++; p == pe {
		goto _test_eof184
	}
st_case_184:
	if data[p] == 84 {
		goto st185
	}
{
	goto st0

}
st185:
	if p++; p == pe {
		goto _test_eof185
	}
st_case_185:
	switch data[p] {
	case 10:
goto ctr223
	case 13:
goto ctr224
	}
{
	goto st0

}
st186:
	if p++; p == pe {
		goto _test_eof186
	}
st_case_186:
	if data[p] == 84 {
		goto st187
	}
{
	goto st0

}
st187:
	if p++; p == pe {
		goto _test_eof187
	}
st_case_187:
	if data[p] == 65 {
		goto st188
	}
{
	goto st0

}
st188:
	if p++; p == pe {
		goto _test_eof188
	}
st_case_188:
	if data[p] == 82 {
		goto st189
	}
{
	goto st0

}
st189:
	if p++; p == pe {
		goto _test_eof189
	}
st_case_189:
	if data[p] == 84 {
		goto st190
	}
{
	goto st0

}
st190:
	if p++; p == pe {
		goto _test_eof190
	}
st_case_190:
	if data[p] == 84 {
		goto st191
	}
{
	goto st0

}
st191:
	if p++; p == pe {
		goto _test_eof191
	}
st_case_191:
	if data[p] == 76 {
		goto st192
	}
{
	goto st0

}
st192:
	if p++; p == pe {
		goto _test_eof192
	}
st_case_192:
	if data[p] == 83 {
		goto st193
	}
{
	goto st0

}
st193:
	if p++; p == pe {
		goto _test_eof193
	}
st_case_193:
	switch data[p] {
	case 10:
goto ctr232
	case 13:
goto ctr233
	}
{
	goto st0

}
st194:
	if p++; p == pe {
		goto _test_eof194
	}
st_case_194:
	if data[p] == 82 {
		goto st195
	}
{
	goto st0

}
st195:
	if p++; p == pe {
		goto _test_eof195
	}
st_case_195:
	if data[p] == 70 {
		goto st196
	}
{
	goto st0

}
st196:
	if p++; p == pe {
		goto _test_eof196
	}
st_case_196:
	if data[p] == 89 {
		goto st197
	}
{
	goto st0

}
st197:
	if p++; p == pe {
		goto _test_eof197
	}
st_case_197:
	if data[p] == 32 {
		goto st198
	}
{
	goto st0

}
st198:
	if p++; p == pe {
		goto _test_eof198
	}
st_case_198:
	switch data[p] {
	case 10:
goto st0
	case 13:
goto st0
	}
	switch {
	case data[p] > 9:
		if 11 <= data[p] && data[p] <= 12 {
			goto st199
		}
	default:
		goto st199
	}
{
	goto st199

}
st199:
	if p++; p == pe {
		goto _test_eof199
	}
st_case_199:
	switch data[p] {
	case 10:
goto ctr239
	case 13:
goto ctr240
	}
	switch {
	case data[p] > 9:
		if 11 <= data[p] && data[p] <= 12 {
			goto st199
		}
	default:
		goto st199
	}
{
	goto st199

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
	_test_eof200: cs = 200
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
	_test_eof201: cs = 201
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
	_test_eof144: cs = 144
	goto _test_eof
	_test_eof145: cs = 145
	goto _test_eof
	_test_eof146: cs = 146
	goto _test_eof
	_test_eof147: cs = 147
	goto _test_eof
	_test_eof148: cs = 148
	goto _test_eof
	_test_eof149: cs = 149
	goto _test_eof
	_test_eof150: cs = 150
	goto _test_eof
	_test_eof151: cs = 151
	goto _test_eof
	_test_eof152: cs = 152
	goto _test_eof
	_test_eof153: cs = 153
	goto _test_eof
	_test_eof154: cs = 154
	goto _test_eof
	_test_eof155: cs = 155
	goto _test_eof
	_test_eof156: cs = 156
	goto _test_eof
	_test_eof157: cs = 157
	goto _test_eof
	_test_eof158: cs = 158
	goto _test_eof
	_test_eof159: cs = 159
	goto _test_eof
	_test_eof160: cs = 160
	goto _test_eof
	_test_eof161: cs = 161
	goto _test_eof
	_test_eof162: cs = 162
	goto _test_eof
	_test_eof163: cs = 163
	goto _test_eof
	_test_eof164: cs = 164
	goto _test_eof
	_test_eof165: cs = 165
	goto _test_eof
	_test_eof166: cs = 166
	goto _test_eof
	_test_eof167: cs = 167
	goto _test_eof
	_test_eof168: cs = 168
	goto _test_eof
	_test_eof169: cs = 169
	goto _test_eof
	_test_eof170: cs = 170
	goto _test_eof
	_test_eof171: cs = 171
	goto _test_eof
	_test_eof172: cs = 172
	goto _test_eof
	_test_eof173: cs = 173
	goto _test_eof
	_test_eof174: cs = 174
	goto _test_eof
	_test_eof175: cs = 175
	goto _test_eof
	_test_eof176: cs = 176
	goto _test_eof
	_test_eof177: cs = 177
	goto _test_eof
	_test_eof178: cs = 178
	goto _test_eof
	_test_eof179: cs = 179
	goto _test_eof
	_test_eof202: cs = 202
	goto _test_eof
	_test_eof180: cs = 180
	goto _test_eof
	_test_eof181: cs = 181
	goto _test_eof
	_test_eof182: cs = 182
	goto _test_eof
	_test_eof183: cs = 183
	goto _test_eof
	_test_eof184: cs = 184
	goto _test_eof
	_test_eof185: cs = 185
	goto _test_eof
	_test_eof186: cs = 186
	goto _test_eof
	_test_eof187: cs = 187
	goto _test_eof
	_test_eof188: cs = 188
	goto _test_eof
	_test_eof189: cs = 189
	goto _test_eof
	_test_eof190: cs = 190
	goto _test_eof
	_test_eof191: cs = 191
	goto _test_eof
	_test_eof192: cs = 192
	goto _test_eof
	_test_eof193: cs = 193
	goto _test_eof
	_test_eof194: cs = 194
	goto _test_eof
	_test_eof195: cs = 195
	goto _test_eof
	_test_eof196: cs = 196
	goto _test_eof
	_test_eof197: cs = 197
	goto _test_eof
	_test_eof198: cs = 198
	goto _test_eof
	_test_eof199: cs = 199
	goto _test_eof

	_test_eof: {}
	_out: {}
	}

// line 99 "smtp.ragel"

	if cs == smtp_error {
		return data[p:], fmt.Errorf("Invalid character in pos %d: `%c`.", p, data[p])
	}

	// Not yet a full parse, remeber everything from pb to p if we are
	// recording.
	if cs < smtp_first_final {
		if parser.buffer != nil {
			_, err := parser.buffer.Write(data[pb:p])
			if err != nil {
				return data[p:], fmt.Errorf("Could not buffer data: %s", err)
			}
		}
		parser.cs = cs
		return data[p:], Dangling
	}

	// Full parse:
	parser.cs = smtp_en_main
	return data[p:], nil
}
