package main

import "fmt"

func fibonacciNumbers(satisfied func(int) bool, endCondition func(int) bool) chan int {
	c := make(chan int)

	go func() {
		fib0, fib1 := 1, 1

		for endCondition(fib0) {
			if satisfied(fib0) {
				c <- fib0
			}
			fib0, fib1 = fib1, fib0+fib1
		}

		close(c)
	}()

	return c
}

func main() {

	sum := 0
	for fib := range fibonacciNumbers(func(fib int) bool { return (fib % 2) == 0 }, func(fib int) bool { return fib < 4000000 }) {
		fmt.Println(fib)
		sum += fib
	}

	fmt.Printf("The sum of the even-valued terms in the Fibonacci sequence, whose values do not exceed four million, is %d.\n", sum)
}
