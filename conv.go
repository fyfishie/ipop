/*
 * @Author: fyfishie
 * @Date: 2023-05-17:09
 * @LastEditors: fyfishie
 * @LastEditTime: 2023-05-17:09
 * @Description: :)
 * @email: muren.zhuang@outlook.com
 */
package ipop

import (
	"strings"

	"github.com/fyfishie/ipop/lib"
)

// 16843009 -> 1.1.1.1
func IntToString(i int) string {
	a := i & 255
	b := i & 65280 >> 8
	c := i & 16711680 >> 16
	d := i & 4278190080 >> 24
	return lib.IntToStringMap[d] + "." + lib.IntToStringMap[c] + "." + lib.IntToStringMap[b] + "." + lib.IntToStringMap[a]
}

// 1.1.1.1 -> 16843009
func String2Int(s string) int {
	sArray := strings.Split(s, ".")
	return lib.String2IntMap[sArray[3]] + lib.String2IntMap[sArray[2]]*256 + lib.String2IntMap[sArray[1]]*65536 + lib.String2IntMap[sArray[0]]*16777216
}

func String2Uint32(s string) uint32 {
	ss := strings.Split(s, ".")
	return lib.String2Uint32Map[ss[3]] + lib.String2Uint32Map[ss[2]]*256 + lib.String2Uint32Map[ss[1]]*65536 + lib.String2Uint32Map[ss[0]]*16777216
}
func Uint322String(u uint32) string {
	a := u & 255
	b := u & 65280 >> 8
	c := u & 16711680 >> 16
	d := u & 4278190080 >> 24
	return lib.Uint32ToStringMap[d] + "." + lib.Uint32ToStringMap[c] + "." + lib.Uint32ToStringMap[b] + "." + lib.Uint32ToStringMap[a]
}

// attention!!!this function doesn't provide format check!!!
func Strings2Ints(ss []string) []int {
	res := []int{}
	for _, s := range ss {
		sArray := strings.Split(s, ".")
		res = append(res, lib.String2IntMap[sArray[3]]+lib.String2IntMap[sArray[2]]*256+lib.String2IntMap[sArray[1]]*65536+lib.String2IntMap[sArray[0]]*16777216)
	}
	return res
}

func Ints2Strings(is []int) []string {
	res := []string{}
	for _, i := range is {
		res = append(res,
			lib.IntToStringMap[i&4278190080>>24]+
				"."+lib.IntToStringMap[i&16711680>>16]+
				"."+lib.IntToStringMap[i&65280>>8]+
				"."+lib.IntToStringMap[i&255])
	}
	return res
}
