// Package convert
// @author uangi
// @date 2023/7/3 11:03
package convert

import "fmt"

func AnyToByte(i interface{}) []byte {
	return []byte(fmt.Sprintf("%v", i))
}
