// Code generated from Formula.g4 by ANTLR 4.9. DO NOT EDIT.

package parser // Formula

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

import (
	"github.com/yidane/formula/internal/exp"
	"github.com/yidane/formula/opt"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 43, 304,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 5, 3, 51, 10, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 7, 4, 62, 10, 4, 12, 4, 14, 4, 65, 11, 4, 3, 5, 3, 5, 3,
	5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 7, 5, 76, 10, 5, 12, 5, 14, 5, 79,
	11, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 7, 6, 90,
	10, 6, 12, 6, 14, 6, 93, 11, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3,
	7, 3, 7, 3, 7, 7, 7, 104, 10, 7, 12, 7, 14, 7, 107, 11, 7, 3, 8, 3, 8,
	3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 7, 8, 118, 10, 8, 12, 8, 14,
	8, 121, 11, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3,
	9, 3, 9, 3, 9, 3, 9, 3, 9, 7, 9, 137, 10, 9, 12, 9, 14, 9, 140, 11, 9,
	3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3,
	10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10,
	3, 10, 3, 10, 3, 10, 7, 10, 166, 10, 10, 12, 10, 14, 10, 169, 11, 10, 3,
	11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11,
	3, 11, 3, 11, 3, 11, 7, 11, 185, 10, 11, 12, 11, 14, 11, 188, 11, 11, 3,
	12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12,
	3, 12, 3, 12, 3, 12, 7, 12, 204, 10, 12, 12, 12, 14, 12, 207, 11, 12, 3,
	13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13,
	3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 7, 13, 228, 10,
	13, 12, 13, 14, 13, 231, 11, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3,
	14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 5, 14,
	248, 10, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3,
	15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 7, 15, 266, 10, 15,
	12, 15, 14, 15, 269, 11, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15,
	5, 15, 277, 10, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3,
	16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 5, 16, 294, 10, 16,
	3, 17, 3, 17, 3, 17, 3, 17, 5, 17, 300, 10, 17, 3, 18, 3, 18, 3, 18, 2,
	12, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 19, 2, 4, 6, 8, 10, 12, 14, 16,
	18, 20, 22, 24, 26, 28, 30, 32, 34, 2, 7, 3, 2, 5, 6, 3, 2, 7, 8, 3, 2,
	12, 13, 3, 2, 14, 15, 3, 2, 27, 28, 2, 319, 2, 36, 3, 2, 2, 2, 4, 50, 3,
	2, 2, 2, 6, 52, 3, 2, 2, 2, 8, 66, 3, 2, 2, 2, 10, 80, 3, 2, 2, 2, 12,
	94, 3, 2, 2, 2, 14, 108, 3, 2, 2, 2, 16, 122, 3, 2, 2, 2, 18, 141, 3, 2,
	2, 2, 20, 170, 3, 2, 2, 2, 22, 189, 3, 2, 2, 2, 24, 208, 3, 2, 2, 2, 26,
	247, 3, 2, 2, 2, 28, 276, 3, 2, 2, 2, 30, 293, 3, 2, 2, 2, 32, 299, 3,
	2, 2, 2, 34, 301, 3, 2, 2, 2, 36, 37, 5, 4, 3, 2, 37, 38, 7, 2, 2, 3, 38,
	39, 8, 2, 1, 2, 39, 3, 3, 2, 2, 2, 40, 41, 5, 6, 4, 2, 41, 42, 7, 3, 2,
	2, 42, 43, 5, 4, 3, 2, 43, 44, 7, 4, 2, 2, 44, 45, 5, 4, 3, 2, 45, 46,
	8, 3, 1, 2, 46, 51, 3, 2, 2, 2, 47, 48, 5, 6, 4, 2, 48, 49, 8, 3, 1, 2,
	49, 51, 3, 2, 2, 2, 50, 40, 3, 2, 2, 2, 50, 47, 3, 2, 2, 2, 51, 5, 3, 2,
	2, 2, 52, 53, 8, 4, 1, 2, 53, 54, 5, 8, 5, 2, 54, 55, 8, 4, 1, 2, 55, 63,
	3, 2, 2, 2, 56, 57, 12, 4, 2, 2, 57, 58, 9, 2, 2, 2, 58, 59, 5, 8, 5, 2,
	59, 60, 8, 4, 1, 2, 60, 62, 3, 2, 2, 2, 61, 56, 3, 2, 2, 2, 62, 65, 3,
	2, 2, 2, 63, 61, 3, 2, 2, 2, 63, 64, 3, 2, 2, 2, 64, 7, 3, 2, 2, 2, 65,
	63, 3, 2, 2, 2, 66, 67, 8, 5, 1, 2, 67, 68, 5, 10, 6, 2, 68, 69, 8, 5,
	1, 2, 69, 77, 3, 2, 2, 2, 70, 71, 12, 4, 2, 2, 71, 72, 9, 3, 2, 2, 72,
	73, 5, 10, 6, 2, 73, 74, 8, 5, 1, 2, 74, 76, 3, 2, 2, 2, 75, 70, 3, 2,
	2, 2, 76, 79, 3, 2, 2, 2, 77, 75, 3, 2, 2, 2, 77, 78, 3, 2, 2, 2, 78, 9,
	3, 2, 2, 2, 79, 77, 3, 2, 2, 2, 80, 81, 8, 6, 1, 2, 81, 82, 5, 12, 7, 2,
	82, 83, 8, 6, 1, 2, 83, 91, 3, 2, 2, 2, 84, 85, 12, 4, 2, 2, 85, 86, 7,
	9, 2, 2, 86, 87, 5, 12, 7, 2, 87, 88, 8, 6, 1, 2, 88, 90, 3, 2, 2, 2, 89,
	84, 3, 2, 2, 2, 90, 93, 3, 2, 2, 2, 91, 89, 3, 2, 2, 2, 91, 92, 3, 2, 2,
	2, 92, 11, 3, 2, 2, 2, 93, 91, 3, 2, 2, 2, 94, 95, 8, 7, 1, 2, 95, 96,
	5, 14, 8, 2, 96, 97, 8, 7, 1, 2, 97, 105, 3, 2, 2, 2, 98, 99, 12, 4, 2,
	2, 99, 100, 7, 10, 2, 2, 100, 101, 5, 14, 8, 2, 101, 102, 8, 7, 1, 2, 102,
	104, 3, 2, 2, 2, 103, 98, 3, 2, 2, 2, 104, 107, 3, 2, 2, 2, 105, 103, 3,
	2, 2, 2, 105, 106, 3, 2, 2, 2, 106, 13, 3, 2, 2, 2, 107, 105, 3, 2, 2,
	2, 108, 109, 8, 8, 1, 2, 109, 110, 5, 16, 9, 2, 110, 111, 8, 8, 1, 2, 111,
	119, 3, 2, 2, 2, 112, 113, 12, 4, 2, 2, 113, 114, 7, 11, 2, 2, 114, 115,
	5, 16, 9, 2, 115, 116, 8, 8, 1, 2, 116, 118, 3, 2, 2, 2, 117, 112, 3, 2,
	2, 2, 118, 121, 3, 2, 2, 2, 119, 117, 3, 2, 2, 2, 119, 120, 3, 2, 2, 2,
	120, 15, 3, 2, 2, 2, 121, 119, 3, 2, 2, 2, 122, 123, 8, 9, 1, 2, 123, 124,
	5, 18, 10, 2, 124, 125, 8, 9, 1, 2, 125, 138, 3, 2, 2, 2, 126, 127, 12,
	5, 2, 2, 127, 128, 9, 4, 2, 2, 128, 129, 5, 18, 10, 2, 129, 130, 8, 9,
	1, 2, 130, 137, 3, 2, 2, 2, 131, 132, 12, 4, 2, 2, 132, 133, 9, 5, 2, 2,
	133, 134, 5, 18, 10, 2, 134, 135, 8, 9, 1, 2, 135, 137, 3, 2, 2, 2, 136,
	126, 3, 2, 2, 2, 136, 131, 3, 2, 2, 2, 137, 140, 3, 2, 2, 2, 138, 136,
	3, 2, 2, 2, 138, 139, 3, 2, 2, 2, 139, 17, 3, 2, 2, 2, 140, 138, 3, 2,
	2, 2, 141, 142, 8, 10, 1, 2, 142, 143, 5, 20, 11, 2, 143, 144, 8, 10, 1,
	2, 144, 167, 3, 2, 2, 2, 145, 146, 12, 7, 2, 2, 146, 147, 7, 16, 2, 2,
	147, 148, 5, 20, 11, 2, 148, 149, 8, 10, 1, 2, 149, 166, 3, 2, 2, 2, 150,
	151, 12, 6, 2, 2, 151, 152, 7, 17, 2, 2, 152, 153, 5, 20, 11, 2, 153, 154,
	8, 10, 1, 2, 154, 166, 3, 2, 2, 2, 155, 156, 12, 5, 2, 2, 156, 157, 7,
	18, 2, 2, 157, 158, 5, 20, 11, 2, 158, 159, 8, 10, 1, 2, 159, 166, 3, 2,
	2, 2, 160, 161, 12, 4, 2, 2, 161, 162, 7, 19, 2, 2, 162, 163, 5, 20, 11,
	2, 163, 164, 8, 10, 1, 2, 164, 166, 3, 2, 2, 2, 165, 145, 3, 2, 2, 2, 165,
	150, 3, 2, 2, 2, 165, 155, 3, 2, 2, 2, 165, 160, 3, 2, 2, 2, 166, 169,
	3, 2, 2, 2, 167, 165, 3, 2, 2, 2, 167, 168, 3, 2, 2, 2, 168, 19, 3, 2,
	2, 2, 169, 167, 3, 2, 2, 2, 170, 171, 8, 11, 1, 2, 171, 172, 5, 22, 12,
	2, 172, 173, 8, 11, 1, 2, 173, 186, 3, 2, 2, 2, 174, 175, 12, 5, 2, 2,
	175, 176, 7, 20, 2, 2, 176, 177, 5, 22, 12, 2, 177, 178, 8, 11, 1, 2, 178,
	185, 3, 2, 2, 2, 179, 180, 12, 4, 2, 2, 180, 181, 7, 21, 2, 2, 181, 182,
	5, 22, 12, 2, 182, 183, 8, 11, 1, 2, 183, 185, 3, 2, 2, 2, 184, 174, 3,
	2, 2, 2, 184, 179, 3, 2, 2, 2, 185, 188, 3, 2, 2, 2, 186, 184, 3, 2, 2,
	2, 186, 187, 3, 2, 2, 2, 187, 21, 3, 2, 2, 2, 188, 186, 3, 2, 2, 2, 189,
	190, 8, 12, 1, 2, 190, 191, 5, 24, 13, 2, 191, 192, 8, 12, 1, 2, 192, 205,
	3, 2, 2, 2, 193, 194, 12, 5, 2, 2, 194, 195, 7, 22, 2, 2, 195, 196, 5,
	24, 13, 2, 196, 197, 8, 12, 1, 2, 197, 204, 3, 2, 2, 2, 198, 199, 12, 4,
	2, 2, 199, 200, 7, 23, 2, 2, 200, 201, 5, 24, 13, 2, 201, 202, 8, 12, 1,
	2, 202, 204, 3, 2, 2, 2, 203, 193, 3, 2, 2, 2, 203, 198, 3, 2, 2, 2, 204,
	207, 3, 2, 2, 2, 205, 203, 3, 2, 2, 2, 205, 206, 3, 2, 2, 2, 206, 23, 3,
	2, 2, 2, 207, 205, 3, 2, 2, 2, 208, 209, 8, 13, 1, 2, 209, 210, 5, 26,
	14, 2, 210, 211, 8, 13, 1, 2, 211, 229, 3, 2, 2, 2, 212, 213, 12, 6, 2,
	2, 213, 214, 7, 24, 2, 2, 214, 215, 5, 26, 14, 2, 215, 216, 8, 13, 1, 2,
	216, 228, 3, 2, 2, 2, 217, 218, 12, 5, 2, 2, 218, 219, 7, 25, 2, 2, 219,
	220, 5, 26, 14, 2, 220, 221, 8, 13, 1, 2, 221, 228, 3, 2, 2, 2, 222, 223,
	12, 4, 2, 2, 223, 224, 7, 26, 2, 2, 224, 225, 5, 26, 14, 2, 225, 226, 8,
	13, 1, 2, 226, 228, 3, 2, 2, 2, 227, 212, 3, 2, 2, 2, 227, 217, 3, 2, 2,
	2, 227, 222, 3, 2, 2, 2, 228, 231, 3, 2, 2, 2, 229, 227, 3, 2, 2, 2, 229,
	230, 3, 2, 2, 2, 230, 25, 3, 2, 2, 2, 231, 229, 3, 2, 2, 2, 232, 233, 9,
	6, 2, 2, 233, 234, 5, 28, 15, 2, 234, 235, 8, 14, 1, 2, 235, 248, 3, 2,
	2, 2, 236, 237, 7, 29, 2, 2, 237, 238, 5, 28, 15, 2, 238, 239, 8, 14, 1,
	2, 239, 248, 3, 2, 2, 2, 240, 241, 7, 23, 2, 2, 241, 242, 5, 28, 15, 2,
	242, 243, 8, 14, 1, 2, 243, 248, 3, 2, 2, 2, 244, 245, 5, 28, 15, 2, 245,
	246, 8, 14, 1, 2, 246, 248, 3, 2, 2, 2, 247, 232, 3, 2, 2, 2, 247, 236,
	3, 2, 2, 2, 247, 240, 3, 2, 2, 2, 247, 244, 3, 2, 2, 2, 248, 27, 3, 2,
	2, 2, 249, 250, 7, 30, 2, 2, 250, 251, 5, 4, 3, 2, 251, 252, 7, 31, 2,
	2, 252, 253, 8, 15, 1, 2, 253, 277, 3, 2, 2, 2, 254, 255, 5, 30, 16, 2,
	255, 256, 8, 15, 1, 2, 256, 277, 3, 2, 2, 2, 257, 258, 5, 32, 17, 2, 258,
	259, 7, 30, 2, 2, 259, 260, 5, 4, 3, 2, 260, 267, 8, 15, 1, 2, 261, 262,
	7, 32, 2, 2, 262, 263, 5, 4, 3, 2, 263, 264, 8, 15, 1, 2, 264, 266, 3,
	2, 2, 2, 265, 261, 3, 2, 2, 2, 266, 269, 3, 2, 2, 2, 267, 265, 3, 2, 2,
	2, 267, 268, 3, 2, 2, 2, 268, 270, 3, 2, 2, 2, 269, 267, 3, 2, 2, 2, 270,
	271, 7, 31, 2, 2, 271, 272, 8, 15, 1, 2, 272, 277, 3, 2, 2, 2, 273, 274,
	5, 32, 17, 2, 274, 275, 8, 15, 1, 2, 275, 277, 3, 2, 2, 2, 276, 249, 3,
	2, 2, 2, 276, 254, 3, 2, 2, 2, 276, 257, 3, 2, 2, 2, 276, 273, 3, 2, 2,
	2, 277, 29, 3, 2, 2, 2, 278, 279, 7, 37, 2, 2, 279, 294, 8, 16, 1, 2, 280,
	281, 7, 41, 2, 2, 281, 294, 8, 16, 1, 2, 282, 283, 7, 42, 2, 2, 283, 294,
	8, 16, 1, 2, 284, 285, 7, 38, 2, 2, 285, 294, 8, 16, 1, 2, 286, 287, 7,
	34, 2, 2, 287, 294, 8, 16, 1, 2, 288, 289, 7, 35, 2, 2, 289, 294, 8, 16,
	1, 2, 290, 291, 5, 34, 18, 2, 291, 292, 8, 16, 1, 2, 292, 294, 3, 2, 2,
	2, 293, 278, 3, 2, 2, 2, 293, 280, 3, 2, 2, 2, 293, 282, 3, 2, 2, 2, 293,
	284, 3, 2, 2, 2, 293, 286, 3, 2, 2, 2, 293, 288, 3, 2, 2, 2, 293, 290,
	3, 2, 2, 2, 294, 31, 3, 2, 2, 2, 295, 296, 7, 36, 2, 2, 296, 300, 8, 17,
	1, 2, 297, 298, 7, 39, 2, 2, 298, 300, 8, 17, 1, 2, 299, 295, 3, 2, 2,
	2, 299, 297, 3, 2, 2, 2, 300, 33, 3, 2, 2, 2, 301, 302, 7, 33, 2, 2, 302,
	35, 3, 2, 2, 2, 23, 50, 63, 77, 91, 105, 119, 136, 138, 165, 167, 184,
	186, 203, 205, 227, 229, 247, 267, 276, 293, 299,
}
var literalNames = []string{
	"", "'?'", "':'", "'||'", "'or'", "'&&'", "'and'", "'|'", "'^'", "'&'",
	"'=='", "'='", "'!='", "'<>'", "'<'", "'<='", "'>'", "'>='", "'<<'", "'>>'",
	"'+'", "'-'", "'*'", "'/'", "'%'", "'!'", "'not'", "'~'", "'('", "')'",
	"','", "'\u03C0'", "'true'", "'false'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "TRUE", "FALSE",
	"NAME", "INTEGER", "DATETIME", "VAR", "E", "FLOAT", "STRING", "WS",
}

var ruleNames = []string{
	"calc", "expr", "orExpr", "andExpr", "bitOrExpr", "bitXorExpr", "bitAndExpr",
	"eqExpr", "relExpr", "shiftExpr", "addExpr", "multExpr", "unaryExpr", "primaryExpr",
	"value", "id", "π",
}

type FormulaParser struct {
	*antlr.BaseParser
}

// NewFormulaParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *FormulaParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewFormulaParser(input antlr.TokenStream) *FormulaParser {
	this := new(FormulaParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
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
	FormulaParserT__30    = 31
	FormulaParserTRUE     = 32
	FormulaParserFALSE    = 33
	FormulaParserNAME     = 34
	FormulaParserINTEGER  = 35
	FormulaParserDATETIME = 36
	FormulaParserVAR      = 37
	FormulaParserE        = 38
	FormulaParserFLOAT    = 39
	FormulaParserSTRING   = 40
	FormulaParserWS       = 41
)

// FormulaParser rules.
const (
	FormulaParserRULE_calc        = 0
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
	FormulaParserRULE_π           = 16
)

// ICalcContext is an interface to support dynamic dispatch.
type ICalcContext interface {
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

	// IsCalcContext differentiates from other interfaces.
	IsCalcContext()
}

type CalcContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	retValue *opt.LogicalExpression
	_expr    IExprContext
}

func NewEmptyCalcContext() *CalcContext {
	var p = new(CalcContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FormulaParserRULE_calc
	return p
}

func (*CalcContext) IsCalcContext() {}

func NewCalcContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CalcContext {
	var p = new(CalcContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_calc

	return p
}

func (s *CalcContext) GetParser() antlr.Parser { return s.parser }

func (s *CalcContext) Get_expr() IExprContext { return s._expr }

func (s *CalcContext) Set_expr(v IExprContext) { s._expr = v }

func (s *CalcContext) GetRetValue() *opt.LogicalExpression { return s.retValue }

func (s *CalcContext) SetRetValue(v *opt.LogicalExpression) { s.retValue = v }

func (s *CalcContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *CalcContext) EOF() antlr.TerminalNode {
	return s.GetToken(FormulaParserEOF, 0)
}

func (s *CalcContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CalcContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CalcContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaListener); ok {
		listenerT.EnterCalc(s)
	}
}

func (s *CalcContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaListener); ok {
		listenerT.ExitCalc(s)
	}
}

func (p *FormulaParser) Calc() (localctx ICalcContext) {
	localctx = NewCalcContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, FormulaParserRULE_calc)

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
		p.SetState(34)

		var _x = p.Expr()

		localctx.(*CalcContext)._expr = _x
	}
	{
		p.SetState(35)
		p.Match(FormulaParserEOF)
	}
	localctx.(*CalcContext).retValue = localctx.(*CalcContext).Get_expr().GetRetValue()

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

	p.SetState(48)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(38)

			var _x = p.orExpr(0)

			localctx.(*ExprContext).first = _x
		}
		{
			p.SetState(39)
			p.Match(FormulaParserT__0)
		}
		{
			p.SetState(40)

			var _x = p.Expr()

			localctx.(*ExprContext).middle = _x
		}
		{
			p.SetState(41)
			p.Match(FormulaParserT__1)
		}
		{
			p.SetState(42)

			var _x = p.Expr()

			localctx.(*ExprContext).right = _x
		}
		localctx.(*ExprContext).retValue = exp.NewTernaryExpression(localctx.(*ExprContext).GetFirst().GetRetValue(), localctx.(*ExprContext).GetMiddle().GetRetValue(), localctx.(*ExprContext).GetRight().GetRetValue())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(45)

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
		p.SetState(51)

		var _x = p.andExpr(0)

		localctx.(*OrExprContext)._andExpr = _x
	}
	localctx.(*OrExprContext).retValue = localctx.(*OrExprContext).Get_andExpr().GetRetValue()

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(61)
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
			p.SetState(54)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(55)
				_la = p.GetTokenStream().LA(1)

				if !(_la == FormulaParserT__2 || _la == FormulaParserT__3) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(56)

				var _x = p.andExpr(0)

				localctx.(*OrExprContext)._andExpr = _x
			}
			localctx.(*OrExprContext).retValue = exp.NewBinaryExpression("||", localctx.(*OrExprContext).GetFirst().GetRetValue(), localctx.(*OrExprContext).Get_andExpr().GetRetValue())

		}
		p.SetState(63)
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
		p.SetState(65)

		var _x = p.bitOrExpr(0)

		localctx.(*AndExprContext)._bitOrExpr = _x
	}
	localctx.(*AndExprContext).retValue = localctx.(*AndExprContext).Get_bitOrExpr().GetRetValue()

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(75)
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
			p.SetState(68)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(69)
				_la = p.GetTokenStream().LA(1)

				if !(_la == FormulaParserT__4 || _la == FormulaParserT__5) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(70)

				var _x = p.bitOrExpr(0)

				localctx.(*AndExprContext)._bitOrExpr = _x
			}
			localctx.(*AndExprContext).retValue = exp.NewBinaryExpression("&&", localctx.(*AndExprContext).GetFirst().GetRetValue(), localctx.(*AndExprContext).Get_bitOrExpr().GetRetValue())

		}
		p.SetState(77)
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
		p.SetState(79)

		var _x = p.bitXorExpr(0)

		localctx.(*BitOrExprContext)._bitXorExpr = _x
	}
	localctx.(*BitOrExprContext).retValue = localctx.(*BitOrExprContext).Get_bitXorExpr().GetRetValue()

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(89)
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
			p.SetState(82)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(83)
				p.Match(FormulaParserT__6)
			}
			{
				p.SetState(84)

				var _x = p.bitXorExpr(0)

				localctx.(*BitOrExprContext)._bitXorExpr = _x
			}
			localctx.(*BitOrExprContext).retValue = exp.NewBinaryExpression("|", localctx.(*BitOrExprContext).GetFirst().GetRetValue(), localctx.(*BitOrExprContext).Get_bitXorExpr().GetRetValue())

		}
		p.SetState(91)
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
		p.SetState(93)

		var _x = p.bitAndExpr(0)

		localctx.(*BitXorExprContext)._bitAndExpr = _x
	}
	localctx.(*BitXorExprContext).retValue = localctx.(*BitXorExprContext).Get_bitAndExpr().GetRetValue()

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(103)
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
			p.SetState(96)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(97)
				p.Match(FormulaParserT__7)
			}
			{
				p.SetState(98)

				var _x = p.bitAndExpr(0)

				localctx.(*BitXorExprContext)._bitAndExpr = _x
			}
			localctx.(*BitXorExprContext).retValue = exp.NewBinaryExpression("^", localctx.(*BitXorExprContext).GetFirst().GetRetValue(), localctx.(*BitXorExprContext).Get_bitAndExpr().GetRetValue())

		}
		p.SetState(105)
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
		p.SetState(107)

		var _x = p.eqExpr(0)

		localctx.(*BitAndExprContext)._eqExpr = _x
	}
	localctx.(*BitAndExprContext).retValue = localctx.(*BitAndExprContext).Get_eqExpr().GetRetValue()

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(117)
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
			p.SetState(110)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(111)
				p.Match(FormulaParserT__8)
			}
			{
				p.SetState(112)

				var _x = p.eqExpr(0)

				localctx.(*BitAndExprContext)._eqExpr = _x
			}
			localctx.(*BitAndExprContext).retValue = exp.NewBinaryExpression("&", localctx.(*BitAndExprContext).GetFirst().GetRetValue(), localctx.(*BitAndExprContext).Get_eqExpr().GetRetValue())

		}
		p.SetState(119)
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
		p.SetState(121)

		var _x = p.relExpr(0)

		localctx.(*EqExprContext)._relExpr = _x
	}
	localctx.(*EqExprContext).retValue = localctx.(*EqExprContext).Get_relExpr().GetRetValue()

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(136)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(134)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
			case 1:
				localctx = NewEqExprContext(p, _parentctx, _parentState)
				localctx.(*EqExprContext).first = _prevctx
				p.PushNewRecursionContext(localctx, _startState, FormulaParserRULE_eqExpr)
				p.SetState(124)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(125)
					_la = p.GetTokenStream().LA(1)

					if !(_la == FormulaParserT__9 || _la == FormulaParserT__10) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(126)

					var _x = p.relExpr(0)

					localctx.(*EqExprContext)._relExpr = _x
				}
				localctx.(*EqExprContext).retValue = exp.NewBinaryExpression("==", localctx.(*EqExprContext).GetFirst().GetRetValue(), localctx.(*EqExprContext).Get_relExpr().GetRetValue())

			case 2:
				localctx = NewEqExprContext(p, _parentctx, _parentState)
				localctx.(*EqExprContext).first = _prevctx
				p.PushNewRecursionContext(localctx, _startState, FormulaParserRULE_eqExpr)
				p.SetState(129)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(130)
					_la = p.GetTokenStream().LA(1)

					if !(_la == FormulaParserT__11 || _la == FormulaParserT__12) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(131)

					var _x = p.relExpr(0)

					localctx.(*EqExprContext)._relExpr = _x
				}
				localctx.(*EqExprContext).retValue = exp.NewBinaryExpression("!=", localctx.(*EqExprContext).GetFirst().GetRetValue(), localctx.(*EqExprContext).Get_relExpr().GetRetValue())

			}

		}
		p.SetState(138)
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
		p.SetState(140)

		var _x = p.shiftExpr(0)

		localctx.(*RelExprContext)._shiftExpr = _x
	}
	localctx.(*RelExprContext).retValue = localctx.(*RelExprContext).Get_shiftExpr().GetRetValue()

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(165)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(163)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
			case 1:
				localctx = NewRelExprContext(p, _parentctx, _parentState)
				localctx.(*RelExprContext).first = _prevctx
				p.PushNewRecursionContext(localctx, _startState, FormulaParserRULE_relExpr)
				p.SetState(143)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(144)
					p.Match(FormulaParserT__13)
				}
				{
					p.SetState(145)

					var _x = p.shiftExpr(0)

					localctx.(*RelExprContext)._shiftExpr = _x
				}
				localctx.(*RelExprContext).retValue = exp.NewBinaryExpression("<", localctx.(*RelExprContext).GetFirst().GetRetValue(), localctx.(*RelExprContext).Get_shiftExpr().GetRetValue())

			case 2:
				localctx = NewRelExprContext(p, _parentctx, _parentState)
				localctx.(*RelExprContext).first = _prevctx
				p.PushNewRecursionContext(localctx, _startState, FormulaParserRULE_relExpr)
				p.SetState(148)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(149)
					p.Match(FormulaParserT__14)
				}
				{
					p.SetState(150)

					var _x = p.shiftExpr(0)

					localctx.(*RelExprContext)._shiftExpr = _x
				}
				localctx.(*RelExprContext).retValue = exp.NewBinaryExpression("<=", localctx.(*RelExprContext).GetFirst().GetRetValue(), localctx.(*RelExprContext).Get_shiftExpr().GetRetValue())

			case 3:
				localctx = NewRelExprContext(p, _parentctx, _parentState)
				localctx.(*RelExprContext).first = _prevctx
				p.PushNewRecursionContext(localctx, _startState, FormulaParserRULE_relExpr)
				p.SetState(153)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(154)
					p.Match(FormulaParserT__15)
				}
				{
					p.SetState(155)

					var _x = p.shiftExpr(0)

					localctx.(*RelExprContext)._shiftExpr = _x
				}
				localctx.(*RelExprContext).retValue = exp.NewBinaryExpression(">", localctx.(*RelExprContext).GetFirst().GetRetValue(), localctx.(*RelExprContext).Get_shiftExpr().GetRetValue())

			case 4:
				localctx = NewRelExprContext(p, _parentctx, _parentState)
				localctx.(*RelExprContext).first = _prevctx
				p.PushNewRecursionContext(localctx, _startState, FormulaParserRULE_relExpr)
				p.SetState(158)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(159)
					p.Match(FormulaParserT__16)
				}
				{
					p.SetState(160)

					var _x = p.shiftExpr(0)

					localctx.(*RelExprContext)._shiftExpr = _x
				}
				localctx.(*RelExprContext).retValue = exp.NewBinaryExpression(">=", localctx.(*RelExprContext).GetFirst().GetRetValue(), localctx.(*RelExprContext).Get_shiftExpr().GetRetValue())

			}

		}
		p.SetState(167)
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
		p.SetState(169)

		var _x = p.addExpr(0)

		localctx.(*ShiftExprContext)._addExpr = _x
	}
	localctx.(*ShiftExprContext).retValue = localctx.(*ShiftExprContext).Get_addExpr().GetRetValue()

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(184)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(182)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
			case 1:
				localctx = NewShiftExprContext(p, _parentctx, _parentState)
				localctx.(*ShiftExprContext).first = _prevctx
				p.PushNewRecursionContext(localctx, _startState, FormulaParserRULE_shiftExpr)
				p.SetState(172)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(173)
					p.Match(FormulaParserT__17)
				}
				{
					p.SetState(174)

					var _x = p.addExpr(0)

					localctx.(*ShiftExprContext)._addExpr = _x
				}
				localctx.(*ShiftExprContext).retValue = exp.NewBinaryExpression("<<", localctx.(*ShiftExprContext).GetFirst().GetRetValue(), localctx.(*ShiftExprContext).Get_addExpr().GetRetValue())

			case 2:
				localctx = NewShiftExprContext(p, _parentctx, _parentState)
				localctx.(*ShiftExprContext).first = _prevctx
				p.PushNewRecursionContext(localctx, _startState, FormulaParserRULE_shiftExpr)
				p.SetState(177)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(178)
					p.Match(FormulaParserT__18)
				}
				{
					p.SetState(179)

					var _x = p.addExpr(0)

					localctx.(*ShiftExprContext)._addExpr = _x
				}
				localctx.(*ShiftExprContext).retValue = exp.NewBinaryExpression(">>", localctx.(*ShiftExprContext).GetFirst().GetRetValue(), localctx.(*ShiftExprContext).Get_addExpr().GetRetValue())

			}

		}
		p.SetState(186)
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
		p.SetState(188)

		var _x = p.multExpr(0)

		localctx.(*AddExprContext)._multExpr = _x
	}
	localctx.(*AddExprContext).retValue = localctx.(*AddExprContext).Get_multExpr().GetRetValue()

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(203)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(201)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) {
			case 1:
				localctx = NewAddExprContext(p, _parentctx, _parentState)
				localctx.(*AddExprContext).first = _prevctx
				p.PushNewRecursionContext(localctx, _startState, FormulaParserRULE_addExpr)
				p.SetState(191)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(192)
					p.Match(FormulaParserT__19)
				}
				{
					p.SetState(193)

					var _x = p.multExpr(0)

					localctx.(*AddExprContext)._multExpr = _x
				}
				localctx.(*AddExprContext).retValue = exp.NewBinaryExpression("+", localctx.(*AddExprContext).GetFirst().GetRetValue(), localctx.(*AddExprContext).Get_multExpr().GetRetValue())

			case 2:
				localctx = NewAddExprContext(p, _parentctx, _parentState)
				localctx.(*AddExprContext).first = _prevctx
				p.PushNewRecursionContext(localctx, _startState, FormulaParserRULE_addExpr)
				p.SetState(196)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(197)
					p.Match(FormulaParserT__20)
				}
				{
					p.SetState(198)

					var _x = p.multExpr(0)

					localctx.(*AddExprContext)._multExpr = _x
				}
				localctx.(*AddExprContext).retValue = exp.NewBinaryExpression("-", localctx.(*AddExprContext).GetFirst().GetRetValue(), localctx.(*AddExprContext).Get_multExpr().GetRetValue())

			}

		}
		p.SetState(205)
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
		p.SetState(207)

		var _x = p.UnaryExpr()

		localctx.(*MultExprContext)._unaryExpr = _x
	}
	localctx.(*MultExprContext).retValue = localctx.(*MultExprContext).Get_unaryExpr().GetRetValue()

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(227)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(225)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMultExprContext(p, _parentctx, _parentState)
				localctx.(*MultExprContext).first = _prevctx
				p.PushNewRecursionContext(localctx, _startState, FormulaParserRULE_multExpr)
				p.SetState(210)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(211)
					p.Match(FormulaParserT__21)
				}
				{
					p.SetState(212)

					var _x = p.UnaryExpr()

					localctx.(*MultExprContext)._unaryExpr = _x
				}
				localctx.(*MultExprContext).retValue = exp.NewBinaryExpression("*", localctx.(*MultExprContext).GetFirst().GetRetValue(), localctx.(*MultExprContext).Get_unaryExpr().GetRetValue())

			case 2:
				localctx = NewMultExprContext(p, _parentctx, _parentState)
				localctx.(*MultExprContext).first = _prevctx
				p.PushNewRecursionContext(localctx, _startState, FormulaParserRULE_multExpr)
				p.SetState(215)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(216)
					p.Match(FormulaParserT__22)
				}
				{
					p.SetState(217)

					var _x = p.UnaryExpr()

					localctx.(*MultExprContext)._unaryExpr = _x
				}
				localctx.(*MultExprContext).retValue = exp.NewBinaryExpression("/", localctx.(*MultExprContext).GetFirst().GetRetValue(), localctx.(*MultExprContext).Get_unaryExpr().GetRetValue())

			case 3:
				localctx = NewMultExprContext(p, _parentctx, _parentState)
				localctx.(*MultExprContext).first = _prevctx
				p.PushNewRecursionContext(localctx, _startState, FormulaParserRULE_multExpr)
				p.SetState(220)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(221)
					p.Match(FormulaParserT__23)
				}
				{
					p.SetState(222)

					var _x = p.UnaryExpr()

					localctx.(*MultExprContext)._unaryExpr = _x
				}
				localctx.(*MultExprContext).retValue = exp.NewBinaryExpression("%", localctx.(*MultExprContext).GetFirst().GetRetValue(), localctx.(*MultExprContext).Get_unaryExpr().GetRetValue())

			}

		}
		p.SetState(229)
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

	p.SetState(245)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case FormulaParserT__24, FormulaParserT__25:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(230)
			_la = p.GetTokenStream().LA(1)

			if !(_la == FormulaParserT__24 || _la == FormulaParserT__25) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(231)

			var _x = p.PrimaryExpr()

			localctx.(*UnaryExprContext)._primaryExpr = _x
		}
		localctx.(*UnaryExprContext).retValue = exp.NewNotUnaryExpression("!", localctx.(*UnaryExprContext).Get_primaryExpr().GetRetValue())

	case FormulaParserT__26:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(234)
			p.Match(FormulaParserT__26)
		}
		{
			p.SetState(235)

			var _x = p.PrimaryExpr()

			localctx.(*UnaryExprContext)._primaryExpr = _x
		}
		localctx.(*UnaryExprContext).retValue = exp.NewBitwiseNotUnaryExpression("~", localctx.(*UnaryExprContext).Get_primaryExpr().GetRetValue())

	case FormulaParserT__20:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(238)
			p.Match(FormulaParserT__20)
		}
		{
			p.SetState(239)

			var _x = p.PrimaryExpr()

			localctx.(*UnaryExprContext)._primaryExpr = _x
		}
		localctx.(*UnaryExprContext).retValue = exp.NewNegateUnaryExpression("-", localctx.(*UnaryExprContext).Get_primaryExpr().GetRetValue())

	case FormulaParserT__27, FormulaParserT__30, FormulaParserTRUE, FormulaParserFALSE, FormulaParserNAME, FormulaParserINTEGER, FormulaParserDATETIME, FormulaParserVAR, FormulaParserFLOAT, FormulaParserSTRING:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(242)

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

	p.SetState(274)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(247)
			p.Match(FormulaParserT__27)
		}
		{
			p.SetState(248)

			var _x = p.Expr()

			localctx.(*PrimaryExprContext)._expr = _x
		}
		{
			p.SetState(249)
			p.Match(FormulaParserT__28)
		}
		localctx.(*PrimaryExprContext).retValue = localctx.(*PrimaryExprContext).Get_expr().GetRetValue()

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(252)

			var _x = p.Value()

			localctx.(*PrimaryExprContext)._value = _x
		}
		localctx.(*PrimaryExprContext).retValue = localctx.(*PrimaryExprContext).Get_value().GetRetValue()

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(255)

			var _x = p.Id()

			localctx.(*PrimaryExprContext)._id = _x
		}
		{
			p.SetState(256)
			p.Match(FormulaParserT__27)
		}
		{
			p.SetState(257)

			var _x = p.Expr()

			localctx.(*PrimaryExprContext)._expr = _x
		}
		args = append(args, localctx.(*PrimaryExprContext).Get_expr().GetRetValue())
		p.SetState(265)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == FormulaParserT__29 {
			{
				p.SetState(259)
				p.Match(FormulaParserT__29)
			}
			{
				p.SetState(260)

				var _x = p.Expr()

				localctx.(*PrimaryExprContext)._expr = _x
			}
			args = append(args, localctx.(*PrimaryExprContext).Get_expr().GetRetValue())

			p.SetState(267)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(268)
			p.Match(FormulaParserT__28)
		}
		localctx.(*PrimaryExprContext).retValue = exp.NewFunctionExpression(exp.NewIdentifierExpression((func() string {
			if localctx.(*PrimaryExprContext).Get_id() == nil {
				return ""
			} else {
				return p.GetTokenStream().GetTextFromTokens(localctx.(*PrimaryExprContext).Get_id().GetStart(), localctx.(*PrimaryExprContext)._id.GetStop())
			}
		}())), args)

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(271)

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

func (s *ValueContext) Π() IΠContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IΠContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IΠContext)
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

	p.SetState(291)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case FormulaParserINTEGER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(276)

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
			p.SetState(278)

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
			p.SetState(280)

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
			p.SetState(282)

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
			p.SetState(284)
			p.Match(FormulaParserTRUE)
		}
		localctx.(*ValueContext).retValue = exp.NewBooleanValueExpression(true)

	case FormulaParserFALSE:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(286)
			p.Match(FormulaParserFALSE)
		}
		localctx.(*ValueContext).retValue = exp.NewBooleanValueExpression(false)

	case FormulaParserT__30:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(288)
			p.Π()
		}
		localctx.(*ValueContext).retValue = exp.NewPiExpression()

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

	p.SetState(297)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case FormulaParserNAME:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(293)

			var _m = p.Match(FormulaParserNAME)

			localctx.(*IdContext)._NAME = _m
		}
		localctx.(*IdContext).retValue = exp.NewIdentifierExpression((func() string {
			if localctx.(*IdContext).Get_NAME() == nil {
				return ""
			} else {
				return localctx.(*IdContext).Get_NAME().GetText()
			}
		}()))

	case FormulaParserVAR:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(295)

			var _m = p.Match(FormulaParserVAR)

			localctx.(*IdContext)._VAR = _m
		}
		localctx.(*IdContext).retValue = exp.NewVarIdentifierExpression((func() string {
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

// IΠContext is an interface to support dynamic dispatch.
type IΠContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsΠContext differentiates from other interfaces.
	IsΠContext()
}

type ΠContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyΠContext() *ΠContext {
	var p = new(ΠContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FormulaParserRULE_π
	return p
}

func (*ΠContext) IsΠContext() {}

func NewΠContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ΠContext {
	var p = new(ΠContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_π

	return p
}

func (s *ΠContext) GetParser() antlr.Parser { return s.parser }
func (s *ΠContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ΠContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ΠContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaListener); ok {
		listenerT.EnterΠ(s)
	}
}

func (s *ΠContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaListener); ok {
		listenerT.ExitΠ(s)
	}
}

func (p *FormulaParser) Π() (localctx IΠContext) {
	localctx = NewΠContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, FormulaParserRULE_π)

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
		p.SetState(299)
		p.Match(FormulaParserT__30)
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
