package main

import "fmt"

func checkIfPangram(sentence string) bool {
	alph := make(map[string]bool)
	for _, char := range sentence {
		if char >= 97 && char <= 122 {
			alph[string(char)] = true
		}
	}
	return len(alph) == 26
}

func main() {
	fmt.Println(checkIfPangram("thequickbrownfoxjumpsoverthelazydog"))
}
