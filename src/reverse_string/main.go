package main

import "fmt"

func reverseString(s []byte) {
	j := 0
	for i := len(s) - 1; i >= len(s)/2; i-- {
		if j == i {
			break
		}
		buf := s[i]
		s[i] = s[j]
		s[j] = buf
		j++

	}
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c", s[i])
	}
}

func main() {
	str := []byte{'H', 'a', 'n', 'n', 'a', 'h'}
	reverseString(str)
}
