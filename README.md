# ProjectEuler-Golang

Me going through the ProjectEuler problems as I learn Golang.

## Executing

Run a single problem

```
# run a single problem
$ go run . 2
2021/12/19 01:56:48 PE#002 [Elapsed:     0.000 ms] Answer: 4613732

# run multiple problems
$ go run . 1 2
2021/12/19 01:56:48 PE#001 [Elapsed:     0.001 ms] Answer: 233168
2021/12/19 01:56:48 PE#002 [Elapsed:     0.000 ms] Answer: 4613732

# run all benchmarks
$ go test . -bench=.
goos: linux
goarch: amd64
pkg: github.com/lilellia/projecteuler-golang
cpu: Intel(R) Core(TM) i7-6500U CPU @ 2.50GHz
BenchmarkPE001-4        11651794                97.77 ns/op
BenchmarkPE002-4        12716006                94.58 ns/op
PASS
ok      github.com/lilellia/projecteuler-golang 2.559s

# run one benchmark (here, problem 2: give as a 3-digit number)
$ go test . -bench=002
goos: linux
goarch: amd64
pkg: github.com/lilellia/projecteuler-golang
cpu: Intel(R) Core(TM) i7-6500U CPU @ 2.50GHz
BenchmarkPE002-4        12586420                92.37 ns/op
PASS
ok      github.com/lilellia/projecteuler-golang 1.263s
```

## Problems completed (2):

- **PE001:** Multiples of 3 and 5
- **PE002:** Even Fibonacci numbers