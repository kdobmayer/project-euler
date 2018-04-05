package main

import (
	"fmt"

	"github.com/kdobmayer/project-euler/prime"
)

func main() {
	var n int
	for i, count := 1, 0; count <= 10001; i++ {
		if prime.Is(i) {
			n = i
			count++
		}
	}
	fmt.Println(n)
	// Output: 104743
}
