package network

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"time"
)

func OpenClient(host string, port int) {
	ip := ":7777"
	tcpAddr, _ := net.ResolveTCPAddr("tcp4", ip)
	conn, e := net.DialTCP("tcp", nil, tcpAddr)
	if e != nil {
		fmt.Println(e)
	}
	_, _ = conn.Write([]byte("request request"))
	time.Sleep(time.Duration(2) * time.Second)
	bytes, _ := ioutil.ReadAll(conn)
	fmt.Println(string(bytes))
}

func OpenServer(port int) {
	serverIp := ":7777"
	tcpAddr, _ := net.ResolveTCPAddr("tcp4", serverIp)
	listener, _ := net.ListenTCP("tcp", tcpAddr)
	for {
		conn, e := listener.Accept()

		if e != nil {
			continue
		}

		bytes, e := ioutil.ReadAll(conn)
		if e != nil {
			log.Fatal(string(bytes))
		}
		fmt.Println(string(bytes))

		_, _ = conn.Write([]byte("response response"))
		_ = conn.Close()
	}
}