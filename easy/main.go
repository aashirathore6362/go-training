package main

import (
	"fmt"
	"math"
)

func SumOfNumber(n int) {

	var sum int
	for n!=0 {
		ans:= n % 10
		sum = sum + ans
		n=n/10
	}
	fmt.Print(math.Abs(float64(sum)))
}
func main() {

	SumOfNumber(-534)

}
