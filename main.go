package main

import (
	"fmt"
	"math"
)

var (
	sums map[int64][]string
)

func sumDigits(number int64) int64 {
	var sum int64 = 0
	if number == 0 {
		return 0
	}
	sum = number%10 + sumDigits(number/10)
	return sum
}

// Only if you want to see the actual ticket values in the output.
func genCombinations(values []string) []string {
	res := make([]string, 0)
	for i := 0; i < len(values); i++ {
		for j := 0; j < len(values); j++ {
			res = append(res, fmt.Sprintf("%s%s", values[i], values[j]))
		}
	}
	return res
}

func calcCombinations(values []string) int {
	return int(math.Pow(float64(len(values)), 2))
}

func main() {

	sums = make(map[int64][]string)

	var i int64
	for i = 0; i < 1000; i++ {
		sum := sumDigits(i)
		digits, ok := sums[sum]
		if !ok {
			digits = make([]string, 0)
		}
		sums[sum] = append(digits, fmt.Sprintf("%03d", i))
	}

	total := 0
	for _, v := range sums {
		total = total + calcCombinations(v)

		// Uncomment only if we want to print all the possible tickets.
		// tickets := genCombinations(v)
		// fmt.Printf("%d : %v\n\n", k, tickets)
	}

	fmt.Println("Total lucky tickets: ", total)
}
