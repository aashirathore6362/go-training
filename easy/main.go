package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

// sum of numbers
func SumOfNumber(n int) {

	var sum int
	for n != 0 {
		ans := n % 10
		sum = sum + ans
		n = n / 10
	}
	fmt.Print(math.Abs(float64(sum)))
}
func ReverseDigits(n int) {
	var num int = 0
	for n > 0 {
		num = num*10 + n%10
		n = n / 10
	}
	fmt.Print(num)
}
func PowerOfNumber(x, y int) bool {
	var temp int
	temp = 1
	for temp < y {
		temp = temp * x
	}
	return temp == y
}
func RecursionOfPower(x, y float64) bool {
	if x == y {
		return true
	} else if x > y {
		return false
	} else {
		return RecursionOfPower(x, y/x)
	}
}
func takeInputs() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Type something:")
	scanner.Scan()
	input := scanner.Text()
	fmt.Printf("you typed: %q", input)
}
func main() {

	SumOfNumber(-534)
	ReverseDigits(123)
	fmt.Print(PowerOfNumber(5, 24))
	fmt.Print(RecursionOfPower(4, 17))
	takeInputs()
}
