package main

import "fmt"

func main() {
	n := []int{1, 3, 5, 6}
	m := []int{2, 4, 9, 10}
	median := findMedianSortedArrays(n, m)
	fmt.Println(median)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums1 = append(nums1, nums2...)
	n := len(nums1)

	for i := n/2 - 1; i >= 0; i-- {
		heapify(nums1, n, i)
	}

	for i := n - 1; i > 0; i-- {
		nums1[0], nums1[i] = nums1[i], nums1[0]
		heapify(nums1, i, 0)
	}

	l := len(nums1)
	if l%2 == 0 {
		return float64(nums1[l/2]+nums1[l/2-1]) / 2
	}
	return float64(nums1[l/2])
}

func heapify(s []int, n, i int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2

	if left < n && s[left] > s[largest] {
		largest = left
	}

	if right < n && s[right] > s[largest] {
		largest = right
	}

	if largest != i {
		s[i], s[largest] = s[largest], s[i]
		heapify(s, n, largest)
	}
}
