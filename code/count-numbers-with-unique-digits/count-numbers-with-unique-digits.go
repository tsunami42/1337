package main

import "fmt"

func main() {
	fmt.Println(countNumbersWithUniqueDigits(11))
}

func countNumbersWithUniqueDigits(n int) int {
	result := 0
	switch {
	case n == 0:
		return 1
	case n == 1:
		return 10
	default:
		result = 10
		f := 9
		for i := 0; i < n-1 && i < 10-1; i++ {
			f *= (9 - i)
			result += f
		}
	}
	return result
}
