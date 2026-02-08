package main

import (
	"fmt"
	"math"
)

func printPrimes(max int) {
	for n := 2; n <= max; n++ {
		isPrime := true
		sqrtN := int(math.Sqrt(float64(n)))
		for i := 2; i <= sqrtN; i++ {
			if n%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			fmt.Println(n)
		}
	}
}

func test(max int) {
	fmt.Printf("Primes up to %v:\n", max)
	printPrimes(max)
	fmt.Println("===============================================================")
}

func main() {
	test(10)
	test(20)
	test(30)
}
