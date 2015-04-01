package main 

import (
	"fmt"
	"math"
)

func reverse(n int) int {
	divider := 10	
		
	result := 0 

	result += n % divider 
	n /= divider

	// divider *= 10

	fmt.Println(result)
	fmt.Println(n)
	fmt.Println("-----")

	result += (n % divider) * divider
	n /= divider


	fmt.Println(result)
	fmt.Println(n)

	// result += (result % divider) * divider

	// result += n % 1000	

	return result
}

func isPalindrome(n int) bool {
	return n - reverse(n) == 0
}

func findLargestPalindrome(digitNum int) int {
	limit := int(math.Pow(10, float64(digitNum)) - 1);
	for x := limit; x >= 0; x-- {
		for y := limit; y >= 0; y-- {
			if (isPalindrome(x * y)) {
				return x * y
			}
		}
	}

	return 0
}

func main() {
	// fmt.Println(reverse(120))
	reverse(123)
	// fmts.Println(reverse(1))

	// fmt.Printf("The largest palindrome made from the product of two 3-digit numbers is %d.\n", findLargestPalindrome(3))
}