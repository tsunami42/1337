package main

import "fmt"

func main() {
	fmt.Println(reverseString("hello"))
}

func reverseString(s string) string {
	s1 := []byte(s)
	for i := 0; i < len(s1)/2; i++ {
		a := s1[i]
		s1[i] = s1[len(s1)-i-1]
		s1[len(s1)-i-1] = a
	}
	return string(s1)
}
