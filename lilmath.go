package main

import (
	"bufio"
	"errors"
	"log"
	"math"
	"os"
	"strconv"
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

func MinMax(arr []int) (int, int, error) {
	if len(arr) == 0 {
		return 0, 0, errors.New("cannot find min/max of an empty slice")
	}

	min, max := arr[0], arr[0]
	for _, val := range arr {
		if val < min {
			min = val
		}
		if val > max {
			max = val
		}
	}

	return min, max, nil
}


func PrimesLessThan(limit int) []int {
	file, err := os.Open("primes.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	primes := []int{};
	for scanner.Scan() {
		p, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		if p >= limit {
			return primes
		}

		primes = append(primes, p)
	}

	log.Printf("PrimesLessThan was requested primes up to %d, but the primes file was exhausted before that limit", limit)
	return primes
}

// PrimeFactorization returns the prime factorization of n, as a map[int]int
// Example: 28 --> {2:2, 7:1}
func PrimeFactorization(n int) map[int] int {
	pf := make(map[int] int)

	primes := PrimesLessThan(100000)
	for _, p := range primes {
		if n == 1 {
			return pf
		}

		for n % p == 0 {
			pf[p] += 1
			n /= p
		}
	}

	return pf
}