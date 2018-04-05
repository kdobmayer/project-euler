package main

import (
	"fmt"
)

func main() {
	var sum int
	sieved := make(map[int]bool)
	for i := 2; i < 2000000; i++ {
		if !sieved[i] {
			sum += i
			for j := 2; j <= 2000000/i; j++ {
				sieved[i*j] = true
			}
		}
	}
	fmt.Println(sum)
	// Output: 142913828922
}
