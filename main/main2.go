package main

import "let_me_go/network"

func main() {
	println("open client")
	network.OpenClient("", 7777)
}
