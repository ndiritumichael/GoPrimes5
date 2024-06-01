package main

//enable the C import while building the shared object
//import "C"
import (
	"fmt"

	"time"
)

var primesInRange = map[int]int{
	10:     4,
	100:    25,
	1000:   168,
	10000:  1229,
	100000: 9592,
}

func getPrimesTill(end int) int {
	primes := []int{2}
	var factos []int

	for i := 3; i < end; i += 2 {
		isPrime := true

		for _, j := range factos {
			if j == i {
				isPrime = false
				break
			}
		}

		if isPrime {
			for j := 3; j*i <= end; j += 2 {
				factos = append(factos, j*i)
			}
			primes = append(primes, i)
		}
	}

	return len(primes)
}

//export getPrimesCount
func getPrimesCount(end int64) int64 {
	var count int64 = 0
	now := time.Now().UnixNano() / int64(time.Millisecond)
	finish := now + 5000

	for time.Now().UnixNano()/int64(time.Millisecond) < finish {
		result := getPrimesTill(int(end))

		if result == primesInRange[int(end)] {
			count++

		}

	}

	return count
}

func main() {
	end := 100_00
	fmt.Printf("Starting The Test\n")
	count := getPrimesCount(int64(end))
	fmt.Printf("Function ran %d times in 5 seconds\n", count)
}
