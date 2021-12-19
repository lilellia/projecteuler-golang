package main

import (
	"math"
)

// GCD returns the greatest common divisor of two numbers.
func GCD(a int, b int) int {
	for b != 0 {
		a, b = b, a % b
	}
	return a
}

// LCM returns the least common multiple of two numbers.
func LCM(a int, b int) int {
	return a * b / GCD(a, b)
}

// SumOfPowers returns the sum 1^p + 2^p + ... + n^p
func SumOfPowers(n int, p int) int {
	switch p {
	case 0:
		// 1^0 + 2^0 + ... + n^0 = 1 + 1 + ... + 1
		return n
	case 1:
		// 1 + 2 + 3 + ... = n(n+1)/2
		return n * (n + 1) / 2
	case 2:
		// 1^2 + 2^2 + ... + n^2 = n(n+1)(2n+1)/6
		return n * (n + 1) * (2 * n + 1) / 6
	case 3:
		// 1^3 + 2^3 + ... + n^3 = (1 + 2 + ...)^2
		s := SumOfPowers(n, 2)
		return s * s
	default:
		// I don't have a precomputed formula for other exponents,
		// so we just compute manually
		result := 1.0
		for i := 2; i <= n; i++ {
			result += math.Pow(float64(i), float64(p))
		}
		return int(result)
	}
	
}