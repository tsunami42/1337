package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(isHappy(18))
}

func isHappy(n int) bool {
	m := make(map[int]bool)

	for {
		fmt.Println(n)
		if _, ok := m[n]; ok {
			return n == 1
		}

		m[n] = true
		happySum := 0
		for k := n; k > 0; k = k / 10 {
			happySum += int(math.Pow(float64(k%10), 2))
		}
		n = happySum
	}
}
