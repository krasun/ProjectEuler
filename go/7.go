package main

import (
	"fmt"
	"math"
)

func isPrime(number int) bool {
	for divider := 2; divider <= int(math.Sqrt(float64(number))+1); divider++ {
		if number%divider == 0 {
			return false
		}
	}

	return true
}

func main() {
	var position, prime int = 1, 0
	for number := 2; position != 10001; number++ {
		if isPrime(number) {
			position++
			prime = number
		}
	}

	fmt.Printf("The 10 001st prime number is %d.\n", prime)
}
