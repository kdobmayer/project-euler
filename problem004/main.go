package main

import (
	"fmt"
	"strconv"
)

func main() {
	var top int
	for i := 100; i < 1000; i++ {
		for j := i; j < 1000; j++ {
			if n := i * j; isPalindrom(n) && n > top {
				top = n
			}
		}
	}
	fmt.Println(top)
	// Output: 906609
}

func isPalindrom(n int) bool {
	s := strconv.Itoa(n)
	for i := 0; i <= len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}
