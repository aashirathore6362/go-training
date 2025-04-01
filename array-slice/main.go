package main

import (
	"fmt"
	"slices"
)

//print 5 numbers in array, 1 to 5
//Append numbers in array using loops.

func printArray() {
	var arr1 = [10]int{1, 2, 3, 3, 4, 5}
	// fmt.Print(arr1)
	for i := range arr1 {
		fmt.Printf("%d\n", i)
	}
	// arr2 := [...]int{1, 2, 3}
	// arr2 = append([3]int(arr1))

	months := make([]string, 0, 1)
	fmt.Printf("%p\n", months)

	months2 := make([]string, 0, 1)
	fmt.Printf("%p\n", &months2)
}

// write a function to reverse the slice
func reverseSlice() {
	s := []int{0, 1, 2}
	slices.Reverse(s)
	fmt.Println(s)
}

// append the slice

func appendSlice() {
	var s []int
	fmt.Printf("len=%d cap=%d \n", len(s), cap(s))

	//append on 0 index
	s = append(s, 0)
	fmt.Printf("len=%d cap=%d \n", len(s), cap(s))

	//append on 1 index
	s = append(s, 3)
	fmt.Printf("len=%d cap=%d \n", len(s), cap(s))

}

// map data-structure
func mapPrint() {
	ages := map[string]int{
		"alice": 30,
	}
	fmt.Println(ages["alice"])
}

// function declaration
func add(x int, y int) int { return x + y }

func sub(x, y int) (z int) { z = x - y; return }

func first(x int, _ int) int { return x }

func zero(int, int) int { return 0 }

func main() {
	// a := "Help"

	// fmt.Println("Before:", a)
	// modString(a)
	// fmt.Println("After:", &a)

	// printArray()
	reverseSlice()
	appendSlice()
	mapPrint()

	fmt.Printf("%d\n", add(1,2)) // "func(int, int) int"

	fmt.Printf("%T\n", sub) // "func(int, int) int"

	fmt.Printf("%d\n", first(1,1)) // "func(int, int) int"

	fmt.Printf("%T\n", zero) // "func(int, int) int"
}
