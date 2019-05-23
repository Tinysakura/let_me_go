package chapter4

import (
	"sort"
)

// 删除切片中的重复元素
func UniqueInts(original []int)(unique []int) {
	sort.Ints(original)

	// fmt.Println(len(unique))

	for _, element := range original {
		if len(unique) == 0 {
			unique = append(unique, element)
		}

		// fmt.Println(len(unique))
		if unique[len(unique) - 1] != element {
			unique = append(unique, element)
		}
	}

	return unique
}

// 多维切片降维
func Flatten(multiDimension [][]int)(uniDimensional []int) {
	for _, dimension := range multiDimension {
		for _, element := range dimension {
			uniDimensional = append(uniDimensional, element)
		}
	}

	return uniDimensional
}


