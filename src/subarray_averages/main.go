package main

import "fmt"

func getAverages(nums []int, k int) []int {
	buf := make([]int, len(nums))
	avg := make([]int, len(nums))
	buf[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		buf[i] = nums[i] + buf[i-1]
	}
	for i := 0; i < len(nums); i++ {
		if i >= k && i <= len(nums)-k-1 {
			avg[i] = (buf[i+k] - buf[i-k] + nums[i-k]) / (k + k + 1)
		} else {
			avg[i] = -1
		}
	}
	return avg
}

func main() {
	arr := []int{7, 4, 3, 9, 1, 8, 5, 2, 6}
	out := getAverages(arr, 3)
	for i, _ := range out {
		fmt.Printf("%d ", out[i])
	}

}
