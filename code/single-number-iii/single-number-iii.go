package main

import "fmt"

func main() {
	fmt.Println(singleNumber([]int{1, 2, 1, 3, 2, 5, 2, 5}))
}

func singleNumber1(nums []int) []int {
	numMap := make(map[int]int)

	for _, i := range nums {
		numMap[i]++
	}

	result := []int{}
	for k, v := range numMap {
		if v == 1 {
			result = append(result, k)
		}
	}
	return result
}

func singleNumber(nums []int) []int {
	diff := 0
	for _, i := range nums {
		diff ^= i
	}

	// least significant bit, 最低有效位
	//   ...1000...          ...1000...
	// & ...0111... + 1 => & ...1000...
	diff &= -diff

	var a, b int
	for _, i := range nums {
		if diff&i == 0 {
			a ^= i
		} else {
			b ^= i
		}
	}

	return []int{a, b}
}
