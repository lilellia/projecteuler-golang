/*
ProjectEuler #003: Largest prime factor

The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?
*/
package main
import "os"

// PE003 calculates the largest prime factor of n
func PE003(n int) int {
	primeFactors, _ := IntIntMapKV(PrimeFactorization(n))
	_, max, err := MinMax(primeFactors)

	if err != nil {
		os.Exit(1)
	}
	return max
}