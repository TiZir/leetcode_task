package main

import "fmt"

func minStartValue(nums []int) int {
	flag := false
	startValue := 1
	for {
		total := startValue
		for i, _ := range nums {
			total += nums[i]
			if total < 1 {
				flag = false
				break
			} else {
				flag = true
			}
		}
		if flag {
			break
		}
		startValue++
	}
	return startValue
}

func main() {
	arr := []int{1, 2}
	fmt.Printf("%d", minStartValue(arr))
}
