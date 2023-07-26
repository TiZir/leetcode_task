package main

import "fmt"

func reverseString(s []byte) {
	for i, j := len(s)-1, 0; i >= len(s)/2; i-- {
		if j == i {
			break
		}
		s[i], s[j] = s[j], s[i]
		j++
	}
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c", s[i])
	}
}

func main() {
	str := []byte{'H', 'a', 'n', 'a', 'h'}
	reverseString(str)
}
