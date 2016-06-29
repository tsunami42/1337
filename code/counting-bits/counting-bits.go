package main

import "fmt"

func main() {
	fmt.Println(countBits(2))
}

func countBits(num int) []int {
	r := make([]int, num+1)
	r[0] = 0
	for i := 1; i <= num; i++ {
		r[i] = r[i/2] + i&1
	}
	return r
}
