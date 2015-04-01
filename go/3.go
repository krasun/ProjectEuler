package main

import (
	"fmt"
	"math"
)

func hasDivider(n int) bool {
	for divider := 2; divider <= int(math.Sqrt(float64(n))); divider++ {
		if n%divider == 0 {
			return true
		}
	}

	return false
}

func divisibleBy(n, divider int) bool {
	return n%divider == 0
}

func main() {
	n := 600851475143
	maxFactor := int(math.Sqrt(float64(n)))

	for {
		if divisibleBy(n, maxFactor) && !hasDivider(maxFactor) {
			break
		}

		maxFactor--
	}

	fmt.Printf("The largest prime factor of the number %d is %d.\n", n, maxFactor)
}
