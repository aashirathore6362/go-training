package main

import (
	"fmt"
	"math"
)

// Check even or odd program
func even(x int) bool {
	if x%2 == 0 {
		fmt.Println("Given number is even: ")
		return true
	}
	// no var style
	return false
}

// Multiplication Tables
func tables(n int) int {
	var i, ans int
	for i = 1; i <= 10; i++ {
		ans = n * i
		fmt.Printf("%v * %v = %v \n", n, i, ans)
	}
	return ans
}

// Sum of naturals number
func findsum(n int) {
	var i, ans int
	for i = 1; i <= n; i++ {
		ans = ans + i
	}
	fmt.Printf("%v \n", ans)
}

// swap two number, using 3rd variable
func swap(a int, b int) {
	var c int
	fmt.Printf("Before swap a,b : %v %v \n", a, b)
	c = a
	a = b
	b = c
	fmt.Printf("After swap a,b : %v %v \n", a, b)
}

// Find the number closest to n and divisible by m
func divisibles(n, m int) {
	// var a,b int
	a, b := n, n
	for a%m != 0 && b%m != 0 {
		a = a - 1
		b = b + 1
		// fmt.Print(a, b)
	}

	if a%m == 0 && b%m == 0 {

		if math.Abs(float64(a)) < math.Abs(float64(b)) {
			fmt.Print(b)
		} else if math.Abs(float64(a)) > math.Abs(float64(b)) {
			fmt.Print(a)
		}
		// fmt.Printf("%v",a)
	} else if a%m == 0 {
		fmt.Print(a)
	} else if b%m == 0 {
		fmt.Print(b)
	} else {
		fmt.Print(n)
	}
}

// You are given a cubic dice with 6 faces. All the individual faces have a number printed on them.
func oppositeface(n int) {
	if n == 1 || n >= 7 {
		fmt.Printf("Invalid face")
	} else {
		ans := 7 - n
		fmt.Print(ans)
	}
}

// Nth term of AP from First Two Terms
func nthterms(a1, a2, n int) {
	Terms := a1 + (n-1)*(a2-a1)
	print(Terms)
}

func main() {
	fmt.Println(even(2))
	tables(4)
	findsum(5)
	swap(2, 4)
	divisibles(-15, 6)
	oppositeface(2)
	nthterms(2, 3, 4)
}
