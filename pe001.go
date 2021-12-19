/*
ProjectEuler #001: Multiples of 3 and 5

If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.
Find the sum of all the multiples of 3 or 5 below 1000.
*/
package main

// PE001 calculates the result for the generalized problem, finding the sum of multiples of a, b less than n.
func PE001(a int, b int, n int) int {
	/* Observe that the sum we want is precisely
		S = (a + 2a + 3a + ... + xa) + (b + 2b + 3b + ... + yb) - (c + 2c + 3c + ... + zc),
	where c = lcm(a, b) and x, y, z satisfy xa < n <= (x+1)a and similar. This is because all multiples of c
	end up double-counted. But also notice that
		S = a(1 + 2 + ... + x) + b(1 + 2 + ... + y) - c(1 + 2 + ... + z),
	where each of these sums is precisely a triangular number, which has a well-known formula.
	*/
	c := LCM(a, b)
	
	x := (n - 1) / a
	y := (n - 1) / b
	z := (n - 1) / c

	return a * SumOfPowers(x, 1) + b * SumOfPowers(y, 1) - c * SumOfPowers(z, 1)

}