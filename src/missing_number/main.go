package main

import "fmt"

func missingNumber(nums []int) int {
	numEx := make(map[int]interface{}, len(nums))
	for i := 0; i < len(nums); i++ {
		numEx[nums[i]] = nums[i]
	}
	for i := 0; i < len(nums)+1; i++ {
		if _, ok := numEx[i]; !ok {
			return i
		}
	}
	return -1
}

func main() {
	nums := []int{0, 1}
	fmt.Println(missingNumber(nums))
}
