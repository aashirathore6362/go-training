package main

import (
	"math"
)

var evenFn = func(x int) bool {
	return x%2 == 0
}

var oldFun = func(x int) bool {
	return x%2 != 0
}

var primeFun = func(x int) bool {
	if x < 2 {
		return false
	}
	for i := 2; float64(i) <= math.Sqrt(float64(x)); i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

// var oldPrimeFun = func(x int) bool {
// 	filterNumbers(numbers, oldFun)

// }

func filterNumbers(input []int, filterFn func(int) bool) []int {
	var result []int
	for _, num := range input {
		if filterFn(num) {
			result = append(result, num)
		}
	}
	return result
}
