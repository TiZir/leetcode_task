package main

import "fmt"

func sortedSquares(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < ((len(nums) - 1) - i); j++ {
			if nums[j]*nums[j] > nums[j+1]*nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	for i, _ := range nums {
		nums[i] = nums[i] * nums[i]
	}
	return nums
}

func main() {
	arr := []int{-7, -3, 2, 3, 11}
	arr = sortedSquares(arr)
	for _, elem := range arr {
		fmt.Printf("%d ", elem)
	}
}
