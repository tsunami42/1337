package main

import "fmt"

/*
Given an integer array of size n, find all elements that appear more than ⌊ n/3 ⌋ times.
The algorithm should run in linear time and in O(1) space.
*/

func main() {
	fmt.Println(majorityElement([]int{1, 1, 1, 2, 2, 3, 3}))
}

func majorityElement(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	me1, me2, cnt1, cnt2 := 1, 2, 0, 0
	for _, num := range nums {
		switch {
		case me1 == num:
			cnt1++
		case me2 == num:
			cnt2++
		case cnt1 == 0:
			me1, cnt1 = num, 1
		case cnt2 == 0:
			me2, cnt2 = num, 1
		default:
			cnt1--
			cnt2--
		}
	}
	num_cnt := []int{0, 0}
	for _, n := range nums {
		switch {
		case n == me1:
			num_cnt[0]++
		case n == me2:
			num_cnt[1]++
		}
	}
	result := []int{}
	if num_cnt[0] > len(nums)/3 {
		result = append(result, me1)
	}

	if num_cnt[1] > len(nums)/3 {
		result = append(result, me2)
	}
	return result
}
