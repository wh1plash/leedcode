package main

import "fmt"

func main() {
	n := []int{1, 3}
	m := []int{2}
	median := findMedianSortedArrays(n, m)
	fmt.Println(median)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums1 = append(nums1, nums2...)

	for i := 1; i < len(nums1); i++ {
		key := nums1[i]
		j := i - 1

		for j >= 0 && nums1[j] > key {
			nums1[j+1] = nums1[j]
			j--
		}
		nums1[j+1] = key
	}
	fmt.Println(nums1)
	l := len(nums1)
	if l%2 == 0 {
		return float64(nums1[l/2]+nums1[l/2-1]) / 2
	}
	return float64(nums1[l/2])
}
