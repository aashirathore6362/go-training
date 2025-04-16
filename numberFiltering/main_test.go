package main

import (
	"fmt"
	"reflect"
	"slices"
	"testing"
)

// Story 1: Given a list of integers, write a program to return only the even numbers from this list.
func Test_numEven(t *testing.T) {
	inputs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	want := []int{2, 4, 6, 8, 10}
	got := filterNumbers(inputs, evenFn)
	if len(want) == 0 {
		fmt.Println("slice is empty:")
	}
	if !slices.Equal(want, got) {
		t.Errorf("Got %v, want %v", got, want)
	}
}

// Story 2: Given a list of integers, write a program to return only the odd numbers from this list.
func Test_numOdd(t *testing.T) {
	inputs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	want := []int{1, 3, 5, 7, 9}
	got := filterNumbers(inputs, oddFun)
	if len(want) == 0 {
		fmt.Println("slice is empty:")
	}
	if !slices.Equal(want, got) {
		t.Errorf("Got %v, want %v:", got, want)
	}

}

// Story 3: Given a list of integers, write a program to return only the prime numbers from this list.
func Test_numPrime(t *testing.T) {
	inputs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	want := []int{2, 3, 5, 7}
	got := filterNumbers(inputs, primeFun)
	if len(want) == 0 {
		fmt.Println("slice is empty:")
	}
	if !slices.Equal(want, got) {
		t.Errorf("Got %v, want %v", got, want)
	}
}

// Story 4: Given a list of integers, write a program to return only the odd prime numbers from this list.
func Test_numOddPrime(t *testing.T) {
	inputs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	want := []int{3, 5, 7}
	primes := filterNumbers(inputs, primeFun)
	got := filterNumbers(primes, oddFun)
	if len(want) == 0 {
		fmt.Println("slice is empty:")
	}
	if !slices.Equal(want, got) {
		t.Errorf("Got %v, want %v:", got, want)
	}

}

// Story 5: Given a list of integers, write a program to return only the even and multiples of 5 from this list.
func Test_getMultiples(t *testing.T) {
	inputs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	want := []int{10, 20}
	got := filterNumbers(inputs, multiples)
	if len(want) == 0 {
		fmt.Println("slice is empty:")
	}
	if !slices.Equal(want, got) {
		t.Errorf("Got %v, want %v:", got, want)
	}

}

// Story 6: Given a list of integers, write a program to return only the odd and multiples of 3 greater than 10 from this list.
func Test_greater(t *testing.T) {
	inputs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	want := []int{15}
	oddnumbers := filterNumbers(inputs, oddFun)
	got := filterNumbers(oddnumbers, getGreater)
	if len(want) == 0 {
		fmt.Println("slice is empty:")
	}
	if !slices.Equal(want, got) {
		t.Errorf("Got %v, want %v:", got, want)
	}
}

// Story 7: Given a list of integers, and a set of conditions (odd, even, greater than 5, multiple of 3, prime, and many more such custom conditions that may be dynamically defined by user), write a program to return only the integers from the given list that match ALL the conditions.
func Test_allConditions(t *testing.T) {
	tests := []struct {
		nums       []int
		conditions []Condition
		want       []int
	}{
		{
			nums:       []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
			conditions: []Condition{oddFun, greaterThan(5), multipleOf(3)},
			want:       []int{9, 15},
		},
		{
			nums:       []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
			conditions: []Condition{evenFn, lessThan(15), multipleOf(3)},
			want:       []int{6, 12},
		},
	}
	for i := 0; i < len(tests); i++ {
		tt := tests[i]
		got := filterAllConditions(tt.nums, tt.conditions...)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("Test case %d, want %v", got, tt.want)
		}
	}
}

// Story 8: Given a list of integers, and a set of conditions (odd, even, greater than 5, multiple of 3, prime, and many more such custom conditions that may be dynamically defined by user), write a program to return only the integers from the given list that match ANY of the conditions.
func Test_AtOneCondition(t *testing.T) {
	tests := []struct {
		nums       []int
		conditions []Condition
		want       []int
	}{
		{
			nums:       []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
			conditions: []Condition{primeFun, greaterThan(15), multipleOf(5)},
			want:       []int{2, 3, 5, 7, 10, 11, 13, 15, 16, 17, 18, 19, 20},
		},
		{
			nums:       []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
			conditions: []Condition{lessThan(6), multipleOf(3)},
			want:       []int{1, 2, 3, 4, 5, 6, 9, 12, 15, 18},
		},
	}
	for i := 0; i < len(tests); i++ {
		tt := tests[i]
		got := filterAnyConditions(tt.nums, tt.conditions...)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("Test case %d failed:\n Got: %v\n, want: %v", i+1, got, tt.want)
		}
	}
}
