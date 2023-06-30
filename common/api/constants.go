// Package api @author uangi 2023-05
package api

type Method string

const (
	POST   Method = "Post"
	GET    Method = "Get"
	DELETE Method = "Delete"
	PUT    Method = "Put"
	PATCH  Method = "Patch"
)

func (method Method) String() string {
	return string(method)
}
