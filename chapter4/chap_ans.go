package chapter4

import (
	"math"
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
func Flatten(multiDimension [][]int)(uniDimension []int) {
	for _, dimension := range multiDimension {
		for _, element := range dimension {
			uniDimension = append(uniDimension, element)
		}
	}

	return uniDimension
}

// 一维切片升维
func Make2D(uniDimension []int, dimensionNumber int)(multiDimension [][]int) {
	length := int(math.Floor(float64(len(uniDimension)) / float64(dimensionNumber) + 0.5))

	for i := 0; i < dimensionNumber; i ++ {
		if lack := (i + 1) * length - len(uniDimension); lack > 0 {
			multiDimension = append(multiDimension, uniDimension[i * length:])

			for i = 0; i < lack; i++ {
				multiDimension[len(multiDimension) - 1] = append(multiDimension[len(multiDimension)-1], 0)
			}
		} else {
			multiDimension = append(multiDimension, uniDimension[i * length:(i + 1) * length])
		}
	}

	return multiDimension
}


