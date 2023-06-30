// Package character @author uangi 2023-05
package character

// AppendAll 累加字符串
func AppendAll(ss ...string) string {
	var result string
	for _, s := range ss {
		result = result + s
	}
	return result
}

func IsNotBlank(s string) bool {
	return !IsBlank(s)
}

func IsBlank(s string) bool {
	return &s == nil || len(s) == 0
}
