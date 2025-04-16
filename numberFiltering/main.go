package main

import (
	"math"
)

type Condition func(int) bool

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

func greaterThan(n int) Condition {
	return func(x int) bool {
		return x > n
	}
}

func lessThan(n int) Condition {
	return func(x int) bool {
		return x < n
	}
}

func multipleOf(n int) Condition {
	return func(x int) bool {
		return x%n == 0
	}
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

func filterAllConditions(input []int, conditions ...Condition) []int {
	var result []int
	for _, num := range input {
		match := true
		for _, cond := range conditions {
			if !cond(num) {
				match = false
				break
			}
		}
		if match {
			result = append(result, num)
		}
	}
	return result
}
func filterAnyConditions(nums []int, conditions ...Condition) []int {
	var result []int
	for _, num := range nums {
		for _, cond := range conditions {
			if cond(num) {
				result = append(result, num)
				break
			}
		}
	}
	return result
}
