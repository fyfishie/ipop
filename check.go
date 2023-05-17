/*
 * @Author: fyfishie
 * @Date: 2023-05-17:09
 * @LastEditors: fyfishie
 * @LastEditTime: 2023-05-17:09
 * @Description: :)
 * @email: muren.zhuang@outlook.com
 */
package ipop

import "strings"

// simple check
func IsIP(s string) bool {
	return strings.Contains(s, ".")
}

// check by regexp
func IsIPReg(s string) bool {
	match := isIPReg.FindStringSubmatch(s)
	return len(match) == 1
}

// 1.2.3.4/30 ?
func IsCIDR(s string) bool {
	return CIDRReg.MatchString(s)
}
