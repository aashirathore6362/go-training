package main

import "fmt"

func modString(a string) {
	a = "modified"
}
func main() {
	a:= "help"

	fmt.Println("Before:",a)
	modString(a)
	fmt.Println("After:",a)
}