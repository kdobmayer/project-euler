package main

import "fmt"

func main() {
	var sum int
	for i, j := 1, 1; j < 4000000; i, j = j, i+j {
		if j%2 == 0 {
			sum += j
		}
	}
	fmt.Println(sum)
	// Output: 4613732
}
