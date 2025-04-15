package main

import (
	"fmt"
	"slices"
	"testing"
)

// Story 1: Given a list of integers, write a program to return only the even numbers from this list.
func Test_numEven(t *testing.T) {
	want := []int{2, 4, 6, 8, 10}
	got := filterNumbers(want, evenFn)

	if len(want) == 0 {
		fmt.Println("slice is empty:")
	}
	if !slices.Equal(want, got) {
		t.Errorf("Got %v, want %v", got, want)
	}
}

// Story 2: Given a list of integers, write a program to return only the odd numbers from this list.
func Test_numOdd(t *testing.T) {
	want := []int{1, 3, 5, 7, 9}
	got := filterNumbers(want, oddFun)
	if len(want) == 0 {
		fmt.Println("slice is empty:")
	}
	if !slices.Equal(want, got) {
		t.Errorf("Got %v, want %v:", got, want)
	}

}

// Story 3: Given a list of integers, write a program to return only the prime numbers from this list.
func Test_numPrime(t *testing.T) {
	want := []int{2, 3, 5, 7}
	got := filterNumbers(want, primeFun)
	if len(want) == 0 {
		fmt.Println("slice is empty:")
	}
	if !slices.Equal(want, got) {
		t.Errorf("Got %v, want %v", got, want)
	}
}

// Story 4: Given a list of integers, write a program to return only the odd prime numbers from this list.
func Test_numOddPrime(t *testing.T) {
	want := []int{3, 5, 7}
	got := filterNumbers(want, oddFun)
	if len(want) == 0 {
		fmt.Println("slice is empty:")
	}
	if !slices.Equal(want, got) {
		t.Errorf("Got %v, want %v:", got, want)
	}

}

// Story 5: Given a list of integers, write a program to return only the even and multiples of 5 from this list.
func Test_getMultiples(t *testing.T) {
	want := []int{10, 20}
	got := filterNumbers(want, multiples)
	if len(want) == 0 {
		fmt.Println("slice is empty:")
	}
	if !slices.Equal(want, got) {
		t.Errorf("Got %v, want %v:", got, want)
	}

}

// Story 6: Given a list of integers, write a program to return only the odd and multiples of 3 greater than 10 from this list.
func Test_greater(t *testing.T) {
	want := []int{15}
	oddnumbers := filterNumbers(want, oddFun)
	got := filterNumbers(oddnumbers, getGreater)
	if len(want) == 0 {
		fmt.Println("slice is empty:")
	}
	if !slices.Equal(want, got) {
		t.Errorf("Got %v, want %v:", got, want)
	}
}

//Story 7: Given a list of integers, and a set of conditions (odd, even, greater than 5, multiple of 3, prime, and many more such custom conditions that may be dynamically defined by user), write a program to return only the integers from the given list that match ALL the conditions.

func TestFilterNumbers_AllConditions(t *testing.T) {
	want := []int{9, 15}
	got := filterNumbers(want, getGreater5, oddFun)
	if len(want) == 0 {
		fmt.Println("Conditions specified using a set of function:")
	}
	if !slices.Equal(want, got) {
		t.Errorf("Got %v, want %v:", got, want)
	}
}

// Given a list of integers, and a set of conditions (odd, even, greater than 5, multiple of 3, prime, and many more such custom conditions that may be dynamically defined by user), write a program to return only the integers from the given list that match ANY of the conditions.
func TestFilterNumbers_AtOneCondition(t *testing.T) {
	want := []int{19, 5}
	got := filterByAynyCondition(want, primeFun, getGreater5, getGreater)
	if len(want) == 0 {
		fmt.Println("Conditions specified using a set of function:")
	}
	if !slices.Equal(want, got) {
		t.Errorf("Got %v, want %v:", got, want)
	}
}
