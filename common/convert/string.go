// Package convert
// @author uangi
// @date 2023/7/3 11:09
package convert

import "strconv"

func IntToString(i int) string {
	return strconv.Itoa(i)
}

func Int64ToString(i int64) string {
	return strconv.FormatInt(i, 10)
}
