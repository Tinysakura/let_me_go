package main

import (
	"fmt"
	"let_me_go/network"
	"reflect"
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

	// 进行指针偏移
	// 使用unsafe.Sizeof计算结构体每个成员变量所占的地址位
	t := (*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(ip)) + unsafe.Sizeof(ip.GetHost())))
	fmt.Println(*t)
	*t = 80
	fmt.Println(*ip)

	// 使用unsafe 进行string和slice的相互转换
	s := "hello world"
	bytes := string2Bytes(s)
	for _, value := range bytes {
		fmt.Println(value)
	}

	fmt.Println(bytes2String(bytes))
}

func string2Bytes(s string) []byte {
	stringHeader := (*reflect.StringHeader)(unsafe.Pointer(&s))

	sliceHeader := reflect.SliceHeader{
		Data: stringHeader.Data,
		Len:  stringHeader.Len,
		Cap:  stringHeader.Len,
	}

	return *(*[]byte)(unsafe.Pointer(&sliceHeader))
}

func bytes2String(bytes []byte) string {
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&bytes))

	stringHeader := reflect.StringHeader{
		Data: sliceHeader.Data,
		Len:  sliceHeader.Len,
	}

	return *(*string)(unsafe.Pointer(&stringHeader))
}
