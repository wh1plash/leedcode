package main

import "fmt"

func main() {
	nums := []int{3, 2, 4}
	target := 6
	res := twoSum(nums, target)
	fmt.Println(res)
}

func twoSum(nums []int, target int) []int {
	res := []int{}
	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			if target-nums[j] == nums[i] {
				res = append(res, i, j)
				return res
			}
		}
	}
	return res
}
