package main

import (
	"fmt"
	"let_me_go/chapter4"
)

func main() {
	// helloworld
	// helloworld.SayHelloWorld()

	// chapter1
	//if os.Args[1] == "--help" {
	//	chapter1.DrawBigDigits(os.Args[1], "")
	//} else {
	//	chapter1.DrawBigDigits(os.Args[1], os.Args[2])
	//}

	// chapter2
	//s := os.Args[1]
	//split := strings.Split(s, ",")
	//var numbers [] float64
	//
	//for _,value := range split{
	//	x, _ := strconv.ParseFloat(value, 64)
	//	numbers = append(numbers, x)
	//}
	//
	//var stat chapter2.Statistics
	//stat = chapter2.Cal(numbers)
	//fmt.Println(fmt.Sprintf("numbers:%f", stat.Numbers))
	//fmt.Println(fmt.Sprintf("median:%f", stat.Median))
	//fmt.Println(fmt.Sprintf("mean:%f", stat.Mean))

	// chapter3
	//s := os.Args[1]
	//songs := chapter3.ReadPlsPlaylist(s)
	//chapter3.WriteM3uPlaylist(songs)

	// chapter4
	//unique := chapter4.UniqueInts([]int{1, 1, 2, 2, 4, 5})
	//
	//for _, number := range unique {
	//	fmt.Println(number)
	//}

	uniDimensional := chapter4.Flatten([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})

	for _, number := range uniDimensional {
		fmt.Println(number)
	}

}
