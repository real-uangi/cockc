// Package convert
// @author uangi
// @date 2023/9/4 16:15
package convert

import "encoding/json"

func ToJsonBytes(v interface{}) []byte {
	bs, err := json.Marshal(v)
	if err != nil {
		logger.Panic(err)
	}
	return bs
}
