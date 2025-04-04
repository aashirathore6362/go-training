package main

import (
	"fmt"
	"slices"
	"testing"
)

func Test_storyEven(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6}
	want := []int{2,4}
	got := filterNumbers(numbers, evenFn)
	if !slices.Equal(want,got) {
		t.Errorf("Got %v, want %v", got, want)
	}
	// fmt.Println("even number:", got)
}

func Test_storyOdd(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6}
	oldnum := filterNumbers(numbers, oldFun)
	fmt.Println("old numbers:", oldnum)

}

func Test_storyPrime(t *testing.T) {
	number := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	primnum := filterNumbers(number, primeFun)
	fmt.Println("Prime numbers", primnum)
}

// func Test_storyOddPrime(t *testing.T) {
// 	number := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
// 	oldprimenum := filterNumbers(number, oldPrimeFun)
// 	fmt.Println("Prime numbers", oldprimenum)

// }
