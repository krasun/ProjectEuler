package main 

import (
	"fmt"
)

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Println(2520 % i)
	}

	fmt.Printf("The smallest positive number, evenly divisible by all of the numbers from 1 to 20, is %d.")
}