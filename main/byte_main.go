package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func main() {
	fmt.Println(IntToBytes(500))
	fmt.Println("======大端序========")
	bytes := make([]byte, 16)
	binary.BigEndian.PutUint32(bytes, 500)
	fmt.Println(bytes)
	fmt.Printf("value:%d \n", binary.BigEndian.Uint32(bytes[:4]))
	binary.BigEndian.PutUint32(bytes[4:], 500)
	fmt.Println(bytes)
	fmt.Printf("value:%d \n", binary.BigEndian.Uint32(bytes[:4]))

	fmt.Println("======小端序========")

	bytes2 := make([]byte, 16)
	binary.LittleEndian.PutUint32(bytes2, 500)
	fmt.Println(bytes2)
	fmt.Printf("value:%d\n", binary.LittleEndian.Uint64(bytes2))
	binary.LittleEndian.PutUint32(bytes2[4:], 500)
	fmt.Println(bytes2)
}

func IntToBytes(n int) []byte {
	x := uint32(n)

	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}