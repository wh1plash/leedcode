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

	//BubbleSort(nums1)

	n := len(nums1) - 1

	for i := 0; i < n; i++ {
		swapped := false
		for j := 0; j < n-i; j++ {
			if nums1[j] > nums1[j+1] {
				nums1[j], nums1[j+1] = nums1[j+1], nums1[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}

	fmt.Println(nums1)
	l := len(nums1)
	if l%2 == 0 {
		return float64(nums1[l/2]+nums1[l/2-1]) / 2
	}
	return float64(nums1[l/2])
}

func BubbleSort(s []int) {
	n := len(s) - 1
	for i := 0; i < n; i++ {
		swapped := false
		for j := 0; j < n-i; j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}
