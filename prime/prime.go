package prime

// Factors are the prime factors of a number, where the keys are the primes and
// the values are the power of that primes.
type Factors map[int]int

// Factorization provides the prime factors of a number.
func Factorization(n int) Factors {
	fs := make(Factors)
	i := 2
	for i <= n {
		if n%i == 0 {
			fs[i]++
			n /= i
			continue
		}
		i++
	}
	return fs
}

// Number calculates the number from its prime factors.
func (f Factors) Number() int {
	n := 1
	for prime, pow := range f {
		for i := 0; i < pow; i++ {
			n *= prime
		}
	}
	return n
}

// Is tells whether a number is prime.
func Is(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
