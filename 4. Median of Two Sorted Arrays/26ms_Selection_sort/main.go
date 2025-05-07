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

	n := len(nums1)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if nums1[j] < nums1[minIdx] {
				minIdx = j
			}
		}
		nums1[i], nums1[minIdx] = nums1[minIdx], nums1[i]
	}

	fmt.Println(nums1)
	l := len(nums1)
	if l%2 == 0 {
		return float64(nums1[l/2]+nums1[l/2-1]) / 2
	}
	return float64(nums1[l/2])
}
