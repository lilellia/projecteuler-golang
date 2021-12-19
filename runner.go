package main

import (
	"flag"
	"log"
	"strconv"
	"time"
)

func main() {
	flag.Parse()

	var result int;
	for _, problem := range flag.Args() {
		problem, err := strconv.Atoi(problem)
	
		if err != nil {
			log.Printf("Cannot interpret %v. Skipping.\n", problem)
			continue
		}

		start := time.Now()
		switch problem {
		case 1:
			result = PE001(3, 5, 1000)
		case 2:
			isEven := func(n int) bool { return n % 2 == 0 }
			result = PE002(4000000, isEven)
		case 3:
			result = PE003(600851475143)
		default:
			log.Printf("Do not have a solution for #%03d. Skipping.\n", problem)
			continue
		}
	
		elapsed := float64(time.Since(start)) / 1000000.0
		log.Printf("PE#%03d [Elapsed: %9.3f ms] Answer: %d\n", problem, elapsed, result)
	}
	

}