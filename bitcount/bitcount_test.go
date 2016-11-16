package bitcount_test

import (
	"testing"

	. "."
)

func init() {}

func TestTableInitAndSprintf(t *testing.T) {
	table = InitializeBitsPerBytesLookupTable()
	correctTable := "   0:0   1:1   2:1   3:2   4:1   5:2   6:2   7:3\n   8:1   9:2  10:2  11:3  12:2  13:3  14:3  15:4\n  16:1  17:2  18:2  19:3  20:2  21:3  22:3  23:4\n  24:2  25:3  26:3  27:4  28:3  29:4  30:4  31:5\n  32:1  33:2  34:2  35:3  36:2  37:3  38:3  39:4\n  40:2  41:3  42:3  43:4  44:3  45:4  46:4  47:5\n  48:2  49:3  50:3  51:4  52:3  53:4  54:4  55:5\n  56:3  57:4  58:4  59:5  60:4  61:5  62:5  63:6\n  64:1  65:2  66:2  67:3  68:2  69:3  70:3  71:4\n  72:2  73:3  74:3  75:4  76:3  77:4  78:4  79:5\n  80:2  81:3  82:3  83:4  84:3  85:4  86:4  87:5\n  88:3  89:4  90:4  91:5  92:4  93:5  94:5  95:6\n  96:2  97:3  98:3  99:4 100:3 101:4 102:4 103:5\n 104:3 105:4 106:4 107:5 108:4 109:5 110:5 111:6\n 112:3 113:4 114:4 115:5 116:4 117:5 118:5 119:6\n 120:4 121:5 122:5 123:6 124:5 125:6 126:6 127:7\n 128:1 129:2 130:2 131:3 132:2 133:3 134:3 135:4\n 136:2 137:3 138:3 139:4 140:3 141:4 142:4 143:5\n 144:2 145:3 146:3 147:4 148:3 149:4 150:4 151:5\n 152:3 153:4 154:4 155:5 156:4 157:5 158:5 159:6\n 160:2 161:3 162:3 163:4 164:3 165:4 166:4 167:5\n 168:3 169:4 170:4 171:5 172:4 173:5 174:5 175:6\n 176:3 177:4 178:4 179:5 180:4 181:5 182:5 183:6\n 184:4 185:5 186:5 187:6 188:5 189:6 190:6 191:7\n 192:2 193:3 194:3 195:4 196:3 197:4 198:4 199:5\n 200:3 201:4 202:4 203:5 204:4 205:5 206:5 207:6\n 208:3 209:4 210:4 211:5 212:4 213:5 214:5 215:6\n 216:4 217:5 218:5 219:6 220:5 221:6 222:6 223:7\n 224:3 225:4 226:4 227:5 228:4 229:5 230:5 231:6\n 232:4 233:5 234:5 235:6 236:5 237:6 238:6 239:7\n 240:4 241:5 242:5 243:6 244:5 245:6 246:6 247:7\n 248:5 249:6 250:6 251:7 252:6 253:7 254:7 255:8"
	if s := SprintfLookupTable(table); s != correctTable {
		t.Fatalf("generated table not correct: \n%q\n", s)
	}
}

var (
	table     []uint8
	testCases = []*struct {
		seed int64
		n    uint
		bits uint
		data []uint16
	}{
		{seed: 1, n: 1, bits: 3, data: []uint16{1090}},
		{seed: 1, n: 10, bits: 67, data: []uint16{1090, 13342, 39822, 2422, 2818, 41740, 38514, 42328, 65168, 36680}},
		{seed: 1, n: 100000000, bits: 799986009},
	}
)

func TestGenerateDataLength(t *testing.T) {
	for i, tc := range testCases {
		data := GenerateData(tc.seed, tc.n)
		if tc.data != nil {
			if len(data) != len(tc.data) {
				t.Fatalf("%d: incorrect data: %v, expected: %v", i, data, tc.data)
			}
			for j, x := range data {
				expected := tc.data[j]
				if expected != x {
					t.Fatalf("%d: incorrect value at %d: %d, expected: %d", i, j, x, expected)
				}
			}
		}

		if res := len(data); uint(res) != tc.n {
			t.Fatalf("%d: incorrect result: %d, expected: %d", i, res, tc.n)
		}
		tc.data = data
	}
}

func TestCountBitsLookupTable(t *testing.T) {
	for i, tc := range testCases {
		if res := CountBitsLookupTable(table, tc.data); uint(res) != tc.bits {
			t.Fatalf("%d: incorrect result: %d, expected: %d", i, res, tc.n)
		}
	}
}

func TestCountBitsTrivial(t *testing.T) {
	for i, tc := range testCases {
		if res := CountBitsTrivial(tc.data); uint(res) != tc.bits {
			t.Fatalf("%d: incorrect result: %d, expected: %d", i, res, tc.n)
		}
	}
}
