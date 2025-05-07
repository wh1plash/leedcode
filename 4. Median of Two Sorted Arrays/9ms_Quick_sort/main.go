package main

import "fmt"

func main() {
	n := []int{0, 0, 0, 0, 0}
	m := []int{-1, 0, 0, 0, 0, 0, 1}
	median := findMedianSortedArrays(n, m)
	fmt.Println(median)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums1 = append(nums1, nums2...)
	sort(nums1, 0, len(nums1)-1)
	fmt.Println(nums1)
	l := len(nums1)
	if l%2 == 0 {
		return float64(nums1[l/2]+nums1[l/2-1]) / 2
	}
	return float64(nums1[l/2])
}

func sort(s []int, low, high int) {
	if low < high {
		p := partition(s, low, high)
		sort(s, low, p-1)
		sort(s, p+1, high)
	}
}

func partition(s []int, low, high int) int {
	pivot := s[high]
	i := low
	for j := low; j < high; j++ {
		if s[j] < pivot {
			s[i], s[j] = s[j], s[i]
			i++
		}
	}
	s[i], s[high] = s[high], s[i]
	return i
}
