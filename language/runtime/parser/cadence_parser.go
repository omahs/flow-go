// Code generated from parser/Cadence.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Cadence
import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

import "strings"

var _ = strings.Builder{}

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 73, 774,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9,
	39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44, 9, 44,
	4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9, 49, 4,
	50, 9, 50, 4, 51, 9, 51, 4, 52, 9, 52, 4, 53, 9, 53, 4, 54, 9, 54, 4, 55,
	9, 55, 4, 56, 9, 56, 4, 57, 9, 57, 4, 58, 9, 58, 4, 59, 9, 59, 4, 60, 9,
	60, 4, 61, 9, 61, 4, 62, 9, 62, 4, 63, 9, 63, 4, 64, 9, 64, 4, 65, 9, 65,
	4, 66, 9, 66, 4, 67, 9, 67, 4, 68, 9, 68, 4, 69, 9, 69, 4, 70, 9, 70, 4,
	71, 9, 71, 4, 72, 9, 72, 4, 73, 9, 73, 4, 74, 9, 74, 4, 75, 9, 75, 4, 76,
	9, 76, 4, 77, 9, 77, 4, 78, 9, 78, 4, 79, 9, 79, 4, 80, 9, 80, 4, 81, 9,
	81, 4, 82, 9, 82, 4, 83, 9, 83, 4, 84, 9, 84, 4, 85, 9, 85, 4, 86, 9, 86,
	4, 87, 9, 87, 4, 88, 9, 88, 4, 89, 9, 89, 3, 2, 3, 2, 5, 2, 181, 10, 2,
	7, 2, 183, 10, 2, 12, 2, 14, 2, 186, 11, 2, 3, 2, 3, 2, 3, 3, 3, 3, 5,
	3, 192, 10, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 200, 10, 4, 3,
	5, 3, 5, 3, 5, 3, 5, 7, 5, 206, 10, 5, 12, 5, 14, 5, 209, 11, 5, 3, 5,
	3, 5, 5, 5, 213, 10, 5, 3, 5, 3, 5, 5, 5, 217, 10, 5, 3, 6, 3, 6, 3, 6,
	3, 6, 5, 6, 223, 10, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7,
	3, 8, 3, 8, 3, 8, 3, 8, 7, 8, 237, 10, 8, 12, 8, 14, 8, 240, 11, 8, 5,
	8, 242, 10, 8, 3, 9, 3, 9, 3, 10, 3, 10, 5, 10, 248, 10, 10, 3, 10, 3,
	10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11,
	3, 12, 7, 12, 263, 10, 12, 12, 12, 14, 12, 266, 11, 12, 3, 13, 3, 13, 3,
	13, 3, 13, 3, 13, 5, 13, 273, 10, 13, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15,
	5, 15, 280, 10, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3,
	16, 5, 16, 290, 10, 16, 3, 16, 5, 16, 293, 10, 16, 3, 16, 3, 16, 3, 17,
	3, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 18, 7, 18, 305, 10, 18, 12,
	18, 14, 18, 308, 11, 18, 5, 18, 310, 10, 18, 3, 18, 3, 18, 3, 19, 5, 19,
	315, 10, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 20, 5, 20, 322, 10, 20, 3,
	20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 7, 21, 332, 10, 21,
	12, 21, 14, 21, 335, 11, 21, 5, 21, 337, 10, 21, 3, 22, 3, 22, 3, 22, 3,
	22, 3, 22, 5, 22, 344, 10, 22, 3, 23, 3, 23, 3, 24, 3, 24, 3, 24, 3, 24,
	3, 24, 7, 24, 353, 10, 24, 12, 24, 14, 24, 356, 11, 24, 5, 24, 358, 10,
	24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 25, 3, 25, 3, 25, 3, 25, 3, 26,
	3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3,
	27, 3, 28, 3, 28, 3, 28, 3, 28, 3, 29, 3, 29, 5, 29, 387, 10, 29, 3, 29,
	5, 29, 390, 10, 29, 3, 29, 3, 29, 3, 29, 3, 30, 3, 30, 3, 30, 3, 30, 3,
	30, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 32, 3, 32, 3, 32, 7, 32, 408,
	10, 32, 12, 32, 14, 32, 411, 11, 32, 3, 33, 3, 33, 3, 33, 5, 33, 416, 10,
	33, 3, 34, 3, 34, 3, 34, 7, 34, 421, 10, 34, 12, 34, 14, 34, 424, 11, 34,
	3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 5,
	35, 436, 10, 35, 3, 36, 3, 36, 3, 36, 5, 36, 441, 10, 36, 3, 37, 3, 37,
	3, 38, 3, 38, 3, 39, 3, 39, 3, 39, 5, 39, 450, 10, 39, 3, 39, 3, 39, 3,
	39, 3, 39, 5, 39, 456, 10, 39, 5, 39, 458, 10, 39, 3, 40, 3, 40, 3, 40,
	3, 40, 3, 41, 3, 41, 3, 41, 3, 41, 3, 42, 3, 42, 3, 42, 3, 42, 3, 42, 5,
	42, 473, 10, 42, 3, 42, 3, 42, 3, 42, 3, 42, 3, 42, 5, 42, 480, 10, 42,
	3, 43, 3, 43, 7, 43, 484, 10, 43, 12, 43, 14, 43, 487, 11, 43, 3, 43, 3,
	43, 3, 43, 3, 44, 3, 44, 3, 44, 3, 44, 3, 45, 3, 45, 3, 46, 3, 46, 3, 47,
	3, 47, 3, 47, 3, 47, 3, 47, 3, 47, 5, 47, 506, 10, 47, 3, 48, 3, 48, 3,
	48, 3, 48, 3, 48, 3, 48, 7, 48, 514, 10, 48, 12, 48, 14, 48, 517, 11, 48,
	3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 7, 49, 525, 10, 49, 12, 49, 14,
	49, 528, 11, 49, 3, 50, 3, 50, 3, 50, 3, 50, 3, 50, 3, 50, 3, 50, 7, 50,
	537, 10, 50, 12, 50, 14, 50, 540, 11, 50, 3, 51, 3, 51, 3, 51, 3, 51, 3,
	51, 3, 51, 3, 51, 7, 51, 549, 10, 51, 12, 51, 14, 51, 552, 11, 51, 3, 52,
	3, 52, 3, 52, 5, 52, 557, 10, 52, 3, 53, 3, 53, 3, 53, 3, 53, 3, 53, 3,
	53, 7, 53, 565, 10, 53, 12, 53, 14, 53, 568, 11, 53, 3, 54, 3, 54, 3, 54,
	3, 54, 3, 54, 3, 54, 7, 54, 576, 10, 54, 12, 54, 14, 54, 579, 11, 54, 3,
	55, 3, 55, 3, 55, 3, 55, 3, 55, 3, 55, 3, 55, 7, 55, 588, 10, 55, 12, 55,
	14, 55, 591, 11, 55, 3, 56, 3, 56, 3, 56, 3, 56, 3, 56, 3, 56, 3, 56, 7,
	56, 600, 10, 56, 12, 56, 14, 56, 603, 11, 56, 3, 57, 3, 57, 6, 57, 607,
	10, 57, 13, 57, 14, 57, 608, 3, 57, 3, 57, 5, 57, 613, 10, 57, 3, 58, 3,
	58, 3, 58, 3, 58, 5, 58, 619, 10, 58, 3, 59, 3, 59, 7, 59, 623, 10, 59,
	12, 59, 14, 59, 626, 11, 59, 3, 60, 3, 60, 5, 60, 630, 10, 60, 3, 61, 3,
	61, 3, 62, 3, 62, 3, 63, 3, 63, 3, 64, 3, 64, 3, 65, 3, 65, 3, 66, 3, 66,
	3, 66, 3, 66, 5, 66, 646, 10, 66, 3, 67, 3, 67, 3, 67, 3, 67, 3, 68, 3,
	68, 3, 68, 3, 69, 3, 69, 3, 69, 3, 69, 3, 69, 3, 70, 3, 70, 3, 71, 3, 71,
	3, 72, 3, 72, 3, 72, 3, 72, 5, 72, 668, 10, 72, 3, 72, 3, 72, 3, 73, 3,
	73, 3, 73, 3, 73, 3, 74, 3, 74, 5, 74, 678, 10, 74, 3, 75, 3, 75, 3, 75,
	3, 76, 3, 76, 3, 76, 5, 76, 686, 10, 76, 3, 76, 3, 76, 3, 77, 3, 77, 3,
	77, 3, 77, 7, 77, 694, 10, 77, 12, 77, 14, 77, 697, 11, 77, 5, 77, 699,
	10, 77, 3, 77, 3, 77, 3, 78, 3, 78, 3, 78, 5, 78, 706, 10, 78, 3, 78, 3,
	78, 3, 79, 3, 79, 3, 79, 3, 79, 3, 79, 3, 79, 5, 79, 716, 10, 79, 3, 80,
	3, 80, 3, 81, 3, 81, 3, 82, 3, 82, 3, 83, 5, 83, 725, 10, 83, 3, 83, 3,
	83, 3, 84, 3, 84, 3, 84, 3, 84, 3, 84, 5, 84, 734, 10, 84, 3, 85, 3, 85,
	3, 85, 3, 85, 7, 85, 740, 10, 85, 12, 85, 14, 85, 743, 11, 85, 5, 85, 745,
	10, 85, 3, 85, 3, 85, 3, 86, 3, 86, 3, 86, 3, 86, 7, 86, 753, 10, 86, 12,
	86, 14, 86, 756, 11, 86, 5, 86, 758, 10, 86, 3, 86, 3, 86, 3, 87, 3, 87,
	3, 87, 3, 87, 3, 88, 3, 88, 3, 89, 3, 89, 3, 89, 3, 89, 5, 89, 772, 10,
	89, 3, 89, 2, 10, 94, 96, 98, 100, 104, 106, 108, 110, 90, 2, 4, 6, 8,
	10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44,
	46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80,
	82, 84, 86, 88, 90, 92, 94, 96, 98, 100, 102, 104, 106, 108, 110, 112,
	114, 116, 118, 120, 122, 124, 126, 128, 130, 132, 134, 136, 138, 140, 142,
	144, 146, 148, 150, 152, 154, 156, 158, 160, 162, 164, 166, 168, 170, 172,
	174, 176, 2, 12, 3, 2, 51, 52, 3, 2, 36, 38, 4, 2, 11, 11, 28, 28, 3, 2,
	15, 16, 3, 2, 17, 20, 3, 2, 21, 22, 3, 2, 23, 25, 4, 2, 22, 22, 27, 28,
	3, 2, 56, 57, 3, 2, 60, 63, 2, 786, 2, 184, 3, 2, 2, 2, 4, 191, 3, 2, 2,
	2, 6, 199, 3, 2, 2, 2, 8, 201, 3, 2, 2, 2, 10, 222, 3, 2, 2, 2, 12, 224,
	3, 2, 2, 2, 14, 241, 3, 2, 2, 2, 16, 243, 3, 2, 2, 2, 18, 245, 3, 2, 2,
	2, 20, 253, 3, 2, 2, 2, 22, 264, 3, 2, 2, 2, 24, 272, 3, 2, 2, 2, 26, 274,
	3, 2, 2, 2, 28, 276, 3, 2, 2, 2, 30, 283, 3, 2, 2, 2, 32, 296, 3, 2, 2,
	2, 34, 300, 3, 2, 2, 2, 36, 314, 3, 2, 2, 2, 38, 321, 3, 2, 2, 2, 40, 336,
	3, 2, 2, 2, 42, 343, 3, 2, 2, 2, 44, 345, 3, 2, 2, 2, 46, 347, 3, 2, 2,
	2, 48, 364, 3, 2, 2, 2, 50, 368, 3, 2, 2, 2, 52, 374, 3, 2, 2, 2, 54, 380,
	3, 2, 2, 2, 56, 384, 3, 2, 2, 2, 58, 394, 3, 2, 2, 2, 60, 399, 3, 2, 2,
	2, 62, 409, 3, 2, 2, 2, 64, 412, 3, 2, 2, 2, 66, 422, 3, 2, 2, 2, 68, 435,
	3, 2, 2, 2, 70, 437, 3, 2, 2, 2, 72, 442, 3, 2, 2, 2, 74, 444, 3, 2, 2,
	2, 76, 446, 3, 2, 2, 2, 78, 459, 3, 2, 2, 2, 80, 463, 3, 2, 2, 2, 82, 467,
	3, 2, 2, 2, 84, 481, 3, 2, 2, 2, 86, 491, 3, 2, 2, 2, 88, 495, 3, 2, 2,
	2, 90, 497, 3, 2, 2, 2, 92, 499, 3, 2, 2, 2, 94, 507, 3, 2, 2, 2, 96, 518,
	3, 2, 2, 2, 98, 529, 3, 2, 2, 2, 100, 541, 3, 2, 2, 2, 102, 553, 3, 2,
	2, 2, 104, 558, 3, 2, 2, 2, 106, 569, 3, 2, 2, 2, 108, 580, 3, 2, 2, 2,
	110, 592, 3, 2, 2, 2, 112, 612, 3, 2, 2, 2, 114, 618, 3, 2, 2, 2, 116,
	620, 3, 2, 2, 2, 118, 629, 3, 2, 2, 2, 120, 631, 3, 2, 2, 2, 122, 633,
	3, 2, 2, 2, 124, 635, 3, 2, 2, 2, 126, 637, 3, 2, 2, 2, 128, 639, 3, 2,
	2, 2, 130, 645, 3, 2, 2, 2, 132, 647, 3, 2, 2, 2, 134, 651, 3, 2, 2, 2,
	136, 654, 3, 2, 2, 2, 138, 659, 3, 2, 2, 2, 140, 661, 3, 2, 2, 2, 142,
	663, 3, 2, 2, 2, 144, 671, 3, 2, 2, 2, 146, 677, 3, 2, 2, 2, 148, 679,
	3, 2, 2, 2, 150, 682, 3, 2, 2, 2, 152, 689, 3, 2, 2, 2, 154, 705, 3, 2,
	2, 2, 156, 715, 3, 2, 2, 2, 158, 717, 3, 2, 2, 2, 160, 719, 3, 2, 2, 2,
	162, 721, 3, 2, 2, 2, 164, 724, 3, 2, 2, 2, 166, 733, 3, 2, 2, 2, 168,
	735, 3, 2, 2, 2, 170, 748, 3, 2, 2, 2, 172, 761, 3, 2, 2, 2, 174, 765,
	3, 2, 2, 2, 176, 771, 3, 2, 2, 2, 178, 180, 5, 6, 4, 2, 179, 181, 7, 3,
	2, 2, 180, 179, 3, 2, 2, 2, 180, 181, 3, 2, 2, 2, 181, 183, 3, 2, 2, 2,
	182, 178, 3, 2, 2, 2, 183, 186, 3, 2, 2, 2, 184, 182, 3, 2, 2, 2, 184,
	185, 3, 2, 2, 2, 185, 187, 3, 2, 2, 2, 186, 184, 3, 2, 2, 2, 187, 188,
	7, 2, 2, 3, 188, 3, 3, 2, 2, 2, 189, 192, 5, 2, 2, 2, 190, 192, 5, 66,
	34, 2, 191, 189, 3, 2, 2, 2, 191, 190, 3, 2, 2, 2, 192, 5, 3, 2, 2, 2,
	193, 200, 5, 12, 7, 2, 194, 200, 5, 20, 11, 2, 195, 200, 5, 30, 16, 2,
	196, 200, 5, 82, 42, 2, 197, 200, 5, 8, 5, 2, 198, 200, 5, 32, 17, 2, 199,
	193, 3, 2, 2, 2, 199, 194, 3, 2, 2, 2, 199, 195, 3, 2, 2, 2, 199, 196,
	3, 2, 2, 2, 199, 197, 3, 2, 2, 2, 199, 198, 3, 2, 2, 2, 200, 7, 3, 2, 2,
	2, 201, 212, 7, 59, 2, 2, 202, 207, 5, 174, 88, 2, 203, 204, 7, 4, 2, 2,
	204, 206, 5, 174, 88, 2, 205, 203, 3, 2, 2, 2, 206, 209, 3, 2, 2, 2, 207,
	205, 3, 2, 2, 2, 207, 208, 3, 2, 2, 2, 208, 210, 3, 2, 2, 2, 209, 207,
	3, 2, 2, 2, 210, 211, 7, 60, 2, 2, 211, 213, 3, 2, 2, 2, 212, 202, 3, 2,
	2, 2, 212, 213, 3, 2, 2, 2, 213, 216, 3, 2, 2, 2, 214, 217, 5, 162, 82,
	2, 215, 217, 7, 67, 2, 2, 216, 214, 3, 2, 2, 2, 216, 215, 3, 2, 2, 2, 217,
	9, 3, 2, 2, 2, 218, 223, 3, 2, 2, 2, 219, 223, 7, 45, 2, 2, 220, 223, 7,
	46, 2, 2, 221, 223, 7, 47, 2, 2, 222, 218, 3, 2, 2, 2, 222, 219, 3, 2,
	2, 2, 222, 220, 3, 2, 2, 2, 222, 221, 3, 2, 2, 2, 223, 11, 3, 2, 2, 2,
	224, 225, 5, 10, 6, 2, 225, 226, 5, 26, 14, 2, 226, 227, 5, 174, 88, 2,
	227, 228, 5, 14, 8, 2, 228, 229, 7, 5, 2, 2, 229, 230, 5, 22, 12, 2, 230,
	231, 7, 6, 2, 2, 231, 13, 3, 2, 2, 2, 232, 233, 7, 7, 2, 2, 233, 238, 5,
	174, 88, 2, 234, 235, 7, 4, 2, 2, 235, 237, 5, 174, 88, 2, 236, 234, 3,
	2, 2, 2, 237, 240, 3, 2, 2, 2, 238, 236, 3, 2, 2, 2, 238, 239, 3, 2, 2,
	2, 239, 242, 3, 2, 2, 2, 240, 238, 3, 2, 2, 2, 241, 232, 3, 2, 2, 2, 241,
	242, 3, 2, 2, 2, 242, 15, 3, 2, 2, 2, 243, 244, 9, 2, 2, 2, 244, 17, 3,
	2, 2, 2, 245, 247, 5, 10, 6, 2, 246, 248, 5, 16, 9, 2, 247, 246, 3, 2,
	2, 2, 247, 248, 3, 2, 2, 2, 248, 249, 3, 2, 2, 2, 249, 250, 5, 174, 88,
	2, 250, 251, 7, 7, 2, 2, 251, 252, 5, 38, 20, 2, 252, 19, 3, 2, 2, 2, 253,
	254, 5, 10, 6, 2, 254, 255, 5, 26, 14, 2, 255, 256, 7, 39, 2, 2, 256, 257,
	5, 174, 88, 2, 257, 258, 7, 5, 2, 2, 258, 259, 5, 22, 12, 2, 259, 260,
	7, 6, 2, 2, 260, 21, 3, 2, 2, 2, 261, 263, 5, 24, 13, 2, 262, 261, 3, 2,
	2, 2, 263, 266, 3, 2, 2, 2, 264, 262, 3, 2, 2, 2, 264, 265, 3, 2, 2, 2,
	265, 23, 3, 2, 2, 2, 266, 264, 3, 2, 2, 2, 267, 273, 5, 18, 10, 2, 268,
	273, 5, 28, 15, 2, 269, 273, 5, 30, 16, 2, 270, 273, 5, 20, 11, 2, 271,
	273, 5, 12, 7, 2, 272, 267, 3, 2, 2, 2, 272, 268, 3, 2, 2, 2, 272, 269,
	3, 2, 2, 2, 272, 270, 3, 2, 2, 2, 272, 271, 3, 2, 2, 2, 273, 25, 3, 2,
	2, 2, 274, 275, 9, 3, 2, 2, 275, 27, 3, 2, 2, 2, 276, 277, 5, 174, 88,
	2, 277, 279, 5, 34, 18, 2, 278, 280, 5, 56, 29, 2, 279, 278, 3, 2, 2, 2,
	279, 280, 3, 2, 2, 2, 280, 281, 3, 2, 2, 2, 281, 282, 6, 15, 2, 3, 282,
	29, 3, 2, 2, 2, 283, 284, 5, 10, 6, 2, 284, 285, 7, 40, 2, 2, 285, 286,
	5, 174, 88, 2, 286, 289, 5, 34, 18, 2, 287, 288, 7, 7, 2, 2, 288, 290,
	5, 38, 20, 2, 289, 287, 3, 2, 2, 2, 289, 290, 3, 2, 2, 2, 290, 292, 3,
	2, 2, 2, 291, 293, 5, 56, 29, 2, 292, 291, 3, 2, 2, 2, 292, 293, 3, 2,
	2, 2, 293, 294, 3, 2, 2, 2, 294, 295, 6, 16, 3, 3, 295, 31, 3, 2, 2, 2,
	296, 297, 7, 41, 2, 2, 297, 298, 5, 174, 88, 2, 298, 299, 5, 34, 18, 2,
	299, 33, 3, 2, 2, 2, 300, 309, 7, 33, 2, 2, 301, 306, 5, 36, 19, 2, 302,
	303, 7, 4, 2, 2, 303, 305, 5, 36, 19, 2, 304, 302, 3, 2, 2, 2, 305, 308,
	3, 2, 2, 2, 306, 304, 3, 2, 2, 2, 306, 307, 3, 2, 2, 2, 307, 310, 3, 2,
	2, 2, 308, 306, 3, 2, 2, 2, 309, 301, 3, 2, 2, 2, 309, 310, 3, 2, 2, 2,
	310, 311, 3, 2, 2, 2, 311, 312, 7, 34, 2, 2, 312, 35, 3, 2, 2, 2, 313,
	315, 5, 174, 88, 2, 314, 313, 3, 2, 2, 2, 314, 315, 3, 2, 2, 2, 315, 316,
	3, 2, 2, 2, 316, 317, 5, 174, 88, 2, 317, 318, 7, 7, 2, 2, 318, 319, 5,
	38, 20, 2, 319, 37, 3, 2, 2, 2, 320, 322, 7, 28, 2, 2, 321, 320, 3, 2,
	2, 2, 321, 322, 3, 2, 2, 2, 322, 323, 3, 2, 2, 2, 323, 324, 5, 40, 21,
	2, 324, 39, 3, 2, 2, 2, 325, 326, 7, 26, 2, 2, 326, 327, 6, 21, 4, 2, 327,
	337, 5, 42, 22, 2, 328, 333, 5, 42, 22, 2, 329, 330, 6, 21, 5, 2, 330,
	332, 7, 29, 2, 2, 331, 329, 3, 2, 2, 2, 332, 335, 3, 2, 2, 2, 333, 331,
	3, 2, 2, 2, 333, 334, 3, 2, 2, 2, 334, 337, 3, 2, 2, 2, 335, 333, 3, 2,
	2, 2, 336, 325, 3, 2, 2, 2, 336, 328, 3, 2, 2, 2, 337, 41, 3, 2, 2, 2,
	338, 344, 5, 44, 23, 2, 339, 344, 5, 46, 24, 2, 340, 344, 5, 48, 25, 2,
	341, 344, 5, 50, 26, 2, 342, 344, 5, 52, 27, 2, 343, 338, 3, 2, 2, 2, 343,
	339, 3, 2, 2, 2, 343, 340, 3, 2, 2, 2, 343, 341, 3, 2, 2, 2, 343, 342,
	3, 2, 2, 2, 344, 43, 3, 2, 2, 2, 345, 346, 5, 174, 88, 2, 346, 45, 3, 2,
	2, 2, 347, 348, 7, 33, 2, 2, 348, 357, 7, 33, 2, 2, 349, 354, 5, 38, 20,
	2, 350, 351, 7, 4, 2, 2, 351, 353, 5, 38, 20, 2, 352, 350, 3, 2, 2, 2,
	353, 356, 3, 2, 2, 2, 354, 352, 3, 2, 2, 2, 354, 355, 3, 2, 2, 2, 355,
	358, 3, 2, 2, 2, 356, 354, 3, 2, 2, 2, 357, 349, 3, 2, 2, 2, 357, 358,
	3, 2, 2, 2, 358, 359, 3, 2, 2, 2, 359, 360, 7, 34, 2, 2, 360, 361, 7, 7,
	2, 2, 361, 362, 5, 38, 20, 2, 362, 363, 7, 34, 2, 2, 363, 47, 3, 2, 2,
	2, 364, 365, 7, 8, 2, 2, 365, 366, 5, 40, 21, 2, 366, 367, 7, 9, 2, 2,
	367, 49, 3, 2, 2, 2, 368, 369, 7, 8, 2, 2, 369, 370, 5, 40, 21, 2, 370,
	371, 7, 3, 2, 2, 371, 372, 7, 64, 2, 2, 372, 373, 7, 9, 2, 2, 373, 51,
	3, 2, 2, 2, 374, 375, 7, 5, 2, 2, 375, 376, 5, 40, 21, 2, 376, 377, 7,
	7, 2, 2, 377, 378, 5, 40, 21, 2, 378, 379, 7, 6, 2, 2, 379, 53, 3, 2, 2,
	2, 380, 381, 7, 5, 2, 2, 381, 382, 5, 66, 34, 2, 382, 383, 7, 6, 2, 2,
	383, 55, 3, 2, 2, 2, 384, 386, 7, 5, 2, 2, 385, 387, 5, 58, 30, 2, 386,
	385, 3, 2, 2, 2, 386, 387, 3, 2, 2, 2, 387, 389, 3, 2, 2, 2, 388, 390,
	5, 60, 31, 2, 389, 388, 3, 2, 2, 2, 389, 390, 3, 2, 2, 2, 390, 391, 3,
	2, 2, 2, 391, 392, 5, 66, 34, 2, 392, 393, 7, 6, 2, 2, 393, 57, 3, 2, 2,
	2, 394, 395, 7, 43, 2, 2, 395, 396, 7, 5, 2, 2, 396, 397, 5, 62, 32, 2,
	397, 398, 7, 6, 2, 2, 398, 59, 3, 2, 2, 2, 399, 400, 7, 44, 2, 2, 400,
	401, 7, 5, 2, 2, 401, 402, 5, 62, 32, 2, 402, 403, 7, 6, 2, 2, 403, 61,
	3, 2, 2, 2, 404, 405, 5, 64, 33, 2, 405, 406, 5, 176, 89, 2, 406, 408,
	3, 2, 2, 2, 407, 404, 3, 2, 2, 2, 408, 411, 3, 2, 2, 2, 409, 407, 3, 2,
	2, 2, 409, 410, 3, 2, 2, 2, 410, 63, 3, 2, 2, 2, 411, 409, 3, 2, 2, 2,
	412, 415, 5, 90, 46, 2, 413, 414, 7, 7, 2, 2, 414, 416, 5, 90, 46, 2, 415,
	413, 3, 2, 2, 2, 415, 416, 3, 2, 2, 2, 416, 65, 3, 2, 2, 2, 417, 418, 5,
	68, 35, 2, 418, 419, 5, 176, 89, 2, 419, 421, 3, 2, 2, 2, 420, 417, 3,
	2, 2, 2, 421, 424, 3, 2, 2, 2, 422, 420, 3, 2, 2, 2, 422, 423, 3, 2, 2,
	2, 423, 67, 3, 2, 2, 2, 424, 422, 3, 2, 2, 2, 425, 436, 5, 70, 36, 2, 426,
	436, 5, 72, 37, 2, 427, 436, 5, 74, 38, 2, 428, 436, 5, 76, 39, 2, 429,
	436, 5, 78, 40, 2, 430, 436, 5, 80, 41, 2, 431, 436, 5, 6, 4, 2, 432, 436,
	5, 84, 43, 2, 433, 436, 5, 86, 44, 2, 434, 436, 5, 90, 46, 2, 435, 425,
	3, 2, 2, 2, 435, 426, 3, 2, 2, 2, 435, 427, 3, 2, 2, 2, 435, 428, 3, 2,
	2, 2, 435, 429, 3, 2, 2, 2, 435, 430, 3, 2, 2, 2, 435, 431, 3, 2, 2, 2,
	435, 432, 3, 2, 2, 2, 435, 433, 3, 2, 2, 2, 435, 434, 3, 2, 2, 2, 436,
	69, 3, 2, 2, 2, 437, 440, 7, 48, 2, 2, 438, 439, 6, 36, 6, 2, 439, 441,
	5, 90, 46, 2, 440, 438, 3, 2, 2, 2, 440, 441, 3, 2, 2, 2, 441, 71, 3, 2,
	2, 2, 442, 443, 7, 49, 2, 2, 443, 73, 3, 2, 2, 2, 444, 445, 7, 50, 2, 2,
	445, 75, 3, 2, 2, 2, 446, 449, 7, 53, 2, 2, 447, 450, 5, 90, 46, 2, 448,
	450, 5, 82, 42, 2, 449, 447, 3, 2, 2, 2, 449, 448, 3, 2, 2, 2, 450, 451,
	3, 2, 2, 2, 451, 457, 5, 54, 28, 2, 452, 455, 7, 54, 2, 2, 453, 456, 5,
	76, 39, 2, 454, 456, 5, 54, 28, 2, 455, 453, 3, 2, 2, 2, 455, 454, 3, 2,
	2, 2, 456, 458, 3, 2, 2, 2, 457, 452, 3, 2, 2, 2, 457, 458, 3, 2, 2, 2,
	458, 77, 3, 2, 2, 2, 459, 460, 7, 55, 2, 2, 460, 461, 5, 90, 46, 2, 461,
	462, 5, 54, 28, 2, 462, 79, 3, 2, 2, 2, 463, 464, 7, 42, 2, 2, 464, 465,
	5, 174, 88, 2, 465, 466, 5, 152, 77, 2, 466, 81, 3, 2, 2, 2, 467, 468,
	5, 10, 6, 2, 468, 469, 5, 16, 9, 2, 469, 472, 5, 174, 88, 2, 470, 471,
	7, 7, 2, 2, 471, 473, 5, 38, 20, 2, 472, 470, 3, 2, 2, 2, 472, 473, 3,
	2, 2, 2, 473, 474, 3, 2, 2, 2, 474, 475, 5, 88, 45, 2, 475, 479, 5, 90,
	46, 2, 476, 477, 5, 88, 45, 2, 477, 478, 5, 90, 46, 2, 478, 480, 3, 2,
	2, 2, 479, 476, 3, 2, 2, 2, 479, 480, 3, 2, 2, 2, 480, 83, 3, 2, 2, 2,
	481, 485, 5, 174, 88, 2, 482, 484, 5, 146, 74, 2, 483, 482, 3, 2, 2, 2,
	484, 487, 3, 2, 2, 2, 485, 483, 3, 2, 2, 2, 485, 486, 3, 2, 2, 2, 486,
	488, 3, 2, 2, 2, 487, 485, 3, 2, 2, 2, 488, 489, 5, 88, 45, 2, 489, 490,
	5, 90, 46, 2, 490, 85, 3, 2, 2, 2, 491, 492, 5, 90, 46, 2, 492, 493, 7,
	10, 2, 2, 493, 494, 5, 90, 46, 2, 494, 87, 3, 2, 2, 2, 495, 496, 9, 4,
	2, 2, 496, 89, 3, 2, 2, 2, 497, 498, 5, 92, 47, 2, 498, 91, 3, 2, 2, 2,
	499, 505, 5, 94, 48, 2, 500, 501, 7, 29, 2, 2, 501, 502, 5, 90, 46, 2,
	502, 503, 7, 7, 2, 2, 503, 504, 5, 90, 46, 2, 504, 506, 3, 2, 2, 2, 505,
	500, 3, 2, 2, 2, 505, 506, 3, 2, 2, 2, 506, 93, 3, 2, 2, 2, 507, 508, 8,
	48, 1, 2, 508, 509, 5, 96, 49, 2, 509, 515, 3, 2, 2, 2, 510, 511, 12, 3,
	2, 2, 511, 512, 7, 12, 2, 2, 512, 514, 5, 96, 49, 2, 513, 510, 3, 2, 2,
	2, 514, 517, 3, 2, 2, 2, 515, 513, 3, 2, 2, 2, 515, 516, 3, 2, 2, 2, 516,
	95, 3, 2, 2, 2, 517, 515, 3, 2, 2, 2, 518, 519, 8, 49, 1, 2, 519, 520,
	5, 98, 50, 2, 520, 526, 3, 2, 2, 2, 521, 522, 12, 3, 2, 2, 522, 523, 7,
	13, 2, 2, 523, 525, 5, 98, 50, 2, 524, 521, 3, 2, 2, 2, 525, 528, 3, 2,
	2, 2, 526, 524, 3, 2, 2, 2, 526, 527, 3, 2, 2, 2, 527, 97, 3, 2, 2, 2,
	528, 526, 3, 2, 2, 2, 529, 530, 8, 50, 1, 2, 530, 531, 5, 100, 51, 2, 531,
	538, 3, 2, 2, 2, 532, 533, 12, 3, 2, 2, 533, 534, 5, 120, 61, 2, 534, 535,
	5, 100, 51, 2, 535, 537, 3, 2, 2, 2, 536, 532, 3, 2, 2, 2, 537, 540, 3,
	2, 2, 2, 538, 536, 3, 2, 2, 2, 538, 539, 3, 2, 2, 2, 539, 99, 3, 2, 2,
	2, 540, 538, 3, 2, 2, 2, 541, 542, 8, 51, 1, 2, 542, 543, 5, 102, 52, 2,
	543, 550, 3, 2, 2, 2, 544, 545, 12, 3, 2, 2, 545, 546, 5, 122, 62, 2, 546,
	547, 5, 102, 52, 2, 547, 549, 3, 2, 2, 2, 548, 544, 3, 2, 2, 2, 549, 552,
	3, 2, 2, 2, 550, 548, 3, 2, 2, 2, 550, 551, 3, 2, 2, 2, 551, 101, 3, 2,
	2, 2, 552, 550, 3, 2, 2, 2, 553, 556, 5, 104, 53, 2, 554, 555, 7, 30, 2,
	2, 555, 557, 5, 102, 52, 2, 556, 554, 3, 2, 2, 2, 556, 557, 3, 2, 2, 2,
	557, 103, 3, 2, 2, 2, 558, 559, 8, 53, 1, 2, 559, 560, 5, 106, 54, 2, 560,
	566, 3, 2, 2, 2, 561, 562, 12, 3, 2, 2, 562, 563, 7, 32, 2, 2, 563, 565,
	5, 38, 20, 2, 564, 561, 3, 2, 2, 2, 565, 568, 3, 2, 2, 2, 566, 564, 3,
	2, 2, 2, 566, 567, 3, 2, 2, 2, 567, 105, 3, 2, 2, 2, 568, 566, 3, 2, 2,
	2, 569, 570, 8, 54, 1, 2, 570, 571, 5, 108, 55, 2, 571, 577, 3, 2, 2, 2,
	572, 573, 12, 3, 2, 2, 573, 574, 7, 26, 2, 2, 574, 576, 5, 108, 55, 2,
	575, 572, 3, 2, 2, 2, 576, 579, 3, 2, 2, 2, 577, 575, 3, 2, 2, 2, 577,
	578, 3, 2, 2, 2, 578, 107, 3, 2, 2, 2, 579, 577, 3, 2, 2, 2, 580, 581,
	8, 55, 1, 2, 581, 582, 5, 110, 56, 2, 582, 589, 3, 2, 2, 2, 583, 584, 12,
	3, 2, 2, 584, 585, 5, 124, 63, 2, 585, 586, 5, 110, 56, 2, 586, 588, 3,
	2, 2, 2, 587, 583, 3, 2, 2, 2, 588, 591, 3, 2, 2, 2, 589, 587, 3, 2, 2,
	2, 589, 590, 3, 2, 2, 2, 590, 109, 3, 2, 2, 2, 591, 589, 3, 2, 2, 2, 592,
	593, 8, 56, 1, 2, 593, 594, 5, 112, 57, 2, 594, 601, 3, 2, 2, 2, 595, 596,
	12, 3, 2, 2, 596, 597, 5, 126, 64, 2, 597, 598, 5, 112, 57, 2, 598, 600,
	3, 2, 2, 2, 599, 595, 3, 2, 2, 2, 600, 603, 3, 2, 2, 2, 601, 599, 3, 2,
	2, 2, 601, 602, 3, 2, 2, 2, 602, 111, 3, 2, 2, 2, 603, 601, 3, 2, 2, 2,
	604, 613, 5, 114, 58, 2, 605, 607, 5, 128, 65, 2, 606, 605, 3, 2, 2, 2,
	607, 608, 3, 2, 2, 2, 608, 606, 3, 2, 2, 2, 608, 609, 3, 2, 2, 2, 609,
	610, 3, 2, 2, 2, 610, 611, 5, 112, 57, 2, 611, 613, 3, 2, 2, 2, 612, 604,
	3, 2, 2, 2, 612, 606, 3, 2, 2, 2, 613, 113, 3, 2, 2, 2, 614, 619, 5, 132,
	67, 2, 615, 619, 5, 134, 68, 2, 616, 619, 5, 136, 69, 2, 617, 619, 5, 116,
	59, 2, 618, 614, 3, 2, 2, 2, 618, 615, 3, 2, 2, 2, 618, 616, 3, 2, 2, 2,
	618, 617, 3, 2, 2, 2, 619, 115, 3, 2, 2, 2, 620, 624, 5, 130, 66, 2, 621,
	623, 5, 118, 60, 2, 622, 621, 3, 2, 2, 2, 623, 626, 3, 2, 2, 2, 624, 622,
	3, 2, 2, 2, 624, 625, 3, 2, 2, 2, 625, 117, 3, 2, 2, 2, 626, 624, 3, 2,
	2, 2, 627, 630, 5, 146, 74, 2, 628, 630, 5, 152, 77, 2, 629, 627, 3, 2,
	2, 2, 629, 628, 3, 2, 2, 2, 630, 119, 3, 2, 2, 2, 631, 632, 9, 5, 2, 2,
	632, 121, 3, 2, 2, 2, 633, 634, 9, 6, 2, 2, 634, 123, 3, 2, 2, 2, 635,
	636, 9, 7, 2, 2, 636, 125, 3, 2, 2, 2, 637, 638, 9, 8, 2, 2, 638, 127,
	3, 2, 2, 2, 639, 640, 9, 9, 2, 2, 640, 129, 3, 2, 2, 2, 641, 646, 5, 138,
	70, 2, 642, 646, 5, 140, 71, 2, 643, 646, 5, 142, 72, 2, 644, 646, 5, 144,
	73, 2, 645, 641, 3, 2, 2, 2, 645, 642, 3, 2, 2, 2, 645, 643, 3, 2, 2, 2,
	645, 644, 3, 2, 2, 2, 646, 131, 3, 2, 2, 2, 647, 648, 7, 61, 2, 2, 648,
	649, 5, 174, 88, 2, 649, 650, 5, 152, 77, 2, 650, 133, 3, 2, 2, 2, 651,
	652, 7, 62, 2, 2, 652, 653, 5, 90, 46, 2, 653, 135, 3, 2, 2, 2, 654, 655,
	7, 26, 2, 2, 655, 656, 5, 90, 46, 2, 656, 657, 7, 31, 2, 2, 657, 658, 5,
	40, 21, 2, 658, 137, 3, 2, 2, 2, 659, 660, 5, 174, 88, 2, 660, 139, 3,
	2, 2, 2, 661, 662, 5, 156, 79, 2, 662, 141, 3, 2, 2, 2, 663, 664, 7, 40,
	2, 2, 664, 667, 5, 34, 18, 2, 665, 666, 7, 7, 2, 2, 666, 668, 5, 38, 20,
	2, 667, 665, 3, 2, 2, 2, 667, 668, 3, 2, 2, 2, 668, 669, 3, 2, 2, 2, 669,
	670, 5, 56, 29, 2, 670, 143, 3, 2, 2, 2, 671, 672, 7, 33, 2, 2, 672, 673,
	5, 90, 46, 2, 673, 674, 7, 34, 2, 2, 674, 145, 3, 2, 2, 2, 675, 678, 5,
	148, 75, 2, 676, 678, 5, 150, 76, 2, 677, 675, 3, 2, 2, 2, 677, 676, 3,
	2, 2, 2, 678, 147, 3, 2, 2, 2, 679, 680, 7, 14, 2, 2, 680, 681, 5, 174,
	88, 2, 681, 149, 3, 2, 2, 2, 682, 685, 7, 8, 2, 2, 683, 686, 5, 90, 46,
	2, 684, 686, 5, 40, 21, 2, 685, 683, 3, 2, 2, 2, 685, 684, 3, 2, 2, 2,
	686, 687, 3, 2, 2, 2, 687, 688, 7, 9, 2, 2, 688, 151, 3, 2, 2, 2, 689,
	698, 7, 33, 2, 2, 690, 695, 5, 154, 78, 2, 691, 692, 7, 4, 2, 2, 692, 694,
	5, 154, 78, 2, 693, 691, 3, 2, 2, 2, 694, 697, 3, 2, 2, 2, 695, 693, 3,
	2, 2, 2, 695, 696, 3, 2, 2, 2, 696, 699, 3, 2, 2, 2, 697, 695, 3, 2, 2,
	2, 698, 690, 3, 2, 2, 2, 698, 699, 3, 2, 2, 2, 699, 700, 3, 2, 2, 2, 700,
	701, 7, 34, 2, 2, 701, 153, 3, 2, 2, 2, 702, 703, 5, 174, 88, 2, 703, 704,
	7, 7, 2, 2, 704, 706, 3, 2, 2, 2, 705, 702, 3, 2, 2, 2, 705, 706, 3, 2,
	2, 2, 706, 707, 3, 2, 2, 2, 707, 708, 5, 90, 46, 2, 708, 155, 3, 2, 2,
	2, 709, 716, 5, 164, 83, 2, 710, 716, 5, 158, 80, 2, 711, 716, 5, 168,
	85, 2, 712, 716, 5, 170, 86, 2, 713, 716, 5, 162, 82, 2, 714, 716, 5, 160,
	81, 2, 715, 709, 3, 2, 2, 2, 715, 710, 3, 2, 2, 2, 715, 711, 3, 2, 2, 2,
	715, 712, 3, 2, 2, 2, 715, 713, 3, 2, 2, 2, 715, 714, 3, 2, 2, 2, 716,
	157, 3, 2, 2, 2, 717, 718, 9, 10, 2, 2, 718, 159, 3, 2, 2, 2, 719, 720,
	7, 58, 2, 2, 720, 161, 3, 2, 2, 2, 721, 722, 7, 69, 2, 2, 722, 163, 3,
	2, 2, 2, 723, 725, 7, 22, 2, 2, 724, 723, 3, 2, 2, 2, 724, 725, 3, 2, 2,
	2, 725, 726, 3, 2, 2, 2, 726, 727, 5, 166, 84, 2, 727, 165, 3, 2, 2, 2,
	728, 734, 7, 64, 2, 2, 729, 734, 7, 65, 2, 2, 730, 734, 7, 66, 2, 2, 731,
	734, 7, 67, 2, 2, 732, 734, 7, 68, 2, 2, 733, 728, 3, 2, 2, 2, 733, 729,
	3, 2, 2, 2, 733, 730, 3, 2, 2, 2, 733, 731, 3, 2, 2, 2, 733, 732, 3, 2,
	2, 2, 734, 167, 3, 2, 2, 2, 735, 744, 7, 8, 2, 2, 736, 741, 5, 90, 46,
	2, 737, 738, 7, 4, 2, 2, 738, 740, 5, 90, 46, 2, 739, 737, 3, 2, 2, 2,
	740, 743, 3, 2, 2, 2, 741, 739, 3, 2, 2, 2, 741, 742, 3, 2, 2, 2, 742,
	745, 3, 2, 2, 2, 743, 741, 3, 2, 2, 2, 744, 736, 3, 2, 2, 2, 744, 745,
	3, 2, 2, 2, 745, 746, 3, 2, 2, 2, 746, 747, 7, 9, 2, 2, 747, 169, 3, 2,
	2, 2, 748, 757, 7, 5, 2, 2, 749, 754, 5, 172, 87, 2, 750, 751, 7, 4, 2,
	2, 751, 753, 5, 172, 87, 2, 752, 750, 3, 2, 2, 2, 753, 756, 3, 2, 2, 2,
	754, 752, 3, 2, 2, 2, 754, 755, 3, 2, 2, 2, 755, 758, 3, 2, 2, 2, 756,
	754, 3, 2, 2, 2, 757, 749, 3, 2, 2, 2, 757, 758, 3, 2, 2, 2, 758, 759,
	3, 2, 2, 2, 759, 760, 7, 6, 2, 2, 760, 171, 3, 2, 2, 2, 761, 762, 5, 90,
	46, 2, 762, 763, 7, 7, 2, 2, 763, 764, 5, 90, 46, 2, 764, 173, 3, 2, 2,
	2, 765, 766, 9, 11, 2, 2, 766, 175, 3, 2, 2, 2, 767, 772, 7, 3, 2, 2, 768,
	772, 7, 2, 2, 3, 769, 772, 6, 89, 15, 2, 770, 772, 6, 89, 16, 2, 771, 767,
	3, 2, 2, 2, 771, 768, 3, 2, 2, 2, 771, 769, 3, 2, 2, 2, 771, 770, 3, 2,
	2, 2, 772, 177, 3, 2, 2, 2, 70, 180, 184, 191, 199, 207, 212, 216, 222,
	238, 241, 247, 264, 272, 279, 289, 292, 306, 309, 314, 321, 333, 336, 343,
	354, 357, 386, 389, 409, 415, 422, 435, 440, 449, 455, 457, 472, 479, 485,
	505, 515, 526, 538, 550, 556, 566, 577, 589, 601, 608, 612, 618, 624, 629,
	645, 667, 677, 685, 695, 698, 705, 715, 724, 733, 741, 744, 754, 757, 771,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "';'", "','", "'{'", "'}'", "':'", "'['", "']'", "'<->'", "'='", "'||'",
	"'&&'", "'.'", "'=='", "'!='", "'<'", "'>'", "'<='", "'>='", "'+'", "'-'",
	"'*'", "'/'", "'%'", "'&'", "'!'", "'<-'", "'?'", "", "'as'", "'as?'",
	"'('", "')'", "'transaction'", "'struct'", "'resource'", "'contract'",
	"'interface'", "'fun'", "'event'", "'emit'", "'pre'", "'post'", "'priv'",
	"'pub'", "'pub(set)'", "'return'", "'break'", "'continue'", "'let'", "'var'",
	"'if'", "'else'", "'while'", "'true'", "'false'", "'nil'", "'import'",
	"'from'", "'create'", "'destroy'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "Equal", "Unequal",
	"Less", "Greater", "LessEqual", "GreaterEqual", "Plus", "Minus", "Mul",
	"Div", "Mod", "Ampersand", "Negate", "Move", "Optional", "NilCoalescing",
	"Downcasting", "FailableDowncasting", "OpenParen", "CloseParen", "Transaction",
	"Struct", "Resource", "Contract", "Interface", "Fun", "Event", "Emit",
	"Pre", "Post", "Priv", "Pub", "PubSet", "Return", "Break", "Continue",
	"Let", "Var", "If", "Else", "While", "True", "False", "Nil", "Import",
	"From", "Create", "Destroy", "Identifier", "DecimalLiteral", "BinaryLiteral",
	"OctalLiteral", "HexadecimalLiteral", "InvalidNumberLiteral", "StringLiteral",
	"WS", "Terminator", "BlockComment", "LineComment",
}

var ruleNames = []string{
	"program", "replInput", "declaration", "importDeclaration", "access", "compositeDeclaration",
	"conformances", "variableKind", "field", "interfaceDeclaration", "members",
	"member", "compositeKind", "specialFunctionDeclaration", "functionDeclaration",
	"eventDeclaration", "parameterList", "parameter", "typeAnnotation", "fullType",
	"baseType", "nominalType", "functionType", "variableSizedType", "constantSizedType",
	"dictionaryType", "block", "functionBlock", "preConditions", "postConditions",
	"conditions", "condition", "statements", "statement", "returnStatement",
	"breakStatement", "continueStatement", "ifStatement", "whileStatement",
	"emitStatement", "variableDeclaration", "assignment", "swap", "transfer",
	"expression", "conditionalExpression", "orExpression", "andExpression",
	"equalityExpression", "relationalExpression", "nilCoalescingExpression",
	"failableDowncastingExpression", "concatenatingExpression", "additiveExpression",
	"multiplicativeExpression", "unaryExpression", "primaryExpression", "composedExpression",
	"primaryExpressionSuffix", "equalityOp", "relationalOp", "additiveOp",
	"multiplicativeOp", "unaryOp", "primaryExpressionStart", "createExpression",
	"destroyExpression", "referenceExpression", "identifierExpression", "literalExpression",
	"functionExpression", "nestedExpression", "expressionAccess", "memberAccess",
	"bracketExpression", "invocation", "argument", "literal", "booleanLiteral",
	"nilLiteral", "stringLiteral", "integerLiteral", "positiveIntegerLiteral",
	"arrayLiteral", "dictionaryLiteral", "dictionaryEntry", "identifier", "eos",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type CadenceParser struct {
	*antlr.BaseParser
}

func NewCadenceParser(input antlr.TokenStream) *CadenceParser {
	this := new(CadenceParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Cadence.g4"

	return this
}

// Returns true if on the current index of the parser's
// token stream a token exists on the Hidden channel which
// either is a line terminator, or is a multi line comment that
// contains a line terminator.
func (p *CadenceParser) lineTerminatorAhead() bool {
	// Get the token ahead of the current index.
	possibleIndexEosToken := p.GetCurrentToken().GetTokenIndex() - 1
	ahead := p.GetTokenStream().Get(possibleIndexEosToken)

	if ahead.GetChannel() != antlr.LexerHidden {
		// We're only interested in tokens on the HIDDEN channel.
		return true
	}

	if ahead.GetTokenType() == CadenceParserTerminator {
		// There is definitely a line terminator ahead.
		return true
	}

	if ahead.GetTokenType() == CadenceParserWS {
		// Get the token ahead of the current whitespaces.
		possibleIndexEosToken = p.GetCurrentToken().GetTokenIndex() - 2
		ahead = p.GetTokenStream().Get(possibleIndexEosToken)
	}

	// Get the token's text and type.
	text := ahead.GetText()
	_type := ahead.GetTokenType()

	// Check if the token is, or contains a line terminator.
	return (_type == CadenceParserBlockComment && (strings.Contains(text, "\r") || strings.Contains(text, "\n"))) ||
		(_type == CadenceParserTerminator)
}

func (p *CadenceParser) noWhitespace() bool {
	index := p.GetCurrentToken().GetTokenIndex()
	return p.GetTokenStream().Get(index-1).GetTokenType() != CadenceParserWS
}

// CadenceParser tokens.
const (
	CadenceParserEOF                  = antlr.TokenEOF
	CadenceParserT__0                 = 1
	CadenceParserT__1                 = 2
	CadenceParserT__2                 = 3
	CadenceParserT__3                 = 4
	CadenceParserT__4                 = 5
	CadenceParserT__5                 = 6
	CadenceParserT__6                 = 7
	CadenceParserT__7                 = 8
	CadenceParserT__8                 = 9
	CadenceParserT__9                 = 10
	CadenceParserT__10                = 11
	CadenceParserT__11                = 12
	CadenceParserEqual                = 13
	CadenceParserUnequal              = 14
	CadenceParserLess                 = 15
	CadenceParserGreater              = 16
	CadenceParserLessEqual            = 17
	CadenceParserGreaterEqual         = 18
	CadenceParserPlus                 = 19
	CadenceParserMinus                = 20
	CadenceParserMul                  = 21
	CadenceParserDiv                  = 22
	CadenceParserMod                  = 23
	CadenceParserAmpersand            = 24
	CadenceParserNegate               = 25
	CadenceParserMove                 = 26
	CadenceParserOptional             = 27
	CadenceParserNilCoalescing        = 28
	CadenceParserDowncasting          = 29
	CadenceParserFailableDowncasting  = 30
	CadenceParserOpenParen            = 31
	CadenceParserCloseParen           = 32
	CadenceParserTransaction          = 33
	CadenceParserStruct               = 34
	CadenceParserResource             = 35
	CadenceParserContract             = 36
	CadenceParserInterface            = 37
	CadenceParserFun                  = 38
	CadenceParserEvent                = 39
	CadenceParserEmit                 = 40
	CadenceParserPre                  = 41
	CadenceParserPost                 = 42
	CadenceParserPriv                 = 43
	CadenceParserPub                  = 44
	CadenceParserPubSet               = 45
	CadenceParserReturn               = 46
	CadenceParserBreak                = 47
	CadenceParserContinue             = 48
	CadenceParserLet                  = 49
	CadenceParserVar                  = 50
	CadenceParserIf                   = 51
	CadenceParserElse                 = 52
	CadenceParserWhile                = 53
	CadenceParserTrue                 = 54
	CadenceParserFalse                = 55
	CadenceParserNil                  = 56
	CadenceParserImport               = 57
	CadenceParserFrom                 = 58
	CadenceParserCreate               = 59
	CadenceParserDestroy              = 60
	CadenceParserIdentifier           = 61
	CadenceParserDecimalLiteral       = 62
	CadenceParserBinaryLiteral        = 63
	CadenceParserOctalLiteral         = 64
	CadenceParserHexadecimalLiteral   = 65
	CadenceParserInvalidNumberLiteral = 66
	CadenceParserStringLiteral        = 67
	CadenceParserWS                   = 68
	CadenceParserTerminator           = 69
	CadenceParserBlockComment         = 70
	CadenceParserLineComment          = 71
)

// CadenceParser rules.
const (
	CadenceParserRULE_program                       = 0
	CadenceParserRULE_replInput                     = 1
	CadenceParserRULE_declaration                   = 2
	CadenceParserRULE_importDeclaration             = 3
	CadenceParserRULE_access                        = 4
	CadenceParserRULE_compositeDeclaration          = 5
	CadenceParserRULE_conformances                  = 6
	CadenceParserRULE_variableKind                  = 7
	CadenceParserRULE_field                         = 8
	CadenceParserRULE_interfaceDeclaration          = 9
	CadenceParserRULE_members                       = 10
	CadenceParserRULE_member                        = 11
	CadenceParserRULE_compositeKind                 = 12
	CadenceParserRULE_specialFunctionDeclaration    = 13
	CadenceParserRULE_functionDeclaration           = 14
	CadenceParserRULE_eventDeclaration              = 15
	CadenceParserRULE_parameterList                 = 16
	CadenceParserRULE_parameter                     = 17
	CadenceParserRULE_typeAnnotation                = 18
	CadenceParserRULE_fullType                      = 19
	CadenceParserRULE_baseType                      = 20
	CadenceParserRULE_nominalType                   = 21
	CadenceParserRULE_functionType                  = 22
	CadenceParserRULE_variableSizedType             = 23
	CadenceParserRULE_constantSizedType             = 24
	CadenceParserRULE_dictionaryType                = 25
	CadenceParserRULE_block                         = 26
	CadenceParserRULE_functionBlock                 = 27
	CadenceParserRULE_preConditions                 = 28
	CadenceParserRULE_postConditions                = 29
	CadenceParserRULE_conditions                    = 30
	CadenceParserRULE_condition                     = 31
	CadenceParserRULE_statements                    = 32
	CadenceParserRULE_statement                     = 33
	CadenceParserRULE_returnStatement               = 34
	CadenceParserRULE_breakStatement                = 35
	CadenceParserRULE_continueStatement             = 36
	CadenceParserRULE_ifStatement                   = 37
	CadenceParserRULE_whileStatement                = 38
	CadenceParserRULE_emitStatement                 = 39
	CadenceParserRULE_variableDeclaration           = 40
	CadenceParserRULE_assignment                    = 41
	CadenceParserRULE_swap                          = 42
	CadenceParserRULE_transfer                      = 43
	CadenceParserRULE_expression                    = 44
	CadenceParserRULE_conditionalExpression         = 45
	CadenceParserRULE_orExpression                  = 46
	CadenceParserRULE_andExpression                 = 47
	CadenceParserRULE_equalityExpression            = 48
	CadenceParserRULE_relationalExpression          = 49
	CadenceParserRULE_nilCoalescingExpression       = 50
	CadenceParserRULE_failableDowncastingExpression = 51
	CadenceParserRULE_concatenatingExpression       = 52
	CadenceParserRULE_additiveExpression            = 53
	CadenceParserRULE_multiplicativeExpression      = 54
	CadenceParserRULE_unaryExpression               = 55
	CadenceParserRULE_primaryExpression             = 56
	CadenceParserRULE_composedExpression            = 57
	CadenceParserRULE_primaryExpressionSuffix       = 58
	CadenceParserRULE_equalityOp                    = 59
	CadenceParserRULE_relationalOp                  = 60
	CadenceParserRULE_additiveOp                    = 61
	CadenceParserRULE_multiplicativeOp              = 62
	CadenceParserRULE_unaryOp                       = 63
	CadenceParserRULE_primaryExpressionStart        = 64
	CadenceParserRULE_createExpression              = 65
	CadenceParserRULE_destroyExpression             = 66
	CadenceParserRULE_referenceExpression           = 67
	CadenceParserRULE_identifierExpression          = 68
	CadenceParserRULE_literalExpression             = 69
	CadenceParserRULE_functionExpression            = 70
	CadenceParserRULE_nestedExpression              = 71
	CadenceParserRULE_expressionAccess              = 72
	CadenceParserRULE_memberAccess                  = 73
	CadenceParserRULE_bracketExpression             = 74
	CadenceParserRULE_invocation                    = 75
	CadenceParserRULE_argument                      = 76
	CadenceParserRULE_literal                       = 77
	CadenceParserRULE_booleanLiteral                = 78
	CadenceParserRULE_nilLiteral                    = 79
	CadenceParserRULE_stringLiteral                 = 80
	CadenceParserRULE_integerLiteral                = 81
	CadenceParserRULE_positiveIntegerLiteral        = 82
	CadenceParserRULE_arrayLiteral                  = 83
	CadenceParserRULE_dictionaryLiteral             = 84
	CadenceParserRULE_dictionaryEntry               = 85
	CadenceParserRULE_identifier                    = 86
	CadenceParserRULE_eos                           = 87
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CadenceParserRULE_program
	return p
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CadenceParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(CadenceParserEOF, 0)
}

func (s *ProgramContext) AllDeclaration() []IDeclarationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDeclarationContext)(nil)).Elem())
	var tst = make([]IDeclarationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDeclarationContext)
		}
	}

	return tst
}

func (s *ProgramContext) Declaration(i int) IDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclarationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDeclarationContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CadenceVisitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CadenceParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, CadenceParserRULE_program)
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

	p.EnterOuterAlt(localctx, 1)
	p.SetState(182)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la-34)&-(0x1f+1)) == 0 && ((1<<uint((_la-34)))&((1<<(CadenceParserStruct-34))|(1<<(CadenceParserResource-34))|(1<<(CadenceParserContract-34))|(1<<(CadenceParserFun-34))|(1<<(CadenceParserEvent-34))|(1<<(CadenceParserPriv-34))|(1<<(CadenceParserPub-34))|(1<<(CadenceParserPubSet-34))|(1<<(CadenceParserLet-34))|(1<<(CadenceParserVar-34))|(1<<(CadenceParserImport-34)))) != 0 {
		{
			p.SetState(176)
			p.Declaration()
		}
		p.SetState(178)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CadenceParserT__0 {
			{
				p.SetState(177)
				p.Match(CadenceParserT__0)
			}

		}

		p.SetState(184)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(185)
		p.Match(CadenceParserEOF)
	}

	return localctx
}

// IReplInputContext is an interface to support dynamic dispatch.
type IReplInputContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReplInputContext differentiates from other interfaces.
	IsReplInputContext()
}

type ReplInputContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReplInputContext() *ReplInputContext {
	var p = new(ReplInputContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CadenceParserRULE_replInput
	return p
}

func (*ReplInputContext) IsReplInputContext() {}

func NewReplInputContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReplInputContext {
	var p = new(ReplInputContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CadenceParserRULE_replInput

	return p
}

func (s *ReplInputContext) GetParser() antlr.Parser { return s.parser }

func (s *ReplInputContext) Program() IProgramContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IProgramContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IProgramContext)
}

func (s *ReplInputContext) Statements() IStatementsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementsContext)
}

func (s *ReplInputContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReplInputContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReplInputContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.EnterReplInput(s)
	}
}

func (s *ReplInputContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.ExitReplInput(s)
	}
}

func (s *ReplInputContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CadenceVisitor:
		return t.VisitReplInput(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CadenceParser) ReplInput() (localctx IReplInputContext) {
	localctx = NewReplInputContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, CadenceParserRULE_replInput)

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

	p.SetState(189)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(187)
			p.Program()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(188)
			p.Statements()
		}

	}

	return localctx
}

// IDeclarationContext is an interface to support dynamic dispatch.
type IDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDeclarationContext differentiates from other interfaces.
	IsDeclarationContext()
}

type DeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclarationContext() *DeclarationContext {
	var p = new(DeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CadenceParserRULE_declaration
	return p
}

func (*DeclarationContext) IsDeclarationContext() {}

func NewDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationContext {
	var p = new(DeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CadenceParserRULE_declaration

	return p
}

func (s *DeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarationContext) CompositeDeclaration() ICompositeDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICompositeDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICompositeDeclarationContext)
}

func (s *DeclarationContext) InterfaceDeclaration() IInterfaceDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInterfaceDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInterfaceDeclarationContext)
}

func (s *DeclarationContext) FunctionDeclaration() IFunctionDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionDeclarationContext)
}

func (s *DeclarationContext) VariableDeclaration() IVariableDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableDeclarationContext)
}

func (s *DeclarationContext) ImportDeclaration() IImportDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IImportDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IImportDeclarationContext)
}

func (s *DeclarationContext) EventDeclaration() IEventDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEventDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEventDeclarationContext)
}

func (s *DeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.EnterDeclaration(s)
	}
}

func (s *DeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.ExitDeclaration(s)
	}
}

func (s *DeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CadenceVisitor:
		return t.VisitDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CadenceParser) Declaration() (localctx IDeclarationContext) {
	localctx = NewDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, CadenceParserRULE_declaration)

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

	p.SetState(197)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(191)
			p.CompositeDeclaration()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(192)
			p.InterfaceDeclaration()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(193)
			p.FunctionDeclaration(true)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(194)
			p.VariableDeclaration()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(195)
			p.ImportDeclaration()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(196)
			p.EventDeclaration()
		}

	}

	return localctx
}

// IImportDeclarationContext is an interface to support dynamic dispatch.
type IImportDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsImportDeclarationContext differentiates from other interfaces.
	IsImportDeclarationContext()
}

type ImportDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImportDeclarationContext() *ImportDeclarationContext {
	var p = new(ImportDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CadenceParserRULE_importDeclaration
	return p
}

func (*ImportDeclarationContext) IsImportDeclarationContext() {}

func NewImportDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportDeclarationContext {
	var p = new(ImportDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CadenceParserRULE_importDeclaration

	return p
}

func (s *ImportDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *ImportDeclarationContext) Import() antlr.TerminalNode {
	return s.GetToken(CadenceParserImport, 0)
}

func (s *ImportDeclarationContext) StringLiteral() IStringLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStringLiteralContext)
}

func (s *ImportDeclarationContext) HexadecimalLiteral() antlr.TerminalNode {
	return s.GetToken(CadenceParserHexadecimalLiteral, 0)
}

func (s *ImportDeclarationContext) AllIdentifier() []IIdentifierContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIdentifierContext)(nil)).Elem())
	var tst = make([]IIdentifierContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIdentifierContext)
		}
	}

	return tst
}

func (s *ImportDeclarationContext) Identifier(i int) IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *ImportDeclarationContext) From() antlr.TerminalNode {
	return s.GetToken(CadenceParserFrom, 0)
}

func (s *ImportDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImportDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImportDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.EnterImportDeclaration(s)
	}
}

func (s *ImportDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.ExitImportDeclaration(s)
	}
}

func (s *ImportDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CadenceVisitor:
		return t.VisitImportDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CadenceParser) ImportDeclaration() (localctx IImportDeclarationContext) {
	localctx = NewImportDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, CadenceParserRULE_importDeclaration)
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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(199)
		p.Match(CadenceParserImport)
	}
	p.SetState(210)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la-58)&-(0x1f+1)) == 0 && ((1<<uint((_la-58)))&((1<<(CadenceParserFrom-58))|(1<<(CadenceParserCreate-58))|(1<<(CadenceParserDestroy-58))|(1<<(CadenceParserIdentifier-58)))) != 0 {
		{
			p.SetState(200)
			p.Identifier()
		}
		p.SetState(205)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == CadenceParserT__1 {
			{
				p.SetState(201)
				p.Match(CadenceParserT__1)
			}
			{
				p.SetState(202)
				p.Identifier()
			}

			p.SetState(207)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(208)
			p.Match(CadenceParserFrom)
		}

	}
	p.SetState(214)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CadenceParserStringLiteral:
		{
			p.SetState(212)
			p.StringLiteral()
		}

	case CadenceParserHexadecimalLiteral:
		{
			p.SetState(213)
			p.Match(CadenceParserHexadecimalLiteral)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IAccessContext is an interface to support dynamic dispatch.
type IAccessContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAccessContext differentiates from other interfaces.
	IsAccessContext()
}

type AccessContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAccessContext() *AccessContext {
	var p = new(AccessContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CadenceParserRULE_access
	return p
}

func (*AccessContext) IsAccessContext() {}

func NewAccessContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AccessContext {
	var p = new(AccessContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CadenceParserRULE_access

	return p
}

func (s *AccessContext) GetParser() antlr.Parser { return s.parser }

func (s *AccessContext) Priv() antlr.TerminalNode {
	return s.GetToken(CadenceParserPriv, 0)
}

func (s *AccessContext) Pub() antlr.TerminalNode {
	return s.GetToken(CadenceParserPub, 0)
}

func (s *AccessContext) PubSet() antlr.TerminalNode {
	return s.GetToken(CadenceParserPubSet, 0)
}

func (s *AccessContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AccessContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AccessContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.EnterAccess(s)
	}
}

func (s *AccessContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.ExitAccess(s)
	}
}

func (s *AccessContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CadenceVisitor:
		return t.VisitAccess(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CadenceParser) Access() (localctx IAccessContext) {
	localctx = NewAccessContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, CadenceParserRULE_access)

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

	p.SetState(220)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CadenceParserStruct, CadenceParserResource, CadenceParserContract, CadenceParserFun, CadenceParserLet, CadenceParserVar, CadenceParserFrom, CadenceParserCreate, CadenceParserDestroy, CadenceParserIdentifier:
		p.EnterOuterAlt(localctx, 1)

	case CadenceParserPriv:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(217)
			p.Match(CadenceParserPriv)
		}

	case CadenceParserPub:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(218)
			p.Match(CadenceParserPub)
		}

	case CadenceParserPubSet:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(219)
			p.Match(CadenceParserPubSet)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ICompositeDeclarationContext is an interface to support dynamic dispatch.
type ICompositeDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCompositeDeclarationContext differentiates from other interfaces.
	IsCompositeDeclarationContext()
}

type CompositeDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCompositeDeclarationContext() *CompositeDeclarationContext {
	var p = new(CompositeDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CadenceParserRULE_compositeDeclaration
	return p
}

func (*CompositeDeclarationContext) IsCompositeDeclarationContext() {}

func NewCompositeDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CompositeDeclarationContext {
	var p = new(CompositeDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CadenceParserRULE_compositeDeclaration

	return p
}

func (s *CompositeDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *CompositeDeclarationContext) Access() IAccessContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAccessContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAccessContext)
}

func (s *CompositeDeclarationContext) CompositeKind() ICompositeKindContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICompositeKindContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICompositeKindContext)
}

func (s *CompositeDeclarationContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *CompositeDeclarationContext) Conformances() IConformancesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConformancesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConformancesContext)
}

func (s *CompositeDeclarationContext) Members() IMembersContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMembersContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMembersContext)
}

func (s *CompositeDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompositeDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CompositeDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.EnterCompositeDeclaration(s)
	}
}

func (s *CompositeDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.ExitCompositeDeclaration(s)
	}
}

func (s *CompositeDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CadenceVisitor:
		return t.VisitCompositeDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CadenceParser) CompositeDeclaration() (localctx ICompositeDeclarationContext) {
	localctx = NewCompositeDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, CadenceParserRULE_compositeDeclaration)

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
		p.SetState(222)
		p.Access()
	}
	{
		p.SetState(223)
		p.CompositeKind()
	}
	{
		p.SetState(224)
		p.Identifier()
	}
	{
		p.SetState(225)
		p.Conformances()
	}
	{
		p.SetState(226)
		p.Match(CadenceParserT__2)
	}
	{
		p.SetState(227)
		p.Members(true)
	}
	{
		p.SetState(228)
		p.Match(CadenceParserT__3)
	}

	return localctx
}

// IConformancesContext is an interface to support dynamic dispatch.
type IConformancesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConformancesContext differentiates from other interfaces.
	IsConformancesContext()
}

type ConformancesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConformancesContext() *ConformancesContext {
	var p = new(ConformancesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CadenceParserRULE_conformances
	return p
}

func (*ConformancesContext) IsConformancesContext() {}

func NewConformancesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConformancesContext {
	var p = new(ConformancesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CadenceParserRULE_conformances

	return p
}

func (s *ConformancesContext) GetParser() antlr.Parser { return s.parser }

func (s *ConformancesContext) AllIdentifier() []IIdentifierContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIdentifierContext)(nil)).Elem())
	var tst = make([]IIdentifierContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIdentifierContext)
		}
	}

	return tst
}

func (s *ConformancesContext) Identifier(i int) IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *ConformancesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConformancesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConformancesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.EnterConformances(s)
	}
}

func (s *ConformancesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.ExitConformances(s)
	}
}

func (s *ConformancesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CadenceVisitor:
		return t.VisitConformances(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CadenceParser) Conformances() (localctx IConformancesContext) {
	localctx = NewConformancesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, CadenceParserRULE_conformances)
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

	p.EnterOuterAlt(localctx, 1)
	p.SetState(239)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CadenceParserT__4 {
		{
			p.SetState(230)
			p.Match(CadenceParserT__4)
		}
		{
			p.SetState(231)
			p.Identifier()
		}
		p.SetState(236)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == CadenceParserT__1 {
			{
				p.SetState(232)
				p.Match(CadenceParserT__1)
			}
			{
				p.SetState(233)
				p.Identifier()
			}

			p.SetState(238)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}

	return localctx
}

// IVariableKindContext is an interface to support dynamic dispatch.
type IVariableKindContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableKindContext differentiates from other interfaces.
	IsVariableKindContext()
}

type VariableKindContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableKindContext() *VariableKindContext {
	var p = new(VariableKindContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CadenceParserRULE_variableKind
	return p
}

func (*VariableKindContext) IsVariableKindContext() {}

func NewVariableKindContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableKindContext {
	var p = new(VariableKindContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CadenceParserRULE_variableKind

	return p
}

func (s *VariableKindContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableKindContext) Let() antlr.TerminalNode {
	return s.GetToken(CadenceParserLet, 0)
}

func (s *VariableKindContext) Var() antlr.TerminalNode {
	return s.GetToken(CadenceParserVar, 0)
}

func (s *VariableKindContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableKindContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableKindContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.EnterVariableKind(s)
	}
}

func (s *VariableKindContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.ExitVariableKind(s)
	}
}

func (s *VariableKindContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CadenceVisitor:
		return t.VisitVariableKind(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CadenceParser) VariableKind() (localctx IVariableKindContext) {
	localctx = NewVariableKindContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, CadenceParserRULE_variableKind)
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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(241)
		_la = p.GetTokenStream().LA(1)

		if !(_la == CadenceParserLet || _la == CadenceParserVar) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IFieldContext is an interface to support dynamic dispatch.
type IFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldContext differentiates from other interfaces.
	IsFieldContext()
}

type FieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldContext() *FieldContext {
	var p = new(FieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CadenceParserRULE_field
	return p
}

func (*FieldContext) IsFieldContext() {}

func NewFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldContext {
	var p = new(FieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CadenceParserRULE_field

	return p
}

func (s *FieldContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldContext) Access() IAccessContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAccessContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAccessContext)
}

func (s *FieldContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *FieldContext) TypeAnnotation() ITypeAnnotationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeAnnotationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeAnnotationContext)
}

func (s *FieldContext) VariableKind() IVariableKindContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableKindContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableKindContext)
}

func (s *FieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.EnterField(s)
	}
}

func (s *FieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.ExitField(s)
	}
}

func (s *FieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CadenceVisitor:
		return t.VisitField(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CadenceParser) Field() (localctx IFieldContext) {
	localctx = NewFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, CadenceParserRULE_field)
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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(243)
		p.Access()
	}
	p.SetState(245)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CadenceParserLet || _la == CadenceParserVar {
		{
			p.SetState(244)
			p.VariableKind()
		}

	}
	{
		p.SetState(247)
		p.Identifier()
	}
	{
		p.SetState(248)
		p.Match(CadenceParserT__4)
	}
	{
		p.SetState(249)
		p.TypeAnnotation()
	}

	return localctx
}

// IInterfaceDeclarationContext is an interface to support dynamic dispatch.
type IInterfaceDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInterfaceDeclarationContext differentiates from other interfaces.
	IsInterfaceDeclarationContext()
}

type InterfaceDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInterfaceDeclarationContext() *InterfaceDeclarationContext {
	var p = new(InterfaceDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CadenceParserRULE_interfaceDeclaration
	return p
}

func (*InterfaceDeclarationContext) IsInterfaceDeclarationContext() {}

func NewInterfaceDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InterfaceDeclarationContext {
	var p = new(InterfaceDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CadenceParserRULE_interfaceDeclaration

	return p
}

func (s *InterfaceDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *InterfaceDeclarationContext) Access() IAccessContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAccessContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAccessContext)
}

func (s *InterfaceDeclarationContext) CompositeKind() ICompositeKindContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICompositeKindContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICompositeKindContext)
}

func (s *InterfaceDeclarationContext) Interface() antlr.TerminalNode {
	return s.GetToken(CadenceParserInterface, 0)
}

func (s *InterfaceDeclarationContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *InterfaceDeclarationContext) Members() IMembersContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMembersContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMembersContext)
}

func (s *InterfaceDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InterfaceDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InterfaceDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.EnterInterfaceDeclaration(s)
	}
}

func (s *InterfaceDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.ExitInterfaceDeclaration(s)
	}
}

func (s *InterfaceDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CadenceVisitor:
		return t.VisitInterfaceDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CadenceParser) InterfaceDeclaration() (localctx IInterfaceDeclarationContext) {
	localctx = NewInterfaceDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, CadenceParserRULE_interfaceDeclaration)

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
		p.SetState(251)
		p.Access()
	}
	{
		p.SetState(252)
		p.CompositeKind()
	}
	{
		p.SetState(253)
		p.Match(CadenceParserInterface)
	}
	{
		p.SetState(254)
		p.Identifier()
	}
	{
		p.SetState(255)
		p.Match(CadenceParserT__2)
	}
	{
		p.SetState(256)
		p.Members(false)
	}
	{
		p.SetState(257)
		p.Match(CadenceParserT__3)
	}

	return localctx
}

// IMembersContext is an interface to support dynamic dispatch.
type IMembersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetFunctionBlockRequired returns the functionBlockRequired attribute.
	GetFunctionBlockRequired() bool

	// SetFunctionBlockRequired sets the functionBlockRequired attribute.
	SetFunctionBlockRequired(bool)

	// IsMembersContext differentiates from other interfaces.
	IsMembersContext()
}

type MembersContext struct {
	*antlr.BaseParserRuleContext
	parser                antlr.Parser
	functionBlockRequired bool
}

func NewEmptyMembersContext() *MembersContext {
	var p = new(MembersContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CadenceParserRULE_members
	return p
}

func (*MembersContext) IsMembersContext() {}

func NewMembersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int, functionBlockRequired bool) *MembersContext {
	var p = new(MembersContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CadenceParserRULE_members

	p.functionBlockRequired = functionBlockRequired

	return p
}

func (s *MembersContext) GetParser() antlr.Parser { return s.parser }

func (s *MembersContext) GetFunctionBlockRequired() bool { return s.functionBlockRequired }

func (s *MembersContext) SetFunctionBlockRequired(v bool) { s.functionBlockRequired = v }

func (s *MembersContext) AllMember() []IMemberContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMemberContext)(nil)).Elem())
	var tst = make([]IMemberContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMemberContext)
		}
	}

	return tst
}

func (s *MembersContext) Member(i int) IMemberContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMemberContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMemberContext)
}

func (s *MembersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MembersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MembersContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.EnterMembers(s)
	}
}

func (s *MembersContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.ExitMembers(s)
	}
}

func (s *MembersContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CadenceVisitor:
		return t.VisitMembers(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CadenceParser) Members(functionBlockRequired bool) (localctx IMembersContext) {
	localctx = NewMembersContext(p, p.GetParserRuleContext(), p.GetState(), functionBlockRequired)
	p.EnterRule(localctx, 20, CadenceParserRULE_members)
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

	p.EnterOuterAlt(localctx, 1)
	p.SetState(262)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la-34)&-(0x1f+1)) == 0 && ((1<<uint((_la-34)))&((1<<(CadenceParserStruct-34))|(1<<(CadenceParserResource-34))|(1<<(CadenceParserContract-34))|(1<<(CadenceParserFun-34))|(1<<(CadenceParserPriv-34))|(1<<(CadenceParserPub-34))|(1<<(CadenceParserPubSet-34))|(1<<(CadenceParserLet-34))|(1<<(CadenceParserVar-34))|(1<<(CadenceParserFrom-34))|(1<<(CadenceParserCreate-34))|(1<<(CadenceParserDestroy-34))|(1<<(CadenceParserIdentifier-34)))) != 0 {
		{
			p.SetState(259)
			p.Member(functionBlockRequired)
		}

		p.SetState(264)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IMemberContext is an interface to support dynamic dispatch.
type IMemberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetFunctionBlockRequired returns the functionBlockRequired attribute.
	GetFunctionBlockRequired() bool

	// SetFunctionBlockRequired sets the functionBlockRequired attribute.
	SetFunctionBlockRequired(bool)

	// IsMemberContext differentiates from other interfaces.
	IsMemberContext()
}

type MemberContext struct {
	*antlr.BaseParserRuleContext
	parser                antlr.Parser
	functionBlockRequired bool
}

func NewEmptyMemberContext() *MemberContext {
	var p = new(MemberContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CadenceParserRULE_member
	return p
}

func (*MemberContext) IsMemberContext() {}

func NewMemberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int, functionBlockRequired bool) *MemberContext {
	var p = new(MemberContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CadenceParserRULE_member

	p.functionBlockRequired = functionBlockRequired

	return p
}

func (s *MemberContext) GetParser() antlr.Parser { return s.parser }

func (s *MemberContext) GetFunctionBlockRequired() bool { return s.functionBlockRequired }

func (s *MemberContext) SetFunctionBlockRequired(v bool) { s.functionBlockRequired = v }

func (s *MemberContext) Field() IFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
}

func (s *MemberContext) SpecialFunctionDeclaration() ISpecialFunctionDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISpecialFunctionDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISpecialFunctionDeclarationContext)
}

func (s *MemberContext) FunctionDeclaration() IFunctionDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionDeclarationContext)
}

func (s *MemberContext) InterfaceDeclaration() IInterfaceDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInterfaceDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInterfaceDeclarationContext)
}

func (s *MemberContext) CompositeDeclaration() ICompositeDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICompositeDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICompositeDeclarationContext)
}

func (s *MemberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MemberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MemberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.EnterMember(s)
	}
}

func (s *MemberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.ExitMember(s)
	}
}

func (s *MemberContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CadenceVisitor:
		return t.VisitMember(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CadenceParser) Member(functionBlockRequired bool) (localctx IMemberContext) {
	localctx = NewMemberContext(p, p.GetParserRuleContext(), p.GetState(), functionBlockRequired)
	p.EnterRule(localctx, 22, CadenceParserRULE_member)

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

	p.SetState(270)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(265)
			p.Field()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(266)
			p.SpecialFunctionDeclaration(functionBlockRequired)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(267)
			p.FunctionDeclaration(functionBlockRequired)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(268)
			p.InterfaceDeclaration()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(269)
			p.CompositeDeclaration()
		}

	}

	return localctx
}

// ICompositeKindContext is an interface to support dynamic dispatch.
type ICompositeKindContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCompositeKindContext differentiates from other interfaces.
	IsCompositeKindContext()
}

type CompositeKindContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCompositeKindContext() *CompositeKindContext {
	var p = new(CompositeKindContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CadenceParserRULE_compositeKind
	return p
}

func (*CompositeKindContext) IsCompositeKindContext() {}

func NewCompositeKindContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CompositeKindContext {
	var p = new(CompositeKindContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CadenceParserRULE_compositeKind

	return p
}

func (s *CompositeKindContext) GetParser() antlr.Parser { return s.parser }

func (s *CompositeKindContext) Struct() antlr.TerminalNode {
	return s.GetToken(CadenceParserStruct, 0)
}

func (s *CompositeKindContext) Resource() antlr.TerminalNode {
	return s.GetToken(CadenceParserResource, 0)
}

func (s *CompositeKindContext) Contract() antlr.TerminalNode {
	return s.GetToken(CadenceParserContract, 0)
}

func (s *CompositeKindContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompositeKindContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CompositeKindContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.EnterCompositeKind(s)
	}
}

func (s *CompositeKindContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.ExitCompositeKind(s)
	}
}

func (s *CompositeKindContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CadenceVisitor:
		return t.VisitCompositeKind(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CadenceParser) CompositeKind() (localctx ICompositeKindContext) {
	localctx = NewCompositeKindContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, CadenceParserRULE_compositeKind)
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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(272)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-34)&-(0x1f+1)) == 0 && ((1<<uint((_la-34)))&((1<<(CadenceParserStruct-34))|(1<<(CadenceParserResource-34))|(1<<(CadenceParserContract-34)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ISpecialFunctionDeclarationContext is an interface to support dynamic dispatch.
type ISpecialFunctionDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetB returns the b rule contexts.
	GetB() IFunctionBlockContext

	// SetB sets the b rule contexts.
	SetB(IFunctionBlockContext)

	// GetFunctionBlockRequired returns the functionBlockRequired attribute.
	GetFunctionBlockRequired() bool

	// SetFunctionBlockRequired sets the functionBlockRequired attribute.
	SetFunctionBlockRequired(bool)

	// IsSpecialFunctionDeclarationContext differentiates from other interfaces.
	IsSpecialFunctionDeclarationContext()
}

type SpecialFunctionDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser                antlr.Parser
	functionBlockRequired bool
	b                     IFunctionBlockContext
}

func NewEmptySpecialFunctionDeclarationContext() *SpecialFunctionDeclarationContext {
	var p = new(SpecialFunctionDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CadenceParserRULE_specialFunctionDeclaration
	return p
}

func (*SpecialFunctionDeclarationContext) IsSpecialFunctionDeclarationContext() {}

func NewSpecialFunctionDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int, functionBlockRequired bool) *SpecialFunctionDeclarationContext {
	var p = new(SpecialFunctionDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CadenceParserRULE_specialFunctionDeclaration

	p.functionBlockRequired = functionBlockRequired

	return p
}

func (s *SpecialFunctionDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *SpecialFunctionDeclarationContext) GetB() IFunctionBlockContext { return s.b }

func (s *SpecialFunctionDeclarationContext) SetB(v IFunctionBlockContext) { s.b = v }

func (s *SpecialFunctionDeclarationContext) GetFunctionBlockRequired() bool {
	return s.functionBlockRequired
}

func (s *SpecialFunctionDeclarationContext) SetFunctionBlockRequired(v bool) {
	s.functionBlockRequired = v
}

func (s *SpecialFunctionDeclarationContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *SpecialFunctionDeclarationContext) ParameterList() IParameterListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParameterListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParameterListContext)
}

func (s *SpecialFunctionDeclarationContext) FunctionBlock() IFunctionBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionBlockContext)
}

func (s *SpecialFunctionDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SpecialFunctionDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SpecialFunctionDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.EnterSpecialFunctionDeclaration(s)
	}
}

func (s *SpecialFunctionDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.ExitSpecialFunctionDeclaration(s)
	}
}

func (s *SpecialFunctionDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CadenceVisitor:
		return t.VisitSpecialFunctionDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CadenceParser) SpecialFunctionDeclaration(functionBlockRequired bool) (localctx ISpecialFunctionDeclarationContext) {
	localctx = NewSpecialFunctionDeclarationContext(p, p.GetParserRuleContext(), p.GetState(), functionBlockRequired)
	p.EnterRule(localctx, 26, CadenceParserRULE_specialFunctionDeclaration)

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
		p.SetState(274)
		p.Identifier()
	}
	{
		p.SetState(275)
		p.ParameterList()
	}
	p.SetState(277)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(276)

			var _x = p.FunctionBlock()

			localctx.(*SpecialFunctionDeclarationContext).b = _x
		}

	}
	p.SetState(279)

	if !(!localctx.(*SpecialFunctionDeclarationContext).functionBlockRequired || localctx.(*SpecialFunctionDeclarationContext).b != nil) {
		panic(antlr.NewFailedPredicateException(p, " !$functionBlockRequired || $ctx.b != nil ", ""))
	}

	return localctx
}

// IFunctionDeclarationContext is an interface to support dynamic dispatch.
type IFunctionDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetReturnType returns the returnType rule contexts.
	GetReturnType() ITypeAnnotationContext

	// GetB returns the b rule contexts.
	GetB() IFunctionBlockContext

	// SetReturnType sets the returnType rule contexts.
	SetReturnType(ITypeAnnotationContext)

	// SetB sets the b rule contexts.
	SetB(IFunctionBlockContext)

	// GetFunctionBlockRequired returns the functionBlockRequired attribute.
	GetFunctionBlockRequired() bool

	// SetFunctionBlockRequired sets the functionBlockRequired attribute.
	SetFunctionBlockRequired(bool)

	// IsFunctionDeclarationContext differentiates from other interfaces.
	IsFunctionDeclarationContext()
}

type FunctionDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser                antlr.Parser
	functionBlockRequired bool
	returnType            ITypeAnnotationContext
	b                     IFunctionBlockContext
}

func NewEmptyFunctionDeclarationContext() *FunctionDeclarationContext {
	var p = new(FunctionDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CadenceParserRULE_functionDeclaration
	return p
}

func (*FunctionDeclarationContext) IsFunctionDeclarationContext() {}

func NewFunctionDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int, functionBlockRequired bool) *FunctionDeclarationContext {
	var p = new(FunctionDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CadenceParserRULE_functionDeclaration

	p.functionBlockRequired = functionBlockRequired

	return p
}

func (s *FunctionDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionDeclarationContext) GetReturnType() ITypeAnnotationContext { return s.returnType }

func (s *FunctionDeclarationContext) GetB() IFunctionBlockContext { return s.b }

func (s *FunctionDeclarationContext) SetReturnType(v ITypeAnnotationContext) { s.returnType = v }

func (s *FunctionDeclarationContext) SetB(v IFunctionBlockContext) { s.b = v }

func (s *FunctionDeclarationContext) GetFunctionBlockRequired() bool { return s.functionBlockRequired }

func (s *FunctionDeclarationContext) SetFunctionBlockRequired(v bool) { s.functionBlockRequired = v }

func (s *FunctionDeclarationContext) Access() IAccessContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAccessContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAccessContext)
}

func (s *FunctionDeclarationContext) Fun() antlr.TerminalNode {
	return s.GetToken(CadenceParserFun, 0)
}

func (s *FunctionDeclarationContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *FunctionDeclarationContext) ParameterList() IParameterListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParameterListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParameterListContext)
}

func (s *FunctionDeclarationContext) TypeAnnotation() ITypeAnnotationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeAnnotationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeAnnotationContext)
}

func (s *FunctionDeclarationContext) FunctionBlock() IFunctionBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionBlockContext)
}

func (s *FunctionDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.EnterFunctionDeclaration(s)
	}
}

func (s *FunctionDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.ExitFunctionDeclaration(s)
	}
}

func (s *FunctionDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CadenceVisitor:
		return t.VisitFunctionDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CadenceParser) FunctionDeclaration(functionBlockRequired bool) (localctx IFunctionDeclarationContext) {
	localctx = NewFunctionDeclarationContext(p, p.GetParserRuleContext(), p.GetState(), functionBlockRequired)
	p.EnterRule(localctx, 28, CadenceParserRULE_functionDeclaration)

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
		p.SetState(281)
		p.Access()
	}
	{
		p.SetState(282)
		p.Match(CadenceParserFun)
	}
	{
		p.SetState(283)
		p.Identifier()
	}
	{
		p.SetState(284)
		p.ParameterList()
	}
	p.SetState(287)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(285)
			p.Match(CadenceParserT__4)
		}
		{
			p.SetState(286)

			var _x = p.TypeAnnotation()

			localctx.(*FunctionDeclarationContext).returnType = _x
		}

	}
	p.SetState(290)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(289)

			var _x = p.FunctionBlock()

			localctx.(*FunctionDeclarationContext).b = _x
		}

	}
	p.SetState(292)

	if !(!localctx.(*FunctionDeclarationContext).functionBlockRequired || localctx.(*FunctionDeclarationContext).b != nil) {
		panic(antlr.NewFailedPredicateException(p, " !$functionBlockRequired || $ctx.b != nil ", ""))
	}

	return localctx
}

// IEventDeclarationContext is an interface to support dynamic dispatch.
type IEventDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEventDeclarationContext differentiates from other interfaces.
	IsEventDeclarationContext()
}

type EventDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEventDeclarationContext() *EventDeclarationContext {
	var p = new(EventDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CadenceParserRULE_eventDeclaration
	return p
}

func (*EventDeclarationContext) IsEventDeclarationContext() {}

func NewEventDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EventDeclarationContext {
	var p = new(EventDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CadenceParserRULE_eventDeclaration

	return p
}

func (s *EventDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *EventDeclarationContext) Event() antlr.TerminalNode {
	return s.GetToken(CadenceParserEvent, 0)
}

func (s *EventDeclarationContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *EventDeclarationContext) ParameterList() IParameterListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParameterListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParameterListContext)
}

func (s *EventDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EventDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EventDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.EnterEventDeclaration(s)
	}
}

func (s *EventDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.ExitEventDeclaration(s)
	}
}

func (s *EventDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CadenceVisitor:
		return t.VisitEventDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CadenceParser) EventDeclaration() (localctx IEventDeclarationContext) {
	localctx = NewEventDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, CadenceParserRULE_eventDeclaration)

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
		p.SetState(294)
		p.Match(CadenceParserEvent)
	}
	{
		p.SetState(295)
		p.Identifier()
	}
	{
		p.SetState(296)
		p.ParameterList()
	}

	return localctx
}

// IParameterListContext is an interface to support dynamic dispatch.
type IParameterListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParameterListContext differentiates from other interfaces.
	IsParameterListContext()
}

type ParameterListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParameterListContext() *ParameterListContext {
	var p = new(ParameterListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CadenceParserRULE_parameterList
	return p
}

func (*ParameterListContext) IsParameterListContext() {}

func NewParameterListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParameterListContext {
	var p = new(ParameterListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CadenceParserRULE_parameterList

	return p
}

func (s *ParameterListContext) GetParser() antlr.Parser { return s.parser }

func (s *ParameterListContext) OpenParen() antlr.TerminalNode {
	return s.GetToken(CadenceParserOpenParen, 0)
}

func (s *ParameterListContext) CloseParen() antlr.TerminalNode {
	return s.GetToken(CadenceParserCloseParen, 0)
}

func (s *ParameterListContext) AllParameter() []IParameterContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IParameterContext)(nil)).Elem())
	var tst = make([]IParameterContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IParameterContext)
		}
	}

	return tst
}

func (s *ParameterListContext) Parameter(i int) IParameterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParameterContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IParameterContext)
}

func (s *ParameterListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParameterListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParameterListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.EnterParameterList(s)
	}
}

func (s *ParameterListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.ExitParameterList(s)
	}
}

func (s *ParameterListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CadenceVisitor:
		return t.VisitParameterList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CadenceParser) ParameterList() (localctx IParameterListContext) {
	localctx = NewParameterListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, CadenceParserRULE_parameterList)
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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(298)
		p.Match(CadenceParserOpenParen)
	}
	p.SetState(307)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la-58)&-(0x1f+1)) == 0 && ((1<<uint((_la-58)))&((1<<(CadenceParserFrom-58))|(1<<(CadenceParserCreate-58))|(1<<(CadenceParserDestroy-58))|(1<<(CadenceParserIdentifier-58)))) != 0 {
		{
			p.SetState(299)
			p.Parameter()
		}
		p.SetState(304)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == CadenceParserT__1 {
			{
				p.SetState(300)
				p.Match(CadenceParserT__1)
			}
			{
				p.SetState(301)
				p.Parameter()
			}

			p.SetState(306)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(309)
		p.Match(CadenceParserCloseParen)
	}

	return localctx
}

// IParameterContext is an interface to support dynamic dispatch.
type IParameterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetArgumentLabel returns the argumentLabel rule contexts.
	GetArgumentLabel() IIdentifierContext

	// GetParameterName returns the parameterName rule contexts.
	GetParameterName() IIdentifierContext

	// SetArgumentLabel sets the argumentLabel rule contexts.
	SetArgumentLabel(IIdentifierContext)

	// SetParameterName sets the parameterName rule contexts.
	SetParameterName(IIdentifierContext)

	// IsParameterContext differentiates from other interfaces.
	IsParameterContext()
}

type ParameterContext struct {
	*antlr.BaseParserRuleContext
	parser        antlr.Parser
	argumentLabel IIdentifierContext
	parameterName IIdentifierContext
}

func NewEmptyParameterContext() *ParameterContext {
	var p = new(ParameterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CadenceParserRULE_parameter
	return p
}

func (*ParameterContext) IsParameterContext() {}

func NewParameterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParameterContext {
	var p = new(ParameterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CadenceParserRULE_parameter

	return p
}

func (s *ParameterContext) GetParser() antlr.Parser { return s.parser }

func (s *ParameterContext) GetArgumentLabel() IIdentifierContext { return s.argumentLabel }

func (s *ParameterContext) GetParameterName() IIdentifierContext { return s.parameterName }

func (s *ParameterContext) SetArgumentLabel(v IIdentifierContext) { s.argumentLabel = v }

func (s *ParameterContext) SetParameterName(v IIdentifierContext) { s.parameterName = v }

func (s *ParameterContext) TypeAnnotation() ITypeAnnotationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeAnnotationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeAnnotationContext)
}

func (s *ParameterContext) AllIdentifier() []IIdentifierContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIdentifierContext)(nil)).Elem())
	var tst = make([]IIdentifierContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIdentifierContext)
		}
	}

	return tst
}

func (s *ParameterContext) Identifier(i int) IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *ParameterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParameterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParameterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.EnterParameter(s)
	}
}

func (s *ParameterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.ExitParameter(s)
	}
}

func (s *ParameterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CadenceVisitor:
		return t.VisitParameter(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CadenceParser) Parameter() (localctx IParameterContext) {
	localctx = NewParameterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, CadenceParserRULE_parameter)

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
	p.SetState(312)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(311)

			var _x = p.Identifier()

			localctx.(*ParameterContext).argumentLabel = _x
		}

	}
	{
		p.SetState(314)

		var _x = p.Identifier()

		localctx.(*ParameterContext).parameterName = _x
	}
	{
		p.SetState(315)
		p.Match(CadenceParserT__4)
	}
	{
		p.SetState(316)
		p.TypeAnnotation()
	}

	return localctx
}

// ITypeAnnotationContext is an interface to support dynamic dispatch.
type ITypeAnnotationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeAnnotationContext differentiates from other interfaces.
	IsTypeAnnotationContext()
}

type TypeAnnotationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeAnnotationContext() *TypeAnnotationContext {
	var p = new(TypeAnnotationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CadenceParserRULE_typeAnnotation
	return p
}

func (*TypeAnnotationContext) IsTypeAnnotationContext() {}

func NewTypeAnnotationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeAnnotationContext {
	var p = new(TypeAnnotationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CadenceParserRULE_typeAnnotation

	return p
}

func (s *TypeAnnotationContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeAnnotationContext) FullType() IFullTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFullTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFullTypeContext)
}

func (s *TypeAnnotationContext) Move() antlr.TerminalNode {
	return s.GetToken(CadenceParserMove, 0)
}

func (s *TypeAnnotationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeAnnotationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeAnnotationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.EnterTypeAnnotation(s)
	}
}

func (s *TypeAnnotationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.ExitTypeAnnotation(s)
	}
}

func (s *TypeAnnotationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CadenceVisitor:
		return t.VisitTypeAnnotation(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CadenceParser) TypeAnnotation() (localctx ITypeAnnotationContext) {
	localctx = NewTypeAnnotationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, CadenceParserRULE_typeAnnotation)
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

	p.EnterOuterAlt(localctx, 1)
	p.SetState(319)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CadenceParserMove {
		{
			p.SetState(318)
			p.Match(CadenceParserMove)
		}

	}
	{
		p.SetState(321)
		p.FullType()
	}

	return localctx
}

// IFullTypeContext is an interface to support dynamic dispatch.
type IFullTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetReference returns the reference token.
	GetReference() antlr.Token

	// Get_Optional returns the _Optional token.
	Get_Optional() antlr.Token

	// SetReference sets the reference token.
	SetReference(antlr.Token)

	// Set_Optional sets the _Optional token.
	Set_Optional(antlr.Token)

	// GetOptionals returns the optionals token list.
	GetOptionals() []antlr.Token

	// SetOptionals sets the optionals token list.
	SetOptionals([]antlr.Token)

	// IsFullTypeContext differentiates from other interfaces.
	IsFullTypeContext()
}

type FullTypeContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	reference antlr.Token
	_Optional antlr.Token
	optionals []antlr.Token
}

func NewEmptyFullTypeContext() *FullTypeContext {
	var p = new(FullTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CadenceParserRULE_fullType
	return p
}

func (*FullTypeContext) IsFullTypeContext() {}

func NewFullTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FullTypeContext {
	var p = new(FullTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CadenceParserRULE_fullType

	return p
}

func (s *FullTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *FullTypeContext) GetReference() antlr.Token { return s.reference }

func (s *FullTypeContext) Get_Optional() antlr.Token { return s._Optional }

func (s *FullTypeContext) SetReference(v antlr.Token) { s.reference = v }

func (s *FullTypeContext) Set_Optional(v antlr.Token) { s._Optional = v }

func (s *FullTypeContext) GetOptionals() []antlr.Token { return s.optionals }

func (s *FullTypeContext) SetOptionals(v []antlr.Token) { s.optionals = v }

func (s *FullTypeContext) BaseType() IBaseTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBaseTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBaseTypeContext)
}

func (s *FullTypeContext) Ampersand() antlr.TerminalNode {
	return s.GetToken(CadenceParserAmpersand, 0)
}

func (s *FullTypeContext) AllOptional() []antlr.TerminalNode {
	return s.GetTokens(CadenceParserOptional)
}

func (s *FullTypeContext) Optional(i int) antlr.TerminalNode {
	return s.GetToken(CadenceParserOptional, i)
}

func (s *FullTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FullTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FullTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.EnterFullType(s)
	}
}

func (s *FullTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.ExitFullType(s)
	}
}

func (s *FullTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CadenceVisitor:
		return t.VisitFullType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CadenceParser) FullType() (localctx IFullTypeContext) {
	localctx = NewFullTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, CadenceParserRULE_fullType)

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

	var _alt int

	p.SetState(334)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CadenceParserAmpersand:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(323)

			var _m = p.Match(CadenceParserAmpersand)

			localctx.(*FullTypeContext).reference = _m
		}
		p.SetState(324)

		if !(p.noWhitespace()) {
			panic(antlr.NewFailedPredicateException(p, "p.noWhitespace()", ""))
		}
		{
			p.SetState(325)
			p.BaseType()
		}

	case CadenceParserT__2, CadenceParserT__5, CadenceParserOpenParen, CadenceParserFrom, CadenceParserCreate, CadenceParserDestroy, CadenceParserIdentifier:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(326)
			p.BaseType()
		}
		p.SetState(331)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				p.SetState(327)

				if !(p.noWhitespace()) {
					panic(antlr.NewFailedPredicateException(p, "p.noWhitespace()", ""))
				}
				{
					p.SetState(328)

					var _m = p.Match(CadenceParserOptional)

					localctx.(*FullTypeContext)._Optional = _m
				}
				localctx.(*FullTypeContext).optionals = append(localctx.(*FullTypeContext).optionals, localctx.(*FullTypeContext)._Optional)

			}
			p.SetState(333)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext())
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IBaseTypeContext is an interface to support dynamic dispatch.
type IBaseTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBaseTypeContext differentiates from other interfaces.
	IsBaseTypeContext()
}

type BaseTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBaseTypeContext() *BaseTypeContext {
	var p = new(BaseTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CadenceParserRULE_baseType
	return p
}

func (*BaseTypeContext) IsBaseTypeContext() {}

func NewBaseTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BaseTypeContext {
	var p = new(BaseTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CadenceParserRULE_baseType

	return p
}

func (s *BaseTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *BaseTypeContext) NominalType() INominalTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INominalTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INominalTypeContext)
}

func (s *BaseTypeContext) FunctionType() IFunctionTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionTypeContext)
}

func (s *BaseTypeContext) VariableSizedType() IVariableSizedTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableSizedTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableSizedTypeContext)
}

func (s *BaseTypeContext) ConstantSizedType() IConstantSizedTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstantSizedTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstantSizedTypeContext)
}

func (s *BaseTypeContext) DictionaryType() IDictionaryTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDictionaryTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDictionaryTypeContext)
}

func (s *BaseTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BaseTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BaseTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.EnterBaseType(s)
	}
}

func (s *BaseTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.ExitBaseType(s)
	}
}

func (s *BaseTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CadenceVisitor:
		return t.VisitBaseType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CadenceParser) BaseType() (localctx IBaseTypeContext) {
	localctx = NewBaseTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, CadenceParserRULE_baseType)

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

	p.SetState(341)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(336)
			p.NominalType()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(337)
			p.FunctionType()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(338)
			p.VariableSizedType()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(339)
			p.ConstantSizedType()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(340)
			p.DictionaryType()
		}

	}

	return localctx
}

// INominalTypeContext is an interface to support dynamic dispatch.
type INominalTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNominalTypeContext differentiates from other interfaces.
	IsNominalTypeContext()
}

type NominalTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNominalTypeContext() *NominalTypeContext {
	var p = new(NominalTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CadenceParserRULE_nominalType
	return p
}

func (*NominalTypeContext) IsNominalTypeContext() {}

func NewNominalTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NominalTypeContext {
	var p = new(NominalTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CadenceParserRULE_nominalType

	return p
}

func (s *NominalTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *NominalTypeContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *NominalTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NominalTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NominalTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.EnterNominalType(s)
	}
}

func (s *NominalTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.ExitNominalType(s)
	}
}

func (s *NominalTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CadenceVisitor:
		return t.VisitNominalType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CadenceParser) NominalType() (localctx INominalTypeContext) {
	localctx = NewNominalTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, CadenceParserRULE_nominalType)

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
		p.SetState(343)
		p.Identifier()
	}

	return localctx
}

// IFunctionTypeContext is an interface to support dynamic dispatch.
type IFunctionTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_typeAnnotation returns the _typeAnnotation rule contexts.
	Get_typeAnnotation() ITypeAnnotationContext

	// GetReturnType returns the returnType rule contexts.
	GetReturnType() ITypeAnnotationContext

	// Set_typeAnnotation sets the _typeAnnotation rule contexts.
	Set_typeAnnotation(ITypeAnnotationContext)

	// SetReturnType sets the returnType rule contexts.
	SetReturnType(ITypeAnnotationContext)

	// GetParameterTypes returns the parameterTypes rule context list.
	GetParameterTypes() []ITypeAnnotationContext

	// SetParameterTypes sets the parameterTypes rule context list.
	SetParameterTypes([]ITypeAnnotationContext)

	// IsFunctionTypeContext differentiates from other interfaces.
	IsFunctionTypeContext()
}

type FunctionTypeContext struct {
	*antlr.BaseParserRuleContext
	parser          antlr.Parser
	_typeAnnotation ITypeAnnotationContext
	parameterTypes  []ITypeAnnotationContext
	returnType      ITypeAnnotationContext
}

func NewEmptyFunctionTypeContext() *FunctionTypeContext {
	var p = new(FunctionTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CadenceParserRULE_functionType
	return p
}

func (*FunctionTypeContext) IsFunctionTypeContext() {}

func NewFunctionTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionTypeContext {
	var p = new(FunctionTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CadenceParserRULE_functionType

	return p
}

func (s *FunctionTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionTypeContext) Get_typeAnnotation() ITypeAnnotationContext { return s._typeAnnotation }

func (s *FunctionTypeContext) GetReturnType() ITypeAnnotationContext { return s.returnType }

func (s *FunctionTypeContext) Set_typeAnnotation(v ITypeAnnotationContext) { s._typeAnnotation = v }

func (s *FunctionTypeContext) SetReturnType(v ITypeAnnotationContext) { s.returnType = v }

func (s *FunctionTypeContext) GetParameterTypes() []ITypeAnnotationContext { return s.parameterTypes }

func (s *FunctionTypeContext) SetParameterTypes(v []ITypeAnnotationContext) { s.parameterTypes = v }

func (s *FunctionTypeContext) AllOpenParen() []antlr.TerminalNode {
	return s.GetTokens(CadenceParserOpenParen)
}

func (s *FunctionTypeContext) OpenParen(i int) antlr.TerminalNode {
	return s.GetToken(CadenceParserOpenParen, i)
}

func (s *FunctionTypeContext) AllCloseParen() []antlr.TerminalNode {
	return s.GetTokens(CadenceParserCloseParen)
}

func (s *FunctionTypeContext) CloseParen(i int) antlr.TerminalNode {
	return s.GetToken(CadenceParserCloseParen, i)
}

func (s *FunctionTypeContext) AllTypeAnnotation() []ITypeAnnotationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITypeAnnotationContext)(nil)).Elem())
	var tst = make([]ITypeAnnotationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITypeAnnotationContext)
		}
	}

	return tst
}

func (s *FunctionTypeContext) TypeAnnotation(i int) ITypeAnnotationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeAnnotationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITypeAnnotationContext)
}

func (s *FunctionTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.EnterFunctionType(s)
	}
}

func (s *FunctionTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.ExitFunctionType(s)
	}
}

func (s *FunctionTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CadenceVisitor:
		return t.VisitFunctionType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CadenceParser) FunctionType() (localctx IFunctionTypeContext) {
	localctx = NewFunctionTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, CadenceParserRULE_functionType)
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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(345)
		p.Match(CadenceParserOpenParen)
	}
	{
		p.SetState(346)
		p.Match(CadenceParserOpenParen)
	}
	p.SetState(355)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<CadenceParserT__2)|(1<<CadenceParserT__5)|(1<<CadenceParserAmpersand)|(1<<CadenceParserMove)|(1<<CadenceParserOpenParen))) != 0) || (((_la-58)&-(0x1f+1)) == 0 && ((1<<uint((_la-58)))&((1<<(CadenceParserFrom-58))|(1<<(CadenceParserCreate-58))|(1<<(CadenceParserDestroy-58))|(1<<(CadenceParserIdentifier-58)))) != 0) {
		{
			p.SetState(347)

			var _x = p.TypeAnnotation()

			localctx.(*FunctionTypeContext)._typeAnnotation = _x
		}
		localctx.(*FunctionTypeContext).parameterTypes = append(localctx.(*FunctionTypeContext).parameterTypes, localctx.(*FunctionTypeContext)._typeAnnotation)
		p.SetState(352)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == CadenceParserT__1 {
			{
				p.SetState(348)
				p.Match(CadenceParserT__1)
			}
			{
				p.SetState(349)

				var _x = p.TypeAnnotation()

				localctx.(*FunctionTypeContext)._typeAnnotation = _x
			}
			localctx.(*FunctionTypeContext).parameterTypes = append(localctx.(*FunctionTypeContext).parameterTypes, localctx.(*FunctionTypeContext)._typeAnnotation)

			p.SetState(354)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(357)
		p.Match(CadenceParserCloseParen)
	}
	{
		p.SetState(358)
		p.Match(CadenceParserT__4)
	}
	{
		p.SetState(359)

		var _x = p.TypeAnnotation()

		localctx.(*FunctionTypeContext).returnType = _x
	}
	{
		p.SetState(360)
		p.Match(CadenceParserCloseParen)
	}

	return localctx
}

// IVariableSizedTypeContext is an interface to support dynamic dispatch.
type IVariableSizedTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableSizedTypeContext differentiates from other interfaces.
	IsVariableSizedTypeContext()
}

type VariableSizedTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableSizedTypeContext() *VariableSizedTypeContext {
	var p = new(VariableSizedTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CadenceParserRULE_variableSizedType
	return p
}

func (*VariableSizedTypeContext) IsVariableSizedTypeContext() {}

func NewVariableSizedTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableSizedTypeContext {
	var p = new(VariableSizedTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CadenceParserRULE_variableSizedType

	return p
}

func (s *VariableSizedTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableSizedTypeContext) FullType() IFullTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFullTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFullTypeContext)
}

func (s *VariableSizedTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableSizedTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableSizedTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.EnterVariableSizedType(s)
	}
}

func (s *VariableSizedTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.ExitVariableSizedType(s)
	}
}

func (s *VariableSizedTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CadenceVisitor:
		return t.VisitVariableSizedType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CadenceParser) VariableSizedType() (localctx IVariableSizedTypeContext) {
	localctx = NewVariableSizedTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, CadenceParserRULE_variableSizedType)

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
		p.SetState(362)
		p.Match(CadenceParserT__5)
	}
	{
		p.SetState(363)
		p.FullType()
	}
	{
		p.SetState(364)
		p.Match(CadenceParserT__6)
	}

	return localctx
}

// IConstantSizedTypeContext is an interface to support dynamic dispatch.
type IConstantSizedTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetSize returns the size token.
	GetSize() antlr.Token

	// SetSize sets the size token.
	SetSize(antlr.Token)

	// IsConstantSizedTypeContext differentiates from other interfaces.
	IsConstantSizedTypeContext()
}

type ConstantSizedTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	size   antlr.Token
}

func NewEmptyConstantSizedTypeContext() *ConstantSizedTypeContext {
	var p = new(ConstantSizedTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CadenceParserRULE_constantSizedType
	return p
}

func (*ConstantSizedTypeContext) IsConstantSizedTypeContext() {}

func NewConstantSizedTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstantSizedTypeContext {
	var p = new(ConstantSizedTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CadenceParserRULE_constantSizedType

	return p
}

func (s *ConstantSizedTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstantSizedTypeContext) GetSize() antlr.Token { return s.size }

func (s *ConstantSizedTypeContext) SetSize(v antlr.Token) { s.size = v }

func (s *ConstantSizedTypeContext) FullType() IFullTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFullTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFullTypeContext)
}

func (s *ConstantSizedTypeContext) DecimalLiteral() antlr.TerminalNode {
	return s.GetToken(CadenceParserDecimalLiteral, 0)
}

func (s *ConstantSizedTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantSizedTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstantSizedTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.EnterConstantSizedType(s)
	}
}

func (s *ConstantSizedTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.ExitConstantSizedType(s)
	}
}

func (s *ConstantSizedTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CadenceVisitor:
		return t.VisitConstantSizedType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CadenceParser) ConstantSizedType() (localctx IConstantSizedTypeContext) {
	localctx = NewConstantSizedTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, CadenceParserRULE_constantSizedType)

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
		p.SetState(366)
		p.Match(CadenceParserT__5)
	}
	{
		p.SetState(367)
		p.FullType()
	}
	{
		p.SetState(368)
		p.Match(CadenceParserT__0)
	}
	{
		p.SetState(369)

		var _m = p.Match(CadenceParserDecimalLiteral)

		localctx.(*ConstantSizedTypeContext).size = _m
	}
	{
		p.SetState(370)
		p.Match(CadenceParserT__6)
	}

	return localctx
}

// IDictionaryTypeContext is an interface to support dynamic dispatch.
type IDictionaryTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetKeyType returns the keyType rule contexts.
	GetKeyType() IFullTypeContext

	// GetValueType returns the valueType rule contexts.
	GetValueType() IFullTypeContext

	// SetKeyType sets the keyType rule contexts.
	SetKeyType(IFullTypeContext)

	// SetValueType sets the valueType rule contexts.
	SetValueType(IFullTypeContext)

	// IsDictionaryTypeContext differentiates from other interfaces.
	IsDictionaryTypeContext()
}

type DictionaryTypeContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	keyType   IFullTypeContext
	valueType IFullTypeContext
}

func NewEmptyDictionaryTypeContext() *DictionaryTypeContext {
	var p = new(DictionaryTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CadenceParserRULE_dictionaryType
	return p
}

func (*DictionaryTypeContext) IsDictionaryTypeContext() {}

func NewDictionaryTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DictionaryTypeContext {
	var p = new(DictionaryTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CadenceParserRULE_dictionaryType

	return p
}

func (s *DictionaryTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *DictionaryTypeContext) GetKeyType() IFullTypeContext { return s.keyType }

func (s *DictionaryTypeContext) GetValueType() IFullTypeContext { return s.valueType }

func (s *DictionaryTypeContext) SetKeyType(v IFullTypeContext) { s.keyType = v }

func (s *DictionaryTypeContext) SetValueType(v IFullTypeContext) { s.valueType = v }

func (s *DictionaryTypeContext) AllFullType() []IFullTypeContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFullTypeContext)(nil)).Elem())
	var tst = make([]IFullTypeContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFullTypeContext)
		}
	}

	return tst
}

func (s *DictionaryTypeContext) FullType(i int) IFullTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFullTypeContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFullTypeContext)
}

func (s *DictionaryTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DictionaryTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DictionaryTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.EnterDictionaryType(s)
	}
}

func (s *DictionaryTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.ExitDictionaryType(s)
	}
}

func (s *DictionaryTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CadenceVisitor:
		return t.VisitDictionaryType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CadenceParser) DictionaryType() (localctx IDictionaryTypeContext) {
	localctx = NewDictionaryTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, CadenceParserRULE_dictionaryType)

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
		p.SetState(372)
		p.Match(CadenceParserT__2)
	}
	{
		p.SetState(373)

		var _x = p.FullType()

		localctx.(*DictionaryTypeContext).keyType = _x
	}
	{
		p.SetState(374)
		p.Match(CadenceParserT__4)
	}
	{
		p.SetState(375)

		var _x = p.FullType()

		localctx.(*DictionaryTypeContext).valueType = _x
	}
	{
		p.SetState(376)
		p.Match(CadenceParserT__3)
	}

	return localctx
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CadenceParserRULE_block
	return p
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CadenceParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) Statements() IStatementsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementsContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (s *BlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CadenceVisitor:
		return t.VisitBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CadenceParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, CadenceParserRULE_block)

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
		p.SetState(378)
		p.Match(CadenceParserT__2)
	}
	{
		p.SetState(379)
		p.Statements()
	}
	{
		p.SetState(380)
		p.Match(CadenceParserT__3)
	}

	return localctx
}

// IFunctionBlockContext is an interface to support dynamic dispatch.
type IFunctionBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionBlockContext differentiates from other interfaces.
	IsFunctionBlockContext()
}

type FunctionBlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionBlockContext() *FunctionBlockContext {
	var p = new(FunctionBlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CadenceParserRULE_functionBlock
	return p
}

func (*FunctionBlockContext) IsFunctionBlockContext() {}

func NewFunctionBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionBlockContext {
	var p = new(FunctionBlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CadenceParserRULE_functionBlock

	return p
}

func (s *FunctionBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionBlockContext) Statements() IStatementsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementsContext)
}

func (s *FunctionBlockContext) PreConditions() IPreConditionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPreConditionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPreConditionsContext)
}

func (s *FunctionBlockContext) PostConditions() IPostConditionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPostConditionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPostConditionsContext)
}

func (s *FunctionBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.EnterFunctionBlock(s)
	}
}

func (s *FunctionBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.ExitFunctionBlock(s)
	}
}

func (s *FunctionBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CadenceVisitor:
		return t.VisitFunctionBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CadenceParser) FunctionBlock() (localctx IFunctionBlockContext) {
	localctx = NewFunctionBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, CadenceParserRULE_functionBlock)
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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(382)
		p.Match(CadenceParserT__2)
	}
	p.SetState(384)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CadenceParserPre {
		{
			p.SetState(383)
			p.PreConditions()
		}

	}
	p.SetState(387)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CadenceParserPost {
		{
			p.SetState(386)
			p.PostConditions()
		}

	}
	{
		p.SetState(389)
		p.Statements()
	}
	{
		p.SetState(390)
		p.Match(CadenceParserT__3)
	}

	return localctx
}

// IPreConditionsContext is an interface to support dynamic dispatch.
type IPreConditionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPreConditionsContext differentiates from other interfaces.
	IsPreConditionsContext()
}

type PreConditionsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPreConditionsContext() *PreConditionsContext {
	var p = new(PreConditionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CadenceParserRULE_preConditions
	return p
}

func (*PreConditionsContext) IsPreConditionsContext() {}

func NewPreConditionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PreConditionsContext {
	var p = new(PreConditionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CadenceParserRULE_preConditions

	return p
}

func (s *PreConditionsContext) GetParser() antlr.Parser { return s.parser }

func (s *PreConditionsContext) Pre() antlr.TerminalNode {
	return s.GetToken(CadenceParserPre, 0)
}

func (s *PreConditionsContext) Conditions() IConditionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConditionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConditionsContext)
}

func (s *PreConditionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PreConditionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PreConditionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.EnterPreConditions(s)
	}
}

func (s *PreConditionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.ExitPreConditions(s)
	}
}

func (s *PreConditionsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CadenceVisitor:
		return t.VisitPreConditions(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CadenceParser) PreConditions() (localctx IPreConditionsContext) {
	localctx = NewPreConditionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, CadenceParserRULE_preConditions)

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
		p.SetState(392)
		p.Match(CadenceParserPre)
	}
	{
		p.SetState(393)
		p.Match(CadenceParserT__2)
	}
	{
		p.SetState(394)
		p.Conditions()
	}
	{
		p.SetState(395)
		p.Match(CadenceParserT__3)
	}

	return localctx
}

// IPostConditionsContext is an interface to support dynamic dispatch.
type IPostConditionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPostConditionsContext differentiates from other interfaces.
	IsPostConditionsContext()
}

type PostConditionsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPostConditionsContext() *PostConditionsContext {
	var p = new(PostConditionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CadenceParserRULE_postConditions
	return p
}

func (*PostConditionsContext) IsPostConditionsContext() {}

func NewPostConditionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PostConditionsContext {
	var p = new(PostConditionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CadenceParserRULE_postConditions

	return p
}

func (s *PostConditionsContext) GetParser() antlr.Parser { return s.parser }

func (s *PostConditionsContext) Post() antlr.TerminalNode {
	return s.GetToken(CadenceParserPost, 0)
}

func (s *PostConditionsContext) Conditions() IConditionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConditionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConditionsContext)
}

func (s *PostConditionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PostConditionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PostConditionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.EnterPostConditions(s)
	}
}

func (s *PostConditionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.ExitPostConditions(s)
	}
}

func (s *PostConditionsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CadenceVisitor:
		return t.VisitPostConditions(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CadenceParser) PostConditions() (localctx IPostConditionsContext) {
	localctx = NewPostConditionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, CadenceParserRULE_postConditions)

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
		p.SetState(397)
		p.Match(CadenceParserPost)
	}
	{
		p.SetState(398)
		p.Match(CadenceParserT__2)
	}
	{
		p.SetState(399)
		p.Conditions()
	}
	{
		p.SetState(400)
		p.Match(CadenceParserT__3)
	}

	return localctx
}

// IConditionsContext is an interface to support dynamic dispatch.
type IConditionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConditionsContext differentiates from other interfaces.
	IsConditionsContext()
}

type ConditionsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionsContext() *ConditionsContext {
	var p = new(ConditionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CadenceParserRULE_conditions
	return p
}

func (*ConditionsContext) IsConditionsContext() {}

func NewConditionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionsContext {
	var p = new(ConditionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CadenceParserRULE_conditions

	return p
}

func (s *ConditionsContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionsContext) AllCondition() []IConditionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IConditionContext)(nil)).Elem())
	var tst = make([]IConditionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IConditionContext)
		}
	}

	return tst
}

func (s *ConditionsContext) Condition(i int) IConditionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConditionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IConditionContext)
}

func (s *ConditionsContext) AllEos() []IEosContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEosContext)(nil)).Elem())
	var tst = make([]IEosContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEosContext)
		}
	}

	return tst
}

func (s *ConditionsContext) Eos(i int) IEosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEosContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *ConditionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.EnterConditions(s)
	}
}

func (s *ConditionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.ExitConditions(s)
	}
}

func (s *ConditionsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CadenceVisitor:
		return t.VisitConditions(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CadenceParser) Conditions() (localctx IConditionsContext) {
	localctx = NewConditionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, CadenceParserRULE_conditions)
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

	p.EnterOuterAlt(localctx, 1)
	p.SetState(407)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<CadenceParserT__2)|(1<<CadenceParserT__5)|(1<<CadenceParserMinus)|(1<<CadenceParserAmpersand)|(1<<CadenceParserNegate)|(1<<CadenceParserMove)|(1<<CadenceParserOpenParen))) != 0) || (((_la-38)&-(0x1f+1)) == 0 && ((1<<uint((_la-38)))&((1<<(CadenceParserFun-38))|(1<<(CadenceParserTrue-38))|(1<<(CadenceParserFalse-38))|(1<<(CadenceParserNil-38))|(1<<(CadenceParserFrom-38))|(1<<(CadenceParserCreate-38))|(1<<(CadenceParserDestroy-38))|(1<<(CadenceParserIdentifier-38))|(1<<(CadenceParserDecimalLiteral-38))|(1<<(CadenceParserBinaryLiteral-38))|(1<<(CadenceParserOctalLiteral-38))|(1<<(CadenceParserHexadecimalLiteral-38))|(1<<(CadenceParserInvalidNumberLiteral-38))|(1<<(CadenceParserStringLiteral-38)))) != 0) {
		{
			p.SetState(402)
			p.Condition()
		}
		{
			p.SetState(403)
			p.Eos()
		}

		p.SetState(409)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IConditionContext is an interface to support dynamic dispatch.
type IConditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetTest returns the test rule contexts.
	GetTest() IExpressionContext

	// GetMessage returns the message rule contexts.
	GetMessage() IExpressionContext

	// SetTest sets the test rule contexts.
	SetTest(IExpressionContext)

	// SetMessage sets the message rule contexts.
	SetMessage(IExpressionContext)

	// IsConditionContext differentiates from other interfaces.
	IsConditionContext()
}

type ConditionContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	test    IExpressionContext
	message IExpressionContext
}

func NewEmptyConditionContext() *ConditionContext {
	var p = new(ConditionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CadenceParserRULE_condition
	return p
}

func (*ConditionContext) IsConditionContext() {}

func NewConditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionContext {
	var p = new(ConditionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CadenceParserRULE_condition

	return p
}

func (s *ConditionContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionContext) GetTest() IExpressionContext { return s.test }

func (s *ConditionContext) GetMessage() IExpressionContext { return s.message }

func (s *ConditionContext) SetTest(v IExpressionContext) { s.test = v }

func (s *ConditionContext) SetMessage(v IExpressionContext) { s.message = v }

func (s *ConditionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ConditionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.EnterCondition(s)
	}
}

func (s *ConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.ExitCondition(s)
	}
}

func (s *ConditionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CadenceVisitor:
		return t.VisitCondition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CadenceParser) Condition() (localctx IConditionContext) {
	localctx = NewConditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, CadenceParserRULE_condition)

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
		p.SetState(410)

		var _x = p.Expression()

		localctx.(*ConditionContext).test = _x
	}
	p.SetState(413)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(411)
			p.Match(CadenceParserT__4)
		}
		{
			p.SetState(412)

			var _x = p.Expression()

			localctx.(*ConditionContext).message = _x
		}

	}

	return localctx
}

// IStatementsContext is an interface to support dynamic dispatch.
type IStatementsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementsContext differentiates from other interfaces.
	IsStatementsContext()
}

type StatementsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementsContext() *StatementsContext {
	var p = new(StatementsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CadenceParserRULE_statements
	return p
}

func (*StatementsContext) IsStatementsContext() {}

func NewStatementsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementsContext {
	var p = new(StatementsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CadenceParserRULE_statements

	return p
}

func (s *StatementsContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementsContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *StatementsContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *StatementsContext) AllEos() []IEosContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEosContext)(nil)).Elem())
	var tst = make([]IEosContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEosContext)
		}
	}

	return tst
}

func (s *StatementsContext) Eos(i int) IEosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEosContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *StatementsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.EnterStatements(s)
	}
}

func (s *StatementsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.ExitStatements(s)
	}
}

func (s *StatementsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CadenceVisitor:
		return t.VisitStatements(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CadenceParser) Statements() (localctx IStatementsContext) {
	localctx = NewStatementsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, CadenceParserRULE_statements)
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

	p.EnterOuterAlt(localctx, 1)
	p.SetState(420)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<CadenceParserT__2)|(1<<CadenceParserT__5)|(1<<CadenceParserMinus)|(1<<CadenceParserAmpersand)|(1<<CadenceParserNegate)|(1<<CadenceParserMove)|(1<<CadenceParserOpenParen))) != 0) || (((_la-34)&-(0x1f+1)) == 0 && ((1<<uint((_la-34)))&((1<<(CadenceParserStruct-34))|(1<<(CadenceParserResource-34))|(1<<(CadenceParserContract-34))|(1<<(CadenceParserFun-34))|(1<<(CadenceParserEvent-34))|(1<<(CadenceParserEmit-34))|(1<<(CadenceParserPriv-34))|(1<<(CadenceParserPub-34))|(1<<(CadenceParserPubSet-34))|(1<<(CadenceParserReturn-34))|(1<<(CadenceParserBreak-34))|(1<<(CadenceParserContinue-34))|(1<<(CadenceParserLet-34))|(1<<(CadenceParserVar-34))|(1<<(CadenceParserIf-34))|(1<<(CadenceParserWhile-34))|(1<<(CadenceParserTrue-34))|(1<<(CadenceParserFalse-34))|(1<<(CadenceParserNil-34))|(1<<(CadenceParserImport-34))|(1<<(CadenceParserFrom-34))|(1<<(CadenceParserCreate-34))|(1<<(CadenceParserDestroy-34))|(1<<(CadenceParserIdentifier-34))|(1<<(CadenceParserDecimalLiteral-34))|(1<<(CadenceParserBinaryLiteral-34))|(1<<(CadenceParserOctalLiteral-34))|(1<<(CadenceParserHexadecimalLiteral-34)))) != 0) || _la == CadenceParserInvalidNumberLiteral || _la == CadenceParserStringLiteral {
		{
			p.SetState(415)
			p.Statement()
		}
		{
			p.SetState(416)
			p.Eos()
		}

		p.SetState(422)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CadenceParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CadenceParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) ReturnStatement() IReturnStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReturnStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReturnStatementContext)
}

func (s *StatementContext) BreakStatement() IBreakStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBreakStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBreakStatementContext)
}

func (s *StatementContext) ContinueStatement() IContinueStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IContinueStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IContinueStatementContext)
}

func (s *StatementContext) IfStatement() IIfStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIfStatementContext)
}

func (s *StatementContext) WhileStatement() IWhileStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWhileStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWhileStatementContext)
}

func (s *StatementContext) EmitStatement() IEmitStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEmitStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEmitStatementContext)
}

func (s *StatementContext) Declaration() IDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDeclarationContext)
}

func (s *StatementContext) Assignment() IAssignmentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignmentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignmentContext)
}

func (s *StatementContext) Swap() ISwapContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISwapContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISwapContext)
}

func (s *StatementContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CadenceVisitor:
		return t.VisitStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CadenceParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, CadenceParserRULE_statement)

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

	p.SetState(433)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(423)
			p.ReturnStatement()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(424)
			p.BreakStatement()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(425)
			p.ContinueStatement()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(426)
			p.IfStatement()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(427)
			p.WhileStatement()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(428)
			p.EmitStatement()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(429)
			p.Declaration()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(430)
			p.Assignment()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(431)
			p.Swap()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(432)
			p.Expression()
		}

	}

	return localctx
}

// IReturnStatementContext is an interface to support dynamic dispatch.
type IReturnStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReturnStatementContext differentiates from other interfaces.
	IsReturnStatementContext()
}

type ReturnStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturnStatementContext() *ReturnStatementContext {
	var p = new(ReturnStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CadenceParserRULE_returnStatement
	return p
}

func (*ReturnStatementContext) IsReturnStatementContext() {}

func NewReturnStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnStatementContext {
	var p = new(ReturnStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CadenceParserRULE_returnStatement

	return p
}

func (s *ReturnStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnStatementContext) Return() antlr.TerminalNode {
	return s.GetToken(CadenceParserReturn, 0)
}

func (s *ReturnStatementContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ReturnStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReturnStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.EnterReturnStatement(s)
	}
}

func (s *ReturnStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.ExitReturnStatement(s)
	}
}

func (s *ReturnStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CadenceVisitor:
		return t.VisitReturnStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CadenceParser) ReturnStatement() (localctx IReturnStatementContext) {
	localctx = NewReturnStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, CadenceParserRULE_returnStatement)

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
		p.SetState(435)
		p.Match(CadenceParserReturn)
	}
	p.SetState(438)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext()) == 1 {
		p.SetState(436)

		if !(!p.lineTerminatorAhead()) {
			panic(antlr.NewFailedPredicateException(p, "!p.lineTerminatorAhead()", ""))
		}
		{
			p.SetState(437)
			p.Expression()
		}

	}

	return localctx
}

// IBreakStatementContext is an interface to support dynamic dispatch.
type IBreakStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBreakStatementContext differentiates from other interfaces.
	IsBreakStatementContext()
}

type BreakStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBreakStatementContext() *BreakStatementContext {
	var p = new(BreakStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CadenceParserRULE_breakStatement
	return p
}

func (*BreakStatementContext) IsBreakStatementContext() {}

func NewBreakStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BreakStatementContext {
	var p = new(BreakStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CadenceParserRULE_breakStatement

	return p
}

func (s *BreakStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *BreakStatementContext) Break() antlr.TerminalNode {
	return s.GetToken(CadenceParserBreak, 0)
}

func (s *BreakStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BreakStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BreakStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.EnterBreakStatement(s)
	}
}

func (s *BreakStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.ExitBreakStatement(s)
	}
}

func (s *BreakStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CadenceVisitor:
		return t.VisitBreakStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CadenceParser) BreakStatement() (localctx IBreakStatementContext) {
	localctx = NewBreakStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, CadenceParserRULE_breakStatement)

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
		p.SetState(440)
		p.Match(CadenceParserBreak)
	}

	return localctx
}

// IContinueStatementContext is an interface to support dynamic dispatch.
type IContinueStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsContinueStatementContext differentiates from other interfaces.
	IsContinueStatementContext()
}

type ContinueStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyContinueStatementContext() *ContinueStatementContext {
	var p = new(ContinueStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CadenceParserRULE_continueStatement
	return p
}

func (*ContinueStatementContext) IsContinueStatementContext() {}

func NewContinueStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ContinueStatementContext {
	var p = new(ContinueStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CadenceParserRULE_continueStatement

	return p
}

func (s *ContinueStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ContinueStatementContext) Continue() antlr.TerminalNode {
	return s.GetToken(CadenceParserContinue, 0)
}

func (s *ContinueStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContinueStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ContinueStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.EnterContinueStatement(s)
	}
}

func (s *ContinueStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.ExitContinueStatement(s)
	}
}

func (s *ContinueStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CadenceVisitor:
		return t.VisitContinueStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CadenceParser) ContinueStatement() (localctx IContinueStatementContext) {
	localctx = NewContinueStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, CadenceParserRULE_continueStatement)

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
		p.SetState(442)
		p.Match(CadenceParserContinue)
	}

	return localctx
}

// IIfStatementContext is an interface to support dynamic dispatch.
type IIfStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetTestExpression returns the testExpression rule contexts.
	GetTestExpression() IExpressionContext

	// GetTestDeclaration returns the testDeclaration rule contexts.
	GetTestDeclaration() IVariableDeclarationContext

	// GetThen returns the then rule contexts.
	GetThen() IBlockContext

	// GetAlt returns the alt rule contexts.
	GetAlt() IBlockContext

	// SetTestExpression sets the testExpression rule contexts.
	SetTestExpression(IExpressionContext)

	// SetTestDeclaration sets the testDeclaration rule contexts.
	SetTestDeclaration(IVariableDeclarationContext)

	// SetThen sets the then rule contexts.
	SetThen(IBlockContext)

	// SetAlt sets the alt rule contexts.
	SetAlt(IBlockContext)

	// IsIfStatementContext differentiates from other interfaces.
	IsIfStatementContext()
}

type IfStatementContext struct {
	*antlr.BaseParserRuleContext
	parser          antlr.Parser
	testExpression  IExpressionContext
	testDeclaration IVariableDeclarationContext
	then            IBlockContext
	alt             IBlockContext
}

func NewEmptyIfStatementContext() *IfStatementContext {
	var p = new(IfStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CadenceParserRULE_ifStatement
	return p
}

func (*IfStatementContext) IsIfStatementContext() {}

func NewIfStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStatementContext {
	var p = new(IfStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CadenceParserRULE_ifStatement

	return p
}

func (s *IfStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *IfStatementContext) GetTestExpression() IExpressionContext { return s.testExpression }

func (s *IfStatementContext) GetTestDeclaration() IVariableDeclarationContext {
	return s.testDeclaration
}

func (s *IfStatementContext) GetThen() IBlockContext { return s.then }

func (s *IfStatementContext) GetAlt() IBlockContext { return s.alt }

func (s *IfStatementContext) SetTestExpression(v IExpressionContext) { s.testExpression = v }

func (s *IfStatementContext) SetTestDeclaration(v IVariableDeclarationContext) { s.testDeclaration = v }

func (s *IfStatementContext) SetThen(v IBlockContext) { s.then = v }

func (s *IfStatementContext) SetAlt(v IBlockContext) { s.alt = v }

func (s *IfStatementContext) If() antlr.TerminalNode {
	return s.GetToken(CadenceParserIf, 0)
}

func (s *IfStatementContext) AllBlock() []IBlockContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBlockContext)(nil)).Elem())
	var tst = make([]IBlockContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBlockContext)
		}
	}

	return tst
}

func (s *IfStatementContext) Block(i int) IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *IfStatementContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IfStatementContext) VariableDeclaration() IVariableDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableDeclarationContext)
}

func (s *IfStatementContext) Else() antlr.TerminalNode {
	return s.GetToken(CadenceParserElse, 0)
}

func (s *IfStatementContext) IfStatement() IIfStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIfStatementContext)
}

func (s *IfStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.EnterIfStatement(s)
	}
}

func (s *IfStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.ExitIfStatement(s)
	}
}

func (s *IfStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CadenceVisitor:
		return t.VisitIfStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CadenceParser) IfStatement() (localctx IIfStatementContext) {
	localctx = NewIfStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, CadenceParserRULE_ifStatement)

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
		p.SetState(444)
		p.Match(CadenceParserIf)
	}
	p.SetState(447)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CadenceParserT__2, CadenceParserT__5, CadenceParserMinus, CadenceParserAmpersand, CadenceParserNegate, CadenceParserMove, CadenceParserOpenParen, CadenceParserFun, CadenceParserTrue, CadenceParserFalse, CadenceParserNil, CadenceParserFrom, CadenceParserCreate, CadenceParserDestroy, CadenceParserIdentifier, CadenceParserDecimalLiteral, CadenceParserBinaryLiteral, CadenceParserOctalLiteral, CadenceParserHexadecimalLiteral, CadenceParserInvalidNumberLiteral, CadenceParserStringLiteral:
		{
			p.SetState(445)

			var _x = p.Expression()

			localctx.(*IfStatementContext).testExpression = _x
		}

	case CadenceParserPriv, CadenceParserPub, CadenceParserPubSet, CadenceParserLet, CadenceParserVar:
		{
			p.SetState(446)

			var _x = p.VariableDeclaration()

			localctx.(*IfStatementContext).testDeclaration = _x
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(449)

		var _x = p.Block()

		localctx.(*IfStatementContext).then = _x
	}
	p.SetState(455)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 34, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(450)
			p.Match(CadenceParserElse)
		}
		p.SetState(453)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case CadenceParserIf:
			{
				p.SetState(451)
				p.IfStatement()
			}

		case CadenceParserT__2:
			{
				p.SetState(452)

				var _x = p.Block()

				localctx.(*IfStatementContext).alt = _x
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

	}

	return localctx
}

// IWhileStatementContext is an interface to support dynamic dispatch.
type IWhileStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWhileStatementContext differentiates from other interfaces.
	IsWhileStatementContext()
}

type WhileStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhileStatementContext() *WhileStatementContext {
	var p = new(WhileStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CadenceParserRULE_whileStatement
	return p
}

func (*WhileStatementContext) IsWhileStatementContext() {}

func NewWhileStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhileStatementContext {
	var p = new(WhileStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CadenceParserRULE_whileStatement

	return p
}

func (s *WhileStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *WhileStatementContext) While() antlr.TerminalNode {
	return s.GetToken(CadenceParserWhile, 0)
}

func (s *WhileStatementContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *WhileStatementContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *WhileStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhileStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.EnterWhileStatement(s)
	}
}

func (s *WhileStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.ExitWhileStatement(s)
	}
}

func (s *WhileStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CadenceVisitor:
		return t.VisitWhileStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CadenceParser) WhileStatement() (localctx IWhileStatementContext) {
	localctx = NewWhileStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, CadenceParserRULE_whileStatement)

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
		p.SetState(457)
		p.Match(CadenceParserWhile)
	}
	{
		p.SetState(458)
		p.Expression()
	}
	{
		p.SetState(459)
		p.Block()
	}

	return localctx
}

// IEmitStatementContext is an interface to support dynamic dispatch.
type IEmitStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEmitStatementContext differentiates from other interfaces.
	IsEmitStatementContext()
}

type EmitStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEmitStatementContext() *EmitStatementContext {
	var p = new(EmitStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CadenceParserRULE_emitStatement
	return p
}

func (*EmitStatementContext) IsEmitStatementContext() {}

func NewEmitStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EmitStatementContext {
	var p = new(EmitStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CadenceParserRULE_emitStatement

	return p
}

func (s *EmitStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *EmitStatementContext) Emit() antlr.TerminalNode {
	return s.GetToken(CadenceParserEmit, 0)
}

func (s *EmitStatementContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *EmitStatementContext) Invocation() IInvocationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInvocationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInvocationContext)
}

func (s *EmitStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EmitStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EmitStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.EnterEmitStatement(s)
	}
}

func (s *EmitStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.ExitEmitStatement(s)
	}
}

func (s *EmitStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CadenceVisitor:
		return t.VisitEmitStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CadenceParser) EmitStatement() (localctx IEmitStatementContext) {
	localctx = NewEmitStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, CadenceParserRULE_emitStatement)

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
		p.SetState(461)
		p.Match(CadenceParserEmit)
	}
	{
		p.SetState(462)
		p.Identifier()
	}
	{
		p.SetState(463)
		p.Invocation()
	}

	return localctx
}

// IVariableDeclarationContext is an interface to support dynamic dispatch.
type IVariableDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetLeftTransfer returns the leftTransfer rule contexts.
	GetLeftTransfer() ITransferContext

	// GetLeftExpression returns the leftExpression rule contexts.
	GetLeftExpression() IExpressionContext

	// GetRightTransfer returns the rightTransfer rule contexts.
	GetRightTransfer() ITransferContext

	// GetRightExpression returns the rightExpression rule contexts.
	GetRightExpression() IExpressionContext

	// SetLeftTransfer sets the leftTransfer rule contexts.
	SetLeftTransfer(ITransferContext)

	// SetLeftExpression sets the leftExpression rule contexts.
	SetLeftExpression(IExpressionContext)

	// SetRightTransfer sets the rightTransfer rule contexts.
	SetRightTransfer(ITransferContext)

	// SetRightExpression sets the rightExpression rule contexts.
	SetRightExpression(IExpressionContext)

	// IsVariableDeclarationContext differentiates from other interfaces.
	IsVariableDeclarationContext()
}

type VariableDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser          antlr.Parser
	leftTransfer    ITransferContext
	leftExpression  IExpressionContext
	rightTransfer   ITransferContext
	rightExpression IExpressionContext
}

func NewEmptyVariableDeclarationContext() *VariableDeclarationContext {
	var p = new(VariableDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CadenceParserRULE_variableDeclaration
	return p
}

func (*VariableDeclarationContext) IsVariableDeclarationContext() {}

func NewVariableDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableDeclarationContext {
	var p = new(VariableDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CadenceParserRULE_variableDeclaration

	return p
}

func (s *VariableDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableDeclarationContext) GetLeftTransfer() ITransferContext { return s.leftTransfer }

func (s *VariableDeclarationContext) GetLeftExpression() IExpressionContext { return s.leftExpression }

func (s *VariableDeclarationContext) GetRightTransfer() ITransferContext { return s.rightTransfer }

func (s *VariableDeclarationContext) GetRightExpression() IExpressionContext { return s.rightExpression }

func (s *VariableDeclarationContext) SetLeftTransfer(v ITransferContext) { s.leftTransfer = v }

func (s *VariableDeclarationContext) SetLeftExpression(v IExpressionContext) { s.leftExpression = v }

func (s *VariableDeclarationContext) SetRightTransfer(v ITransferContext) { s.rightTransfer = v }

func (s *VariableDeclarationContext) SetRightExpression(v IExpressionContext) { s.rightExpression = v }

func (s *VariableDeclarationContext) Access() IAccessContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAccessContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAccessContext)
}

func (s *VariableDeclarationContext) VariableKind() IVariableKindContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableKindContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableKindContext)
}

func (s *VariableDeclarationContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *VariableDeclarationContext) AllTransfer() []ITransferContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITransferContext)(nil)).Elem())
	var tst = make([]ITransferContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITransferContext)
		}
	}

	return tst
}

func (s *VariableDeclarationContext) Transfer(i int) ITransferContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITransferContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITransferContext)
}

func (s *VariableDeclarationContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *VariableDeclarationContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *VariableDeclarationContext) TypeAnnotation() ITypeAnnotationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeAnnotationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeAnnotationContext)
}

func (s *VariableDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.EnterVariableDeclaration(s)
	}
}

func (s *VariableDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.ExitVariableDeclaration(s)
	}
}

func (s *VariableDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CadenceVisitor:
		return t.VisitVariableDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CadenceParser) VariableDeclaration() (localctx IVariableDeclarationContext) {
	localctx = NewVariableDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, CadenceParserRULE_variableDeclaration)
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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(465)
		p.Access()
	}
	{
		p.SetState(466)
		p.VariableKind()
	}
	{
		p.SetState(467)
		p.Identifier()
	}
	p.SetState(470)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CadenceParserT__4 {
		{
			p.SetState(468)
			p.Match(CadenceParserT__4)
		}
		{
			p.SetState(469)
			p.TypeAnnotation()
		}

	}
	{
		p.SetState(472)

		var _x = p.Transfer()

		localctx.(*VariableDeclarationContext).leftTransfer = _x
	}
	{
		p.SetState(473)

		var _x = p.Expression()

		localctx.(*VariableDeclarationContext).leftExpression = _x
	}
	p.SetState(477)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 36, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(474)

			var _x = p.Transfer()

			localctx.(*VariableDeclarationContext).rightTransfer = _x
		}
		{
			p.SetState(475)

			var _x = p.Expression()

			localctx.(*VariableDeclarationContext).rightExpression = _x
		}

	}

	return localctx
}

// IAssignmentContext is an interface to support dynamic dispatch.
type IAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssignmentContext differentiates from other interfaces.
	IsAssignmentContext()
}

type AssignmentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentContext() *AssignmentContext {
	var p = new(AssignmentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CadenceParserRULE_assignment
	return p
}

func (*AssignmentContext) IsAssignmentContext() {}

func NewAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentContext {
	var p = new(AssignmentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CadenceParserRULE_assignment

	return p
}

func (s *AssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *AssignmentContext) Transfer() ITransferContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITransferContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITransferContext)
}

func (s *AssignmentContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AssignmentContext) AllExpressionAccess() []IExpressionAccessContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionAccessContext)(nil)).Elem())
	var tst = make([]IExpressionAccessContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionAccessContext)
		}
	}

	return tst
}

func (s *AssignmentContext) ExpressionAccess(i int) IExpressionAccessContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionAccessContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionAccessContext)
}

func (s *AssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.EnterAssignment(s)
	}
}

func (s *AssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.ExitAssignment(s)
	}
}

func (s *AssignmentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CadenceVisitor:
		return t.VisitAssignment(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CadenceParser) Assignment() (localctx IAssignmentContext) {
	localctx = NewAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, CadenceParserRULE_assignment)
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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(479)
		p.Identifier()
	}
	p.SetState(483)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CadenceParserT__5 || _la == CadenceParserT__11 {
		{
			p.SetState(480)
			p.ExpressionAccess()
		}

		p.SetState(485)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(486)
		p.Transfer()
	}
	{
		p.SetState(487)
		p.Expression()
	}

	return localctx
}

// ISwapContext is an interface to support dynamic dispatch.
type ISwapContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetLeft returns the left rule contexts.
	GetLeft() IExpressionContext

	// GetRight returns the right rule contexts.
	GetRight() IExpressionContext

	// SetLeft sets the left rule contexts.
	SetLeft(IExpressionContext)

	// SetRight sets the right rule contexts.
	SetRight(IExpressionContext)

	// IsSwapContext differentiates from other interfaces.
	IsSwapContext()
}

type SwapContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	left   IExpressionContext
	right  IExpressionContext
}

func NewEmptySwapContext() *SwapContext {
	var p = new(SwapContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CadenceParserRULE_swap
	return p
}

func (*SwapContext) IsSwapContext() {}

func NewSwapContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SwapContext {
	var p = new(SwapContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CadenceParserRULE_swap

	return p
}

func (s *SwapContext) GetParser() antlr.Parser { return s.parser }

func (s *SwapContext) GetLeft() IExpressionContext { return s.left }

func (s *SwapContext) GetRight() IExpressionContext { return s.right }

func (s *SwapContext) SetLeft(v IExpressionContext) { s.left = v }

func (s *SwapContext) SetRight(v IExpressionContext) { s.right = v }

func (s *SwapContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *SwapContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *SwapContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwapContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SwapContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.EnterSwap(s)
	}
}

func (s *SwapContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.ExitSwap(s)
	}
}

func (s *SwapContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CadenceVisitor:
		return t.VisitSwap(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CadenceParser) Swap() (localctx ISwapContext) {
	localctx = NewSwapContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, CadenceParserRULE_swap)

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
		p.SetState(489)

		var _x = p.Expression()

		localctx.(*SwapContext).left = _x
	}
	{
		p.SetState(490)
		p.Match(CadenceParserT__7)
	}
	{
		p.SetState(491)

		var _x = p.Expression()

		localctx.(*SwapContext).right = _x
	}

	return localctx
}

// ITransferContext is an interface to support dynamic dispatch.
type ITransferContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTransferContext differentiates from other interfaces.
	IsTransferContext()
}

type TransferContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTransferContext() *TransferContext {
	var p = new(TransferContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CadenceParserRULE_transfer
	return p
}

func (*TransferContext) IsTransferContext() {}

func NewTransferContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TransferContext {
	var p = new(TransferContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CadenceParserRULE_transfer

	return p
}

func (s *TransferContext) GetParser() antlr.Parser { return s.parser }

func (s *TransferContext) Move() antlr.TerminalNode {
	return s.GetToken(CadenceParserMove, 0)
}

func (s *TransferContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TransferContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TransferContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.EnterTransfer(s)
	}
}

func (s *TransferContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.ExitTransfer(s)
	}
}

func (s *TransferContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CadenceVisitor:
		return t.VisitTransfer(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CadenceParser) Transfer() (localctx ITransferContext) {
	localctx = NewTransferContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, CadenceParserRULE_transfer)
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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(493)
		_la = p.GetTokenStream().LA(1)

		if !(_la == CadenceParserT__8 || _la == CadenceParserMove) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CadenceParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CadenceParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) ConditionalExpression() IConditionalExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConditionalExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConditionalExpressionContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (s *ExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CadenceVisitor:
		return t.VisitExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CadenceParser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, CadenceParserRULE_expression)

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
		p.SetState(495)
		p.ConditionalExpression()
	}

	return localctx
}

// IConditionalExpressionContext is an interface to support dynamic dispatch.
type IConditionalExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetThen returns the then rule contexts.
	GetThen() IExpressionContext

	// GetAlt returns the alt rule contexts.
	GetAlt() IExpressionContext

	// SetThen sets the then rule contexts.
	SetThen(IExpressionContext)

	// SetAlt sets the alt rule contexts.
	SetAlt(IExpressionContext)

	// IsConditionalExpressionContext differentiates from other interfaces.
	IsConditionalExpressionContext()
}

type ConditionalExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	then   IExpressionContext
	alt    IExpressionContext
}

func NewEmptyConditionalExpressionContext() *ConditionalExpressionContext {
	var p = new(ConditionalExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CadenceParserRULE_conditionalExpression
	return p
}

func (*ConditionalExpressionContext) IsConditionalExpressionContext() {}

func NewConditionalExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionalExpressionContext {
	var p = new(ConditionalExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CadenceParserRULE_conditionalExpression

	return p
}

func (s *ConditionalExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionalExpressionContext) GetThen() IExpressionContext { return s.then }

func (s *ConditionalExpressionContext) GetAlt() IExpressionContext { return s.alt }

func (s *ConditionalExpressionContext) SetThen(v IExpressionContext) { s.then = v }

func (s *ConditionalExpressionContext) SetAlt(v IExpressionContext) { s.alt = v }

func (s *ConditionalExpressionContext) OrExpression() IOrExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOrExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOrExpressionContext)
}

func (s *ConditionalExpressionContext) Optional() antlr.TerminalNode {
	return s.GetToken(CadenceParserOptional, 0)
}

func (s *ConditionalExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ConditionalExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ConditionalExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionalExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionalExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.EnterConditionalExpression(s)
	}
}

func (s *ConditionalExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.ExitConditionalExpression(s)
	}
}

func (s *ConditionalExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CadenceVisitor:
		return t.VisitConditionalExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CadenceParser) ConditionalExpression() (localctx IConditionalExpressionContext) {
	localctx = NewConditionalExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, CadenceParserRULE_conditionalExpression)

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
		p.SetState(497)
		p.orExpression(0)
	}
	p.SetState(503)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 38, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(498)
			p.Match(CadenceParserOptional)
		}
		{
			p.SetState(499)

			var _x = p.Expression()

			localctx.(*ConditionalExpressionContext).then = _x
		}
		{
			p.SetState(500)
			p.Match(CadenceParserT__4)
		}
		{
			p.SetState(501)

			var _x = p.Expression()

			localctx.(*ConditionalExpressionContext).alt = _x
		}

	}

	return localctx
}

// IOrExpressionContext is an interface to support dynamic dispatch.
type IOrExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOrExpressionContext differentiates from other interfaces.
	IsOrExpressionContext()
}

type OrExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOrExpressionContext() *OrExpressionContext {
	var p = new(OrExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CadenceParserRULE_orExpression
	return p
}

func (*OrExpressionContext) IsOrExpressionContext() {}

func NewOrExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OrExpressionContext {
	var p = new(OrExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CadenceParserRULE_orExpression

	return p
}

func (s *OrExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *OrExpressionContext) AndExpression() IAndExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAndExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAndExpressionContext)
}

func (s *OrExpressionContext) OrExpression() IOrExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOrExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOrExpressionContext)
}

func (s *OrExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OrExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.EnterOrExpression(s)
	}
}

func (s *OrExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.ExitOrExpression(s)
	}
}

func (s *OrExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CadenceVisitor:
		return t.VisitOrExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CadenceParser) OrExpression() (localctx IOrExpressionContext) {
	return p.orExpression(0)
}

func (p *CadenceParser) orExpression(_p int) (localctx IOrExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewOrExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IOrExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 92
	p.EnterRecursionRule(localctx, 92, CadenceParserRULE_orExpression, _p)

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
		p.SetState(506)
		p.andExpression(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(513)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 39, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewOrExpressionContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, CadenceParserRULE_orExpression)
			p.SetState(508)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(509)
				p.Match(CadenceParserT__9)
			}
			{
				p.SetState(510)
				p.andExpression(0)
			}

		}
		p.SetState(515)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 39, p.GetParserRuleContext())
	}

	return localctx
}

// IAndExpressionContext is an interface to support dynamic dispatch.
type IAndExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAndExpressionContext differentiates from other interfaces.
	IsAndExpressionContext()
}

type AndExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAndExpressionContext() *AndExpressionContext {
	var p = new(AndExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CadenceParserRULE_andExpression
	return p
}

func (*AndExpressionContext) IsAndExpressionContext() {}

func NewAndExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AndExpressionContext {
	var p = new(AndExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CadenceParserRULE_andExpression

	return p
}

func (s *AndExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *AndExpressionContext) EqualityExpression() IEqualityExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEqualityExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEqualityExpressionContext)
}

func (s *AndExpressionContext) AndExpression() IAndExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAndExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAndExpressionContext)
}

func (s *AndExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AndExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.EnterAndExpression(s)
	}
}

func (s *AndExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.ExitAndExpression(s)
	}
}

func (s *AndExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CadenceVisitor:
		return t.VisitAndExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CadenceParser) AndExpression() (localctx IAndExpressionContext) {
	return p.andExpression(0)
}

func (p *CadenceParser) andExpression(_p int) (localctx IAndExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewAndExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IAndExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 94
	p.EnterRecursionRule(localctx, 94, CadenceParserRULE_andExpression, _p)

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
		p.SetState(517)
		p.equalityExpression(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(524)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 40, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewAndExpressionContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, CadenceParserRULE_andExpression)
			p.SetState(519)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(520)
				p.Match(CadenceParserT__10)
			}
			{
				p.SetState(521)
				p.equalityExpression(0)
			}

		}
		p.SetState(526)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 40, p.GetParserRuleContext())
	}

	return localctx
}

// IEqualityExpressionContext is an interface to support dynamic dispatch.
type IEqualityExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEqualityExpressionContext differentiates from other interfaces.
	IsEqualityExpressionContext()
}

type EqualityExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEqualityExpressionContext() *EqualityExpressionContext {
	var p = new(EqualityExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CadenceParserRULE_equalityExpression
	return p
}

func (*EqualityExpressionContext) IsEqualityExpressionContext() {}

func NewEqualityExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EqualityExpressionContext {
	var p = new(EqualityExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CadenceParserRULE_equalityExpression

	return p
}

func (s *EqualityExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *EqualityExpressionContext) RelationalExpression() IRelationalExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRelationalExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRelationalExpressionContext)
}

func (s *EqualityExpressionContext) EqualityExpression() IEqualityExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEqualityExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEqualityExpressionContext)
}

func (s *EqualityExpressionContext) EqualityOp() IEqualityOpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEqualityOpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEqualityOpContext)
}

func (s *EqualityExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualityExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EqualityExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.EnterEqualityExpression(s)
	}
}

func (s *EqualityExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.ExitEqualityExpression(s)
	}
}

func (s *EqualityExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CadenceVisitor:
		return t.VisitEqualityExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CadenceParser) EqualityExpression() (localctx IEqualityExpressionContext) {
	return p.equalityExpression(0)
}

func (p *CadenceParser) equalityExpression(_p int) (localctx IEqualityExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewEqualityExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IEqualityExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 96
	p.EnterRecursionRule(localctx, 96, CadenceParserRULE_equalityExpression, _p)

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
		p.SetState(528)
		p.relationalExpression(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(536)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 41, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewEqualityExpressionContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, CadenceParserRULE_equalityExpression)
			p.SetState(530)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(531)
				p.EqualityOp()
			}
			{
				p.SetState(532)
				p.relationalExpression(0)
			}

		}
		p.SetState(538)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 41, p.GetParserRuleContext())
	}

	return localctx
}

// IRelationalExpressionContext is an interface to support dynamic dispatch.
type IRelationalExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRelationalExpressionContext differentiates from other interfaces.
	IsRelationalExpressionContext()
}

type RelationalExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRelationalExpressionContext() *RelationalExpressionContext {
	var p = new(RelationalExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CadenceParserRULE_relationalExpression
	return p
}

func (*RelationalExpressionContext) IsRelationalExpressionContext() {}

func NewRelationalExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelationalExpressionContext {
	var p = new(RelationalExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CadenceParserRULE_relationalExpression

	return p
}

func (s *RelationalExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *RelationalExpressionContext) NilCoalescingExpression() INilCoalescingExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INilCoalescingExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INilCoalescingExpressionContext)
}

func (s *RelationalExpressionContext) RelationalExpression() IRelationalExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRelationalExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRelationalExpressionContext)
}

func (s *RelationalExpressionContext) RelationalOp() IRelationalOpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRelationalOpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRelationalOpContext)
}

func (s *RelationalExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationalExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RelationalExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.EnterRelationalExpression(s)
	}
}

func (s *RelationalExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.ExitRelationalExpression(s)
	}
}

func (s *RelationalExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CadenceVisitor:
		return t.VisitRelationalExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CadenceParser) RelationalExpression() (localctx IRelationalExpressionContext) {
	return p.relationalExpression(0)
}

func (p *CadenceParser) relationalExpression(_p int) (localctx IRelationalExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewRelationalExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IRelationalExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 98
	p.EnterRecursionRule(localctx, 98, CadenceParserRULE_relationalExpression, _p)

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
		p.SetState(540)
		p.NilCoalescingExpression()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(548)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 42, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewRelationalExpressionContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, CadenceParserRULE_relationalExpression)
			p.SetState(542)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(543)
				p.RelationalOp()
			}
			{
				p.SetState(544)
				p.NilCoalescingExpression()
			}

		}
		p.SetState(550)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 42, p.GetParserRuleContext())
	}

	return localctx
}

// INilCoalescingExpressionContext is an interface to support dynamic dispatch.
type INilCoalescingExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNilCoalescingExpressionContext differentiates from other interfaces.
	IsNilCoalescingExpressionContext()
}

type NilCoalescingExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNilCoalescingExpressionContext() *NilCoalescingExpressionContext {
	var p = new(NilCoalescingExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CadenceParserRULE_nilCoalescingExpression
	return p
}

func (*NilCoalescingExpressionContext) IsNilCoalescingExpressionContext() {}

func NewNilCoalescingExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NilCoalescingExpressionContext {
	var p = new(NilCoalescingExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CadenceParserRULE_nilCoalescingExpression

	return p
}

func (s *NilCoalescingExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *NilCoalescingExpressionContext) FailableDowncastingExpression() IFailableDowncastingExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFailableDowncastingExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFailableDowncastingExpressionContext)
}

func (s *NilCoalescingExpressionContext) NilCoalescing() antlr.TerminalNode {
	return s.GetToken(CadenceParserNilCoalescing, 0)
}

func (s *NilCoalescingExpressionContext) NilCoalescingExpression() INilCoalescingExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INilCoalescingExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INilCoalescingExpressionContext)
}

func (s *NilCoalescingExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NilCoalescingExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NilCoalescingExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.EnterNilCoalescingExpression(s)
	}
}

func (s *NilCoalescingExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.ExitNilCoalescingExpression(s)
	}
}

func (s *NilCoalescingExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CadenceVisitor:
		return t.VisitNilCoalescingExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CadenceParser) NilCoalescingExpression() (localctx INilCoalescingExpressionContext) {
	localctx = NewNilCoalescingExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 100, CadenceParserRULE_nilCoalescingExpression)

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
		p.SetState(551)
		p.failableDowncastingExpression(0)
	}
	p.SetState(554)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 43, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(552)
			p.Match(CadenceParserNilCoalescing)
		}
		{
			p.SetState(553)
			p.NilCoalescingExpression()
		}

	}

	return localctx
}

// IFailableDowncastingExpressionContext is an interface to support dynamic dispatch.
type IFailableDowncastingExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFailableDowncastingExpressionContext differentiates from other interfaces.
	IsFailableDowncastingExpressionContext()
}

type FailableDowncastingExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFailableDowncastingExpressionContext() *FailableDowncastingExpressionContext {
	var p = new(FailableDowncastingExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CadenceParserRULE_failableDowncastingExpression
	return p
}

func (*FailableDowncastingExpressionContext) IsFailableDowncastingExpressionContext() {}

func NewFailableDowncastingExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FailableDowncastingExpressionContext {
	var p = new(FailableDowncastingExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CadenceParserRULE_failableDowncastingExpression

	return p
}

func (s *FailableDowncastingExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *FailableDowncastingExpressionContext) ConcatenatingExpression() IConcatenatingExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConcatenatingExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConcatenatingExpressionContext)
}

func (s *FailableDowncastingExpressionContext) FailableDowncastingExpression() IFailableDowncastingExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFailableDowncastingExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFailableDowncastingExpressionContext)
}

func (s *FailableDowncastingExpressionContext) FailableDowncasting() antlr.TerminalNode {
	return s.GetToken(CadenceParserFailableDowncasting, 0)
}

func (s *FailableDowncastingExpressionContext) TypeAnnotation() ITypeAnnotationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeAnnotationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeAnnotationContext)
}

func (s *FailableDowncastingExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FailableDowncastingExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FailableDowncastingExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.EnterFailableDowncastingExpression(s)
	}
}

func (s *FailableDowncastingExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.ExitFailableDowncastingExpression(s)
	}
}

func (s *FailableDowncastingExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CadenceVisitor:
		return t.VisitFailableDowncastingExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CadenceParser) FailableDowncastingExpression() (localctx IFailableDowncastingExpressionContext) {
	return p.failableDowncastingExpression(0)
}

func (p *CadenceParser) failableDowncastingExpression(_p int) (localctx IFailableDowncastingExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewFailableDowncastingExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IFailableDowncastingExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 102
	p.EnterRecursionRule(localctx, 102, CadenceParserRULE_failableDowncastingExpression, _p)

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
		p.SetState(557)
		p.concatenatingExpression(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(564)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 44, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewFailableDowncastingExpressionContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, CadenceParserRULE_failableDowncastingExpression)
			p.SetState(559)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(560)
				p.Match(CadenceParserFailableDowncasting)
			}
			{
				p.SetState(561)
				p.TypeAnnotation()
			}

		}
		p.SetState(566)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 44, p.GetParserRuleContext())
	}

	return localctx
}

// IConcatenatingExpressionContext is an interface to support dynamic dispatch.
type IConcatenatingExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConcatenatingExpressionContext differentiates from other interfaces.
	IsConcatenatingExpressionContext()
}

type ConcatenatingExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConcatenatingExpressionContext() *ConcatenatingExpressionContext {
	var p = new(ConcatenatingExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CadenceParserRULE_concatenatingExpression
	return p
}

func (*ConcatenatingExpressionContext) IsConcatenatingExpressionContext() {}

func NewConcatenatingExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConcatenatingExpressionContext {
	var p = new(ConcatenatingExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CadenceParserRULE_concatenatingExpression

	return p
}

func (s *ConcatenatingExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ConcatenatingExpressionContext) AdditiveExpression() IAdditiveExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAdditiveExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAdditiveExpressionContext)
}

func (s *ConcatenatingExpressionContext) ConcatenatingExpression() IConcatenatingExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConcatenatingExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConcatenatingExpressionContext)
}

func (s *ConcatenatingExpressionContext) Ampersand() antlr.TerminalNode {
	return s.GetToken(CadenceParserAmpersand, 0)
}

func (s *ConcatenatingExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConcatenatingExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConcatenatingExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.EnterConcatenatingExpression(s)
	}
}

func (s *ConcatenatingExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.ExitConcatenatingExpression(s)
	}
}

func (s *ConcatenatingExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CadenceVisitor:
		return t.VisitConcatenatingExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CadenceParser) ConcatenatingExpression() (localctx IConcatenatingExpressionContext) {
	return p.concatenatingExpression(0)
}

func (p *CadenceParser) concatenatingExpression(_p int) (localctx IConcatenatingExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewConcatenatingExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IConcatenatingExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 104
	p.EnterRecursionRule(localctx, 104, CadenceParserRULE_concatenatingExpression, _p)

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
		p.SetState(568)
		p.additiveExpression(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(575)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 45, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewConcatenatingExpressionContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, CadenceParserRULE_concatenatingExpression)
			p.SetState(570)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(571)
				p.Match(CadenceParserAmpersand)
			}
			{
				p.SetState(572)
				p.additiveExpression(0)
			}

		}
		p.SetState(577)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 45, p.GetParserRuleContext())
	}

	return localctx
}

// IAdditiveExpressionContext is an interface to support dynamic dispatch.
type IAdditiveExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAdditiveExpressionContext differentiates from other interfaces.
	IsAdditiveExpressionContext()
}

type AdditiveExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAdditiveExpressionContext() *AdditiveExpressionContext {
	var p = new(AdditiveExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CadenceParserRULE_additiveExpression
	return p
}

func (*AdditiveExpressionContext) IsAdditiveExpressionContext() {}

func NewAdditiveExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AdditiveExpressionContext {
	var p = new(AdditiveExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CadenceParserRULE_additiveExpression

	return p
}

func (s *AdditiveExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *AdditiveExpressionContext) MultiplicativeExpression() IMultiplicativeExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMultiplicativeExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMultiplicativeExpressionContext)
}

func (s *AdditiveExpressionContext) AdditiveExpression() IAdditiveExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAdditiveExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAdditiveExpressionContext)
}

func (s *AdditiveExpressionContext) AdditiveOp() IAdditiveOpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAdditiveOpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAdditiveOpContext)
}

func (s *AdditiveExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AdditiveExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AdditiveExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.EnterAdditiveExpression(s)
	}
}

func (s *AdditiveExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.ExitAdditiveExpression(s)
	}
}

func (s *AdditiveExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CadenceVisitor:
		return t.VisitAdditiveExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CadenceParser) AdditiveExpression() (localctx IAdditiveExpressionContext) {
	return p.additiveExpression(0)
}

func (p *CadenceParser) additiveExpression(_p int) (localctx IAdditiveExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewAdditiveExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IAdditiveExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 106
	p.EnterRecursionRule(localctx, 106, CadenceParserRULE_additiveExpression, _p)

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
		p.SetState(579)
		p.multiplicativeExpression(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(587)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 46, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewAdditiveExpressionContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, CadenceParserRULE_additiveExpression)
			p.SetState(581)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(582)
				p.AdditiveOp()
			}
			{
				p.SetState(583)
				p.multiplicativeExpression(0)
			}

		}
		p.SetState(589)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 46, p.GetParserRuleContext())
	}

	return localctx
}

// IMultiplicativeExpressionContext is an interface to support dynamic dispatch.
type IMultiplicativeExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMultiplicativeExpressionContext differentiates from other interfaces.
	IsMultiplicativeExpressionContext()
}

type MultiplicativeExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMultiplicativeExpressionContext() *MultiplicativeExpressionContext {
	var p = new(MultiplicativeExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CadenceParserRULE_multiplicativeExpression
	return p
}

func (*MultiplicativeExpressionContext) IsMultiplicativeExpressionContext() {}

func NewMultiplicativeExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultiplicativeExpressionContext {
	var p = new(MultiplicativeExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CadenceParserRULE_multiplicativeExpression

	return p
}

func (s *MultiplicativeExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *MultiplicativeExpressionContext) UnaryExpression() IUnaryExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnaryExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnaryExpressionContext)
}

func (s *MultiplicativeExpressionContext) MultiplicativeExpression() IMultiplicativeExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMultiplicativeExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMultiplicativeExpressionContext)
}

func (s *MultiplicativeExpressionContext) MultiplicativeOp() IMultiplicativeOpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMultiplicativeOpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMultiplicativeOpContext)
}

func (s *MultiplicativeExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiplicativeExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MultiplicativeExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.EnterMultiplicativeExpression(s)
	}
}

func (s *MultiplicativeExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.ExitMultiplicativeExpression(s)
	}
}

func (s *MultiplicativeExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CadenceVisitor:
		return t.VisitMultiplicativeExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CadenceParser) MultiplicativeExpression() (localctx IMultiplicativeExpressionContext) {
	return p.multiplicativeExpression(0)
}

func (p *CadenceParser) multiplicativeExpression(_p int) (localctx IMultiplicativeExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewMultiplicativeExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IMultiplicativeExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 108
	p.EnterRecursionRule(localctx, 108, CadenceParserRULE_multiplicativeExpression, _p)

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
		p.SetState(591)
		p.UnaryExpression()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(599)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 47, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewMultiplicativeExpressionContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, CadenceParserRULE_multiplicativeExpression)
			p.SetState(593)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(594)
				p.MultiplicativeOp()
			}
			{
				p.SetState(595)
				p.UnaryExpression()
			}

		}
		p.SetState(601)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 47, p.GetParserRuleContext())
	}

	return localctx
}

// IUnaryExpressionContext is an interface to support dynamic dispatch.
type IUnaryExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnaryExpressionContext differentiates from other interfaces.
	IsUnaryExpressionContext()
}

type UnaryExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnaryExpressionContext() *UnaryExpressionContext {
	var p = new(UnaryExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CadenceParserRULE_unaryExpression
	return p
}

func (*UnaryExpressionContext) IsUnaryExpressionContext() {}

func NewUnaryExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnaryExpressionContext {
	var p = new(UnaryExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CadenceParserRULE_unaryExpression

	return p
}

func (s *UnaryExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *UnaryExpressionContext) PrimaryExpression() IPrimaryExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimaryExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimaryExpressionContext)
}

func (s *UnaryExpressionContext) UnaryExpression() IUnaryExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnaryExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnaryExpressionContext)
}

func (s *UnaryExpressionContext) AllUnaryOp() []IUnaryOpContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IUnaryOpContext)(nil)).Elem())
	var tst = make([]IUnaryOpContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IUnaryOpContext)
		}
	}

	return tst
}

func (s *UnaryExpressionContext) UnaryOp(i int) IUnaryOpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnaryOpContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IUnaryOpContext)
}

func (s *UnaryExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnaryExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.EnterUnaryExpression(s)
	}
}

func (s *UnaryExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.ExitUnaryExpression(s)
	}
}

func (s *UnaryExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CadenceVisitor:
		return t.VisitUnaryExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CadenceParser) UnaryExpression() (localctx IUnaryExpressionContext) {
	localctx = NewUnaryExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 110, CadenceParserRULE_unaryExpression)

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

	var _alt int

	p.SetState(610)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 49, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(602)
			p.PrimaryExpression()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(604)
		p.GetErrorHandler().Sync(p)
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(603)
					p.UnaryOp()
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(606)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 48, p.GetParserRuleContext())
		}
		{
			p.SetState(608)
			p.UnaryExpression()
		}

	}

	return localctx
}

// IPrimaryExpressionContext is an interface to support dynamic dispatch.
type IPrimaryExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrimaryExpressionContext differentiates from other interfaces.
	IsPrimaryExpressionContext()
}

type PrimaryExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryExpressionContext() *PrimaryExpressionContext {
	var p = new(PrimaryExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CadenceParserRULE_primaryExpression
	return p
}

func (*PrimaryExpressionContext) IsPrimaryExpressionContext() {}

func NewPrimaryExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryExpressionContext {
	var p = new(PrimaryExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CadenceParserRULE_primaryExpression

	return p
}

func (s *PrimaryExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryExpressionContext) CreateExpression() ICreateExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICreateExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICreateExpressionContext)
}

func (s *PrimaryExpressionContext) DestroyExpression() IDestroyExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDestroyExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDestroyExpressionContext)
}

func (s *PrimaryExpressionContext) ReferenceExpression() IReferenceExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReferenceExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReferenceExpressionContext)
}

func (s *PrimaryExpressionContext) ComposedExpression() IComposedExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComposedExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IComposedExpressionContext)
}

func (s *PrimaryExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimaryExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.EnterPrimaryExpression(s)
	}
}

func (s *PrimaryExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.ExitPrimaryExpression(s)
	}
}

func (s *PrimaryExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CadenceVisitor:
		return t.VisitPrimaryExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CadenceParser) PrimaryExpression() (localctx IPrimaryExpressionContext) {
	localctx = NewPrimaryExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 112, CadenceParserRULE_primaryExpression)

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

	p.SetState(616)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 50, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(612)
			p.CreateExpression()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(613)
			p.DestroyExpression()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(614)
			p.ReferenceExpression()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(615)
			p.ComposedExpression()
		}

	}

	return localctx
}

// IComposedExpressionContext is an interface to support dynamic dispatch.
type IComposedExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsComposedExpressionContext differentiates from other interfaces.
	IsComposedExpressionContext()
}

type ComposedExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComposedExpressionContext() *ComposedExpressionContext {
	var p = new(ComposedExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CadenceParserRULE_composedExpression
	return p
}

func (*ComposedExpressionContext) IsComposedExpressionContext() {}

func NewComposedExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComposedExpressionContext {
	var p = new(ComposedExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CadenceParserRULE_composedExpression

	return p
}

func (s *ComposedExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ComposedExpressionContext) PrimaryExpressionStart() IPrimaryExpressionStartContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimaryExpressionStartContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimaryExpressionStartContext)
}

func (s *ComposedExpressionContext) AllPrimaryExpressionSuffix() []IPrimaryExpressionSuffixContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPrimaryExpressionSuffixContext)(nil)).Elem())
	var tst = make([]IPrimaryExpressionSuffixContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPrimaryExpressionSuffixContext)
		}
	}

	return tst
}

func (s *ComposedExpressionContext) PrimaryExpressionSuffix(i int) IPrimaryExpressionSuffixContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimaryExpressionSuffixContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPrimaryExpressionSuffixContext)
}

func (s *ComposedExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComposedExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComposedExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.EnterComposedExpression(s)
	}
}

func (s *ComposedExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.ExitComposedExpression(s)
	}
}

func (s *ComposedExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CadenceVisitor:
		return t.VisitComposedExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CadenceParser) ComposedExpression() (localctx IComposedExpressionContext) {
	localctx = NewComposedExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 114, CadenceParserRULE_composedExpression)

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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(618)
		p.PrimaryExpressionStart()
	}
	p.SetState(622)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 51, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(619)
				p.PrimaryExpressionSuffix()
			}

		}
		p.SetState(624)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 51, p.GetParserRuleContext())
	}

	return localctx
}

// IPrimaryExpressionSuffixContext is an interface to support dynamic dispatch.
type IPrimaryExpressionSuffixContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrimaryExpressionSuffixContext differentiates from other interfaces.
	IsPrimaryExpressionSuffixContext()
}

type PrimaryExpressionSuffixContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryExpressionSuffixContext() *PrimaryExpressionSuffixContext {
	var p = new(PrimaryExpressionSuffixContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CadenceParserRULE_primaryExpressionSuffix
	return p
}

func (*PrimaryExpressionSuffixContext) IsPrimaryExpressionSuffixContext() {}

func NewPrimaryExpressionSuffixContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryExpressionSuffixContext {
	var p = new(PrimaryExpressionSuffixContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CadenceParserRULE_primaryExpressionSuffix

	return p
}

func (s *PrimaryExpressionSuffixContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryExpressionSuffixContext) ExpressionAccess() IExpressionAccessContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionAccessContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionAccessContext)
}

func (s *PrimaryExpressionSuffixContext) Invocation() IInvocationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInvocationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInvocationContext)
}

func (s *PrimaryExpressionSuffixContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryExpressionSuffixContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimaryExpressionSuffixContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.EnterPrimaryExpressionSuffix(s)
	}
}

func (s *PrimaryExpressionSuffixContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.ExitPrimaryExpressionSuffix(s)
	}
}

func (s *PrimaryExpressionSuffixContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CadenceVisitor:
		return t.VisitPrimaryExpressionSuffix(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CadenceParser) PrimaryExpressionSuffix() (localctx IPrimaryExpressionSuffixContext) {
	localctx = NewPrimaryExpressionSuffixContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 116, CadenceParserRULE_primaryExpressionSuffix)

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

	p.SetState(627)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CadenceParserT__5, CadenceParserT__11:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(625)
			p.ExpressionAccess()
		}

	case CadenceParserOpenParen:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(626)
			p.Invocation()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IEqualityOpContext is an interface to support dynamic dispatch.
type IEqualityOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEqualityOpContext differentiates from other interfaces.
	IsEqualityOpContext()
}

type EqualityOpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEqualityOpContext() *EqualityOpContext {
	var p = new(EqualityOpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CadenceParserRULE_equalityOp
	return p
}

func (*EqualityOpContext) IsEqualityOpContext() {}

func NewEqualityOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EqualityOpContext {
	var p = new(EqualityOpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CadenceParserRULE_equalityOp

	return p
}

func (s *EqualityOpContext) GetParser() antlr.Parser { return s.parser }

func (s *EqualityOpContext) Equal() antlr.TerminalNode {
	return s.GetToken(CadenceParserEqual, 0)
}

func (s *EqualityOpContext) Unequal() antlr.TerminalNode {
	return s.GetToken(CadenceParserUnequal, 0)
}

func (s *EqualityOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualityOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EqualityOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.EnterEqualityOp(s)
	}
}

func (s *EqualityOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.ExitEqualityOp(s)
	}
}

func (s *EqualityOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CadenceVisitor:
		return t.VisitEqualityOp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CadenceParser) EqualityOp() (localctx IEqualityOpContext) {
	localctx = NewEqualityOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 118, CadenceParserRULE_equalityOp)
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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(629)
		_la = p.GetTokenStream().LA(1)

		if !(_la == CadenceParserEqual || _la == CadenceParserUnequal) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IRelationalOpContext is an interface to support dynamic dispatch.
type IRelationalOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRelationalOpContext differentiates from other interfaces.
	IsRelationalOpContext()
}

type RelationalOpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRelationalOpContext() *RelationalOpContext {
	var p = new(RelationalOpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CadenceParserRULE_relationalOp
	return p
}

func (*RelationalOpContext) IsRelationalOpContext() {}

func NewRelationalOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelationalOpContext {
	var p = new(RelationalOpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CadenceParserRULE_relationalOp

	return p
}

func (s *RelationalOpContext) GetParser() antlr.Parser { return s.parser }

func (s *RelationalOpContext) Less() antlr.TerminalNode {
	return s.GetToken(CadenceParserLess, 0)
}

func (s *RelationalOpContext) Greater() antlr.TerminalNode {
	return s.GetToken(CadenceParserGreater, 0)
}

func (s *RelationalOpContext) LessEqual() antlr.TerminalNode {
	return s.GetToken(CadenceParserLessEqual, 0)
}

func (s *RelationalOpContext) GreaterEqual() antlr.TerminalNode {
	return s.GetToken(CadenceParserGreaterEqual, 0)
}

func (s *RelationalOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationalOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RelationalOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.EnterRelationalOp(s)
	}
}

func (s *RelationalOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.ExitRelationalOp(s)
	}
}

func (s *RelationalOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CadenceVisitor:
		return t.VisitRelationalOp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CadenceParser) RelationalOp() (localctx IRelationalOpContext) {
	localctx = NewRelationalOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 120, CadenceParserRULE_relationalOp)
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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(631)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<CadenceParserLess)|(1<<CadenceParserGreater)|(1<<CadenceParserLessEqual)|(1<<CadenceParserGreaterEqual))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IAdditiveOpContext is an interface to support dynamic dispatch.
type IAdditiveOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAdditiveOpContext differentiates from other interfaces.
	IsAdditiveOpContext()
}

type AdditiveOpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAdditiveOpContext() *AdditiveOpContext {
	var p = new(AdditiveOpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CadenceParserRULE_additiveOp
	return p
}

func (*AdditiveOpContext) IsAdditiveOpContext() {}

func NewAdditiveOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AdditiveOpContext {
	var p = new(AdditiveOpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CadenceParserRULE_additiveOp

	return p
}

func (s *AdditiveOpContext) GetParser() antlr.Parser { return s.parser }

func (s *AdditiveOpContext) Plus() antlr.TerminalNode {
	return s.GetToken(CadenceParserPlus, 0)
}

func (s *AdditiveOpContext) Minus() antlr.TerminalNode {
	return s.GetToken(CadenceParserMinus, 0)
}

func (s *AdditiveOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AdditiveOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AdditiveOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.EnterAdditiveOp(s)
	}
}

func (s *AdditiveOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.ExitAdditiveOp(s)
	}
}

func (s *AdditiveOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CadenceVisitor:
		return t.VisitAdditiveOp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CadenceParser) AdditiveOp() (localctx IAdditiveOpContext) {
	localctx = NewAdditiveOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 122, CadenceParserRULE_additiveOp)
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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(633)
		_la = p.GetTokenStream().LA(1)

		if !(_la == CadenceParserPlus || _la == CadenceParserMinus) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IMultiplicativeOpContext is an interface to support dynamic dispatch.
type IMultiplicativeOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMultiplicativeOpContext differentiates from other interfaces.
	IsMultiplicativeOpContext()
}

type MultiplicativeOpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMultiplicativeOpContext() *MultiplicativeOpContext {
	var p = new(MultiplicativeOpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CadenceParserRULE_multiplicativeOp
	return p
}

func (*MultiplicativeOpContext) IsMultiplicativeOpContext() {}

func NewMultiplicativeOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultiplicativeOpContext {
	var p = new(MultiplicativeOpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CadenceParserRULE_multiplicativeOp

	return p
}

func (s *MultiplicativeOpContext) GetParser() antlr.Parser { return s.parser }

func (s *MultiplicativeOpContext) Mul() antlr.TerminalNode {
	return s.GetToken(CadenceParserMul, 0)
}

func (s *MultiplicativeOpContext) Div() antlr.TerminalNode {
	return s.GetToken(CadenceParserDiv, 0)
}

func (s *MultiplicativeOpContext) Mod() antlr.TerminalNode {
	return s.GetToken(CadenceParserMod, 0)
}

func (s *MultiplicativeOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiplicativeOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MultiplicativeOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.EnterMultiplicativeOp(s)
	}
}

func (s *MultiplicativeOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.ExitMultiplicativeOp(s)
	}
}

func (s *MultiplicativeOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CadenceVisitor:
		return t.VisitMultiplicativeOp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CadenceParser) MultiplicativeOp() (localctx IMultiplicativeOpContext) {
	localctx = NewMultiplicativeOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 124, CadenceParserRULE_multiplicativeOp)
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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(635)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<CadenceParserMul)|(1<<CadenceParserDiv)|(1<<CadenceParserMod))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IUnaryOpContext is an interface to support dynamic dispatch.
type IUnaryOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnaryOpContext differentiates from other interfaces.
	IsUnaryOpContext()
}

type UnaryOpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnaryOpContext() *UnaryOpContext {
	var p = new(UnaryOpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CadenceParserRULE_unaryOp
	return p
}

func (*UnaryOpContext) IsUnaryOpContext() {}

func NewUnaryOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnaryOpContext {
	var p = new(UnaryOpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CadenceParserRULE_unaryOp

	return p
}

func (s *UnaryOpContext) GetParser() antlr.Parser { return s.parser }

func (s *UnaryOpContext) Minus() antlr.TerminalNode {
	return s.GetToken(CadenceParserMinus, 0)
}

func (s *UnaryOpContext) Negate() antlr.TerminalNode {
	return s.GetToken(CadenceParserNegate, 0)
}

func (s *UnaryOpContext) Move() antlr.TerminalNode {
	return s.GetToken(CadenceParserMove, 0)
}

func (s *UnaryOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnaryOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.EnterUnaryOp(s)
	}
}

func (s *UnaryOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.ExitUnaryOp(s)
	}
}

func (s *UnaryOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CadenceVisitor:
		return t.VisitUnaryOp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CadenceParser) UnaryOp() (localctx IUnaryOpContext) {
	localctx = NewUnaryOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 126, CadenceParserRULE_unaryOp)
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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(637)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<CadenceParserMinus)|(1<<CadenceParserNegate)|(1<<CadenceParserMove))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IPrimaryExpressionStartContext is an interface to support dynamic dispatch.
type IPrimaryExpressionStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrimaryExpressionStartContext differentiates from other interfaces.
	IsPrimaryExpressionStartContext()
}

type PrimaryExpressionStartContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryExpressionStartContext() *PrimaryExpressionStartContext {
	var p = new(PrimaryExpressionStartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CadenceParserRULE_primaryExpressionStart
	return p
}

func (*PrimaryExpressionStartContext) IsPrimaryExpressionStartContext() {}

func NewPrimaryExpressionStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryExpressionStartContext {
	var p = new(PrimaryExpressionStartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CadenceParserRULE_primaryExpressionStart

	return p
}

func (s *PrimaryExpressionStartContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryExpressionStartContext) IdentifierExpression() IIdentifierExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierExpressionContext)
}

func (s *PrimaryExpressionStartContext) LiteralExpression() ILiteralExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteralExpressionContext)
}

func (s *PrimaryExpressionStartContext) FunctionExpression() IFunctionExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionExpressionContext)
}

func (s *PrimaryExpressionStartContext) NestedExpression() INestedExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INestedExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INestedExpressionContext)
}

func (s *PrimaryExpressionStartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryExpressionStartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimaryExpressionStartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.EnterPrimaryExpressionStart(s)
	}
}

func (s *PrimaryExpressionStartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.ExitPrimaryExpressionStart(s)
	}
}

func (s *PrimaryExpressionStartContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CadenceVisitor:
		return t.VisitPrimaryExpressionStart(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CadenceParser) PrimaryExpressionStart() (localctx IPrimaryExpressionStartContext) {
	localctx = NewPrimaryExpressionStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 128, CadenceParserRULE_primaryExpressionStart)

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

	p.SetState(643)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CadenceParserFrom, CadenceParserCreate, CadenceParserDestroy, CadenceParserIdentifier:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(639)
			p.IdentifierExpression()
		}

	case CadenceParserT__2, CadenceParserT__5, CadenceParserMinus, CadenceParserTrue, CadenceParserFalse, CadenceParserNil, CadenceParserDecimalLiteral, CadenceParserBinaryLiteral, CadenceParserOctalLiteral, CadenceParserHexadecimalLiteral, CadenceParserInvalidNumberLiteral, CadenceParserStringLiteral:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(640)
			p.LiteralExpression()
		}

	case CadenceParserFun:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(641)
			p.FunctionExpression()
		}

	case CadenceParserOpenParen:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(642)
			p.NestedExpression()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ICreateExpressionContext is an interface to support dynamic dispatch.
type ICreateExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCreateExpressionContext differentiates from other interfaces.
	IsCreateExpressionContext()
}

type CreateExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCreateExpressionContext() *CreateExpressionContext {
	var p = new(CreateExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CadenceParserRULE_createExpression
	return p
}

func (*CreateExpressionContext) IsCreateExpressionContext() {}

func NewCreateExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CreateExpressionContext {
	var p = new(CreateExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CadenceParserRULE_createExpression

	return p
}

func (s *CreateExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *CreateExpressionContext) Create() antlr.TerminalNode {
	return s.GetToken(CadenceParserCreate, 0)
}

func (s *CreateExpressionContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *CreateExpressionContext) Invocation() IInvocationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInvocationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInvocationContext)
}

func (s *CreateExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CreateExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CreateExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.EnterCreateExpression(s)
	}
}

func (s *CreateExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.ExitCreateExpression(s)
	}
}

func (s *CreateExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CadenceVisitor:
		return t.VisitCreateExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CadenceParser) CreateExpression() (localctx ICreateExpressionContext) {
	localctx = NewCreateExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 130, CadenceParserRULE_createExpression)

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
		p.SetState(645)
		p.Match(CadenceParserCreate)
	}
	{
		p.SetState(646)
		p.Identifier()
	}
	{
		p.SetState(647)
		p.Invocation()
	}

	return localctx
}

// IDestroyExpressionContext is an interface to support dynamic dispatch.
type IDestroyExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDestroyExpressionContext differentiates from other interfaces.
	IsDestroyExpressionContext()
}

type DestroyExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDestroyExpressionContext() *DestroyExpressionContext {
	var p = new(DestroyExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CadenceParserRULE_destroyExpression
	return p
}

func (*DestroyExpressionContext) IsDestroyExpressionContext() {}

func NewDestroyExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DestroyExpressionContext {
	var p = new(DestroyExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CadenceParserRULE_destroyExpression

	return p
}

func (s *DestroyExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *DestroyExpressionContext) Destroy() antlr.TerminalNode {
	return s.GetToken(CadenceParserDestroy, 0)
}

func (s *DestroyExpressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *DestroyExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DestroyExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DestroyExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.EnterDestroyExpression(s)
	}
}

func (s *DestroyExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.ExitDestroyExpression(s)
	}
}

func (s *DestroyExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CadenceVisitor:
		return t.VisitDestroyExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CadenceParser) DestroyExpression() (localctx IDestroyExpressionContext) {
	localctx = NewDestroyExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 132, CadenceParserRULE_destroyExpression)

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
		p.SetState(649)
		p.Match(CadenceParserDestroy)
	}
	{
		p.SetState(650)
		p.Expression()
	}

	return localctx
}

// IReferenceExpressionContext is an interface to support dynamic dispatch.
type IReferenceExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReferenceExpressionContext differentiates from other interfaces.
	IsReferenceExpressionContext()
}

type ReferenceExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReferenceExpressionContext() *ReferenceExpressionContext {
	var p = new(ReferenceExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CadenceParserRULE_referenceExpression
	return p
}

func (*ReferenceExpressionContext) IsReferenceExpressionContext() {}

func NewReferenceExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReferenceExpressionContext {
	var p = new(ReferenceExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CadenceParserRULE_referenceExpression

	return p
}

func (s *ReferenceExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ReferenceExpressionContext) Ampersand() antlr.TerminalNode {
	return s.GetToken(CadenceParserAmpersand, 0)
}

func (s *ReferenceExpressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ReferenceExpressionContext) Downcasting() antlr.TerminalNode {
	return s.GetToken(CadenceParserDowncasting, 0)
}

func (s *ReferenceExpressionContext) FullType() IFullTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFullTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFullTypeContext)
}

func (s *ReferenceExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReferenceExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReferenceExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.EnterReferenceExpression(s)
	}
}

func (s *ReferenceExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.ExitReferenceExpression(s)
	}
}

func (s *ReferenceExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CadenceVisitor:
		return t.VisitReferenceExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CadenceParser) ReferenceExpression() (localctx IReferenceExpressionContext) {
	localctx = NewReferenceExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 134, CadenceParserRULE_referenceExpression)

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
		p.SetState(652)
		p.Match(CadenceParserAmpersand)
	}
	{
		p.SetState(653)
		p.Expression()
	}
	{
		p.SetState(654)
		p.Match(CadenceParserDowncasting)
	}
	{
		p.SetState(655)
		p.FullType()
	}

	return localctx
}

// IIdentifierExpressionContext is an interface to support dynamic dispatch.
type IIdentifierExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIdentifierExpressionContext differentiates from other interfaces.
	IsIdentifierExpressionContext()
}

type IdentifierExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifierExpressionContext() *IdentifierExpressionContext {
	var p = new(IdentifierExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CadenceParserRULE_identifierExpression
	return p
}

func (*IdentifierExpressionContext) IsIdentifierExpressionContext() {}

func NewIdentifierExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierExpressionContext {
	var p = new(IdentifierExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CadenceParserRULE_identifierExpression

	return p
}

func (s *IdentifierExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierExpressionContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *IdentifierExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.EnterIdentifierExpression(s)
	}
}

func (s *IdentifierExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.ExitIdentifierExpression(s)
	}
}

func (s *IdentifierExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CadenceVisitor:
		return t.VisitIdentifierExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CadenceParser) IdentifierExpression() (localctx IIdentifierExpressionContext) {
	localctx = NewIdentifierExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 136, CadenceParserRULE_identifierExpression)

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
		p.SetState(657)
		p.Identifier()
	}

	return localctx
}

// ILiteralExpressionContext is an interface to support dynamic dispatch.
type ILiteralExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLiteralExpressionContext differentiates from other interfaces.
	IsLiteralExpressionContext()
}

type LiteralExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralExpressionContext() *LiteralExpressionContext {
	var p = new(LiteralExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CadenceParserRULE_literalExpression
	return p
}

func (*LiteralExpressionContext) IsLiteralExpressionContext() {}

func NewLiteralExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralExpressionContext {
	var p = new(LiteralExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CadenceParserRULE_literalExpression

	return p
}

func (s *LiteralExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralExpressionContext) Literal() ILiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *LiteralExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.EnterLiteralExpression(s)
	}
}

func (s *LiteralExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.ExitLiteralExpression(s)
	}
}

func (s *LiteralExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CadenceVisitor:
		return t.VisitLiteralExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CadenceParser) LiteralExpression() (localctx ILiteralExpressionContext) {
	localctx = NewLiteralExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 138, CadenceParserRULE_literalExpression)

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
		p.SetState(659)
		p.Literal()
	}

	return localctx
}

// IFunctionExpressionContext is an interface to support dynamic dispatch.
type IFunctionExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetReturnType returns the returnType rule contexts.
	GetReturnType() ITypeAnnotationContext

	// SetReturnType sets the returnType rule contexts.
	SetReturnType(ITypeAnnotationContext)

	// IsFunctionExpressionContext differentiates from other interfaces.
	IsFunctionExpressionContext()
}

type FunctionExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	returnType ITypeAnnotationContext
}

func NewEmptyFunctionExpressionContext() *FunctionExpressionContext {
	var p = new(FunctionExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CadenceParserRULE_functionExpression
	return p
}

func (*FunctionExpressionContext) IsFunctionExpressionContext() {}

func NewFunctionExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionExpressionContext {
	var p = new(FunctionExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CadenceParserRULE_functionExpression

	return p
}

func (s *FunctionExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionExpressionContext) GetReturnType() ITypeAnnotationContext { return s.returnType }

func (s *FunctionExpressionContext) SetReturnType(v ITypeAnnotationContext) { s.returnType = v }

func (s *FunctionExpressionContext) Fun() antlr.TerminalNode {
	return s.GetToken(CadenceParserFun, 0)
}

func (s *FunctionExpressionContext) ParameterList() IParameterListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParameterListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParameterListContext)
}

func (s *FunctionExpressionContext) FunctionBlock() IFunctionBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionBlockContext)
}

func (s *FunctionExpressionContext) TypeAnnotation() ITypeAnnotationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeAnnotationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeAnnotationContext)
}

func (s *FunctionExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.EnterFunctionExpression(s)
	}
}

func (s *FunctionExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.ExitFunctionExpression(s)
	}
}

func (s *FunctionExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CadenceVisitor:
		return t.VisitFunctionExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CadenceParser) FunctionExpression() (localctx IFunctionExpressionContext) {
	localctx = NewFunctionExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 140, CadenceParserRULE_functionExpression)
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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(661)
		p.Match(CadenceParserFun)
	}
	{
		p.SetState(662)
		p.ParameterList()
	}
	p.SetState(665)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CadenceParserT__4 {
		{
			p.SetState(663)
			p.Match(CadenceParserT__4)
		}
		{
			p.SetState(664)

			var _x = p.TypeAnnotation()

			localctx.(*FunctionExpressionContext).returnType = _x
		}

	}
	{
		p.SetState(667)
		p.FunctionBlock()
	}

	return localctx
}

// INestedExpressionContext is an interface to support dynamic dispatch.
type INestedExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNestedExpressionContext differentiates from other interfaces.
	IsNestedExpressionContext()
}

type NestedExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNestedExpressionContext() *NestedExpressionContext {
	var p = new(NestedExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CadenceParserRULE_nestedExpression
	return p
}

func (*NestedExpressionContext) IsNestedExpressionContext() {}

func NewNestedExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NestedExpressionContext {
	var p = new(NestedExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CadenceParserRULE_nestedExpression

	return p
}

func (s *NestedExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *NestedExpressionContext) OpenParen() antlr.TerminalNode {
	return s.GetToken(CadenceParserOpenParen, 0)
}

func (s *NestedExpressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *NestedExpressionContext) CloseParen() antlr.TerminalNode {
	return s.GetToken(CadenceParserCloseParen, 0)
}

func (s *NestedExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NestedExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NestedExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.EnterNestedExpression(s)
	}
}

func (s *NestedExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.ExitNestedExpression(s)
	}
}

func (s *NestedExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CadenceVisitor:
		return t.VisitNestedExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CadenceParser) NestedExpression() (localctx INestedExpressionContext) {
	localctx = NewNestedExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 142, CadenceParserRULE_nestedExpression)

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
		p.SetState(669)
		p.Match(CadenceParserOpenParen)
	}
	{
		p.SetState(670)
		p.Expression()
	}
	{
		p.SetState(671)
		p.Match(CadenceParserCloseParen)
	}

	return localctx
}

// IExpressionAccessContext is an interface to support dynamic dispatch.
type IExpressionAccessContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionAccessContext differentiates from other interfaces.
	IsExpressionAccessContext()
}

type ExpressionAccessContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionAccessContext() *ExpressionAccessContext {
	var p = new(ExpressionAccessContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CadenceParserRULE_expressionAccess
	return p
}

func (*ExpressionAccessContext) IsExpressionAccessContext() {}

func NewExpressionAccessContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionAccessContext {
	var p = new(ExpressionAccessContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CadenceParserRULE_expressionAccess

	return p
}

func (s *ExpressionAccessContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionAccessContext) MemberAccess() IMemberAccessContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMemberAccessContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMemberAccessContext)
}

func (s *ExpressionAccessContext) BracketExpression() IBracketExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBracketExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBracketExpressionContext)
}

func (s *ExpressionAccessContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionAccessContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionAccessContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.EnterExpressionAccess(s)
	}
}

func (s *ExpressionAccessContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.ExitExpressionAccess(s)
	}
}

func (s *ExpressionAccessContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CadenceVisitor:
		return t.VisitExpressionAccess(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CadenceParser) ExpressionAccess() (localctx IExpressionAccessContext) {
	localctx = NewExpressionAccessContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 144, CadenceParserRULE_expressionAccess)

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

	p.SetState(675)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CadenceParserT__11:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(673)
			p.MemberAccess()
		}

	case CadenceParserT__5:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(674)
			p.BracketExpression()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IMemberAccessContext is an interface to support dynamic dispatch.
type IMemberAccessContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMemberAccessContext differentiates from other interfaces.
	IsMemberAccessContext()
}

type MemberAccessContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMemberAccessContext() *MemberAccessContext {
	var p = new(MemberAccessContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CadenceParserRULE_memberAccess
	return p
}

func (*MemberAccessContext) IsMemberAccessContext() {}

func NewMemberAccessContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MemberAccessContext {
	var p = new(MemberAccessContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CadenceParserRULE_memberAccess

	return p
}

func (s *MemberAccessContext) GetParser() antlr.Parser { return s.parser }

func (s *MemberAccessContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *MemberAccessContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MemberAccessContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MemberAccessContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.EnterMemberAccess(s)
	}
}

func (s *MemberAccessContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.ExitMemberAccess(s)
	}
}

func (s *MemberAccessContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CadenceVisitor:
		return t.VisitMemberAccess(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CadenceParser) MemberAccess() (localctx IMemberAccessContext) {
	localctx = NewMemberAccessContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 146, CadenceParserRULE_memberAccess)

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
		p.SetState(677)
		p.Match(CadenceParserT__11)
	}
	{
		p.SetState(678)
		p.Identifier()
	}

	return localctx
}

// IBracketExpressionContext is an interface to support dynamic dispatch.
type IBracketExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBracketExpressionContext differentiates from other interfaces.
	IsBracketExpressionContext()
}

type BracketExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBracketExpressionContext() *BracketExpressionContext {
	var p = new(BracketExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CadenceParserRULE_bracketExpression
	return p
}

func (*BracketExpressionContext) IsBracketExpressionContext() {}

func NewBracketExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BracketExpressionContext {
	var p = new(BracketExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CadenceParserRULE_bracketExpression

	return p
}

func (s *BracketExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *BracketExpressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *BracketExpressionContext) FullType() IFullTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFullTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFullTypeContext)
}

func (s *BracketExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BracketExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BracketExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.EnterBracketExpression(s)
	}
}

func (s *BracketExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.ExitBracketExpression(s)
	}
}

func (s *BracketExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CadenceVisitor:
		return t.VisitBracketExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CadenceParser) BracketExpression() (localctx IBracketExpressionContext) {
	localctx = NewBracketExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 148, CadenceParserRULE_bracketExpression)

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
		p.SetState(680)
		p.Match(CadenceParserT__5)
	}
	p.SetState(683)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 56, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(681)
			p.Expression()
		}

	case 2:
		{
			p.SetState(682)
			p.FullType()
		}

	}
	{
		p.SetState(685)
		p.Match(CadenceParserT__6)
	}

	return localctx
}

// IInvocationContext is an interface to support dynamic dispatch.
type IInvocationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInvocationContext differentiates from other interfaces.
	IsInvocationContext()
}

type InvocationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInvocationContext() *InvocationContext {
	var p = new(InvocationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CadenceParserRULE_invocation
	return p
}

func (*InvocationContext) IsInvocationContext() {}

func NewInvocationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InvocationContext {
	var p = new(InvocationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CadenceParserRULE_invocation

	return p
}

func (s *InvocationContext) GetParser() antlr.Parser { return s.parser }

func (s *InvocationContext) OpenParen() antlr.TerminalNode {
	return s.GetToken(CadenceParserOpenParen, 0)
}

func (s *InvocationContext) CloseParen() antlr.TerminalNode {
	return s.GetToken(CadenceParserCloseParen, 0)
}

func (s *InvocationContext) AllArgument() []IArgumentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IArgumentContext)(nil)).Elem())
	var tst = make([]IArgumentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IArgumentContext)
		}
	}

	return tst
}

func (s *InvocationContext) Argument(i int) IArgumentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgumentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IArgumentContext)
}

func (s *InvocationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InvocationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InvocationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.EnterInvocation(s)
	}
}

func (s *InvocationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.ExitInvocation(s)
	}
}

func (s *InvocationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CadenceVisitor:
		return t.VisitInvocation(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CadenceParser) Invocation() (localctx IInvocationContext) {
	localctx = NewInvocationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 150, CadenceParserRULE_invocation)
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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(687)
		p.Match(CadenceParserOpenParen)
	}
	p.SetState(696)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<CadenceParserT__2)|(1<<CadenceParserT__5)|(1<<CadenceParserMinus)|(1<<CadenceParserAmpersand)|(1<<CadenceParserNegate)|(1<<CadenceParserMove)|(1<<CadenceParserOpenParen))) != 0) || (((_la-38)&-(0x1f+1)) == 0 && ((1<<uint((_la-38)))&((1<<(CadenceParserFun-38))|(1<<(CadenceParserTrue-38))|(1<<(CadenceParserFalse-38))|(1<<(CadenceParserNil-38))|(1<<(CadenceParserFrom-38))|(1<<(CadenceParserCreate-38))|(1<<(CadenceParserDestroy-38))|(1<<(CadenceParserIdentifier-38))|(1<<(CadenceParserDecimalLiteral-38))|(1<<(CadenceParserBinaryLiteral-38))|(1<<(CadenceParserOctalLiteral-38))|(1<<(CadenceParserHexadecimalLiteral-38))|(1<<(CadenceParserInvalidNumberLiteral-38))|(1<<(CadenceParserStringLiteral-38)))) != 0) {
		{
			p.SetState(688)
			p.Argument()
		}
		p.SetState(693)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == CadenceParserT__1 {
			{
				p.SetState(689)
				p.Match(CadenceParserT__1)
			}
			{
				p.SetState(690)
				p.Argument()
			}

			p.SetState(695)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(698)
		p.Match(CadenceParserCloseParen)
	}

	return localctx
}

// IArgumentContext is an interface to support dynamic dispatch.
type IArgumentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgumentContext differentiates from other interfaces.
	IsArgumentContext()
}

type ArgumentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentContext() *ArgumentContext {
	var p = new(ArgumentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CadenceParserRULE_argument
	return p
}

func (*ArgumentContext) IsArgumentContext() {}

func NewArgumentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentContext {
	var p = new(ArgumentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CadenceParserRULE_argument

	return p
}

func (s *ArgumentContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ArgumentContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *ArgumentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.EnterArgument(s)
	}
}

func (s *ArgumentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.ExitArgument(s)
	}
}

func (s *ArgumentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CadenceVisitor:
		return t.VisitArgument(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CadenceParser) Argument() (localctx IArgumentContext) {
	localctx = NewArgumentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 152, CadenceParserRULE_argument)

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
	p.SetState(703)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 59, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(700)
			p.Identifier()
		}
		{
			p.SetState(701)
			p.Match(CadenceParserT__4)
		}

	}
	{
		p.SetState(705)
		p.Expression()
	}

	return localctx
}

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CadenceParserRULE_literal
	return p
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CadenceParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) IntegerLiteral() IIntegerLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntegerLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIntegerLiteralContext)
}

func (s *LiteralContext) BooleanLiteral() IBooleanLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBooleanLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBooleanLiteralContext)
}

func (s *LiteralContext) ArrayLiteral() IArrayLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArrayLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArrayLiteralContext)
}

func (s *LiteralContext) DictionaryLiteral() IDictionaryLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDictionaryLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDictionaryLiteralContext)
}

func (s *LiteralContext) StringLiteral() IStringLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStringLiteralContext)
}

func (s *LiteralContext) NilLiteral() INilLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INilLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INilLiteralContext)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.EnterLiteral(s)
	}
}

func (s *LiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.ExitLiteral(s)
	}
}

func (s *LiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CadenceVisitor:
		return t.VisitLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CadenceParser) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 154, CadenceParserRULE_literal)

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

	p.SetState(713)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CadenceParserMinus, CadenceParserDecimalLiteral, CadenceParserBinaryLiteral, CadenceParserOctalLiteral, CadenceParserHexadecimalLiteral, CadenceParserInvalidNumberLiteral:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(707)
			p.IntegerLiteral()
		}

	case CadenceParserTrue, CadenceParserFalse:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(708)
			p.BooleanLiteral()
		}

	case CadenceParserT__5:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(709)
			p.ArrayLiteral()
		}

	case CadenceParserT__2:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(710)
			p.DictionaryLiteral()
		}

	case CadenceParserStringLiteral:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(711)
			p.StringLiteral()
		}

	case CadenceParserNil:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(712)
			p.NilLiteral()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IBooleanLiteralContext is an interface to support dynamic dispatch.
type IBooleanLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBooleanLiteralContext differentiates from other interfaces.
	IsBooleanLiteralContext()
}

type BooleanLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBooleanLiteralContext() *BooleanLiteralContext {
	var p = new(BooleanLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CadenceParserRULE_booleanLiteral
	return p
}

func (*BooleanLiteralContext) IsBooleanLiteralContext() {}

func NewBooleanLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BooleanLiteralContext {
	var p = new(BooleanLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CadenceParserRULE_booleanLiteral

	return p
}

func (s *BooleanLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *BooleanLiteralContext) True() antlr.TerminalNode {
	return s.GetToken(CadenceParserTrue, 0)
}

func (s *BooleanLiteralContext) False() antlr.TerminalNode {
	return s.GetToken(CadenceParserFalse, 0)
}

func (s *BooleanLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BooleanLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.EnterBooleanLiteral(s)
	}
}

func (s *BooleanLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.ExitBooleanLiteral(s)
	}
}

func (s *BooleanLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CadenceVisitor:
		return t.VisitBooleanLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CadenceParser) BooleanLiteral() (localctx IBooleanLiteralContext) {
	localctx = NewBooleanLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 156, CadenceParserRULE_booleanLiteral)
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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(715)
		_la = p.GetTokenStream().LA(1)

		if !(_la == CadenceParserTrue || _la == CadenceParserFalse) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// INilLiteralContext is an interface to support dynamic dispatch.
type INilLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNilLiteralContext differentiates from other interfaces.
	IsNilLiteralContext()
}

type NilLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNilLiteralContext() *NilLiteralContext {
	var p = new(NilLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CadenceParserRULE_nilLiteral
	return p
}

func (*NilLiteralContext) IsNilLiteralContext() {}

func NewNilLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NilLiteralContext {
	var p = new(NilLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CadenceParserRULE_nilLiteral

	return p
}

func (s *NilLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *NilLiteralContext) Nil() antlr.TerminalNode {
	return s.GetToken(CadenceParserNil, 0)
}

func (s *NilLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NilLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NilLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.EnterNilLiteral(s)
	}
}

func (s *NilLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.ExitNilLiteral(s)
	}
}

func (s *NilLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CadenceVisitor:
		return t.VisitNilLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CadenceParser) NilLiteral() (localctx INilLiteralContext) {
	localctx = NewNilLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 158, CadenceParserRULE_nilLiteral)

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
		p.SetState(717)
		p.Match(CadenceParserNil)
	}

	return localctx
}

// IStringLiteralContext is an interface to support dynamic dispatch.
type IStringLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStringLiteralContext differentiates from other interfaces.
	IsStringLiteralContext()
}

type StringLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStringLiteralContext() *StringLiteralContext {
	var p = new(StringLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CadenceParserRULE_stringLiteral
	return p
}

func (*StringLiteralContext) IsStringLiteralContext() {}

func NewStringLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringLiteralContext {
	var p = new(StringLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CadenceParserRULE_stringLiteral

	return p
}

func (s *StringLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *StringLiteralContext) StringLiteral() antlr.TerminalNode {
	return s.GetToken(CadenceParserStringLiteral, 0)
}

func (s *StringLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StringLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.EnterStringLiteral(s)
	}
}

func (s *StringLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.ExitStringLiteral(s)
	}
}

func (s *StringLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CadenceVisitor:
		return t.VisitStringLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CadenceParser) StringLiteral() (localctx IStringLiteralContext) {
	localctx = NewStringLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 160, CadenceParserRULE_stringLiteral)

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
		p.SetState(719)
		p.Match(CadenceParserStringLiteral)
	}

	return localctx
}

// IIntegerLiteralContext is an interface to support dynamic dispatch.
type IIntegerLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIntegerLiteralContext differentiates from other interfaces.
	IsIntegerLiteralContext()
}

type IntegerLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntegerLiteralContext() *IntegerLiteralContext {
	var p = new(IntegerLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CadenceParserRULE_integerLiteral
	return p
}

func (*IntegerLiteralContext) IsIntegerLiteralContext() {}

func NewIntegerLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntegerLiteralContext {
	var p = new(IntegerLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CadenceParserRULE_integerLiteral

	return p
}

func (s *IntegerLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *IntegerLiteralContext) PositiveIntegerLiteral() IPositiveIntegerLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPositiveIntegerLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPositiveIntegerLiteralContext)
}

func (s *IntegerLiteralContext) Minus() antlr.TerminalNode {
	return s.GetToken(CadenceParserMinus, 0)
}

func (s *IntegerLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntegerLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntegerLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.EnterIntegerLiteral(s)
	}
}

func (s *IntegerLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.ExitIntegerLiteral(s)
	}
}

func (s *IntegerLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CadenceVisitor:
		return t.VisitIntegerLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CadenceParser) IntegerLiteral() (localctx IIntegerLiteralContext) {
	localctx = NewIntegerLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 162, CadenceParserRULE_integerLiteral)
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

	p.EnterOuterAlt(localctx, 1)
	p.SetState(722)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CadenceParserMinus {
		{
			p.SetState(721)
			p.Match(CadenceParserMinus)
		}

	}
	{
		p.SetState(724)
		p.PositiveIntegerLiteral()
	}

	return localctx
}

// IPositiveIntegerLiteralContext is an interface to support dynamic dispatch.
type IPositiveIntegerLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPositiveIntegerLiteralContext differentiates from other interfaces.
	IsPositiveIntegerLiteralContext()
}

type PositiveIntegerLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPositiveIntegerLiteralContext() *PositiveIntegerLiteralContext {
	var p = new(PositiveIntegerLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CadenceParserRULE_positiveIntegerLiteral
	return p
}

func (*PositiveIntegerLiteralContext) IsPositiveIntegerLiteralContext() {}

func NewPositiveIntegerLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PositiveIntegerLiteralContext {
	var p = new(PositiveIntegerLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CadenceParserRULE_positiveIntegerLiteral

	return p
}

func (s *PositiveIntegerLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *PositiveIntegerLiteralContext) CopyFrom(ctx *PositiveIntegerLiteralContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *PositiveIntegerLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PositiveIntegerLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type BinaryLiteralContext struct {
	*PositiveIntegerLiteralContext
}

func NewBinaryLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BinaryLiteralContext {
	var p = new(BinaryLiteralContext)

	p.PositiveIntegerLiteralContext = NewEmptyPositiveIntegerLiteralContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PositiveIntegerLiteralContext))

	return p
}

func (s *BinaryLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryLiteralContext) BinaryLiteral() antlr.TerminalNode {
	return s.GetToken(CadenceParserBinaryLiteral, 0)
}

func (s *BinaryLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.EnterBinaryLiteral(s)
	}
}

func (s *BinaryLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.ExitBinaryLiteral(s)
	}
}

func (s *BinaryLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CadenceVisitor:
		return t.VisitBinaryLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

type OctalLiteralContext struct {
	*PositiveIntegerLiteralContext
}

func NewOctalLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OctalLiteralContext {
	var p = new(OctalLiteralContext)

	p.PositiveIntegerLiteralContext = NewEmptyPositiveIntegerLiteralContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PositiveIntegerLiteralContext))

	return p
}

func (s *OctalLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OctalLiteralContext) OctalLiteral() antlr.TerminalNode {
	return s.GetToken(CadenceParserOctalLiteral, 0)
}

func (s *OctalLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.EnterOctalLiteral(s)
	}
}

func (s *OctalLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.ExitOctalLiteral(s)
	}
}

func (s *OctalLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CadenceVisitor:
		return t.VisitOctalLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

type InvalidNumberLiteralContext struct {
	*PositiveIntegerLiteralContext
}

func NewInvalidNumberLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InvalidNumberLiteralContext {
	var p = new(InvalidNumberLiteralContext)

	p.PositiveIntegerLiteralContext = NewEmptyPositiveIntegerLiteralContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PositiveIntegerLiteralContext))

	return p
}

func (s *InvalidNumberLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InvalidNumberLiteralContext) InvalidNumberLiteral() antlr.TerminalNode {
	return s.GetToken(CadenceParserInvalidNumberLiteral, 0)
}

func (s *InvalidNumberLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.EnterInvalidNumberLiteral(s)
	}
}

func (s *InvalidNumberLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.ExitInvalidNumberLiteral(s)
	}
}

func (s *InvalidNumberLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CadenceVisitor:
		return t.VisitInvalidNumberLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

type DecimalLiteralContext struct {
	*PositiveIntegerLiteralContext
}

func NewDecimalLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DecimalLiteralContext {
	var p = new(DecimalLiteralContext)

	p.PositiveIntegerLiteralContext = NewEmptyPositiveIntegerLiteralContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PositiveIntegerLiteralContext))

	return p
}

func (s *DecimalLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DecimalLiteralContext) DecimalLiteral() antlr.TerminalNode {
	return s.GetToken(CadenceParserDecimalLiteral, 0)
}

func (s *DecimalLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.EnterDecimalLiteral(s)
	}
}

func (s *DecimalLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.ExitDecimalLiteral(s)
	}
}

func (s *DecimalLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CadenceVisitor:
		return t.VisitDecimalLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

type HexadecimalLiteralContext struct {
	*PositiveIntegerLiteralContext
}

func NewHexadecimalLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *HexadecimalLiteralContext {
	var p = new(HexadecimalLiteralContext)

	p.PositiveIntegerLiteralContext = NewEmptyPositiveIntegerLiteralContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PositiveIntegerLiteralContext))

	return p
}

func (s *HexadecimalLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HexadecimalLiteralContext) HexadecimalLiteral() antlr.TerminalNode {
	return s.GetToken(CadenceParserHexadecimalLiteral, 0)
}

func (s *HexadecimalLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.EnterHexadecimalLiteral(s)
	}
}

func (s *HexadecimalLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.ExitHexadecimalLiteral(s)
	}
}

func (s *HexadecimalLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CadenceVisitor:
		return t.VisitHexadecimalLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CadenceParser) PositiveIntegerLiteral() (localctx IPositiveIntegerLiteralContext) {
	localctx = NewPositiveIntegerLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 164, CadenceParserRULE_positiveIntegerLiteral)

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

	p.SetState(731)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CadenceParserDecimalLiteral:
		localctx = NewDecimalLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(726)
			p.Match(CadenceParserDecimalLiteral)
		}

	case CadenceParserBinaryLiteral:
		localctx = NewBinaryLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(727)
			p.Match(CadenceParserBinaryLiteral)
		}

	case CadenceParserOctalLiteral:
		localctx = NewOctalLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(728)
			p.Match(CadenceParserOctalLiteral)
		}

	case CadenceParserHexadecimalLiteral:
		localctx = NewHexadecimalLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(729)
			p.Match(CadenceParserHexadecimalLiteral)
		}

	case CadenceParserInvalidNumberLiteral:
		localctx = NewInvalidNumberLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(730)
			p.Match(CadenceParserInvalidNumberLiteral)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IArrayLiteralContext is an interface to support dynamic dispatch.
type IArrayLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArrayLiteralContext differentiates from other interfaces.
	IsArrayLiteralContext()
}

type ArrayLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayLiteralContext() *ArrayLiteralContext {
	var p = new(ArrayLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CadenceParserRULE_arrayLiteral
	return p
}

func (*ArrayLiteralContext) IsArrayLiteralContext() {}

func NewArrayLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayLiteralContext {
	var p = new(ArrayLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CadenceParserRULE_arrayLiteral

	return p
}

func (s *ArrayLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayLiteralContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ArrayLiteralContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ArrayLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.EnterArrayLiteral(s)
	}
}

func (s *ArrayLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.ExitArrayLiteral(s)
	}
}

func (s *ArrayLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CadenceVisitor:
		return t.VisitArrayLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CadenceParser) ArrayLiteral() (localctx IArrayLiteralContext) {
	localctx = NewArrayLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 166, CadenceParserRULE_arrayLiteral)
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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(733)
		p.Match(CadenceParserT__5)
	}
	p.SetState(742)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<CadenceParserT__2)|(1<<CadenceParserT__5)|(1<<CadenceParserMinus)|(1<<CadenceParserAmpersand)|(1<<CadenceParserNegate)|(1<<CadenceParserMove)|(1<<CadenceParserOpenParen))) != 0) || (((_la-38)&-(0x1f+1)) == 0 && ((1<<uint((_la-38)))&((1<<(CadenceParserFun-38))|(1<<(CadenceParserTrue-38))|(1<<(CadenceParserFalse-38))|(1<<(CadenceParserNil-38))|(1<<(CadenceParserFrom-38))|(1<<(CadenceParserCreate-38))|(1<<(CadenceParserDestroy-38))|(1<<(CadenceParserIdentifier-38))|(1<<(CadenceParserDecimalLiteral-38))|(1<<(CadenceParserBinaryLiteral-38))|(1<<(CadenceParserOctalLiteral-38))|(1<<(CadenceParserHexadecimalLiteral-38))|(1<<(CadenceParserInvalidNumberLiteral-38))|(1<<(CadenceParserStringLiteral-38)))) != 0) {
		{
			p.SetState(734)
			p.Expression()
		}
		p.SetState(739)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == CadenceParserT__1 {
			{
				p.SetState(735)
				p.Match(CadenceParserT__1)
			}
			{
				p.SetState(736)
				p.Expression()
			}

			p.SetState(741)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(744)
		p.Match(CadenceParserT__6)
	}

	return localctx
}

// IDictionaryLiteralContext is an interface to support dynamic dispatch.
type IDictionaryLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDictionaryLiteralContext differentiates from other interfaces.
	IsDictionaryLiteralContext()
}

type DictionaryLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDictionaryLiteralContext() *DictionaryLiteralContext {
	var p = new(DictionaryLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CadenceParserRULE_dictionaryLiteral
	return p
}

func (*DictionaryLiteralContext) IsDictionaryLiteralContext() {}

func NewDictionaryLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DictionaryLiteralContext {
	var p = new(DictionaryLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CadenceParserRULE_dictionaryLiteral

	return p
}

func (s *DictionaryLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *DictionaryLiteralContext) AllDictionaryEntry() []IDictionaryEntryContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDictionaryEntryContext)(nil)).Elem())
	var tst = make([]IDictionaryEntryContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDictionaryEntryContext)
		}
	}

	return tst
}

func (s *DictionaryLiteralContext) DictionaryEntry(i int) IDictionaryEntryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDictionaryEntryContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDictionaryEntryContext)
}

func (s *DictionaryLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DictionaryLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DictionaryLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.EnterDictionaryLiteral(s)
	}
}

func (s *DictionaryLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.ExitDictionaryLiteral(s)
	}
}

func (s *DictionaryLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CadenceVisitor:
		return t.VisitDictionaryLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CadenceParser) DictionaryLiteral() (localctx IDictionaryLiteralContext) {
	localctx = NewDictionaryLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 168, CadenceParserRULE_dictionaryLiteral)
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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(746)
		p.Match(CadenceParserT__2)
	}
	p.SetState(755)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<CadenceParserT__2)|(1<<CadenceParserT__5)|(1<<CadenceParserMinus)|(1<<CadenceParserAmpersand)|(1<<CadenceParserNegate)|(1<<CadenceParserMove)|(1<<CadenceParserOpenParen))) != 0) || (((_la-38)&-(0x1f+1)) == 0 && ((1<<uint((_la-38)))&((1<<(CadenceParserFun-38))|(1<<(CadenceParserTrue-38))|(1<<(CadenceParserFalse-38))|(1<<(CadenceParserNil-38))|(1<<(CadenceParserFrom-38))|(1<<(CadenceParserCreate-38))|(1<<(CadenceParserDestroy-38))|(1<<(CadenceParserIdentifier-38))|(1<<(CadenceParserDecimalLiteral-38))|(1<<(CadenceParserBinaryLiteral-38))|(1<<(CadenceParserOctalLiteral-38))|(1<<(CadenceParserHexadecimalLiteral-38))|(1<<(CadenceParserInvalidNumberLiteral-38))|(1<<(CadenceParserStringLiteral-38)))) != 0) {
		{
			p.SetState(747)
			p.DictionaryEntry()
		}
		p.SetState(752)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == CadenceParserT__1 {
			{
				p.SetState(748)
				p.Match(CadenceParserT__1)
			}
			{
				p.SetState(749)
				p.DictionaryEntry()
			}

			p.SetState(754)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(757)
		p.Match(CadenceParserT__3)
	}

	return localctx
}

// IDictionaryEntryContext is an interface to support dynamic dispatch.
type IDictionaryEntryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetKey returns the key rule contexts.
	GetKey() IExpressionContext

	// GetValue returns the value rule contexts.
	GetValue() IExpressionContext

	// SetKey sets the key rule contexts.
	SetKey(IExpressionContext)

	// SetValue sets the value rule contexts.
	SetValue(IExpressionContext)

	// IsDictionaryEntryContext differentiates from other interfaces.
	IsDictionaryEntryContext()
}

type DictionaryEntryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	key    IExpressionContext
	value  IExpressionContext
}

func NewEmptyDictionaryEntryContext() *DictionaryEntryContext {
	var p = new(DictionaryEntryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CadenceParserRULE_dictionaryEntry
	return p
}

func (*DictionaryEntryContext) IsDictionaryEntryContext() {}

func NewDictionaryEntryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DictionaryEntryContext {
	var p = new(DictionaryEntryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CadenceParserRULE_dictionaryEntry

	return p
}

func (s *DictionaryEntryContext) GetParser() antlr.Parser { return s.parser }

func (s *DictionaryEntryContext) GetKey() IExpressionContext { return s.key }

func (s *DictionaryEntryContext) GetValue() IExpressionContext { return s.value }

func (s *DictionaryEntryContext) SetKey(v IExpressionContext) { s.key = v }

func (s *DictionaryEntryContext) SetValue(v IExpressionContext) { s.value = v }

func (s *DictionaryEntryContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *DictionaryEntryContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *DictionaryEntryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DictionaryEntryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DictionaryEntryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.EnterDictionaryEntry(s)
	}
}

func (s *DictionaryEntryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.ExitDictionaryEntry(s)
	}
}

func (s *DictionaryEntryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CadenceVisitor:
		return t.VisitDictionaryEntry(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CadenceParser) DictionaryEntry() (localctx IDictionaryEntryContext) {
	localctx = NewDictionaryEntryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 170, CadenceParserRULE_dictionaryEntry)

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
		p.SetState(759)

		var _x = p.Expression()

		localctx.(*DictionaryEntryContext).key = _x
	}
	{
		p.SetState(760)
		p.Match(CadenceParserT__4)
	}
	{
		p.SetState(761)

		var _x = p.Expression()

		localctx.(*DictionaryEntryContext).value = _x
	}

	return localctx
}

// IIdentifierContext is an interface to support dynamic dispatch.
type IIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIdentifierContext differentiates from other interfaces.
	IsIdentifierContext()
}

type IdentifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifierContext() *IdentifierContext {
	var p = new(IdentifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CadenceParserRULE_identifier
	return p
}

func (*IdentifierContext) IsIdentifierContext() {}

func NewIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierContext {
	var p = new(IdentifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CadenceParserRULE_identifier

	return p
}

func (s *IdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierContext) Identifier() antlr.TerminalNode {
	return s.GetToken(CadenceParserIdentifier, 0)
}

func (s *IdentifierContext) From() antlr.TerminalNode {
	return s.GetToken(CadenceParserFrom, 0)
}

func (s *IdentifierContext) Create() antlr.TerminalNode {
	return s.GetToken(CadenceParserCreate, 0)
}

func (s *IdentifierContext) Destroy() antlr.TerminalNode {
	return s.GetToken(CadenceParserDestroy, 0)
}

func (s *IdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.EnterIdentifier(s)
	}
}

func (s *IdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.ExitIdentifier(s)
	}
}

func (s *IdentifierContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CadenceVisitor:
		return t.VisitIdentifier(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CadenceParser) Identifier() (localctx IIdentifierContext) {
	localctx = NewIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 172, CadenceParserRULE_identifier)
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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(763)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-58)&-(0x1f+1)) == 0 && ((1<<uint((_la-58)))&((1<<(CadenceParserFrom-58))|(1<<(CadenceParserCreate-58))|(1<<(CadenceParserDestroy-58))|(1<<(CadenceParserIdentifier-58)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IEosContext is an interface to support dynamic dispatch.
type IEosContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEosContext differentiates from other interfaces.
	IsEosContext()
}

type EosContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEosContext() *EosContext {
	var p = new(EosContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CadenceParserRULE_eos
	return p
}

func (*EosContext) IsEosContext() {}

func NewEosContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EosContext {
	var p = new(EosContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CadenceParserRULE_eos

	return p
}

func (s *EosContext) GetParser() antlr.Parser { return s.parser }

func (s *EosContext) EOF() antlr.TerminalNode {
	return s.GetToken(CadenceParserEOF, 0)
}

func (s *EosContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EosContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EosContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.EnterEos(s)
	}
}

func (s *EosContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadenceListener); ok {
		listenerT.ExitEos(s)
	}
}

func (s *EosContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CadenceVisitor:
		return t.VisitEos(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CadenceParser) Eos() (localctx IEosContext) {
	localctx = NewEosContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 174, CadenceParserRULE_eos)

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

	p.SetState(769)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 67, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(765)
			p.Match(CadenceParserT__0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(766)
			p.Match(CadenceParserEOF)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		p.SetState(767)

		if !(p.lineTerminatorAhead()) {
			panic(antlr.NewFailedPredicateException(p, "p.lineTerminatorAhead()", ""))
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		p.SetState(768)

		if !(p.GetTokenStream().LT(1).GetText() == "}") {
			panic(antlr.NewFailedPredicateException(p, "p.GetTokenStream().LT(1).GetText() == \"}\"", ""))
		}

	}

	return localctx
}

func (p *CadenceParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 13:
		var t *SpecialFunctionDeclarationContext = nil
		if localctx != nil {
			t = localctx.(*SpecialFunctionDeclarationContext)
		}
		return p.SpecialFunctionDeclaration_Sempred(t, predIndex)

	case 14:
		var t *FunctionDeclarationContext = nil
		if localctx != nil {
			t = localctx.(*FunctionDeclarationContext)
		}
		return p.FunctionDeclaration_Sempred(t, predIndex)

	case 19:
		var t *FullTypeContext = nil
		if localctx != nil {
			t = localctx.(*FullTypeContext)
		}
		return p.FullType_Sempred(t, predIndex)

	case 34:
		var t *ReturnStatementContext = nil
		if localctx != nil {
			t = localctx.(*ReturnStatementContext)
		}
		return p.ReturnStatement_Sempred(t, predIndex)

	case 46:
		var t *OrExpressionContext = nil
		if localctx != nil {
			t = localctx.(*OrExpressionContext)
		}
		return p.OrExpression_Sempred(t, predIndex)

	case 47:
		var t *AndExpressionContext = nil
		if localctx != nil {
			t = localctx.(*AndExpressionContext)
		}
		return p.AndExpression_Sempred(t, predIndex)

	case 48:
		var t *EqualityExpressionContext = nil
		if localctx != nil {
			t = localctx.(*EqualityExpressionContext)
		}
		return p.EqualityExpression_Sempred(t, predIndex)

	case 49:
		var t *RelationalExpressionContext = nil
		if localctx != nil {
			t = localctx.(*RelationalExpressionContext)
		}
		return p.RelationalExpression_Sempred(t, predIndex)

	case 51:
		var t *FailableDowncastingExpressionContext = nil
		if localctx != nil {
			t = localctx.(*FailableDowncastingExpressionContext)
		}
		return p.FailableDowncastingExpression_Sempred(t, predIndex)

	case 52:
		var t *ConcatenatingExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ConcatenatingExpressionContext)
		}
		return p.ConcatenatingExpression_Sempred(t, predIndex)

	case 53:
		var t *AdditiveExpressionContext = nil
		if localctx != nil {
			t = localctx.(*AdditiveExpressionContext)
		}
		return p.AdditiveExpression_Sempred(t, predIndex)

	case 54:
		var t *MultiplicativeExpressionContext = nil
		if localctx != nil {
			t = localctx.(*MultiplicativeExpressionContext)
		}
		return p.MultiplicativeExpression_Sempred(t, predIndex)

	case 87:
		var t *EosContext = nil
		if localctx != nil {
			t = localctx.(*EosContext)
		}
		return p.Eos_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *CadenceParser) SpecialFunctionDeclaration_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return !localctx.(*SpecialFunctionDeclarationContext).functionBlockRequired || localctx.(*SpecialFunctionDeclarationContext).b != nil

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *CadenceParser) FunctionDeclaration_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return !localctx.(*FunctionDeclarationContext).functionBlockRequired || localctx.(*FunctionDeclarationContext).b != nil

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *CadenceParser) FullType_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 2:
		return p.noWhitespace()

	case 3:
		return p.noWhitespace()

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *CadenceParser) ReturnStatement_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 4:
		return !p.lineTerminatorAhead()

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *CadenceParser) OrExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 5:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *CadenceParser) AndExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 6:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *CadenceParser) EqualityExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 7:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *CadenceParser) RelationalExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 8:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *CadenceParser) FailableDowncastingExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 9:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *CadenceParser) ConcatenatingExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 10:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *CadenceParser) AdditiveExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 11:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *CadenceParser) MultiplicativeExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 12:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *CadenceParser) Eos_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 13:
		return p.lineTerminatorAhead()

	case 14:
		return p.GetTokenStream().LT(1).GetText() == "}"

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
