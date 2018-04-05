package main

import (
	"fmt"

	"github.com/kdobmayer/project-euler/prime"
)

func main() {
	var top int
	for prime := range prime.Factorization(600851475143) {
		if prime > top {
			top = prime
		}
	}
	fmt.Println(top)
	// Output: 6857
}
