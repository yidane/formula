// Generated from Formula.g4 by ANTLR 4.7.

package parser // Formula

import (
	"fmt"
	"github.com/yidane/formula/internal/exp"
	"github.com/yidane/formula/opt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 42, 297,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 5, 3, 49, 10, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 7, 4, 60, 10, 4, 12, 4, 14, 4, 63, 11, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3,
	5, 3, 5, 3, 5, 3, 5, 3, 5, 7, 5, 74, 10, 5, 12, 5, 14, 5, 77, 11, 5, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 7, 6, 88, 10, 6, 12,
	6, 14, 6, 91, 11, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3,
	7, 7, 7, 102, 10, 7, 12, 7, 14, 7, 105, 11, 7, 3, 8, 3, 8, 3, 8, 3, 8,
	3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 7, 8, 116, 10, 8, 12, 8, 14, 8, 119, 11,
	8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3,
	9, 3, 9, 3, 9, 7, 9, 135, 10, 9, 12, 9, 14, 9, 138, 11, 9, 3, 10, 3, 10,
	3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3,
	10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10,
	3, 10, 7, 10, 164, 10, 10, 12, 10, 14, 10, 167, 11, 10, 3, 11, 3, 11, 3,
	11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11,
	3, 11, 7, 11, 183, 10, 11, 12, 11, 14, 11, 186, 11, 11, 3, 12, 3, 12, 3,
	12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12,
	3, 12, 7, 12, 202, 10, 12, 12, 12, 14, 12, 205, 11, 12, 3, 13, 3, 13, 3,
	13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13,
	3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 7, 13, 226, 10, 13, 12, 13, 14,
	13, 229, 11, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14,
	3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 5, 14, 246, 10, 14, 3,
	15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15,
	3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 7, 15, 264, 10, 15, 12, 15, 14, 15,
	267, 11, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 5, 15, 275, 10,
	15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16,
	3, 16, 3, 16, 5, 16, 289, 10, 16, 3, 17, 3, 17, 3, 17, 3, 17, 5, 17, 295,
	10, 17, 3, 17, 2, 12, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 18, 2, 4, 6,
	8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 2, 7, 3, 2, 5, 6, 3,
	2, 7, 8, 3, 2, 12, 13, 3, 2, 14, 15, 3, 2, 27, 28, 2, 312, 2, 34, 3, 2,
	2, 2, 4, 48, 3, 2, 2, 2, 6, 50, 3, 2, 2, 2, 8, 64, 3, 2, 2, 2, 10, 78,
	3, 2, 2, 2, 12, 92, 3, 2, 2, 2, 14, 106, 3, 2, 2, 2, 16, 120, 3, 2, 2,
	2, 18, 139, 3, 2, 2, 2, 20, 168, 3, 2, 2, 2, 22, 187, 3, 2, 2, 2, 24, 206,
	3, 2, 2, 2, 26, 245, 3, 2, 2, 2, 28, 274, 3, 2, 2, 2, 30, 288, 3, 2, 2,
	2, 32, 294, 3, 2, 2, 2, 34, 35, 5, 4, 3, 2, 35, 36, 7, 2, 2, 3, 36, 37,
	8, 2, 1, 2, 37, 3, 3, 2, 2, 2, 38, 39, 5, 6, 4, 2, 39, 40, 7, 3, 2, 2,
	40, 41, 5, 4, 3, 2, 41, 42, 7, 4, 2, 2, 42, 43, 5, 4, 3, 2, 43, 44, 8,
	3, 1, 2, 44, 49, 3, 2, 2, 2, 45, 46, 5, 6, 4, 2, 46, 47, 8, 3, 1, 2, 47,
	49, 3, 2, 2, 2, 48, 38, 3, 2, 2, 2, 48, 45, 3, 2, 2, 2, 49, 5, 3, 2, 2,
	2, 50, 51, 8, 4, 1, 2, 51, 52, 5, 8, 5, 2, 52, 53, 8, 4, 1, 2, 53, 61,
	3, 2, 2, 2, 54, 55, 12, 4, 2, 2, 55, 56, 9, 2, 2, 2, 56, 57, 5, 8, 5, 2,
	57, 58, 8, 4, 1, 2, 58, 60, 3, 2, 2, 2, 59, 54, 3, 2, 2, 2, 60, 63, 3,
	2, 2, 2, 61, 59, 3, 2, 2, 2, 61, 62, 3, 2, 2, 2, 62, 7, 3, 2, 2, 2, 63,
	61, 3, 2, 2, 2, 64, 65, 8, 5, 1, 2, 65, 66, 5, 10, 6, 2, 66, 67, 8, 5,
	1, 2, 67, 75, 3, 2, 2, 2, 68, 69, 12, 4, 2, 2, 69, 70, 9, 3, 2, 2, 70,
	71, 5, 10, 6, 2, 71, 72, 8, 5, 1, 2, 72, 74, 3, 2, 2, 2, 73, 68, 3, 2,
	2, 2, 74, 77, 3, 2, 2, 2, 75, 73, 3, 2, 2, 2, 75, 76, 3, 2, 2, 2, 76, 9,
	3, 2, 2, 2, 77, 75, 3, 2, 2, 2, 78, 79, 8, 6, 1, 2, 79, 80, 5, 12, 7, 2,
	80, 81, 8, 6, 1, 2, 81, 89, 3, 2, 2, 2, 82, 83, 12, 4, 2, 2, 83, 84, 7,
	9, 2, 2, 84, 85, 5, 12, 7, 2, 85, 86, 8, 6, 1, 2, 86, 88, 3, 2, 2, 2, 87,
	82, 3, 2, 2, 2, 88, 91, 3, 2, 2, 2, 89, 87, 3, 2, 2, 2, 89, 90, 3, 2, 2,
	2, 90, 11, 3, 2, 2, 2, 91, 89, 3, 2, 2, 2, 92, 93, 8, 7, 1, 2, 93, 94,
	5, 14, 8, 2, 94, 95, 8, 7, 1, 2, 95, 103, 3, 2, 2, 2, 96, 97, 12, 4, 2,
	2, 97, 98, 7, 10, 2, 2, 98, 99, 5, 14, 8, 2, 99, 100, 8, 7, 1, 2, 100,
	102, 3, 2, 2, 2, 101, 96, 3, 2, 2, 2, 102, 105, 3, 2, 2, 2, 103, 101, 3,
	2, 2, 2, 103, 104, 3, 2, 2, 2, 104, 13, 3, 2, 2, 2, 105, 103, 3, 2, 2,
	2, 106, 107, 8, 8, 1, 2, 107, 108, 5, 16, 9, 2, 108, 109, 8, 8, 1, 2, 109,
	117, 3, 2, 2, 2, 110, 111, 12, 4, 2, 2, 111, 112, 7, 11, 2, 2, 112, 113,
	5, 16, 9, 2, 113, 114, 8, 8, 1, 2, 114, 116, 3, 2, 2, 2, 115, 110, 3, 2,
	2, 2, 116, 119, 3, 2, 2, 2, 117, 115, 3, 2, 2, 2, 117, 118, 3, 2, 2, 2,
	118, 15, 3, 2, 2, 2, 119, 117, 3, 2, 2, 2, 120, 121, 8, 9, 1, 2, 121, 122,
	5, 18, 10, 2, 122, 123, 8, 9, 1, 2, 123, 136, 3, 2, 2, 2, 124, 125, 12,
	5, 2, 2, 125, 126, 9, 4, 2, 2, 126, 127, 5, 18, 10, 2, 127, 128, 8, 9,
	1, 2, 128, 135, 3, 2, 2, 2, 129, 130, 12, 4, 2, 2, 130, 131, 9, 5, 2, 2,
	131, 132, 5, 18, 10, 2, 132, 133, 8, 9, 1, 2, 133, 135, 3, 2, 2, 2, 134,
	124, 3, 2, 2, 2, 134, 129, 3, 2, 2, 2, 135, 138, 3, 2, 2, 2, 136, 134,
	3, 2, 2, 2, 136, 137, 3, 2, 2, 2, 137, 17, 3, 2, 2, 2, 138, 136, 3, 2,
	2, 2, 139, 140, 8, 10, 1, 2, 140, 141, 5, 20, 11, 2, 141, 142, 8, 10, 1,
	2, 142, 165, 3, 2, 2, 2, 143, 144, 12, 7, 2, 2, 144, 145, 7, 16, 2, 2,
	145, 146, 5, 20, 11, 2, 146, 147, 8, 10, 1, 2, 147, 164, 3, 2, 2, 2, 148,
	149, 12, 6, 2, 2, 149, 150, 7, 17, 2, 2, 150, 151, 5, 20, 11, 2, 151, 152,
	8, 10, 1, 2, 152, 164, 3, 2, 2, 2, 153, 154, 12, 5, 2, 2, 154, 155, 7,
	18, 2, 2, 155, 156, 5, 20, 11, 2, 156, 157, 8, 10, 1, 2, 157, 164, 3, 2,
	2, 2, 158, 159, 12, 4, 2, 2, 159, 160, 7, 19, 2, 2, 160, 161, 5, 20, 11,
	2, 161, 162, 8, 10, 1, 2, 162, 164, 3, 2, 2, 2, 163, 143, 3, 2, 2, 2, 163,
	148, 3, 2, 2, 2, 163, 153, 3, 2, 2, 2, 163, 158, 3, 2, 2, 2, 164, 167,
	3, 2, 2, 2, 165, 163, 3, 2, 2, 2, 165, 166, 3, 2, 2, 2, 166, 19, 3, 2,
	2, 2, 167, 165, 3, 2, 2, 2, 168, 169, 8, 11, 1, 2, 169, 170, 5, 22, 12,
	2, 170, 171, 8, 11, 1, 2, 171, 184, 3, 2, 2, 2, 172, 173, 12, 5, 2, 2,
	173, 174, 7, 20, 2, 2, 174, 175, 5, 22, 12, 2, 175, 176, 8, 11, 1, 2, 176,
	183, 3, 2, 2, 2, 177, 178, 12, 4, 2, 2, 178, 179, 7, 21, 2, 2, 179, 180,
	5, 22, 12, 2, 180, 181, 8, 11, 1, 2, 181, 183, 3, 2, 2, 2, 182, 172, 3,
	2, 2, 2, 182, 177, 3, 2, 2, 2, 183, 186, 3, 2, 2, 2, 184, 182, 3, 2, 2,
	2, 184, 185, 3, 2, 2, 2, 185, 21, 3, 2, 2, 2, 186, 184, 3, 2, 2, 2, 187,
	188, 8, 12, 1, 2, 188, 189, 5, 24, 13, 2, 189, 190, 8, 12, 1, 2, 190, 203,
	3, 2, 2, 2, 191, 192, 12, 5, 2, 2, 192, 193, 7, 22, 2, 2, 193, 194, 5,
	24, 13, 2, 194, 195, 8, 12, 1, 2, 195, 202, 3, 2, 2, 2, 196, 197, 12, 4,
	2, 2, 197, 198, 7, 23, 2, 2, 198, 199, 5, 24, 13, 2, 199, 200, 8, 12, 1,
	2, 200, 202, 3, 2, 2, 2, 201, 191, 3, 2, 2, 2, 201, 196, 3, 2, 2, 2, 202,
	205, 3, 2, 2, 2, 203, 201, 3, 2, 2, 2, 203, 204, 3, 2, 2, 2, 204, 23, 3,
	2, 2, 2, 205, 203, 3, 2, 2, 2, 206, 207, 8, 13, 1, 2, 207, 208, 5, 26,
	14, 2, 208, 209, 8, 13, 1, 2, 209, 227, 3, 2, 2, 2, 210, 211, 12, 6, 2,
	2, 211, 212, 7, 24, 2, 2, 212, 213, 5, 26, 14, 2, 213, 214, 8, 13, 1, 2,
	214, 226, 3, 2, 2, 2, 215, 216, 12, 5, 2, 2, 216, 217, 7, 25, 2, 2, 217,
	218, 5, 26, 14, 2, 218, 219, 8, 13, 1, 2, 219, 226, 3, 2, 2, 2, 220, 221,
	12, 4, 2, 2, 221, 222, 7, 26, 2, 2, 222, 223, 5, 26, 14, 2, 223, 224, 8,
	13, 1, 2, 224, 226, 3, 2, 2, 2, 225, 210, 3, 2, 2, 2, 225, 215, 3, 2, 2,
	2, 225, 220, 3, 2, 2, 2, 226, 229, 3, 2, 2, 2, 227, 225, 3, 2, 2, 2, 227,
	228, 3, 2, 2, 2, 228, 25, 3, 2, 2, 2, 229, 227, 3, 2, 2, 2, 230, 231, 9,
	6, 2, 2, 231, 232, 5, 28, 15, 2, 232, 233, 8, 14, 1, 2, 233, 246, 3, 2,
	2, 2, 234, 235, 7, 29, 2, 2, 235, 236, 5, 28, 15, 2, 236, 237, 8, 14, 1,
	2, 237, 246, 3, 2, 2, 2, 238, 239, 7, 23, 2, 2, 239, 240, 5, 28, 15, 2,
	240, 241, 8, 14, 1, 2, 241, 246, 3, 2, 2, 2, 242, 243, 5, 28, 15, 2, 243,
	244, 8, 14, 1, 2, 244, 246, 3, 2, 2, 2, 245, 230, 3, 2, 2, 2, 245, 234,
	3, 2, 2, 2, 245, 238, 3, 2, 2, 2, 245, 242, 3, 2, 2, 2, 246, 27, 3, 2,
	2, 2, 247, 248, 7, 30, 2, 2, 248, 249, 5, 4, 3, 2, 249, 250, 7, 31, 2,
	2, 250, 251, 8, 15, 1, 2, 251, 275, 3, 2, 2, 2, 252, 253, 5, 30, 16, 2,
	253, 254, 8, 15, 1, 2, 254, 275, 3, 2, 2, 2, 255, 256, 5, 32, 17, 2, 256,
	257, 7, 30, 2, 2, 257, 258, 5, 4, 3, 2, 258, 265, 8, 15, 1, 2, 259, 260,
	7, 32, 2, 2, 260, 261, 5, 4, 3, 2, 261, 262, 8, 15, 1, 2, 262, 264, 3,
	2, 2, 2, 263, 259, 3, 2, 2, 2, 264, 267, 3, 2, 2, 2, 265, 263, 3, 2, 2,
	2, 265, 266, 3, 2, 2, 2, 266, 268, 3, 2, 2, 2, 267, 265, 3, 2, 2, 2, 268,
	269, 7, 31, 2, 2, 269, 270, 8, 15, 1, 2, 270, 275, 3, 2, 2, 2, 271, 272,
	5, 32, 17, 2, 272, 273, 8, 15, 1, 2, 273, 275, 3, 2, 2, 2, 274, 247, 3,
	2, 2, 2, 274, 252, 3, 2, 2, 2, 274, 255, 3, 2, 2, 2, 274, 271, 3, 2, 2,
	2, 275, 29, 3, 2, 2, 2, 276, 277, 7, 36, 2, 2, 277, 289, 8, 16, 1, 2, 278,
	279, 7, 40, 2, 2, 279, 289, 8, 16, 1, 2, 280, 281, 7, 41, 2, 2, 281, 289,
	8, 16, 1, 2, 282, 283, 7, 37, 2, 2, 283, 289, 8, 16, 1, 2, 284, 285, 7,
	33, 2, 2, 285, 289, 8, 16, 1, 2, 286, 287, 7, 34, 2, 2, 287, 289, 8, 16,
	1, 2, 288, 276, 3, 2, 2, 2, 288, 278, 3, 2, 2, 2, 288, 280, 3, 2, 2, 2,
	288, 282, 3, 2, 2, 2, 288, 284, 3, 2, 2, 2, 288, 286, 3, 2, 2, 2, 289,
	31, 3, 2, 2, 2, 290, 291, 7, 35, 2, 2, 291, 295, 8, 17, 1, 2, 292, 293,
	7, 38, 2, 2, 293, 295, 8, 17, 1, 2, 294, 290, 3, 2, 2, 2, 294, 292, 3,
	2, 2, 2, 295, 33, 3, 2, 2, 2, 23, 48, 61, 75, 89, 103, 117, 134, 136, 163,
	165, 182, 184, 201, 203, 225, 227, 245, 265, 274, 288, 294,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'?'", "':'", "'||'", "'or'", "'&&'", "'and'", "'|'", "'^'", "'&'",
	"'=='", "'='", "'!='", "'<>'", "'<'", "'<='", "'>'", "'>='", "'<<'", "'>>'",
	"'+'", "'-'", "'*'", "'/'", "'%'", "'!'", "'not'", "'~'", "'('", "')'",
	"','", "'true'", "'false'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "TRUE", "FALSE", "NAME",
	"INTEGER", "DATETIME", "VAR", "E", "FLOAT", "STRING", "WS",
}

var ruleNames = []string{
	"ncalc", "expr", "orExpr", "andExpr", "bitOrExpr", "bitXorExpr", "bitAndExpr",
	"eqExpr", "relExpr", "shiftExpr", "addExpr", "multExpr", "unaryExpr", "primaryExpr",
	"value", "id",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type FormulaParser struct {
	*antlr.BaseParser
}

func NewFormulaParser(input antlr.TokenStream) *FormulaParser {
	this := new(FormulaParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Formula.g4"

	return this
}

// FormulaParser tokens.
const (
	FormulaParserEOF      = antlr.TokenEOF
	FormulaParserT__0     = 1
	FormulaParserT__1     = 2
	FormulaParserT__2     = 3
	FormulaParserT__3     = 4
	FormulaParserT__4     = 5
	FormulaParserT__5     = 6
	FormulaParserT__6     = 7
	FormulaParserT__7     = 8
	FormulaParserT__8     = 9
	FormulaParserT__9     = 10
	FormulaParserT__10    = 11
	FormulaParserT__11    = 12
	FormulaParserT__12    = 13
	FormulaParserT__13    = 14
	FormulaParserT__14    = 15
	FormulaParserT__15    = 16
	FormulaParserT__16    = 17
	FormulaParserT__17    = 18
	FormulaParserT__18    = 19
	FormulaParserT__19    = 20
	FormulaParserT__20    = 21
	FormulaParserT__21    = 22
	FormulaParserT__22    = 23
	FormulaParserT__23    = 24
	FormulaParserT__24    = 25
	FormulaParserT__25    = 26
	FormulaParserT__26    = 27
	FormulaParserT__27    = 28
	FormulaParserT__28    = 29
	FormulaParserT__29    = 30
	FormulaParserTRUE     = 31
	FormulaParserFALSE    = 32
	FormulaParserNAME     = 33
	FormulaParserINTEGER  = 34
	FormulaParserDATETIME = 35
	FormulaParserVAR      = 36
	FormulaParserE        = 37
	FormulaParserFLOAT    = 38
	FormulaParserSTRING   = 39
	FormulaParserWS       = 40
)

// FormulaParser rules.
const (
	FormulaParserRULE_ncalc       = 0
	FormulaParserRULE_expr        = 1
	FormulaParserRULE_orExpr      = 2
	FormulaParserRULE_andExpr     = 3
	FormulaParserRULE_bitOrExpr   = 4
	FormulaParserRULE_bitXorExpr  = 5
	FormulaParserRULE_bitAndExpr  = 6
	FormulaParserRULE_eqExpr      = 7
	FormulaParserRULE_relExpr     = 8
	FormulaParserRULE_shiftExpr   = 9
	FormulaParserRULE_addExpr     = 10
	FormulaParserRULE_multExpr    = 11
	FormulaParserRULE_unaryExpr   = 12
	FormulaParserRULE_primaryExpr = 13
	FormulaParserRULE_value       = 14
	FormulaParserRULE_id          = 15
)

// INcalcContext is an interface to support dynamic dispatch.
type INcalcContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// GetRetValue returns the retValue attribute.
	GetRetValue() *opt.LogicalExpression

	// SetRetValue sets the retValue attribute.
	SetRetValue(*opt.LogicalExpression)

	// IsNcalcContext differentiates from other interfaces.
	IsNcalcContext()
}

type NcalcContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	retValue *opt.LogicalExpression
	_expr    IExprContext
}

func NewEmptyNcalcContext() *NcalcContext {
	var p = new(NcalcContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FormulaParserRULE_ncalc
	return p
}

func (*NcalcContext) IsNcalcContext() {}

func NewNcalcContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NcalcContext {
	var p = new(NcalcContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_ncalc

	return p
}

func (s *NcalcContext) GetParser() antlr.Parser { return s.parser }

func (s *NcalcContext) Get_expr() IExprContext { return s._expr }

func (s *NcalcContext) Set_expr(v IExprContext) { s._expr = v }

func (s *NcalcContext) GetRetValue() *opt.LogicalExpression { return s.retValue }

func (s *NcalcContext) SetRetValue(v *opt.LogicalExpression) { s.retValue = v }

func (s *NcalcContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *NcalcContext) EOF() antlr.TerminalNode {
	return s.GetToken(FormulaParserEOF, 0)
}

func (s *NcalcContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NcalcContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NcalcContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaListener); ok {
		listenerT.EnterNcalc(s)
	}
}

func (s *NcalcContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaListener); ok {
		listenerT.ExitNcalc(s)
	}
}

func (p *FormulaParser) Ncalc() (localctx INcalcContext) {
	localctx = NewNcalcContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, FormulaParserRULE_ncalc)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(32)

		var _x = p.Expr()

		localctx.(*NcalcContext)._expr = _x
	}
	{
		p.SetState(33)
		p.Match(FormulaParserEOF)
	}
	localctx.(*NcalcContext).retValue = localctx.(*NcalcContext).Get_expr().GetRetValue()

	return localctx
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetFirst returns the first rule contexts.
	GetFirst() IOrExprContext

	// GetMiddle returns the middle rule contexts.
	GetMiddle() IExprContext

	// GetRight returns the right rule contexts.
	GetRight() IExprContext

	// Get_orExpr returns the _orExpr rule contexts.
	Get_orExpr() IOrExprContext

	// SetFirst sets the first rule contexts.
	SetFirst(IOrExprContext)

	// SetMiddle sets the middle rule contexts.
	SetMiddle(IExprContext)

	// SetRight sets the right rule contexts.
	SetRight(IExprContext)

	// Set_orExpr sets the _orExpr rule contexts.
	Set_orExpr(IOrExprContext)

	// GetRetValue returns the retValue attribute.
	GetRetValue() *opt.LogicalExpression

	// SetRetValue sets the retValue attribute.
	SetRetValue(*opt.LogicalExpression)

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	retValue *opt.LogicalExpression
	first    IOrExprContext
	middle   IExprContext
	right    IExprContext
	_orExpr  IOrExprContext
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FormulaParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) GetFirst() IOrExprContext { return s.first }

func (s *ExprContext) GetMiddle() IExprContext { return s.middle }

func (s *ExprContext) GetRight() IExprContext { return s.right }

func (s *ExprContext) Get_orExpr() IOrExprContext { return s._orExpr }

func (s *ExprContext) SetFirst(v IOrExprContext) { s.first = v }

func (s *ExprContext) SetMiddle(v IExprContext) { s.middle = v }

func (s *ExprContext) SetRight(v IExprContext) { s.right = v }

func (s *ExprContext) Set_orExpr(v IOrExprContext) { s._orExpr = v }

func (s *ExprContext) GetRetValue() *opt.LogicalExpression { return s.retValue }

func (s *ExprContext) SetRetValue(v *opt.LogicalExpression) { s.retValue = v }

func (s *ExprContext) OrExpr() IOrExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOrExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOrExprContext)
}

func (s *ExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *ExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaListener); ok {
		listenerT.EnterExpr(s)
	}
}

func (s *ExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaListener); ok {
		listenerT.ExitExpr(s)
	}
}

func (p *FormulaParser) Expr() (localctx IExprContext) {
	localctx = NewExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, FormulaParserRULE_expr)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(46)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(36)

			var _x = p.orExpr(0)

			localctx.(*ExprContext).first = _x
		}
		{
			p.SetState(37)
			p.Match(FormulaParserT__0)
		}
		{
			p.SetState(38)

			var _x = p.Expr()

			localctx.(*ExprContext).middle = _x
		}
		{
			p.SetState(39)
			p.Match(FormulaParserT__1)
		}
		{
			p.SetState(40)

			var _x = p.Expr()

			localctx.(*ExprContext).right = _x
		}
		localctx.(*ExprContext).retValue = exp.NewTernaryExpression(localctx.(*ExprContext).GetFirst().GetRetValue(), localctx.(*ExprContext).GetMiddle().GetRetValue(), localctx.(*ExprContext).GetRight().GetRetValue())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(43)

			var _x = p.orExpr(0)

			localctx.(*ExprContext)._orExpr = _x
		}
		localctx.(*ExprContext).retValue = localctx.(*ExprContext).Get_orExpr().GetRetValue()

	}

	return localctx
}

// IOrExprContext is an interface to support dynamic dispatch.
type IOrExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetFirst returns the first rule contexts.
	GetFirst() IOrExprContext

	// Get_andExpr returns the _andExpr rule contexts.
	Get_andExpr() IAndExprContext

	// SetFirst sets the first rule contexts.
	SetFirst(IOrExprContext)

	// Set_andExpr sets the _andExpr rule contexts.
	Set_andExpr(IAndExprContext)

	// GetRetValue returns the retValue attribute.
	GetRetValue() *opt.LogicalExpression

	// SetRetValue sets the retValue attribute.
	SetRetValue(*opt.LogicalExpression)

	// IsOrExprContext differentiates from other interfaces.
	IsOrExprContext()
}

type OrExprContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	retValue *opt.LogicalExpression
	first    IOrExprContext
	_andExpr IAndExprContext
}

func NewEmptyOrExprContext() *OrExprContext {
	var p = new(OrExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FormulaParserRULE_orExpr
	return p
}

func (*OrExprContext) IsOrExprContext() {}

func NewOrExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OrExprContext {
	var p = new(OrExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_orExpr

	return p
}

func (s *OrExprContext) GetParser() antlr.Parser { return s.parser }

func (s *OrExprContext) GetFirst() IOrExprContext { return s.first }

func (s *OrExprContext) Get_andExpr() IAndExprContext { return s._andExpr }

func (s *OrExprContext) SetFirst(v IOrExprContext) { s.first = v }

func (s *OrExprContext) Set_andExpr(v IAndExprContext) { s._andExpr = v }

func (s *OrExprContext) GetRetValue() *opt.LogicalExpression { return s.retValue }

func (s *OrExprContext) SetRetValue(v *opt.LogicalExpression) { s.retValue = v }

func (s *OrExprContext) AndExpr() IAndExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAndExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAndExprContext)
}

func (s *OrExprContext) OrExpr() IOrExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOrExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOrExprContext)
}

func (s *OrExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OrExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaListener); ok {
		listenerT.EnterOrExpr(s)
	}
}

func (s *OrExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaListener); ok {
		listenerT.ExitOrExpr(s)
	}
}

func (p *FormulaParser) OrExpr() (localctx IOrExprContext) {
	return p.orExpr(0)
}

func (p *FormulaParser) orExpr(_p int) (localctx IOrExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewOrExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IOrExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 4
	p.EnterRecursionRule(localctx, 4, FormulaParserRULE_orExpr, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(49)

		var _x = p.andExpr(0)

		localctx.(*OrExprContext)._andExpr = _x
	}
	localctx.(*OrExprContext).retValue = localctx.(*OrExprContext).Get_andExpr().GetRetValue()

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(59)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewOrExprContext(p, _parentctx, _parentState)
			localctx.(*OrExprContext).first = _prevctx
			p.PushNewRecursionContext(localctx, _startState, FormulaParserRULE_orExpr)
			p.SetState(52)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			p.SetState(53)
			_la = p.GetTokenStream().LA(1)

			if !(_la == FormulaParserT__2 || _la == FormulaParserT__3) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
			{
				p.SetState(54)

				var _x = p.andExpr(0)

				localctx.(*OrExprContext)._andExpr = _x
			}
			localctx.(*OrExprContext).retValue = exp.NewBinaryExpression("||", localctx.(*OrExprContext).GetFirst().GetRetValue(), localctx.(*OrExprContext).Get_andExpr().GetRetValue())

		}
		p.SetState(61)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())
	}

	return localctx
}

// IAndExprContext is an interface to support dynamic dispatch.
type IAndExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetFirst returns the first rule contexts.
	GetFirst() IAndExprContext

	// Get_bitOrExpr returns the _bitOrExpr rule contexts.
	Get_bitOrExpr() IBitOrExprContext

	// SetFirst sets the first rule contexts.
	SetFirst(IAndExprContext)

	// Set_bitOrExpr sets the _bitOrExpr rule contexts.
	Set_bitOrExpr(IBitOrExprContext)

	// GetRetValue returns the retValue attribute.
	GetRetValue() *opt.LogicalExpression

	// SetRetValue sets the retValue attribute.
	SetRetValue(*opt.LogicalExpression)

	// IsAndExprContext differentiates from other interfaces.
	IsAndExprContext()
}

type AndExprContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	retValue   *opt.LogicalExpression
	first      IAndExprContext
	_bitOrExpr IBitOrExprContext
}

func NewEmptyAndExprContext() *AndExprContext {
	var p = new(AndExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FormulaParserRULE_andExpr
	return p
}

func (*AndExprContext) IsAndExprContext() {}

func NewAndExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AndExprContext {
	var p = new(AndExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_andExpr

	return p
}

func (s *AndExprContext) GetParser() antlr.Parser { return s.parser }

func (s *AndExprContext) GetFirst() IAndExprContext { return s.first }

func (s *AndExprContext) Get_bitOrExpr() IBitOrExprContext { return s._bitOrExpr }

func (s *AndExprContext) SetFirst(v IAndExprContext) { s.first = v }

func (s *AndExprContext) Set_bitOrExpr(v IBitOrExprContext) { s._bitOrExpr = v }

func (s *AndExprContext) GetRetValue() *opt.LogicalExpression { return s.retValue }

func (s *AndExprContext) SetRetValue(v *opt.LogicalExpression) { s.retValue = v }

func (s *AndExprContext) BitOrExpr() IBitOrExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBitOrExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBitOrExprContext)
}

func (s *AndExprContext) AndExpr() IAndExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAndExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAndExprContext)
}

func (s *AndExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AndExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaListener); ok {
		listenerT.EnterAndExpr(s)
	}
}

func (s *AndExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaListener); ok {
		listenerT.ExitAndExpr(s)
	}
}

func (p *FormulaParser) AndExpr() (localctx IAndExprContext) {
	return p.andExpr(0)
}

func (p *FormulaParser) andExpr(_p int) (localctx IAndExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewAndExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IAndExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 6
	p.EnterRecursionRule(localctx, 6, FormulaParserRULE_andExpr, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(63)

		var _x = p.bitOrExpr(0)

		localctx.(*AndExprContext)._bitOrExpr = _x
	}
	localctx.(*AndExprContext).retValue = localctx.(*AndExprContext).Get_bitOrExpr().GetRetValue()

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(73)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewAndExprContext(p, _parentctx, _parentState)
			localctx.(*AndExprContext).first = _prevctx
			p.PushNewRecursionContext(localctx, _startState, FormulaParserRULE_andExpr)
			p.SetState(66)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			p.SetState(67)
			_la = p.GetTokenStream().LA(1)

			if !(_la == FormulaParserT__4 || _la == FormulaParserT__5) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
			{
				p.SetState(68)

				var _x = p.bitOrExpr(0)

				localctx.(*AndExprContext)._bitOrExpr = _x
			}
			localctx.(*AndExprContext).retValue = exp.NewBinaryExpression("&&", localctx.(*AndExprContext).GetFirst().GetRetValue(), localctx.(*AndExprContext).Get_bitOrExpr().GetRetValue())

		}
		p.SetState(75)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
	}

	return localctx
}

// IBitOrExprContext is an interface to support dynamic dispatch.
type IBitOrExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetFirst returns the first rule contexts.
	GetFirst() IBitOrExprContext

	// Get_bitXorExpr returns the _bitXorExpr rule contexts.
	Get_bitXorExpr() IBitXorExprContext

	// SetFirst sets the first rule contexts.
	SetFirst(IBitOrExprContext)

	// Set_bitXorExpr sets the _bitXorExpr rule contexts.
	Set_bitXorExpr(IBitXorExprContext)

	// GetRetValue returns the retValue attribute.
	GetRetValue() *opt.LogicalExpression

	// SetRetValue sets the retValue attribute.
	SetRetValue(*opt.LogicalExpression)

	// IsBitOrExprContext differentiates from other interfaces.
	IsBitOrExprContext()
}

type BitOrExprContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	retValue    *opt.LogicalExpression
	first       IBitOrExprContext
	_bitXorExpr IBitXorExprContext
}

func NewEmptyBitOrExprContext() *BitOrExprContext {
	var p = new(BitOrExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FormulaParserRULE_bitOrExpr
	return p
}

func (*BitOrExprContext) IsBitOrExprContext() {}

func NewBitOrExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BitOrExprContext {
	var p = new(BitOrExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_bitOrExpr

	return p
}

func (s *BitOrExprContext) GetParser() antlr.Parser { return s.parser }

func (s *BitOrExprContext) GetFirst() IBitOrExprContext { return s.first }

func (s *BitOrExprContext) Get_bitXorExpr() IBitXorExprContext { return s._bitXorExpr }

func (s *BitOrExprContext) SetFirst(v IBitOrExprContext) { s.first = v }

func (s *BitOrExprContext) Set_bitXorExpr(v IBitXorExprContext) { s._bitXorExpr = v }

func (s *BitOrExprContext) GetRetValue() *opt.LogicalExpression { return s.retValue }

func (s *BitOrExprContext) SetRetValue(v *opt.LogicalExpression) { s.retValue = v }

func (s *BitOrExprContext) BitXorExpr() IBitXorExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBitXorExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBitXorExprContext)
}

func (s *BitOrExprContext) BitOrExpr() IBitOrExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBitOrExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBitOrExprContext)
}

func (s *BitOrExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BitOrExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BitOrExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaListener); ok {
		listenerT.EnterBitOrExpr(s)
	}
}

func (s *BitOrExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaListener); ok {
		listenerT.ExitBitOrExpr(s)
	}
}

func (p *FormulaParser) BitOrExpr() (localctx IBitOrExprContext) {
	return p.bitOrExpr(0)
}

func (p *FormulaParser) bitOrExpr(_p int) (localctx IBitOrExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewBitOrExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IBitOrExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 8
	p.EnterRecursionRule(localctx, 8, FormulaParserRULE_bitOrExpr, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(77)

		var _x = p.bitXorExpr(0)

		localctx.(*BitOrExprContext)._bitXorExpr = _x
	}
	localctx.(*BitOrExprContext).retValue = localctx.(*BitOrExprContext).Get_bitXorExpr().GetRetValue()

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(87)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewBitOrExprContext(p, _parentctx, _parentState)
			localctx.(*BitOrExprContext).first = _prevctx
			p.PushNewRecursionContext(localctx, _startState, FormulaParserRULE_bitOrExpr)
			p.SetState(80)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(81)
				p.Match(FormulaParserT__6)
			}
			{
				p.SetState(82)

				var _x = p.bitXorExpr(0)

				localctx.(*BitOrExprContext)._bitXorExpr = _x
			}
			localctx.(*BitOrExprContext).retValue = exp.NewBinaryExpression("|", localctx.(*BitOrExprContext).GetFirst().GetRetValue(), localctx.(*BitOrExprContext).Get_bitXorExpr().GetRetValue())

		}
		p.SetState(89)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())
	}

	return localctx
}

// IBitXorExprContext is an interface to support dynamic dispatch.
type IBitXorExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetFirst returns the first rule contexts.
	GetFirst() IBitXorExprContext

	// Get_bitAndExpr returns the _bitAndExpr rule contexts.
	Get_bitAndExpr() IBitAndExprContext

	// SetFirst sets the first rule contexts.
	SetFirst(IBitXorExprContext)

	// Set_bitAndExpr sets the _bitAndExpr rule contexts.
	Set_bitAndExpr(IBitAndExprContext)

	// GetRetValue returns the retValue attribute.
	GetRetValue() *opt.LogicalExpression

	// SetRetValue sets the retValue attribute.
	SetRetValue(*opt.LogicalExpression)

	// IsBitXorExprContext differentiates from other interfaces.
	IsBitXorExprContext()
}

type BitXorExprContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	retValue    *opt.LogicalExpression
	first       IBitXorExprContext
	_bitAndExpr IBitAndExprContext
}

func NewEmptyBitXorExprContext() *BitXorExprContext {
	var p = new(BitXorExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FormulaParserRULE_bitXorExpr
	return p
}

func (*BitXorExprContext) IsBitXorExprContext() {}

func NewBitXorExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BitXorExprContext {
	var p = new(BitXorExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_bitXorExpr

	return p
}

func (s *BitXorExprContext) GetParser() antlr.Parser { return s.parser }

func (s *BitXorExprContext) GetFirst() IBitXorExprContext { return s.first }

func (s *BitXorExprContext) Get_bitAndExpr() IBitAndExprContext { return s._bitAndExpr }

func (s *BitXorExprContext) SetFirst(v IBitXorExprContext) { s.first = v }

func (s *BitXorExprContext) Set_bitAndExpr(v IBitAndExprContext) { s._bitAndExpr = v }

func (s *BitXorExprContext) GetRetValue() *opt.LogicalExpression { return s.retValue }

func (s *BitXorExprContext) SetRetValue(v *opt.LogicalExpression) { s.retValue = v }

func (s *BitXorExprContext) BitAndExpr() IBitAndExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBitAndExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBitAndExprContext)
}

func (s *BitXorExprContext) BitXorExpr() IBitXorExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBitXorExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBitXorExprContext)
}

func (s *BitXorExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BitXorExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BitXorExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaListener); ok {
		listenerT.EnterBitXorExpr(s)
	}
}

func (s *BitXorExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaListener); ok {
		listenerT.ExitBitXorExpr(s)
	}
}

func (p *FormulaParser) BitXorExpr() (localctx IBitXorExprContext) {
	return p.bitXorExpr(0)
}

func (p *FormulaParser) bitXorExpr(_p int) (localctx IBitXorExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewBitXorExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IBitXorExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 10
	p.EnterRecursionRule(localctx, 10, FormulaParserRULE_bitXorExpr, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(91)

		var _x = p.bitAndExpr(0)

		localctx.(*BitXorExprContext)._bitAndExpr = _x
	}
	localctx.(*BitXorExprContext).retValue = localctx.(*BitXorExprContext).Get_bitAndExpr().GetRetValue()

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(101)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewBitXorExprContext(p, _parentctx, _parentState)
			localctx.(*BitXorExprContext).first = _prevctx
			p.PushNewRecursionContext(localctx, _startState, FormulaParserRULE_bitXorExpr)
			p.SetState(94)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(95)
				p.Match(FormulaParserT__7)
			}
			{
				p.SetState(96)

				var _x = p.bitAndExpr(0)

				localctx.(*BitXorExprContext)._bitAndExpr = _x
			}
			localctx.(*BitXorExprContext).retValue = exp.NewBinaryExpression("^", localctx.(*BitXorExprContext).GetFirst().GetRetValue(), localctx.(*BitXorExprContext).Get_bitAndExpr().GetRetValue())

		}
		p.SetState(103)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())
	}

	return localctx
}

// IBitAndExprContext is an interface to support dynamic dispatch.
type IBitAndExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetFirst returns the first rule contexts.
	GetFirst() IBitAndExprContext

	// Get_eqExpr returns the _eqExpr rule contexts.
	Get_eqExpr() IEqExprContext

	// SetFirst sets the first rule contexts.
	SetFirst(IBitAndExprContext)

	// Set_eqExpr sets the _eqExpr rule contexts.
	Set_eqExpr(IEqExprContext)

	// GetRetValue returns the retValue attribute.
	GetRetValue() *opt.LogicalExpression

	// SetRetValue sets the retValue attribute.
	SetRetValue(*opt.LogicalExpression)

	// IsBitAndExprContext differentiates from other interfaces.
	IsBitAndExprContext()
}

type BitAndExprContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	retValue *opt.LogicalExpression
	first    IBitAndExprContext
	_eqExpr  IEqExprContext
}

func NewEmptyBitAndExprContext() *BitAndExprContext {
	var p = new(BitAndExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FormulaParserRULE_bitAndExpr
	return p
}

func (*BitAndExprContext) IsBitAndExprContext() {}

func NewBitAndExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BitAndExprContext {
	var p = new(BitAndExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_bitAndExpr

	return p
}

func (s *BitAndExprContext) GetParser() antlr.Parser { return s.parser }

func (s *BitAndExprContext) GetFirst() IBitAndExprContext { return s.first }

func (s *BitAndExprContext) Get_eqExpr() IEqExprContext { return s._eqExpr }

func (s *BitAndExprContext) SetFirst(v IBitAndExprContext) { s.first = v }

func (s *BitAndExprContext) Set_eqExpr(v IEqExprContext) { s._eqExpr = v }

func (s *BitAndExprContext) GetRetValue() *opt.LogicalExpression { return s.retValue }

func (s *BitAndExprContext) SetRetValue(v *opt.LogicalExpression) { s.retValue = v }

func (s *BitAndExprContext) EqExpr() IEqExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEqExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEqExprContext)
}

func (s *BitAndExprContext) BitAndExpr() IBitAndExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBitAndExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBitAndExprContext)
}

func (s *BitAndExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BitAndExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BitAndExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaListener); ok {
		listenerT.EnterBitAndExpr(s)
	}
}

func (s *BitAndExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaListener); ok {
		listenerT.ExitBitAndExpr(s)
	}
}

func (p *FormulaParser) BitAndExpr() (localctx IBitAndExprContext) {
	return p.bitAndExpr(0)
}

func (p *FormulaParser) bitAndExpr(_p int) (localctx IBitAndExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewBitAndExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IBitAndExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 12
	p.EnterRecursionRule(localctx, 12, FormulaParserRULE_bitAndExpr, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(105)

		var _x = p.eqExpr(0)

		localctx.(*BitAndExprContext)._eqExpr = _x
	}
	localctx.(*BitAndExprContext).retValue = localctx.(*BitAndExprContext).Get_eqExpr().GetRetValue()

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(115)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewBitAndExprContext(p, _parentctx, _parentState)
			localctx.(*BitAndExprContext).first = _prevctx
			p.PushNewRecursionContext(localctx, _startState, FormulaParserRULE_bitAndExpr)
			p.SetState(108)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(109)
				p.Match(FormulaParserT__8)
			}
			{
				p.SetState(110)

				var _x = p.eqExpr(0)

				localctx.(*BitAndExprContext)._eqExpr = _x
			}
			localctx.(*BitAndExprContext).retValue = exp.NewBinaryExpression("&", localctx.(*BitAndExprContext).GetFirst().GetRetValue(), localctx.(*BitAndExprContext).Get_eqExpr().GetRetValue())

		}
		p.SetState(117)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())
	}

	return localctx
}

// IEqExprContext is an interface to support dynamic dispatch.
type IEqExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetFirst returns the first rule contexts.
	GetFirst() IEqExprContext

	// Get_relExpr returns the _relExpr rule contexts.
	Get_relExpr() IRelExprContext

	// SetFirst sets the first rule contexts.
	SetFirst(IEqExprContext)

	// Set_relExpr sets the _relExpr rule contexts.
	Set_relExpr(IRelExprContext)

	// GetRetValue returns the retValue attribute.
	GetRetValue() *opt.LogicalExpression

	// SetRetValue sets the retValue attribute.
	SetRetValue(*opt.LogicalExpression)

	// IsEqExprContext differentiates from other interfaces.
	IsEqExprContext()
}

type EqExprContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	retValue *opt.LogicalExpression
	first    IEqExprContext
	_relExpr IRelExprContext
}

func NewEmptyEqExprContext() *EqExprContext {
	var p = new(EqExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FormulaParserRULE_eqExpr
	return p
}

func (*EqExprContext) IsEqExprContext() {}

func NewEqExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EqExprContext {
	var p = new(EqExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_eqExpr

	return p
}

func (s *EqExprContext) GetParser() antlr.Parser { return s.parser }

func (s *EqExprContext) GetFirst() IEqExprContext { return s.first }

func (s *EqExprContext) Get_relExpr() IRelExprContext { return s._relExpr }

func (s *EqExprContext) SetFirst(v IEqExprContext) { s.first = v }

func (s *EqExprContext) Set_relExpr(v IRelExprContext) { s._relExpr = v }

func (s *EqExprContext) GetRetValue() *opt.LogicalExpression { return s.retValue }

func (s *EqExprContext) SetRetValue(v *opt.LogicalExpression) { s.retValue = v }

func (s *EqExprContext) RelExpr() IRelExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRelExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRelExprContext)
}

func (s *EqExprContext) EqExpr() IEqExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEqExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEqExprContext)
}

func (s *EqExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EqExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaListener); ok {
		listenerT.EnterEqExpr(s)
	}
}

func (s *EqExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaListener); ok {
		listenerT.ExitEqExpr(s)
	}
}

func (p *FormulaParser) EqExpr() (localctx IEqExprContext) {
	return p.eqExpr(0)
}

func (p *FormulaParser) eqExpr(_p int) (localctx IEqExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewEqExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IEqExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 14
	p.EnterRecursionRule(localctx, 14, FormulaParserRULE_eqExpr, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(119)

		var _x = p.relExpr(0)

		localctx.(*EqExprContext)._relExpr = _x
	}
	localctx.(*EqExprContext).retValue = localctx.(*EqExprContext).Get_relExpr().GetRetValue()

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(134)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(132)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
			case 1:
				localctx = NewEqExprContext(p, _parentctx, _parentState)
				localctx.(*EqExprContext).first = _prevctx
				p.PushNewRecursionContext(localctx, _startState, FormulaParserRULE_eqExpr)
				p.SetState(122)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				p.SetState(123)
				_la = p.GetTokenStream().LA(1)

				if !(_la == FormulaParserT__9 || _la == FormulaParserT__10) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
				{
					p.SetState(124)

					var _x = p.relExpr(0)

					localctx.(*EqExprContext)._relExpr = _x
				}
				localctx.(*EqExprContext).retValue = exp.NewBinaryExpression("==", localctx.(*EqExprContext).GetFirst().GetRetValue(), localctx.(*EqExprContext).Get_relExpr().GetRetValue())

			case 2:
				localctx = NewEqExprContext(p, _parentctx, _parentState)
				localctx.(*EqExprContext).first = _prevctx
				p.PushNewRecursionContext(localctx, _startState, FormulaParserRULE_eqExpr)
				p.SetState(127)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				p.SetState(128)
				_la = p.GetTokenStream().LA(1)

				if !(_la == FormulaParserT__11 || _la == FormulaParserT__12) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
				{
					p.SetState(129)

					var _x = p.relExpr(0)

					localctx.(*EqExprContext)._relExpr = _x
				}
				localctx.(*EqExprContext).retValue = exp.NewBinaryExpression("!=", localctx.(*EqExprContext).GetFirst().GetRetValue(), localctx.(*EqExprContext).Get_relExpr().GetRetValue())

			}

		}
		p.SetState(136)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())
	}

	return localctx
}

// IRelExprContext is an interface to support dynamic dispatch.
type IRelExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetFirst returns the first rule contexts.
	GetFirst() IRelExprContext

	// Get_shiftExpr returns the _shiftExpr rule contexts.
	Get_shiftExpr() IShiftExprContext

	// SetFirst sets the first rule contexts.
	SetFirst(IRelExprContext)

	// Set_shiftExpr sets the _shiftExpr rule contexts.
	Set_shiftExpr(IShiftExprContext)

	// GetRetValue returns the retValue attribute.
	GetRetValue() *opt.LogicalExpression

	// SetRetValue sets the retValue attribute.
	SetRetValue(*opt.LogicalExpression)

	// IsRelExprContext differentiates from other interfaces.
	IsRelExprContext()
}

type RelExprContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	retValue   *opt.LogicalExpression
	first      IRelExprContext
	_shiftExpr IShiftExprContext
}

func NewEmptyRelExprContext() *RelExprContext {
	var p = new(RelExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FormulaParserRULE_relExpr
	return p
}

func (*RelExprContext) IsRelExprContext() {}

func NewRelExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelExprContext {
	var p = new(RelExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_relExpr

	return p
}

func (s *RelExprContext) GetParser() antlr.Parser { return s.parser }

func (s *RelExprContext) GetFirst() IRelExprContext { return s.first }

func (s *RelExprContext) Get_shiftExpr() IShiftExprContext { return s._shiftExpr }

func (s *RelExprContext) SetFirst(v IRelExprContext) { s.first = v }

func (s *RelExprContext) Set_shiftExpr(v IShiftExprContext) { s._shiftExpr = v }

func (s *RelExprContext) GetRetValue() *opt.LogicalExpression { return s.retValue }

func (s *RelExprContext) SetRetValue(v *opt.LogicalExpression) { s.retValue = v }

func (s *RelExprContext) ShiftExpr() IShiftExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IShiftExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IShiftExprContext)
}

func (s *RelExprContext) RelExpr() IRelExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRelExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRelExprContext)
}

func (s *RelExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RelExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaListener); ok {
		listenerT.EnterRelExpr(s)
	}
}

func (s *RelExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaListener); ok {
		listenerT.ExitRelExpr(s)
	}
}

func (p *FormulaParser) RelExpr() (localctx IRelExprContext) {
	return p.relExpr(0)
}

func (p *FormulaParser) relExpr(_p int) (localctx IRelExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewRelExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IRelExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 16
	p.EnterRecursionRule(localctx, 16, FormulaParserRULE_relExpr, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(138)

		var _x = p.shiftExpr(0)

		localctx.(*RelExprContext)._shiftExpr = _x
	}
	localctx.(*RelExprContext).retValue = localctx.(*RelExprContext).Get_shiftExpr().GetRetValue()

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(163)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(161)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
			case 1:
				localctx = NewRelExprContext(p, _parentctx, _parentState)
				localctx.(*RelExprContext).first = _prevctx
				p.PushNewRecursionContext(localctx, _startState, FormulaParserRULE_relExpr)
				p.SetState(141)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(142)
					p.Match(FormulaParserT__13)
				}
				{
					p.SetState(143)

					var _x = p.shiftExpr(0)

					localctx.(*RelExprContext)._shiftExpr = _x
				}
				localctx.(*RelExprContext).retValue = exp.NewBinaryExpression("<", localctx.(*RelExprContext).GetFirst().GetRetValue(), localctx.(*RelExprContext).Get_shiftExpr().GetRetValue())

			case 2:
				localctx = NewRelExprContext(p, _parentctx, _parentState)
				localctx.(*RelExprContext).first = _prevctx
				p.PushNewRecursionContext(localctx, _startState, FormulaParserRULE_relExpr)
				p.SetState(146)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(147)
					p.Match(FormulaParserT__14)
				}
				{
					p.SetState(148)

					var _x = p.shiftExpr(0)

					localctx.(*RelExprContext)._shiftExpr = _x
				}
				localctx.(*RelExprContext).retValue = exp.NewBinaryExpression("<=", localctx.(*RelExprContext).GetFirst().GetRetValue(), localctx.(*RelExprContext).Get_shiftExpr().GetRetValue())

			case 3:
				localctx = NewRelExprContext(p, _parentctx, _parentState)
				localctx.(*RelExprContext).first = _prevctx
				p.PushNewRecursionContext(localctx, _startState, FormulaParserRULE_relExpr)
				p.SetState(151)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(152)
					p.Match(FormulaParserT__15)
				}
				{
					p.SetState(153)

					var _x = p.shiftExpr(0)

					localctx.(*RelExprContext)._shiftExpr = _x
				}
				localctx.(*RelExprContext).retValue = exp.NewBinaryExpression(">", localctx.(*RelExprContext).GetFirst().GetRetValue(), localctx.(*RelExprContext).Get_shiftExpr().GetRetValue())

			case 4:
				localctx = NewRelExprContext(p, _parentctx, _parentState)
				localctx.(*RelExprContext).first = _prevctx
				p.PushNewRecursionContext(localctx, _startState, FormulaParserRULE_relExpr)
				p.SetState(156)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(157)
					p.Match(FormulaParserT__16)
				}
				{
					p.SetState(158)

					var _x = p.shiftExpr(0)

					localctx.(*RelExprContext)._shiftExpr = _x
				}
				localctx.(*RelExprContext).retValue = exp.NewBinaryExpression(">=", localctx.(*RelExprContext).GetFirst().GetRetValue(), localctx.(*RelExprContext).Get_shiftExpr().GetRetValue())

			}

		}
		p.SetState(165)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())
	}

	return localctx
}

// IShiftExprContext is an interface to support dynamic dispatch.
type IShiftExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetFirst returns the first rule contexts.
	GetFirst() IShiftExprContext

	// Get_addExpr returns the _addExpr rule contexts.
	Get_addExpr() IAddExprContext

	// SetFirst sets the first rule contexts.
	SetFirst(IShiftExprContext)

	// Set_addExpr sets the _addExpr rule contexts.
	Set_addExpr(IAddExprContext)

	// GetRetValue returns the retValue attribute.
	GetRetValue() *opt.LogicalExpression

	// SetRetValue sets the retValue attribute.
	SetRetValue(*opt.LogicalExpression)

	// IsShiftExprContext differentiates from other interfaces.
	IsShiftExprContext()
}

type ShiftExprContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	retValue *opt.LogicalExpression
	first    IShiftExprContext
	_addExpr IAddExprContext
}

func NewEmptyShiftExprContext() *ShiftExprContext {
	var p = new(ShiftExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FormulaParserRULE_shiftExpr
	return p
}

func (*ShiftExprContext) IsShiftExprContext() {}

func NewShiftExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ShiftExprContext {
	var p = new(ShiftExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_shiftExpr

	return p
}

func (s *ShiftExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ShiftExprContext) GetFirst() IShiftExprContext { return s.first }

func (s *ShiftExprContext) Get_addExpr() IAddExprContext { return s._addExpr }

func (s *ShiftExprContext) SetFirst(v IShiftExprContext) { s.first = v }

func (s *ShiftExprContext) Set_addExpr(v IAddExprContext) { s._addExpr = v }

func (s *ShiftExprContext) GetRetValue() *opt.LogicalExpression { return s.retValue }

func (s *ShiftExprContext) SetRetValue(v *opt.LogicalExpression) { s.retValue = v }

func (s *ShiftExprContext) AddExpr() IAddExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAddExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAddExprContext)
}

func (s *ShiftExprContext) ShiftExpr() IShiftExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IShiftExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IShiftExprContext)
}

func (s *ShiftExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ShiftExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ShiftExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaListener); ok {
		listenerT.EnterShiftExpr(s)
	}
}

func (s *ShiftExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaListener); ok {
		listenerT.ExitShiftExpr(s)
	}
}

func (p *FormulaParser) ShiftExpr() (localctx IShiftExprContext) {
	return p.shiftExpr(0)
}

func (p *FormulaParser) shiftExpr(_p int) (localctx IShiftExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewShiftExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IShiftExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 18
	p.EnterRecursionRule(localctx, 18, FormulaParserRULE_shiftExpr, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(167)

		var _x = p.addExpr(0)

		localctx.(*ShiftExprContext)._addExpr = _x
	}
	localctx.(*ShiftExprContext).retValue = localctx.(*ShiftExprContext).Get_addExpr().GetRetValue()

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(182)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(180)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
			case 1:
				localctx = NewShiftExprContext(p, _parentctx, _parentState)
				localctx.(*ShiftExprContext).first = _prevctx
				p.PushNewRecursionContext(localctx, _startState, FormulaParserRULE_shiftExpr)
				p.SetState(170)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(171)
					p.Match(FormulaParserT__17)
				}
				{
					p.SetState(172)

					var _x = p.addExpr(0)

					localctx.(*ShiftExprContext)._addExpr = _x
				}
				localctx.(*ShiftExprContext).retValue = exp.NewBinaryExpression("<<", localctx.(*ShiftExprContext).GetFirst().GetRetValue(), localctx.(*ShiftExprContext).Get_addExpr().GetRetValue())

			case 2:
				localctx = NewShiftExprContext(p, _parentctx, _parentState)
				localctx.(*ShiftExprContext).first = _prevctx
				p.PushNewRecursionContext(localctx, _startState, FormulaParserRULE_shiftExpr)
				p.SetState(175)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(176)
					p.Match(FormulaParserT__18)
				}
				{
					p.SetState(177)

					var _x = p.addExpr(0)

					localctx.(*ShiftExprContext)._addExpr = _x
				}
				localctx.(*ShiftExprContext).retValue = exp.NewBinaryExpression(">>", localctx.(*ShiftExprContext).GetFirst().GetRetValue(), localctx.(*ShiftExprContext).Get_addExpr().GetRetValue())

			}

		}
		p.SetState(184)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())
	}

	return localctx
}

// IAddExprContext is an interface to support dynamic dispatch.
type IAddExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetFirst returns the first rule contexts.
	GetFirst() IAddExprContext

	// Get_multExpr returns the _multExpr rule contexts.
	Get_multExpr() IMultExprContext

	// SetFirst sets the first rule contexts.
	SetFirst(IAddExprContext)

	// Set_multExpr sets the _multExpr rule contexts.
	Set_multExpr(IMultExprContext)

	// GetRetValue returns the retValue attribute.
	GetRetValue() *opt.LogicalExpression

	// SetRetValue sets the retValue attribute.
	SetRetValue(*opt.LogicalExpression)

	// IsAddExprContext differentiates from other interfaces.
	IsAddExprContext()
}

type AddExprContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	retValue  *opt.LogicalExpression
	first     IAddExprContext
	_multExpr IMultExprContext
}

func NewEmptyAddExprContext() *AddExprContext {
	var p = new(AddExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FormulaParserRULE_addExpr
	return p
}

func (*AddExprContext) IsAddExprContext() {}

func NewAddExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AddExprContext {
	var p = new(AddExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_addExpr

	return p
}

func (s *AddExprContext) GetParser() antlr.Parser { return s.parser }

func (s *AddExprContext) GetFirst() IAddExprContext { return s.first }

func (s *AddExprContext) Get_multExpr() IMultExprContext { return s._multExpr }

func (s *AddExprContext) SetFirst(v IAddExprContext) { s.first = v }

func (s *AddExprContext) Set_multExpr(v IMultExprContext) { s._multExpr = v }

func (s *AddExprContext) GetRetValue() *opt.LogicalExpression { return s.retValue }

func (s *AddExprContext) SetRetValue(v *opt.LogicalExpression) { s.retValue = v }

func (s *AddExprContext) MultExpr() IMultExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMultExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMultExprContext)
}

func (s *AddExprContext) AddExpr() IAddExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAddExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAddExprContext)
}

func (s *AddExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AddExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaListener); ok {
		listenerT.EnterAddExpr(s)
	}
}

func (s *AddExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaListener); ok {
		listenerT.ExitAddExpr(s)
	}
}

func (p *FormulaParser) AddExpr() (localctx IAddExprContext) {
	return p.addExpr(0)
}

func (p *FormulaParser) addExpr(_p int) (localctx IAddExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewAddExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IAddExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 20
	p.EnterRecursionRule(localctx, 20, FormulaParserRULE_addExpr, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(186)

		var _x = p.multExpr(0)

		localctx.(*AddExprContext)._multExpr = _x
	}
	localctx.(*AddExprContext).retValue = localctx.(*AddExprContext).Get_multExpr().GetRetValue()

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(201)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(199)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) {
			case 1:
				localctx = NewAddExprContext(p, _parentctx, _parentState)
				localctx.(*AddExprContext).first = _prevctx
				p.PushNewRecursionContext(localctx, _startState, FormulaParserRULE_addExpr)
				p.SetState(189)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(190)
					p.Match(FormulaParserT__19)
				}
				{
					p.SetState(191)

					var _x = p.multExpr(0)

					localctx.(*AddExprContext)._multExpr = _x
				}
				localctx.(*AddExprContext).retValue = exp.NewBinaryExpression("+", localctx.(*AddExprContext).GetFirst().GetRetValue(), localctx.(*AddExprContext).Get_multExpr().GetRetValue())

			case 2:
				localctx = NewAddExprContext(p, _parentctx, _parentState)
				localctx.(*AddExprContext).first = _prevctx
				p.PushNewRecursionContext(localctx, _startState, FormulaParserRULE_addExpr)
				p.SetState(194)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(195)
					p.Match(FormulaParserT__20)
				}
				{
					p.SetState(196)

					var _x = p.multExpr(0)

					localctx.(*AddExprContext)._multExpr = _x
				}
				localctx.(*AddExprContext).retValue = exp.NewBinaryExpression("-", localctx.(*AddExprContext).GetFirst().GetRetValue(), localctx.(*AddExprContext).Get_multExpr().GetRetValue())

			}

		}
		p.SetState(203)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext())
	}

	return localctx
}

// IMultExprContext is an interface to support dynamic dispatch.
type IMultExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetFirst returns the first rule contexts.
	GetFirst() IMultExprContext

	// Get_unaryExpr returns the _unaryExpr rule contexts.
	Get_unaryExpr() IUnaryExprContext

	// SetFirst sets the first rule contexts.
	SetFirst(IMultExprContext)

	// Set_unaryExpr sets the _unaryExpr rule contexts.
	Set_unaryExpr(IUnaryExprContext)

	// GetRetValue returns the retValue attribute.
	GetRetValue() *opt.LogicalExpression

	// SetRetValue sets the retValue attribute.
	SetRetValue(*opt.LogicalExpression)

	// IsMultExprContext differentiates from other interfaces.
	IsMultExprContext()
}

type MultExprContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	retValue   *opt.LogicalExpression
	first      IMultExprContext
	_unaryExpr IUnaryExprContext
}

func NewEmptyMultExprContext() *MultExprContext {
	var p = new(MultExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FormulaParserRULE_multExpr
	return p
}

func (*MultExprContext) IsMultExprContext() {}

func NewMultExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultExprContext {
	var p = new(MultExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_multExpr

	return p
}

func (s *MultExprContext) GetParser() antlr.Parser { return s.parser }

func (s *MultExprContext) GetFirst() IMultExprContext { return s.first }

func (s *MultExprContext) Get_unaryExpr() IUnaryExprContext { return s._unaryExpr }

func (s *MultExprContext) SetFirst(v IMultExprContext) { s.first = v }

func (s *MultExprContext) Set_unaryExpr(v IUnaryExprContext) { s._unaryExpr = v }

func (s *MultExprContext) GetRetValue() *opt.LogicalExpression { return s.retValue }

func (s *MultExprContext) SetRetValue(v *opt.LogicalExpression) { s.retValue = v }

func (s *MultExprContext) UnaryExpr() IUnaryExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnaryExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnaryExprContext)
}

func (s *MultExprContext) MultExpr() IMultExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMultExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMultExprContext)
}

func (s *MultExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MultExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaListener); ok {
		listenerT.EnterMultExpr(s)
	}
}

func (s *MultExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaListener); ok {
		listenerT.ExitMultExpr(s)
	}
}

func (p *FormulaParser) MultExpr() (localctx IMultExprContext) {
	return p.multExpr(0)
}

func (p *FormulaParser) multExpr(_p int) (localctx IMultExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewMultExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IMultExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 22
	p.EnterRecursionRule(localctx, 22, FormulaParserRULE_multExpr, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(205)

		var _x = p.UnaryExpr()

		localctx.(*MultExprContext)._unaryExpr = _x
	}
	localctx.(*MultExprContext).retValue = localctx.(*MultExprContext).Get_unaryExpr().GetRetValue()

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(225)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(223)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMultExprContext(p, _parentctx, _parentState)
				localctx.(*MultExprContext).first = _prevctx
				p.PushNewRecursionContext(localctx, _startState, FormulaParserRULE_multExpr)
				p.SetState(208)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(209)
					p.Match(FormulaParserT__21)
				}
				{
					p.SetState(210)

					var _x = p.UnaryExpr()

					localctx.(*MultExprContext)._unaryExpr = _x
				}
				localctx.(*MultExprContext).retValue = exp.NewBinaryExpression("*", localctx.(*MultExprContext).GetFirst().GetRetValue(), localctx.(*MultExprContext).Get_unaryExpr().GetRetValue())

			case 2:
				localctx = NewMultExprContext(p, _parentctx, _parentState)
				localctx.(*MultExprContext).first = _prevctx
				p.PushNewRecursionContext(localctx, _startState, FormulaParserRULE_multExpr)
				p.SetState(213)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(214)
					p.Match(FormulaParserT__22)
				}
				{
					p.SetState(215)

					var _x = p.UnaryExpr()

					localctx.(*MultExprContext)._unaryExpr = _x
				}
				localctx.(*MultExprContext).retValue = exp.NewBinaryExpression("/", localctx.(*MultExprContext).GetFirst().GetRetValue(), localctx.(*MultExprContext).Get_unaryExpr().GetRetValue())

			case 3:
				localctx = NewMultExprContext(p, _parentctx, _parentState)
				localctx.(*MultExprContext).first = _prevctx
				p.PushNewRecursionContext(localctx, _startState, FormulaParserRULE_multExpr)
				p.SetState(218)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(219)
					p.Match(FormulaParserT__23)
				}
				{
					p.SetState(220)

					var _x = p.UnaryExpr()

					localctx.(*MultExprContext)._unaryExpr = _x
				}
				localctx.(*MultExprContext).retValue = exp.NewBinaryExpression("%", localctx.(*MultExprContext).GetFirst().GetRetValue(), localctx.(*MultExprContext).Get_unaryExpr().GetRetValue())

			}

		}
		p.SetState(227)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext())
	}

	return localctx
}

// IUnaryExprContext is an interface to support dynamic dispatch.
type IUnaryExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_primaryExpr returns the _primaryExpr rule contexts.
	Get_primaryExpr() IPrimaryExprContext

	// Set_primaryExpr sets the _primaryExpr rule contexts.
	Set_primaryExpr(IPrimaryExprContext)

	// GetRetValue returns the retValue attribute.
	GetRetValue() *opt.LogicalExpression

	// SetRetValue sets the retValue attribute.
	SetRetValue(*opt.LogicalExpression)

	// IsUnaryExprContext differentiates from other interfaces.
	IsUnaryExprContext()
}

type UnaryExprContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	retValue     *opt.LogicalExpression
	_primaryExpr IPrimaryExprContext
}

func NewEmptyUnaryExprContext() *UnaryExprContext {
	var p = new(UnaryExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FormulaParserRULE_unaryExpr
	return p
}

func (*UnaryExprContext) IsUnaryExprContext() {}

func NewUnaryExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnaryExprContext {
	var p = new(UnaryExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_unaryExpr

	return p
}

func (s *UnaryExprContext) GetParser() antlr.Parser { return s.parser }

func (s *UnaryExprContext) Get_primaryExpr() IPrimaryExprContext { return s._primaryExpr }

func (s *UnaryExprContext) Set_primaryExpr(v IPrimaryExprContext) { s._primaryExpr = v }

func (s *UnaryExprContext) GetRetValue() *opt.LogicalExpression { return s.retValue }

func (s *UnaryExprContext) SetRetValue(v *opt.LogicalExpression) { s.retValue = v }

func (s *UnaryExprContext) PrimaryExpr() IPrimaryExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimaryExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimaryExprContext)
}

func (s *UnaryExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnaryExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaListener); ok {
		listenerT.EnterUnaryExpr(s)
	}
}

func (s *UnaryExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaListener); ok {
		listenerT.ExitUnaryExpr(s)
	}
}

func (p *FormulaParser) UnaryExpr() (localctx IUnaryExprContext) {
	localctx = NewUnaryExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, FormulaParserRULE_unaryExpr)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(243)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case FormulaParserT__24, FormulaParserT__25:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(228)
		_la = p.GetTokenStream().LA(1)

		if !(_la == FormulaParserT__24 || _la == FormulaParserT__25) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
		{
			p.SetState(229)

			var _x = p.PrimaryExpr()

			localctx.(*UnaryExprContext)._primaryExpr = _x
		}
		localctx.(*UnaryExprContext).retValue = exp.NewNotUnaryExpression("!", localctx.(*UnaryExprContext).Get_primaryExpr().GetRetValue())

	case FormulaParserT__26:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(232)
			p.Match(FormulaParserT__26)
		}
		{
			p.SetState(233)

			var _x = p.PrimaryExpr()

			localctx.(*UnaryExprContext)._primaryExpr = _x
		}
		localctx.(*UnaryExprContext).retValue = exp.NewBitwiseNotUnaryExpression("~", localctx.(*UnaryExprContext).Get_primaryExpr().GetRetValue())

	case FormulaParserT__20:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(236)
			p.Match(FormulaParserT__20)
		}
		{
			p.SetState(237)

			var _x = p.PrimaryExpr()

			localctx.(*UnaryExprContext)._primaryExpr = _x
		}
		localctx.(*UnaryExprContext).retValue = exp.NewNegateUnaryExpression("-", localctx.(*UnaryExprContext).Get_primaryExpr().GetRetValue())

	case FormulaParserT__27, FormulaParserTRUE, FormulaParserFALSE, FormulaParserNAME, FormulaParserINTEGER, FormulaParserDATETIME, FormulaParserVAR, FormulaParserFLOAT, FormulaParserSTRING:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(240)

			var _x = p.PrimaryExpr()

			localctx.(*UnaryExprContext)._primaryExpr = _x
		}
		localctx.(*UnaryExprContext).retValue = localctx.(*UnaryExprContext).Get_primaryExpr().GetRetValue()

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IPrimaryExprContext is an interface to support dynamic dispatch.
type IPrimaryExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// Get_value returns the _value rule contexts.
	Get_value() IValueContext

	// Get_id returns the _id rule contexts.
	Get_id() IIdContext

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// Set_value sets the _value rule contexts.
	Set_value(IValueContext)

	// Set_id sets the _id rule contexts.
	Set_id(IIdContext)

	// GetRetValue returns the retValue attribute.
	GetRetValue() *opt.LogicalExpression

	// SetRetValue sets the retValue attribute.
	SetRetValue(*opt.LogicalExpression)

	// IsPrimaryExprContext differentiates from other interfaces.
	IsPrimaryExprContext()
}

type PrimaryExprContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	retValue *opt.LogicalExpression
	_expr    IExprContext
	_value   IValueContext
	_id      IIdContext
}

func NewEmptyPrimaryExprContext() *PrimaryExprContext {
	var p = new(PrimaryExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FormulaParserRULE_primaryExpr
	return p
}

func (*PrimaryExprContext) IsPrimaryExprContext() {}

func NewPrimaryExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryExprContext {
	var p = new(PrimaryExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_primaryExpr

	return p
}

func (s *PrimaryExprContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryExprContext) Get_expr() IExprContext { return s._expr }

func (s *PrimaryExprContext) Get_value() IValueContext { return s._value }

func (s *PrimaryExprContext) Get_id() IIdContext { return s._id }

func (s *PrimaryExprContext) Set_expr(v IExprContext) { s._expr = v }

func (s *PrimaryExprContext) Set_value(v IValueContext) { s._value = v }

func (s *PrimaryExprContext) Set_id(v IIdContext) { s._id = v }

func (s *PrimaryExprContext) GetRetValue() *opt.LogicalExpression { return s.retValue }

func (s *PrimaryExprContext) SetRetValue(v *opt.LogicalExpression) { s.retValue = v }

func (s *PrimaryExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *PrimaryExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *PrimaryExprContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *PrimaryExprContext) Id() IIdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdContext)
}

func (s *PrimaryExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimaryExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaListener); ok {
		listenerT.EnterPrimaryExpr(s)
	}
}

func (s *PrimaryExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaListener); ok {
		listenerT.ExitPrimaryExpr(s)
	}
}

func (p *FormulaParser) PrimaryExpr() (localctx IPrimaryExprContext) {
	localctx = NewPrimaryExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, FormulaParserRULE_primaryExpr)
	var args = make([]*opt.LogicalExpression, 0)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(272)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(245)
			p.Match(FormulaParserT__27)
		}
		{
			p.SetState(246)

			var _x = p.Expr()

			localctx.(*PrimaryExprContext)._expr = _x
		}
		{
			p.SetState(247)
			p.Match(FormulaParserT__28)
		}
		localctx.(*PrimaryExprContext).retValue = localctx.(*PrimaryExprContext).Get_expr().GetRetValue()

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(250)

			var _x = p.Value()

			localctx.(*PrimaryExprContext)._value = _x
		}
		localctx.(*PrimaryExprContext).retValue = localctx.(*PrimaryExprContext).Get_value().GetRetValue()

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(253)

			var _x = p.Id()

			localctx.(*PrimaryExprContext)._id = _x
		}
		{
			p.SetState(254)
			p.Match(FormulaParserT__27)
		}
		{
			p.SetState(255)

			var _x = p.Expr()

			localctx.(*PrimaryExprContext)._expr = _x
		}
		args = append(args, localctx.(*PrimaryExprContext).Get_expr().GetRetValue())
		p.SetState(263)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == FormulaParserT__29 {
			{
				p.SetState(257)
				p.Match(FormulaParserT__29)
			}
			{
				p.SetState(258)

				var _x = p.Expr()

				localctx.(*PrimaryExprContext)._expr = _x
			}
			args = append(args, localctx.(*PrimaryExprContext).Get_expr().GetRetValue())

			p.SetState(265)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(266)
			p.Match(FormulaParserT__28)
		}
		localctx.(*PrimaryExprContext).retValue = exp.NewFunctionExpression(opt.NewIdentifier((func() string {
			if localctx.(*PrimaryExprContext).Get_id() == nil {
				return ""
			} else {
				return p.GetTokenStream().GetTextFromTokens(localctx.(*PrimaryExprContext).Get_id().GetStart(), localctx.(*PrimaryExprContext)._id.GetStop())
			}
		}())), args)

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(269)

			var _x = p.Id()

			localctx.(*PrimaryExprContext)._id = _x
		}
		localctx.(*PrimaryExprContext).retValue = localctx.(*PrimaryExprContext).Get_id().GetRetValue()

	}

	return localctx
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_INTEGER returns the _INTEGER token.
	Get_INTEGER() antlr.Token

	// Get_FLOAT returns the _FLOAT token.
	Get_FLOAT() antlr.Token

	// Get_STRING returns the _STRING token.
	Get_STRING() antlr.Token

	// Get_DATETIME returns the _DATETIME token.
	Get_DATETIME() antlr.Token

	// Set_INTEGER sets the _INTEGER token.
	Set_INTEGER(antlr.Token)

	// Set_FLOAT sets the _FLOAT token.
	Set_FLOAT(antlr.Token)

	// Set_STRING sets the _STRING token.
	Set_STRING(antlr.Token)

	// Set_DATETIME sets the _DATETIME token.
	Set_DATETIME(antlr.Token)

	// GetRetValue returns the retValue attribute.
	GetRetValue() *opt.LogicalExpression

	// SetRetValue sets the retValue attribute.
	SetRetValue(*opt.LogicalExpression)

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	retValue  *opt.LogicalExpression
	_INTEGER  antlr.Token
	_FLOAT    antlr.Token
	_STRING   antlr.Token
	_DATETIME antlr.Token
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FormulaParserRULE_value
	return p
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) Get_INTEGER() antlr.Token { return s._INTEGER }

func (s *ValueContext) Get_FLOAT() antlr.Token { return s._FLOAT }

func (s *ValueContext) Get_STRING() antlr.Token { return s._STRING }

func (s *ValueContext) Get_DATETIME() antlr.Token { return s._DATETIME }

func (s *ValueContext) Set_INTEGER(v antlr.Token) { s._INTEGER = v }

func (s *ValueContext) Set_FLOAT(v antlr.Token) { s._FLOAT = v }

func (s *ValueContext) Set_STRING(v antlr.Token) { s._STRING = v }

func (s *ValueContext) Set_DATETIME(v antlr.Token) { s._DATETIME = v }

func (s *ValueContext) GetRetValue() *opt.LogicalExpression { return s.retValue }

func (s *ValueContext) SetRetValue(v *opt.LogicalExpression) { s.retValue = v }

func (s *ValueContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(FormulaParserINTEGER, 0)
}

func (s *ValueContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(FormulaParserFLOAT, 0)
}

func (s *ValueContext) STRING() antlr.TerminalNode {
	return s.GetToken(FormulaParserSTRING, 0)
}

func (s *ValueContext) DATETIME() antlr.TerminalNode {
	return s.GetToken(FormulaParserDATETIME, 0)
}

func (s *ValueContext) TRUE() antlr.TerminalNode {
	return s.GetToken(FormulaParserTRUE, 0)
}

func (s *ValueContext) FALSE() antlr.TerminalNode {
	return s.GetToken(FormulaParserFALSE, 0)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaListener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaListener); ok {
		listenerT.ExitValue(s)
	}
}

func (p *FormulaParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, FormulaParserRULE_value)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(286)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case FormulaParserINTEGER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(274)

			var _m = p.Match(FormulaParserINTEGER)

			localctx.(*ValueContext)._INTEGER = _m
		}
		localctx.(*ValueContext).retValue = exp.NewIntegerValueExpression((func() string {
			if localctx.(*ValueContext).Get_INTEGER() == nil {
				return ""
			} else {
				return localctx.(*ValueContext).Get_INTEGER().GetText()
			}
		}()))

	case FormulaParserFLOAT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(276)

			var _m = p.Match(FormulaParserFLOAT)

			localctx.(*ValueContext)._FLOAT = _m
		}
		localctx.(*ValueContext).retValue = exp.NewFloatExpression((func() string {
			if localctx.(*ValueContext).Get_FLOAT() == nil {
				return ""
			} else {
				return localctx.(*ValueContext).Get_FLOAT().GetText()
			}
		}()))

	case FormulaParserSTRING:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(278)

			var _m = p.Match(FormulaParserSTRING)

			localctx.(*ValueContext)._STRING = _m
		}
		localctx.(*ValueContext).retValue = exp.NewStringValueExpression((func() string {
			if localctx.(*ValueContext).Get_STRING() == nil {
				return ""
			} else {
				return localctx.(*ValueContext).Get_STRING().GetText()
			}
		}()))

	case FormulaParserDATETIME:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(280)

			var _m = p.Match(FormulaParserDATETIME)

			localctx.(*ValueContext)._DATETIME = _m
		}
		localctx.(*ValueContext).retValue = exp.NewDateTimeExpression((func() string {
			if localctx.(*ValueContext).Get_DATETIME() == nil {
				return ""
			} else {
				return localctx.(*ValueContext).Get_DATETIME().GetText()
			}
		}()))

	case FormulaParserTRUE:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(282)
			p.Match(FormulaParserTRUE)
		}
		localctx.(*ValueContext).retValue = exp.NewBooleanValueExpression(true)

	case FormulaParserFALSE:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(284)
			p.Match(FormulaParserFALSE)
		}
		localctx.(*ValueContext).retValue = exp.NewBooleanValueExpression(false)

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IIdContext is an interface to support dynamic dispatch.
type IIdContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_NAME returns the _NAME token.
	Get_NAME() antlr.Token

	// Get_VAR returns the _VAR token.
	Get_VAR() antlr.Token

	// Set_NAME sets the _NAME token.
	Set_NAME(antlr.Token)

	// Set_VAR sets the _VAR token.
	Set_VAR(antlr.Token)

	// GetRetValue returns the retValue attribute.
	GetRetValue() *opt.LogicalExpression

	// SetRetValue sets the retValue attribute.
	SetRetValue(*opt.LogicalExpression)

	// IsIdContext differentiates from other interfaces.
	IsIdContext()
}

type IdContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	retValue *opt.LogicalExpression
	_NAME    antlr.Token
	_VAR     antlr.Token
}

func NewEmptyIdContext() *IdContext {
	var p = new(IdContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FormulaParserRULE_id
	return p
}

func (*IdContext) IsIdContext() {}

func NewIdContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdContext {
	var p = new(IdContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_id

	return p
}

func (s *IdContext) GetParser() antlr.Parser { return s.parser }

func (s *IdContext) Get_NAME() antlr.Token { return s._NAME }

func (s *IdContext) Get_VAR() antlr.Token { return s._VAR }

func (s *IdContext) Set_NAME(v antlr.Token) { s._NAME = v }

func (s *IdContext) Set_VAR(v antlr.Token) { s._VAR = v }

func (s *IdContext) GetRetValue() *opt.LogicalExpression { return s.retValue }

func (s *IdContext) SetRetValue(v *opt.LogicalExpression) { s.retValue = v }

func (s *IdContext) NAME() antlr.TerminalNode {
	return s.GetToken(FormulaParserNAME, 0)
}

func (s *IdContext) VAR() antlr.TerminalNode {
	return s.GetToken(FormulaParserVAR, 0)
}

func (s *IdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaListener); ok {
		listenerT.EnterId(s)
	}
}

func (s *IdContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaListener); ok {
		listenerT.ExitId(s)
	}
}

func (p *FormulaParser) Id() (localctx IIdContext) {
	localctx = NewIdContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, FormulaParserRULE_id)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(292)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case FormulaParserNAME:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(288)

			var _m = p.Match(FormulaParserNAME)

			localctx.(*IdContext)._NAME = _m
		}
		localctx.(*IdContext).retValue = opt.NewIdentifierExpression((func() string {
			if localctx.(*IdContext).Get_NAME() == nil {
				return ""
			} else {
				return localctx.(*IdContext).Get_NAME().GetText()
			}
		}()))

	case FormulaParserVAR:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(290)

			var _m = p.Match(FormulaParserVAR)

			localctx.(*IdContext)._VAR = _m
		}
		localctx.(*IdContext).retValue = opt.NewVarIdentifierExpression((func() string {
			if localctx.(*IdContext).Get_VAR() == nil {
				return ""
			} else {
				return localctx.(*IdContext).Get_VAR().GetText()
			}
		}()))

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

func (p *FormulaParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 2:
		var t *OrExprContext = nil
		if localctx != nil {
			t = localctx.(*OrExprContext)
		}
		return p.OrExpr_Sempred(t, predIndex)

	case 3:
		var t *AndExprContext = nil
		if localctx != nil {
			t = localctx.(*AndExprContext)
		}
		return p.AndExpr_Sempred(t, predIndex)

	case 4:
		var t *BitOrExprContext = nil
		if localctx != nil {
			t = localctx.(*BitOrExprContext)
		}
		return p.BitOrExpr_Sempred(t, predIndex)

	case 5:
		var t *BitXorExprContext = nil
		if localctx != nil {
			t = localctx.(*BitXorExprContext)
		}
		return p.BitXorExpr_Sempred(t, predIndex)

	case 6:
		var t *BitAndExprContext = nil
		if localctx != nil {
			t = localctx.(*BitAndExprContext)
		}
		return p.BitAndExpr_Sempred(t, predIndex)

	case 7:
		var t *EqExprContext = nil
		if localctx != nil {
			t = localctx.(*EqExprContext)
		}
		return p.EqExpr_Sempred(t, predIndex)

	case 8:
		var t *RelExprContext = nil
		if localctx != nil {
			t = localctx.(*RelExprContext)
		}
		return p.RelExpr_Sempred(t, predIndex)

	case 9:
		var t *ShiftExprContext = nil
		if localctx != nil {
			t = localctx.(*ShiftExprContext)
		}
		return p.ShiftExpr_Sempred(t, predIndex)

	case 10:
		var t *AddExprContext = nil
		if localctx != nil {
			t = localctx.(*AddExprContext)
		}
		return p.AddExpr_Sempred(t, predIndex)

	case 11:
		var t *MultExprContext = nil
		if localctx != nil {
			t = localctx.(*MultExprContext)
		}
		return p.MultExpr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *FormulaParser) OrExpr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *FormulaParser) AndExpr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *FormulaParser) BitOrExpr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 2:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *FormulaParser) BitXorExpr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 3:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *FormulaParser) BitAndExpr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 4:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *FormulaParser) EqExpr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 5:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *FormulaParser) RelExpr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 7:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 10:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *FormulaParser) ShiftExpr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 11:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 12:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *FormulaParser) AddExpr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 13:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 14:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *FormulaParser) MultExpr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 15:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 16:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 17:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
