package main

import "fmt"

func main() {
	g := grayCode(0)
	for idx, val := range g {
		fmt.Printf("%d - %4.4b\n", idx, val)
	}
}

func grayCode(n int) []int {
	if n == 0 {
		return []int{0}
	}
	g := make([]int, 2<<uint(n-1))
	g[0] = 0
	for k := 0; k < n; k++ {
		base := 1 << uint(k)
		for i := 0; i < base; i++ {
			g[base+i] = g[base-i-1] + base
		}
	}
	return g
}
