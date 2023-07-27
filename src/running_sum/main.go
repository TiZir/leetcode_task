package main

import "fmt"

func runningSum(nums []int) []int {
	buf := make([]int, len(nums), len(nums))
	buf[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		buf[i] = nums[i] + buf[i-1]
	}
	return buf
}

func main() {
	arr := []int{1, 1, 1, 1, 1, 1}
	arr = runningSum(arr)
	for i, _ := range arr {
		fmt.Printf("%d", arr[i])
	}
}
