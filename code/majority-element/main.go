package main

import "fmt"

func main() {
	fmt.Println(majorityElement([]int{1, 2, 3, 4, 5, 6, 4, 7}))
}

func majorityElement(nums []int) int {
	var me, cnt int
	me = nums[0]
	cnt = 1
	for _, num := range nums[1:] {
		if num == me {
			cnt++
		} else {
			if cnt == 0 {
				me = num
			} else {
				cnt--
			}
		}
	}
	return me
}
