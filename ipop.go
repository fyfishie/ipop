/*
* @Author: fyfishie
* @Date: 2023-03-15:08
 * @LastEditors: fyfishie
 * @LastEditTime: 2023-04-03:09
* @Description: :)
* @email: muren.zhuang@outlook.com
*/
package ipop

import (
	"regexp"
	"strings"

	"github.com/fyfishie/ipop/lib"
)

var isIPReg = regexp.MustCompile(`\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}`)
var CIDRReg = regexp.MustCompile(`\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}/\d{1,3}`)

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

// simple check
func IsIP(s string) bool {
	return strings.Contains(s, ".")
}

// check by regexp
func IsIPReg(s string) bool {
	match := isIPReg.FindStringSubmatch(s)
	return len(match) == 1
}

// if length of iplist is not greater than 2, this function will return -1
func NetLength(ipList []string) (subnetLength int) {
	length := len(ipList)
	if length < 2 {
		return -1
	}
	ipListInt := []int{}
	for _, ip := range ipList {
		ipListInt = append(ipListInt, String2Int(ip))
	}
	for i := 31; i > -1; i-- {
		mask := lib.NetLengthMask[i]
		maskRes := ipListInt[0] & mask
		j := 1
		for j = 1; j < length; j++ {
			if maskRes != ipListInt[j]&mask {
				break
			}
		}
		if j == length {
			return i
		}
	}
	return -1
}

// 1.2.3.4/30 ?
func IsCIDR(s string) bool {
	return CIDRReg.MatchString(s)
}
