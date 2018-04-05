package main

import "fmt"

func main() {
	fmt.Println(squareOfSum(100) - sumOfSquare(100))
	// Output: 25164150
}

func sumOfSquare(n int) int {
	return (2*n + 1) * (n + 1) * n / 6
}

func squareOfSum(n int) int {
	return n * n * (n + 1) * (n + 1) / 4
}
