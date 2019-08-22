package main

import (
	"fmt"
	"let_me_go/network"
	"unsafe"
)

// 测试unsafe包
func main() {
	ip := new(network.IP)
	ip.SetHost("127.0.0.1")
	ip.SetPort(8080)

	fmt.Println(*ip)

	// 测试unsafe 访问/修改 结构体私有属性
	i := (*string)(unsafe.Pointer(ip))
	fmt.Println(*i)
	*i = "192.168.1.1"
	fmt.Println(*ip)
}
