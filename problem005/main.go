package main

import (
	"fmt"

	"github.com/kdobmayer/project-euler/prime"
)

func main() {
	factors := make(prime.Factors)
	for i := 1; i <= 20; i++ {
		for prime, pow := range prime.Factorization(i) {
			if pow > factors[prime] {
				factors[prime] = pow
			}
		}
	}
	fmt.Println(factors.Number())
	// Output: 232792560
}
