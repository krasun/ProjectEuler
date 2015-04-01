package main

import (
	"fmt"
	"math"
)

func isPalindrome(number int) bool {
	sourceNumber, reverse := number, 0
	for number > 0 {
		residue := number % 10
		reverse = reverse*10 + residue
		number = number / 10
	}

	return reverse == sourceNumber
}

func findLargestPalindrome(powFactor int) int {
	maxFactor := int(math.Pow(10, float64(powFactor)) - 1)
	minFactor := int(math.Pow(10, float64(powFactor-1)))
	largestPalindrome := -1

	for x := maxFactor; x >= minFactor; x-- {
		for y := maxFactor; y >= minFactor; y-- {
			product := x * y
			if isPalindrome(product) && product >= largestPalindrome {
				largestPalindrome = product
			}
		}
	}

	return largestPalindrome
}

func main() {
	fmt.Printf("The largest palindrome made from the product of two 3-digit numbers is %d.\n", findLargestPalindrome(3))
}
