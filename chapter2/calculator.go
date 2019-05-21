package chapter2

import "sort"

type Statistics struct {
	Numbers []float64
	Mean    float64 // 平均数
	Median  float64 // 中位数
}

func Cal(numbers []float64) (statistics Statistics) {
	statistics.Numbers = numbers
	statistics.Mean = mean(numbers)
	statistics.Median = median(numbers)

	return statistics
}

func mean(numbers []float64) (mean float64) {
	for _, x := range numbers {// _用来忽略index
		mean += x
	}

	mean /= float64(len(numbers))

	return mean
}

func median(numbers []float64) (median float64) {
	sort.Float64s(numbers) //有小到大排序
	middleIndex := len(numbers) / 2

	if len(numbers) % 2 != 0 {
		return numbers[int(middleIndex)]
	} else {
		return (numbers[int(middleIndex) - 1] + numbers[int(middleIndex)])/2
	}

}
