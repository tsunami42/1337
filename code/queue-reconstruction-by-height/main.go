package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(reconstructQueue([][]int{
		[]int{7, 0},
		[]int{4, 4},
		[]int{7, 1},
		[]int{5, 0},
		[]int{6, 1},
		[]int{5, 2},
	}))
}

/*
	Pick out tallest group of people and sort them in a subarray (S). Since there's no other groups of people taller than them, therefore each guy's index will be just as same as his k value.
	For 2nd tallest group (and the rest), insert each one of them into (S) by k value. So on and so forth.
*/
func reconstructQueue(people [][]int) [][]int {
	h_arr := []int{}
	for _, p := range people {
		h_arr = append(h_arr, p[0])
	}
	sort.Ints(h_arr)
	i := 0
	for i < len(h_arr)-1 {
		if h_arr[i] == h_arr[i+1] {
			h_arr = append(h_arr[:i], h_arr[i+1:]...)
		} else {
			i = i + 1
		}
	}
	//log.Println(h_arr)

	result := [][]int{}
	for i := len(h_arr) - 1; i >= 0; i-- {
		k_arr := []int{}
		for _, p := range people {
			if p[0] == h_arr[i] {
				k_arr = append(k_arr, p[1])
			}

		}
		sort.Ints(k_arr)

		for _, k := range k_arr {
			result = append(result[:k], append([][]int{[]int{h_arr[i], k}}, result[k:]...)...)
		}

		//log.Println(k_arr)
	}

	return result
}
