package main

import "fmt"

func fizzbuzz() {
	for i := 1; i < 100; i++ {
		isMultipleThree := i%3 == 0
		isMultipleFive := i%5 == 0

		if isMultipleThree && !isMultipleFive {
			fmt.Println("fizz")
		} else if isMultipleFive && !isMultipleThree {
			fmt.Println("buzz")
		} else if isMultipleFive && isMultipleThree {
			fmt.Println("fizzbuzz")
		} else {
			fmt.Println(i)
		}
	}
}

// don't touch below this line

func main() {
	fizzbuzz()
}
