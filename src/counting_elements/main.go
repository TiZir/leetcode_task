package main

import "fmt"

func countElements(arr []int) int {
	couEx := make(map[int]int, len(arr))
	for _, value := range arr {
		couEx[value] += 1
	}
	count := 0
	for key, value := range couEx {
		if couEx[key+1] >= 1 {
			count += value
		}
	}
	return count
}

func main() {
	nums := []int{1, 1, 2, 2}
	fmt.Println(countElements(nums))
}
