package main

import (
	"fmt"
)

func main() {

	difference, sum := 0, 0
	for i := 1; i <= 100; i++ {
		difference += i * i
		sum += i
	}
	difference = (sum * sum) - difference

	fmt.Printf("The difference between the sum of the squares of the first one hundred natural numbers and the square of the sum is %d.\n", difference)
}
