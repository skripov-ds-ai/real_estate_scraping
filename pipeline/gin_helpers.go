package pipeline

import "fmt"

func MakeGinAddr(addr, port string) string {
	return fmt.Sprintf("%s:%s", addr, port)
}
