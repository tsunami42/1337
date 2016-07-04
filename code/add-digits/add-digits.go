package main

import "fmt"

func main() {
	fmt.Println(addDigits(10))
}

//	case: 10
func addDigits(num int) int {
	for num >= 10 {
		num = num%10 + num/10
	}
	return num
}
