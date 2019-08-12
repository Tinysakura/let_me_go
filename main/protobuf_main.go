package main

import (
	p "../protobuf"
	"fmt"
)
import  "github.com/golang/protobuf/proto"

func main() {
	message := "bronya can not say"
	userInfo := p.UserInfo{Message: message, Length: int32(len(message))}

	marshalData, _ := proto.Marshal(&userInfo)
	fmt.Println(marshalData)

	var unMarshalResult p.UserInfo
	e := proto.Unmarshal(marshalData, &unMarshalResult)
	if e != nil {
		fmt.Println(e)
	}
	fmt.Printf("message:%s, length:%d \n", unMarshalResult.Message, unMarshalResult.Length)
}