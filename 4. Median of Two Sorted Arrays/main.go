package main

import "fmt"

func main() {
	n := []int{1, 3}
	m := []int{2}
	median := indMedianSortedArrays(n, m)
	fmt.Println(median)
}

func indMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	res := make([]int, len(nums1)+len(nums2))
	for i := range nums1 {
		for j := range nums2 {
			if nums1[i] < nums2[j] {
				res[i] = nums1[i]
				fmt.Println("case1:", res)
			} else {
				res[i] = nums2[j]
				fmt.Println("case2:", res)
			}
		}

		fmt.Println(res)

	}
	return 2 / 2
}
