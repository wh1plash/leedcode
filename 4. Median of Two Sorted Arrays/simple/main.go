package main

import "fmt"

func main() {
	n := []int{1, 3, 4}
	m := []int{2, 5, 7}
	median := findMedianSortedArrays(n, m)
	fmt.Println(median)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	res := make([]int, len(nums1)+len(nums2))

	copy(res[len(nums1):], res[:len(res)])
	for i := range nums1 {
		res[i] = nums1[i]
	}

	copy(res[len(nums2):], res[:len(res)])
	for i := range nums2 {
		res[i] = nums2[i]
	}

	for i := range res {
		for j := 0; j < len(res)-i-1; j++ {
			if res[j] > res[j+1] {
				res[j], res[j+1] = res[j+1], res[j]
			}
		}
	}

	if len(res)%2 == 0 {
		f := len(res) / 2
		s := f - 1
		mid := float64((res[f] + res[s])) / 2
		return float64(mid)
	}

	mid := len(res) / 2
	n := float64(res[mid])
	return n
}
