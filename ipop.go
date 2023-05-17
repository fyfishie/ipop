/*
* @Author: fyfishie
* @Date: 2023-03-15:08
 * @LastEditors: fyfishie
 * @LastEditTime: 2023-04-26:14
* @Description: :)
* @email: muren.zhuang@outlook.com
*/
package ipop

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/fyfishie/ipop/lib"
)

var isIPReg = regexp.MustCompile(`\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}`)
var CIDRReg = regexp.MustCompile(`\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}/\d{1,3}`)

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

func CIDRRange(cidr string) (start, end int) {
	ss := strings.Split(cidr, "/")
	net := ss[0]
	sub := ss[1]
	intnet := String2Int(net)
	intsub, _ := strconv.Atoi(sub)
	start = intnet & (0xffffffff << (32 - intsub))
	end = start + (1 << (32 - intsub)) - 1
	return start, end
}
