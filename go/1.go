package main

import "fmt"

func main() {
	multiples := make(map[int]bool)
	sum := 0
	for n, by3, by5 := 1, 3, 5; n < 1000; n++ {

		by3 = n * 3
		if by3 >= 1000 {
			break
		}

		_, ok := multiples[by3]
		if !ok {
			sum += by3
			multiples[by3] = true
		}

		by5 = n * 5
		if by5 >= 1000 {
			continue
		}

		_, ok = multiples[by5]
		if !ok {
			sum += by5
			multiples[by5] = true
		}
	}

	fmt.Printf("Sum of all the multiples of 3 or 5 below 1000 is %d.\n", sum)
}
