package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i, countPrimes(i))
	}
}

func countPrimes(n int) int {
	number_list := make([]int, n)
	count := 0
	for i := 2; i < n; i++ {
		if number_list[i] == 0 {
			count++
			for j := 2 * i; j < n; j += i {
				number_list[j] = 1
			}
		}
	}
	return count
}
