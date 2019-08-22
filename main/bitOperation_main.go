package main

import "fmt"

func main() {
	a := 0x43
	fmt.Printf("%b\n", a)
	fmt.Printf("%b\n", ^a)

	b := 7
	fmt.Printf("%b\n", b)
	fmt.Printf("%b\n", b << 2)
}
