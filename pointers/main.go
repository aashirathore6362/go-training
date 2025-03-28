package main

import "fmt"

//print 5 numbers in array, 1 to 5
//Append numbers in array using loops.

func printarray() {
	var arr1 = [10]int{1,2,3,3,4,5}
	// fmt.Print(arr1)
	for i := range arr1 {
		fmt.Printf("%d\n",i)
	}

	// arr2.Append
}

func modString(a string) {
	a = "modified"
}
func main() {
	a := "Help"

	fmt.Println("Before:", a)
	modString(a)
	fmt.Println("After:", &a)

	printarray()
}
