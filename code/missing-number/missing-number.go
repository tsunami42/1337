package main

import "fmt"

func main() {
	fmt.Println(missingNumber([]int{0, 1, 2, 4}))
}

func missingNumber(nums []int) int {
	i := 0
	n := 0
	result := 0
	for i, n = range nums {
		result = result + i - n
	}
	return result + i + 1
}
