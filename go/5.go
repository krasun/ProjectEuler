package main

import (
	"fmt"
)

func generateDivisibleByAllNumbersUpTo(limit int) int {
	primes := findPrimes(limit)
	result := 1
	for _, prime := range primes {
		multiplier := prime
		for {
			if (prime * multiplier) > limit {
				break
			}

			prime *= multiplier
		}

		result *= prime
	}

	return result
}

func findPrimes(limit int) []int {
	// false value determines prime as key
	sieve := make([]bool, limit+1)

	// 2 is the smallest prime
	for p := 2; p <= limit; p++ {
		for i := 2; i*p <= limit; i++ {
			sieve[i*p] = true
		}
	}

	// filter primes
	primes := make([]int, 0)
	for n, isComposite := range sieve {
		if n >= 2 && !isComposite {
			primes = append(primes, n)
		}
	}

	return primes
}

func main() {
	fmt.Printf("The smallest positive number, evenly divisible by all of the numbers from 1 to 20, is %d.\n", generateDivisibleByAllNumbersUpTo(20))
}
