package chapter1

import (
	"fmt"
	"os"
)

var bigdigits = [][]string{
	{"*****","*   *","*   *","*   *","*   *","*   *","*****"},
	{"  *  "," **  ","* *  ","  *  ","  *  ","  *  ","*****"},
	{"*****","    *","    *","*****","*    ","*    ","*****"},
	{"*****","    *","    *","*****","    *","    *","*****"},
	{"  *  "," **  ","* *  ","*****","  *  ","  *  ","  *  "},
	{"*****","*    ","*    ","*****","    *","    *","*****"},
	{"*****","*    ","*    ","*****","*   *","*   *","*****"},
	{"*****","    *","    *","   * ","  *  "," *   ","*    "},
	{"*****","*   *","*   *","*****","*   *","*   *","*****"},
	{"*****","*   *","*   *","*****","    *","    *","*****"},
}

func DrawBigDigits(args1 string, args2 string) {
	if args1 == "--help" {
		fmt.Println("usage:bigdigits [-b|--bar] <whole-number>")
		fmt.Println("-b --bar draw an underbar and an overbar")
		os.Exit(0)
	}

	if args1 == "-b" {
		for row := range bigdigits[0] {
			line := ""
			for column := range args2 {
				digit := args2[column] - '0'
				line += bigdigits[digit][row] + " "
			}
			fmt.Println(line)
		}
	}

	if args1 == "-bar" {
		fmt.Println("=================")
		for row := range bigdigits[0] {
			line := ""
			for column := range args2 {
				digit := args2[column] - '0'
				line += bigdigits[digit][row] + " "
			}
			fmt.Println(line)
		}
		fmt.Println("==================")
	}
}
