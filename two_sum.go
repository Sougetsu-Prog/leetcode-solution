package main

import (
	"fmt"
	"slices"
)

func twoSum(nums []int, target int) []int {
	result := []int{}
	for i, _ := range nums {
		diff := target - nums[i]
		if index := slices.Index(nums, diff); index != i && index != -1 {
			result = append(result, i, index)
			return result
		}
	}
	return result
}

func main() {
	nums := []int{2, 7, 11, 15}
	fmt.Println(twoSum(nums, 9))
}
