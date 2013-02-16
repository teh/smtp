
// line 1 "smtp.ragel"
package smtp

import (
	"fmt"
)

// Good introduction: http://cr.yp.to/smtp/request.html

// line 51 "smtp.ragel"



// line 16 "smtp.go"
const smtp_start = 1;
const smtp_first_final = 153;
const smtp_error = 0;

const smtp_en_main = 1


// line 54 "smtp.ragel"

func ParseEntry(data []byte) error {
	var cs, p, pb, pdomain int
	var current Action
	pe := len(data)
	
// line 31 "smtp.go"
	{
	cs = smtp_start
	}

// line 60 "smtp.ragel"
	
// line 38 "smtp.go"
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
	case 153:
		goto st_case_153
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
	case 154:
		goto st_case_154
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
	case 155:
		goto st_case_155
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
goto st112
	case 82:
goto st115
	case 86:
goto st147
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
goto st153
	case 13:
goto st6
	}
{
	goto st0

}
ctr19:
// line 14 "smtp.ragel"
	{
current.Domain = data[pdomain:p]}
// line 21 "smtp.ragel"
	{
current.ClientHost = data[pb:p]}
// line 41 "smtp.ragel"
	{
emit(ActionEHLO, current)}
	goto st153
ctr41:
// line 24 "smtp.ragel"
	{
current.ClientHost = data[pb:p]}
// line 41 "smtp.ragel"
	{
emit(ActionEHLO, current)}
	goto st153
ctr59:
// line 14 "smtp.ragel"
	{
current.Domain = data[pdomain:p]}
// line 21 "smtp.ragel"
	{
current.ClientHost = data[pb:p]}
// line 40 "smtp.ragel"
	{
emit(ActionHELO, current)}
	goto st153
ctr81:
// line 24 "smtp.ragel"
	{
current.ClientHost = data[pb:p]}
// line 40 "smtp.ragel"
	{
emit(ActionHELO, current)}
	goto st153
ctr114:
// line 37 "smtp.ragel"
	{
emit(ActionMAIL, current)}
	goto st153
ctr158:
// line 38 "smtp.ragel"
	{
emit(ActionRCPT, current)}
	goto st153
st153:
	if p++; p == pe {
		goto _test_eof153
	}
st_case_153:
// line 487 "smtp.go"
{
	goto st0

}
ctr20:
// line 14 "smtp.ragel"
	{
current.Domain = data[pdomain:p]}
// line 21 "smtp.ragel"
	{
current.ClientHost = data[pb:p]}
// line 41 "smtp.ragel"
	{
emit(ActionEHLO, current)}
	goto st6
ctr42:
// line 24 "smtp.ragel"
	{
current.ClientHost = data[pb:p]}
// line 41 "smtp.ragel"
	{
emit(ActionEHLO, current)}
	goto st6
ctr60:
// line 14 "smtp.ragel"
	{
current.Domain = data[pdomain:p]}
// line 21 "smtp.ragel"
	{
current.ClientHost = data[pb:p]}
// line 40 "smtp.ragel"
	{
emit(ActionHELO, current)}
	goto st6
ctr82:
// line 24 "smtp.ragel"
	{
current.ClientHost = data[pb:p]}
// line 40 "smtp.ragel"
	{
emit(ActionHELO, current)}
	goto st6
ctr115:
// line 37 "smtp.ragel"
	{
emit(ActionMAIL, current)}
	goto st6
ctr159:
// line 38 "smtp.ragel"
	{
emit(ActionRCPT, current)}
	goto st6
st6:
	if p++; p == pe {
		goto _test_eof6
	}
st_case_6:
// line 545 "smtp.go"
	if data[p] == 10 {
		goto st153
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
		goto ctr18
	}
	switch {
	case data[p] < 65:
		if 48 <= data[p] && data[p] <= 57 {
			goto ctr17
		}
	case data[p] > 90:
		if 97 <= data[p] && data[p] <= 122 {
			goto ctr17
		}
	default:
		goto ctr17
	}
{
	goto st0

}
ctr17:
// line 22 "smtp.ragel"
	{
pb = p}
// line 15 "smtp.ragel"
	{
pdomain = p}
	goto st12
st12:
	if p++; p == pe {
		goto _test_eof12
	}
st_case_12:
// line 638 "smtp.go"
	switch data[p] {
	case 10:
goto ctr19
	case 13:
goto ctr20
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
ctr18:
// line 25 "smtp.ragel"
	{
pb = p}
	goto st14
st14:
	if p++; p == pe {
		goto _test_eof14
	}
st_case_14:
// line 694 "smtp.go"
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
goto ctr41
	case 13:
goto ctr42
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
// line 22 "smtp.ragel"
	{
pb = p}
// line 15 "smtp.ragel"
	{
pdomain = p}
	goto st48
st48:
	if p++; p == pe {
		goto _test_eof48
	}
st_case_48:
// line 1298 "smtp.go"
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
// line 25 "smtp.ragel"
	{
pb = p}
	goto st50
st50:
	if p++; p == pe {
		goto _test_eof50
	}
st_case_50:
// line 1354 "smtp.go"
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
goto ctr81
	case 13:
goto ctr82
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
goto ctr103
	case 43:
goto ctr104
	case 46:
goto ctr104
	case 64:
goto ctr105
	case 92:
goto ctr106
	}
	switch {
	case data[p] < 65:
		if 48 <= data[p] && data[p] <= 57 {
			goto ctr104
		}
	case data[p] > 90:
		if 97 <= data[p] && data[p] <= 122 {
			goto ctr104
		}
	default:
		goto ctr104
	}
{
	goto st0

}
ctr103:
// line 32 "smtp.ragel"
	{
pb = p}
	goto st90
st90:
	if p++; p == pe {
		goto _test_eof90
	}
st_case_90:
// line 2048 "smtp.go"
	switch data[p] {
	case 43:
goto st91
	case 46:
goto st91
	case 92:
goto st98
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
goto st98
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
ctr104:
// line 32 "smtp.ragel"
	{
pb = p}
	goto st92
st92:
	if p++; p == pe {
		goto _test_eof92
	}
st_case_92:
// line 2114 "smtp.go"
	switch data[p] {
	case 34:
goto st90
	case 43:
goto st92
	case 46:
goto st92
	case 62:
goto ctr111
	case 64:
goto st94
	case 92:
goto st97
	}
	switch {
	case data[p] < 65:
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
ctr111:
// line 31 "smtp.ragel"
	{
current.Address = data[pb:p]}
	goto st93
ctr119:
// line 14 "smtp.ragel"
	{
current.Domain = data[pdomain:p]}
// line 31 "smtp.ragel"
	{
current.Address = data[pb:p]}
	goto st93
st93:
	if p++; p == pe {
		goto _test_eof93
	}
st_case_93:
// line 2163 "smtp.go"
	switch data[p] {
	case 10:
goto ctr114
	case 13:
goto ctr115
	}
{
	goto st0

}
st94:
	if p++; p == pe {
		goto _test_eof94
	}
st_case_94:
	switch data[p] {
	case 34:
goto st90
	case 43:
goto st92
	case 46:
goto st92
	case 62:
goto ctr111
	case 64:
goto st94
	case 92:
goto st97
	}
	switch {
	case data[p] < 65:
		if 48 <= data[p] && data[p] <= 57 {
			goto ctr116
		}
	case data[p] > 90:
		if 97 <= data[p] && data[p] <= 122 {
			goto ctr116
		}
	default:
		goto ctr116
	}
{
	goto st0

}
ctr116:
// line 15 "smtp.ragel"
	{
pdomain = p}
	goto st95
st95:
	if p++; p == pe {
		goto _test_eof95
	}
st_case_95:
// line 2219 "smtp.go"
	switch data[p] {
	case 34:
goto st90
	case 43:
goto st92
	case 46:
goto st96
	case 62:
goto ctr119
	case 64:
goto st94
	case 92:
goto st97
	}
	switch {
	case data[p] < 65:
		if 48 <= data[p] && data[p] <= 57 {
			goto st95
		}
	case data[p] > 90:
		if 97 <= data[p] && data[p] <= 122 {
			goto st95
		}
	default:
		goto st95
	}
{
	goto st0

}
st96:
	if p++; p == pe {
		goto _test_eof96
	}
st_case_96:
	switch data[p] {
	case 34:
goto st90
	case 43:
goto st92
	case 46:
goto st92
	case 62:
goto ctr111
	case 64:
goto st94
	case 92:
goto st97
	}
	switch {
	case data[p] < 65:
		if 48 <= data[p] && data[p] <= 57 {
			goto st95
		}
	case data[p] > 90:
		if 97 <= data[p] && data[p] <= 122 {
			goto st95
		}
	default:
		goto st95
	}
{
	goto st0

}
ctr106:
// line 32 "smtp.ragel"
	{
pb = p}
	goto st97
st97:
	if p++; p == pe {
		goto _test_eof97
	}
st_case_97:
// line 2295 "smtp.go"
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
st98:
	if p++; p == pe {
		goto _test_eof98
	}
st_case_98:
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
ctr105:
// line 32 "smtp.ragel"
	{
pb = p}
	goto st99
st99:
	if p++; p == pe {
		goto _test_eof99
	}
st_case_99:
// line 2363 "smtp.go"
	switch data[p] {
	case 34:
goto st102
	case 43:
goto st104
	case 46:
goto st104
	case 47:
goto st100
	case 58:
goto st0
	case 62:
goto ctr123
	case 63:
goto st100
	case 64:
goto st107
	case 91:
goto st100
	case 92:
goto st110
	}
	switch {
	case data[p] < 48:
		switch {
		case data[p] < 35:
			if data[p] <= 33 {
				goto st100
			}
		case data[p] > 42:
			if 44 <= data[p] && data[p] <= 45 {
				goto st100
			}
		default:
			goto st100
		}
	case data[p] > 57:
		switch {
		case data[p] < 65:
			if 59 <= data[p] && data[p] <= 61 {
				goto st100
			}
		case data[p] > 90:
			switch {
			case data[p] > 96:
				if data[p] <= 122 {
					goto st104
				}
			case data[p] >= 93:
				goto st100
			}
		default:
			goto st104
		}
	default:
		goto st104
	}
{
	goto st100

}
st100:
	if p++; p == pe {
		goto _test_eof100
	}
st_case_100:
	if data[p] == 58 {
		goto st101
	}
	if data[p] <= 57 {
		goto st100
	}
{
	goto st100

}
st101:
	if p++; p == pe {
		goto _test_eof101
	}
st_case_101:
	switch data[p] {
	case 34:
goto ctr103
	case 43:
goto ctr104
	case 46:
goto ctr104
	case 92:
goto ctr106
	}
	switch {
	case data[p] < 64:
		if 48 <= data[p] && data[p] <= 57 {
			goto ctr104
		}
	case data[p] > 90:
		if 97 <= data[p] && data[p] <= 122 {
			goto ctr104
		}
	default:
		goto ctr104
	}
{
	goto st0

}
st102:
	if p++; p == pe {
		goto _test_eof102
	}
st_case_102:
	switch data[p] {
	case 43:
goto st103
	case 46:
goto st103
	case 47:
goto st100
	case 58:
goto st101
	case 91:
goto st100
	case 92:
goto st111
	}
	switch {
	case data[p] < 59:
		switch {
		case data[p] < 44:
			if data[p] <= 42 {
				goto st100
			}
		case data[p] > 45:
			if 48 <= data[p] && data[p] <= 57 {
				goto st103
			}
		default:
			goto st100
		}
	case data[p] > 63:
		switch {
		case data[p] < 93:
			if data[p] <= 90 {
				goto st103
			}
		case data[p] > 96:
			if data[p] <= 122 {
				goto st103
			}
		default:
			goto st100
		}
	default:
		goto st100
	}
{
	goto st100

}
st103:
	if p++; p == pe {
		goto _test_eof103
	}
st_case_103:
	switch data[p] {
	case 34:
goto st104
	case 43:
goto st103
	case 46:
goto st103
	case 47:
goto st100
	case 58:
goto st101
	case 91:
goto st100
	case 92:
goto st111
	}
	switch {
	case data[p] < 48:
		switch {
		case data[p] < 35:
			if data[p] <= 33 {
				goto st100
			}
		case data[p] > 42:
			if 44 <= data[p] && data[p] <= 45 {
				goto st100
			}
		default:
			goto st100
		}
	case data[p] > 57:
		switch {
		case data[p] < 64:
			if 59 <= data[p] {
				goto st100
			}
		case data[p] > 90:
			switch {
			case data[p] > 96:
				if data[p] <= 122 {
					goto st103
				}
			case data[p] >= 93:
				goto st100
			}
		default:
			goto st103
		}
	default:
		goto st103
	}
{
	goto st100

}
st104:
	if p++; p == pe {
		goto _test_eof104
	}
st_case_104:
	switch data[p] {
	case 34:
goto st102
	case 43:
goto st104
	case 46:
goto st104
	case 47:
goto st100
	case 58:
goto st101
	case 62:
goto ctr123
	case 63:
goto st100
	case 64:
goto st107
	case 91:
goto st100
	case 92:
goto st110
	}
	switch {
	case data[p] < 48:
		switch {
		case data[p] < 35:
			if data[p] <= 33 {
				goto st100
			}
		case data[p] > 42:
			if 44 <= data[p] && data[p] <= 45 {
				goto st100
			}
		default:
			goto st100
		}
	case data[p] > 57:
		switch {
		case data[p] < 65:
			if 59 <= data[p] && data[p] <= 61 {
				goto st100
			}
		case data[p] > 90:
			switch {
			case data[p] > 96:
				if data[p] <= 122 {
					goto st104
				}
			case data[p] >= 93:
				goto st100
			}
		default:
			goto st104
		}
	default:
		goto st104
	}
{
	goto st100

}
ctr123:
// line 31 "smtp.ragel"
	{
current.Address = data[pb:p]}
	goto st105
ctr135:
// line 14 "smtp.ragel"
	{
current.Domain = data[pdomain:p]}
// line 31 "smtp.ragel"
	{
current.Address = data[pb:p]}
	goto st105
st105:
	if p++; p == pe {
		goto _test_eof105
	}
st_case_105:
// line 2668 "smtp.go"
	switch data[p] {
	case 10:
goto ctr129
	case 13:
goto ctr130
	case 58:
goto st101
	}
	switch {
	case data[p] < 11:
		if data[p] <= 9 {
			goto st100
		}
	case data[p] > 12:
		if 14 <= data[p] && data[p] <= 57 {
			goto st100
		}
	default:
		goto st100
	}
{
	goto st100

}
ctr129:
// line 37 "smtp.ragel"
	{
emit(ActionMAIL, current)}
	goto st154
st154:
	if p++; p == pe {
		goto _test_eof154
	}
st_case_154:
// line 2703 "smtp.go"
	if data[p] == 58 {
		goto st101
	}
	if data[p] <= 57 {
		goto st100
	}
{
	goto st100

}
ctr130:
// line 37 "smtp.ragel"
	{
emit(ActionMAIL, current)}
	goto st106
st106:
	if p++; p == pe {
		goto _test_eof106
	}
st_case_106:
// line 2724 "smtp.go"
	switch data[p] {
	case 10:
goto st154
	case 58:
goto st101
	}
	switch {
	case data[p] > 9:
		if 11 <= data[p] && data[p] <= 57 {
			goto st100
		}
	default:
		goto st100
	}
{
	goto st100

}
st107:
	if p++; p == pe {
		goto _test_eof107
	}
st_case_107:
	switch data[p] {
	case 34:
goto st102
	case 43:
goto st104
	case 46:
goto st104
	case 47:
goto st100
	case 58:
goto st101
	case 62:
goto ctr123
	case 63:
goto st100
	case 64:
goto st107
	case 91:
goto st100
	case 92:
goto st110
	}
	switch {
	case data[p] < 48:
		switch {
		case data[p] < 35:
			if data[p] <= 33 {
				goto st100
			}
		case data[p] > 42:
			if 44 <= data[p] && data[p] <= 45 {
				goto st100
			}
		default:
			goto st100
		}
	case data[p] > 57:
		switch {
		case data[p] < 65:
			if 59 <= data[p] && data[p] <= 61 {
				goto st100
			}
		case data[p] > 90:
			switch {
			case data[p] > 96:
				if data[p] <= 122 {
					goto ctr132
				}
			case data[p] >= 93:
				goto st100
			}
		default:
			goto ctr132
		}
	default:
		goto ctr132
	}
{
	goto st100

}
ctr132:
// line 15 "smtp.ragel"
	{
pdomain = p}
	goto st108
st108:
	if p++; p == pe {
		goto _test_eof108
	}
st_case_108:
// line 2819 "smtp.go"
	switch data[p] {
	case 34:
goto st102
	case 43:
goto st104
	case 46:
goto st109
	case 47:
goto st100
	case 58:
goto st101
	case 62:
goto ctr135
	case 63:
goto st100
	case 64:
goto st107
	case 91:
goto st100
	case 92:
goto st110
	}
	switch {
	case data[p] < 48:
		switch {
		case data[p] < 35:
			if data[p] <= 33 {
				goto st100
			}
		case data[p] > 42:
			if 44 <= data[p] && data[p] <= 45 {
				goto st100
			}
		default:
			goto st100
		}
	case data[p] > 57:
		switch {
		case data[p] < 65:
			if 59 <= data[p] && data[p] <= 61 {
				goto st100
			}
		case data[p] > 90:
			switch {
			case data[p] > 96:
				if data[p] <= 122 {
					goto st108
				}
			case data[p] >= 93:
				goto st100
			}
		default:
			goto st108
		}
	default:
		goto st108
	}
{
	goto st100

}
st109:
	if p++; p == pe {
		goto _test_eof109
	}
st_case_109:
	switch data[p] {
	case 34:
goto st102
	case 43:
goto st104
	case 46:
goto st104
	case 47:
goto st100
	case 58:
goto st101
	case 62:
goto ctr123
	case 63:
goto st100
	case 64:
goto st107
	case 91:
goto st100
	case 92:
goto st110
	}
	switch {
	case data[p] < 48:
		switch {
		case data[p] < 35:
			if data[p] <= 33 {
				goto st100
			}
		case data[p] > 42:
			if 44 <= data[p] && data[p] <= 45 {
				goto st100
			}
		default:
			goto st100
		}
	case data[p] > 57:
		switch {
		case data[p] < 65:
			if 59 <= data[p] && data[p] <= 61 {
				goto st100
			}
		case data[p] > 90:
			switch {
			case data[p] > 96:
				if data[p] <= 122 {
					goto st108
				}
			case data[p] >= 93:
				goto st100
			}
		default:
			goto st108
		}
	default:
		goto st108
	}
{
	goto st100

}
st110:
	if p++; p == pe {
		goto _test_eof110
	}
st_case_110:
	switch data[p] {
	case 32:
goto st104
	case 33:
goto st100
	case 34:
goto st104
	case 43:
goto st104
	case 46:
goto st104
	case 47:
goto st100
	case 58:
goto st101
	}
	switch {
	case data[p] < 48:
		switch {
		case data[p] < 35:
			if data[p] <= 31 {
				goto st100
			}
		case data[p] > 42:
			if 44 <= data[p] && data[p] <= 45 {
				goto st100
			}
		default:
			goto st100
		}
	case data[p] > 57:
		switch {
		case data[p] < 64:
			if 59 <= data[p] {
				goto st100
			}
		case data[p] > 90:
			switch {
			case data[p] > 96:
				if data[p] <= 122 {
					goto st104
				}
			default:
				goto st100
			}
		default:
			goto st104
		}
	default:
		goto st104
	}
{
	goto st100

}
st111:
	if p++; p == pe {
		goto _test_eof111
	}
st_case_111:
	switch data[p] {
	case 32:
goto st103
	case 33:
goto st100
	case 34:
goto st103
	case 43:
goto st103
	case 46:
goto st103
	case 47:
goto st100
	case 58:
goto st101
	}
	switch {
	case data[p] < 48:
		switch {
		case data[p] < 35:
			if data[p] <= 31 {
				goto st100
			}
		case data[p] > 42:
			if 44 <= data[p] && data[p] <= 45 {
				goto st100
			}
		default:
			goto st100
		}
	case data[p] > 57:
		switch {
		case data[p] < 64:
			if 59 <= data[p] {
				goto st100
			}
		case data[p] > 90:
			switch {
			case data[p] > 96:
				if data[p] <= 122 {
					goto st103
				}
			default:
				goto st100
			}
		default:
			goto st103
		}
	default:
		goto st103
	}
{
	goto st100

}
st112:
	if p++; p == pe {
		goto _test_eof112
	}
st_case_112:
	if data[p] == 85 {
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
	if data[p] == 73 {
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
	if data[p] == 84 {
		goto st5
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
	case 67:
goto st116
	case 83:
goto st146
	}
{
	goto st0

}
st116:
	if p++; p == pe {
		goto _test_eof116
	}
st_case_116:
	if data[p] == 80 {
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
	if data[p] == 84 {
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
	if data[p] == 32 {
		goto st119
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
	case 84:
goto st120
	case 116:
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
	case 79:
goto st121
	case 111:
goto st121
	}
{
	goto st0

}
st121:
	if p++; p == pe {
		goto _test_eof121
	}
st_case_121:
	if data[p] == 58 {
		goto st122
	}
{
	goto st0

}
st122:
	if p++; p == pe {
		goto _test_eof122
	}
st_case_122:
	if data[p] == 60 {
		goto st123
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
	case 34:
goto ctr147
	case 43:
goto ctr148
	case 46:
goto ctr148
	case 64:
goto ctr149
	case 92:
goto ctr150
	}
	switch {
	case data[p] < 65:
		if 48 <= data[p] && data[p] <= 57 {
			goto ctr148
		}
	case data[p] > 90:
		if 97 <= data[p] && data[p] <= 122 {
			goto ctr148
		}
	default:
		goto ctr148
	}
{
	goto st0

}
ctr147:
// line 32 "smtp.ragel"
	{
pb = p}
	goto st124
st124:
	if p++; p == pe {
		goto _test_eof124
	}
st_case_124:
// line 3251 "smtp.go"
	switch data[p] {
	case 43:
goto st125
	case 46:
goto st125
	case 92:
goto st132
	}
	switch {
	case data[p] < 64:
		if 48 <= data[p] && data[p] <= 57 {
			goto st125
		}
	case data[p] > 90:
		if 97 <= data[p] && data[p] <= 122 {
			goto st125
		}
	default:
		goto st125
	}
{
	goto st0

}
st125:
	if p++; p == pe {
		goto _test_eof125
	}
st_case_125:
	switch data[p] {
	case 34:
goto st126
	case 43:
goto st125
	case 46:
goto st125
	case 92:
goto st132
	}
	switch {
	case data[p] < 64:
		if 48 <= data[p] && data[p] <= 57 {
			goto st125
		}
	case data[p] > 90:
		if 97 <= data[p] && data[p] <= 122 {
			goto st125
		}
	default:
		goto st125
	}
{
	goto st0

}
ctr148:
// line 32 "smtp.ragel"
	{
pb = p}
	goto st126
st126:
	if p++; p == pe {
		goto _test_eof126
	}
st_case_126:
// line 3317 "smtp.go"
	switch data[p] {
	case 34:
goto st124
	case 43:
goto st126
	case 46:
goto st126
	case 62:
goto ctr155
	case 64:
goto st128
	case 92:
goto st131
	}
	switch {
	case data[p] < 65:
		if 48 <= data[p] && data[p] <= 57 {
			goto st126
		}
	case data[p] > 90:
		if 97 <= data[p] && data[p] <= 122 {
			goto st126
		}
	default:
		goto st126
	}
{
	goto st0

}
ctr155:
// line 31 "smtp.ragel"
	{
current.Address = data[pb:p]}
	goto st127
ctr163:
// line 14 "smtp.ragel"
	{
current.Domain = data[pdomain:p]}
// line 31 "smtp.ragel"
	{
current.Address = data[pb:p]}
	goto st127
st127:
	if p++; p == pe {
		goto _test_eof127
	}
st_case_127:
// line 3366 "smtp.go"
	switch data[p] {
	case 10:
goto ctr158
	case 13:
goto ctr159
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
	case 34:
goto st124
	case 43:
goto st126
	case 46:
goto st126
	case 62:
goto ctr155
	case 64:
goto st128
	case 92:
goto st131
	}
	switch {
	case data[p] < 65:
		if 48 <= data[p] && data[p] <= 57 {
			goto ctr160
		}
	case data[p] > 90:
		if 97 <= data[p] && data[p] <= 122 {
			goto ctr160
		}
	default:
		goto ctr160
	}
{
	goto st0

}
ctr160:
// line 15 "smtp.ragel"
	{
pdomain = p}
	goto st129
st129:
	if p++; p == pe {
		goto _test_eof129
	}
st_case_129:
// line 3422 "smtp.go"
	switch data[p] {
	case 34:
goto st124
	case 43:
goto st126
	case 46:
goto st130
	case 62:
goto ctr163
	case 64:
goto st128
	case 92:
goto st131
	}
	switch {
	case data[p] < 65:
		if 48 <= data[p] && data[p] <= 57 {
			goto st129
		}
	case data[p] > 90:
		if 97 <= data[p] && data[p] <= 122 {
			goto st129
		}
	default:
		goto st129
	}
{
	goto st0

}
st130:
	if p++; p == pe {
		goto _test_eof130
	}
st_case_130:
	switch data[p] {
	case 34:
goto st124
	case 43:
goto st126
	case 46:
goto st126
	case 62:
goto ctr155
	case 64:
goto st128
	case 92:
goto st131
	}
	switch {
	case data[p] < 65:
		if 48 <= data[p] && data[p] <= 57 {
			goto st129
		}
	case data[p] > 90:
		if 97 <= data[p] && data[p] <= 122 {
			goto st129
		}
	default:
		goto st129
	}
{
	goto st0

}
ctr150:
// line 32 "smtp.ragel"
	{
pb = p}
	goto st131
st131:
	if p++; p == pe {
		goto _test_eof131
	}
st_case_131:
// line 3498 "smtp.go"
	switch data[p] {
	case 32:
goto st126
	case 34:
goto st126
	case 43:
goto st126
	case 46:
goto st126
	}
	switch {
	case data[p] < 64:
		if 48 <= data[p] && data[p] <= 57 {
			goto st126
		}
	case data[p] > 90:
		if 97 <= data[p] && data[p] <= 122 {
			goto st126
		}
	default:
		goto st126
	}
{
	goto st0

}
st132:
	if p++; p == pe {
		goto _test_eof132
	}
st_case_132:
	switch data[p] {
	case 32:
goto st125
	case 34:
goto st125
	case 43:
goto st125
	case 46:
goto st125
	}
	switch {
	case data[p] < 64:
		if 48 <= data[p] && data[p] <= 57 {
			goto st125
		}
	case data[p] > 90:
		if 97 <= data[p] && data[p] <= 122 {
			goto st125
		}
	default:
		goto st125
	}
{
	goto st0

}
ctr149:
// line 32 "smtp.ragel"
	{
pb = p}
	goto st133
st133:
	if p++; p == pe {
		goto _test_eof133
	}
st_case_133:
// line 3566 "smtp.go"
	switch data[p] {
	case 34:
goto st136
	case 43:
goto st138
	case 46:
goto st138
	case 47:
goto st134
	case 58:
goto st0
	case 62:
goto ctr167
	case 63:
goto st134
	case 64:
goto st141
	case 91:
goto st134
	case 92:
goto st144
	}
	switch {
	case data[p] < 48:
		switch {
		case data[p] < 35:
			if data[p] <= 33 {
				goto st134
			}
		case data[p] > 42:
			if 44 <= data[p] && data[p] <= 45 {
				goto st134
			}
		default:
			goto st134
		}
	case data[p] > 57:
		switch {
		case data[p] < 65:
			if 59 <= data[p] && data[p] <= 61 {
				goto st134
			}
		case data[p] > 90:
			switch {
			case data[p] > 96:
				if data[p] <= 122 {
					goto st138
				}
			case data[p] >= 93:
				goto st134
			}
		default:
			goto st138
		}
	default:
		goto st138
	}
{
	goto st134

}
st134:
	if p++; p == pe {
		goto _test_eof134
	}
st_case_134:
	if data[p] == 58 {
		goto st135
	}
	if data[p] <= 57 {
		goto st134
	}
{
	goto st134

}
st135:
	if p++; p == pe {
		goto _test_eof135
	}
st_case_135:
	switch data[p] {
	case 34:
goto ctr147
	case 43:
goto ctr148
	case 46:
goto ctr148
	case 92:
goto ctr150
	}
	switch {
	case data[p] < 64:
		if 48 <= data[p] && data[p] <= 57 {
			goto ctr148
		}
	case data[p] > 90:
		if 97 <= data[p] && data[p] <= 122 {
			goto ctr148
		}
	default:
		goto ctr148
	}
{
	goto st0

}
st136:
	if p++; p == pe {
		goto _test_eof136
	}
st_case_136:
	switch data[p] {
	case 43:
goto st137
	case 46:
goto st137
	case 47:
goto st134
	case 58:
goto st135
	case 91:
goto st134
	case 92:
goto st145
	}
	switch {
	case data[p] < 59:
		switch {
		case data[p] < 44:
			if data[p] <= 42 {
				goto st134
			}
		case data[p] > 45:
			if 48 <= data[p] && data[p] <= 57 {
				goto st137
			}
		default:
			goto st134
		}
	case data[p] > 63:
		switch {
		case data[p] < 93:
			if data[p] <= 90 {
				goto st137
			}
		case data[p] > 96:
			if data[p] <= 122 {
				goto st137
			}
		default:
			goto st134
		}
	default:
		goto st134
	}
{
	goto st134

}
st137:
	if p++; p == pe {
		goto _test_eof137
	}
st_case_137:
	switch data[p] {
	case 34:
goto st138
	case 43:
goto st137
	case 46:
goto st137
	case 47:
goto st134
	case 58:
goto st135
	case 91:
goto st134
	case 92:
goto st145
	}
	switch {
	case data[p] < 48:
		switch {
		case data[p] < 35:
			if data[p] <= 33 {
				goto st134
			}
		case data[p] > 42:
			if 44 <= data[p] && data[p] <= 45 {
				goto st134
			}
		default:
			goto st134
		}
	case data[p] > 57:
		switch {
		case data[p] < 64:
			if 59 <= data[p] {
				goto st134
			}
		case data[p] > 90:
			switch {
			case data[p] > 96:
				if data[p] <= 122 {
					goto st137
				}
			case data[p] >= 93:
				goto st134
			}
		default:
			goto st137
		}
	default:
		goto st137
	}
{
	goto st134

}
st138:
	if p++; p == pe {
		goto _test_eof138
	}
st_case_138:
	switch data[p] {
	case 34:
goto st136
	case 43:
goto st138
	case 46:
goto st138
	case 47:
goto st134
	case 58:
goto st135
	case 62:
goto ctr167
	case 63:
goto st134
	case 64:
goto st141
	case 91:
goto st134
	case 92:
goto st144
	}
	switch {
	case data[p] < 48:
		switch {
		case data[p] < 35:
			if data[p] <= 33 {
				goto st134
			}
		case data[p] > 42:
			if 44 <= data[p] && data[p] <= 45 {
				goto st134
			}
		default:
			goto st134
		}
	case data[p] > 57:
		switch {
		case data[p] < 65:
			if 59 <= data[p] && data[p] <= 61 {
				goto st134
			}
		case data[p] > 90:
			switch {
			case data[p] > 96:
				if data[p] <= 122 {
					goto st138
				}
			case data[p] >= 93:
				goto st134
			}
		default:
			goto st138
		}
	default:
		goto st138
	}
{
	goto st134

}
ctr167:
// line 31 "smtp.ragel"
	{
current.Address = data[pb:p]}
	goto st139
ctr179:
// line 14 "smtp.ragel"
	{
current.Domain = data[pdomain:p]}
// line 31 "smtp.ragel"
	{
current.Address = data[pb:p]}
	goto st139
st139:
	if p++; p == pe {
		goto _test_eof139
	}
st_case_139:
// line 3871 "smtp.go"
	switch data[p] {
	case 10:
goto ctr173
	case 13:
goto ctr174
	case 58:
goto st135
	}
	switch {
	case data[p] < 11:
		if data[p] <= 9 {
			goto st134
		}
	case data[p] > 12:
		if 14 <= data[p] && data[p] <= 57 {
			goto st134
		}
	default:
		goto st134
	}
{
	goto st134

}
ctr173:
// line 38 "smtp.ragel"
	{
emit(ActionRCPT, current)}
	goto st155
st155:
	if p++; p == pe {
		goto _test_eof155
	}
st_case_155:
// line 3906 "smtp.go"
	if data[p] == 58 {
		goto st135
	}
	if data[p] <= 57 {
		goto st134
	}
{
	goto st134

}
ctr174:
// line 38 "smtp.ragel"
	{
emit(ActionRCPT, current)}
	goto st140
st140:
	if p++; p == pe {
		goto _test_eof140
	}
st_case_140:
// line 3927 "smtp.go"
	switch data[p] {
	case 10:
goto st155
	case 58:
goto st135
	}
	switch {
	case data[p] > 9:
		if 11 <= data[p] && data[p] <= 57 {
			goto st134
		}
	default:
		goto st134
	}
{
	goto st134

}
st141:
	if p++; p == pe {
		goto _test_eof141
	}
st_case_141:
	switch data[p] {
	case 34:
goto st136
	case 43:
goto st138
	case 46:
goto st138
	case 47:
goto st134
	case 58:
goto st135
	case 62:
goto ctr167
	case 63:
goto st134
	case 64:
goto st141
	case 91:
goto st134
	case 92:
goto st144
	}
	switch {
	case data[p] < 48:
		switch {
		case data[p] < 35:
			if data[p] <= 33 {
				goto st134
			}
		case data[p] > 42:
			if 44 <= data[p] && data[p] <= 45 {
				goto st134
			}
		default:
			goto st134
		}
	case data[p] > 57:
		switch {
		case data[p] < 65:
			if 59 <= data[p] && data[p] <= 61 {
				goto st134
			}
		case data[p] > 90:
			switch {
			case data[p] > 96:
				if data[p] <= 122 {
					goto ctr176
				}
			case data[p] >= 93:
				goto st134
			}
		default:
			goto ctr176
		}
	default:
		goto ctr176
	}
{
	goto st134

}
ctr176:
// line 15 "smtp.ragel"
	{
pdomain = p}
	goto st142
st142:
	if p++; p == pe {
		goto _test_eof142
	}
st_case_142:
// line 4022 "smtp.go"
	switch data[p] {
	case 34:
goto st136
	case 43:
goto st138
	case 46:
goto st143
	case 47:
goto st134
	case 58:
goto st135
	case 62:
goto ctr179
	case 63:
goto st134
	case 64:
goto st141
	case 91:
goto st134
	case 92:
goto st144
	}
	switch {
	case data[p] < 48:
		switch {
		case data[p] < 35:
			if data[p] <= 33 {
				goto st134
			}
		case data[p] > 42:
			if 44 <= data[p] && data[p] <= 45 {
				goto st134
			}
		default:
			goto st134
		}
	case data[p] > 57:
		switch {
		case data[p] < 65:
			if 59 <= data[p] && data[p] <= 61 {
				goto st134
			}
		case data[p] > 90:
			switch {
			case data[p] > 96:
				if data[p] <= 122 {
					goto st142
				}
			case data[p] >= 93:
				goto st134
			}
		default:
			goto st142
		}
	default:
		goto st142
	}
{
	goto st134

}
st143:
	if p++; p == pe {
		goto _test_eof143
	}
st_case_143:
	switch data[p] {
	case 34:
goto st136
	case 43:
goto st138
	case 46:
goto st138
	case 47:
goto st134
	case 58:
goto st135
	case 62:
goto ctr167
	case 63:
goto st134
	case 64:
goto st141
	case 91:
goto st134
	case 92:
goto st144
	}
	switch {
	case data[p] < 48:
		switch {
		case data[p] < 35:
			if data[p] <= 33 {
				goto st134
			}
		case data[p] > 42:
			if 44 <= data[p] && data[p] <= 45 {
				goto st134
			}
		default:
			goto st134
		}
	case data[p] > 57:
		switch {
		case data[p] < 65:
			if 59 <= data[p] && data[p] <= 61 {
				goto st134
			}
		case data[p] > 90:
			switch {
			case data[p] > 96:
				if data[p] <= 122 {
					goto st142
				}
			case data[p] >= 93:
				goto st134
			}
		default:
			goto st142
		}
	default:
		goto st142
	}
{
	goto st134

}
st144:
	if p++; p == pe {
		goto _test_eof144
	}
st_case_144:
	switch data[p] {
	case 32:
goto st138
	case 33:
goto st134
	case 34:
goto st138
	case 43:
goto st138
	case 46:
goto st138
	case 47:
goto st134
	case 58:
goto st135
	}
	switch {
	case data[p] < 48:
		switch {
		case data[p] < 35:
			if data[p] <= 31 {
				goto st134
			}
		case data[p] > 42:
			if 44 <= data[p] && data[p] <= 45 {
				goto st134
			}
		default:
			goto st134
		}
	case data[p] > 57:
		switch {
		case data[p] < 64:
			if 59 <= data[p] {
				goto st134
			}
		case data[p] > 90:
			switch {
			case data[p] > 96:
				if data[p] <= 122 {
					goto st138
				}
			default:
				goto st134
			}
		default:
			goto st138
		}
	default:
		goto st138
	}
{
	goto st134

}
st145:
	if p++; p == pe {
		goto _test_eof145
	}
st_case_145:
	switch data[p] {
	case 32:
goto st137
	case 33:
goto st134
	case 34:
goto st137
	case 43:
goto st137
	case 46:
goto st137
	case 47:
goto st134
	case 58:
goto st135
	}
	switch {
	case data[p] < 48:
		switch {
		case data[p] < 35:
			if data[p] <= 31 {
				goto st134
			}
		case data[p] > 42:
			if 44 <= data[p] && data[p] <= 45 {
				goto st134
			}
		default:
			goto st134
		}
	case data[p] > 57:
		switch {
		case data[p] < 64:
			if 59 <= data[p] {
				goto st134
			}
		case data[p] > 90:
			switch {
			case data[p] > 96:
				if data[p] <= 122 {
					goto st137
				}
			default:
				goto st134
			}
		default:
			goto st137
		}
	default:
		goto st137
	}
{
	goto st134

}
st146:
	if p++; p == pe {
		goto _test_eof146
	}
st_case_146:
	if data[p] == 69 {
		goto st114
	}
{
	goto st0

}
st147:
	if p++; p == pe {
		goto _test_eof147
	}
st_case_147:
	if data[p] == 82 {
		goto st148
	}
{
	goto st0

}
st148:
	if p++; p == pe {
		goto _test_eof148
	}
st_case_148:
	if data[p] == 70 {
		goto st149
	}
{
	goto st0

}
st149:
	if p++; p == pe {
		goto _test_eof149
	}
st_case_149:
	if data[p] == 89 {
		goto st150
	}
{
	goto st0

}
st150:
	if p++; p == pe {
		goto _test_eof150
	}
st_case_150:
	if data[p] == 32 {
		goto st151
	}
{
	goto st0

}
st151:
	if p++; p == pe {
		goto _test_eof151
	}
st_case_151:
	switch data[p] {
	case 10:
goto st0
	case 13:
goto st0
	}
	switch {
	case data[p] > 9:
		if 11 <= data[p] && data[p] <= 12 {
			goto st152
		}
	default:
		goto st152
	}
{
	goto st152

}
st152:
	if p++; p == pe {
		goto _test_eof152
	}
st_case_152:
	switch data[p] {
	case 10:
goto st153
	case 13:
goto st6
	}
	switch {
	case data[p] > 9:
		if 11 <= data[p] && data[p] <= 12 {
			goto st152
		}
	default:
		goto st152
	}
{
	goto st152

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
	_test_eof153: cs = 153
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
	_test_eof154: cs = 154
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
	_test_eof155: cs = 155
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

	_test_eof: {}
	_out: {}
	}

// line 61 "smtp.ragel"

	if cs == smtp_error {
		return fmt.Errorf("Invalid character in pos %d: `%c`.", p, data[p])
	}
	return nil
}
