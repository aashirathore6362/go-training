package main

import (
	"math"
)

// takes int and return bool.
var evenFn = func(x int) bool {
	return x%2 == 0
}

var oddFun = func(x int) bool {
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

var multiples = func(x int) bool {
	if x%2 == 0 {
		if x%5 == 0 {
			return true
		}
	}
	return false
}
var getGreater = func(x int) bool {
	return x%3 == 0 && x > 10
}
var getGreater5 = func(x int) bool {
	return x%3 == 0 && x > 5
}

func filterNumbers(input []int, filterFn ...func(int) bool) []int {
	var result []int
	for _, num := range input {
		matchALL := true
		for _, fn := range filterFn {
			if !fn(num) {
				matchALL = false
				break
			}
		}
		if matchALL {
			result = append(result, num)
		}
	}
	return result

}

func filterByAynyCondition(input []int, filterFn ...func(int) bool) []int {
	var result []int
	for _, num := range input {
		for _, fn := range filterFn {
			if fn(num) {
				result = append(result, num)
				break
			}
		}
	}
	return result
}