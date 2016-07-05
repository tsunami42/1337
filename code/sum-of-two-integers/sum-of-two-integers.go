package main

import "fmt"

func main() {
	fmt.Printf("%d", getSum(-3, -2))
}

// cases 1, -1
func getSum(a int, b int) int {
	for b != 0 {
		carry := a & b
		a ^= b
		b = carry << 1
	}
	return a
}
