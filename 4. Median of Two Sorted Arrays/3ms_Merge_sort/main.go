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

	sorted := MergeSort(nums1)

	fmt.Println(sorted)
	l := len(sorted)
	if l%2 == 0 {
		return float64(sorted[l/2]+sorted[l/2-1]) / 2
	}
	return float64(sorted[l/2])
}

func MergeSort(s []int) []int {
	if len(s) <= 1 {
		return s
	}

	mid := len(s) / 2
	left := MergeSort(s[:mid])
	right := MergeSort(s[mid:])
	return merge(left, right)
}

func merge(l, r []int) []int {
	res := make([]int, 0, len(l)+len(r))
	i, j := 0, 0

	for i < len(l) && j < len(r) {
		if l[i] <= r[j] {
			res = append(res, l[i])
			i++
		} else {
			res = append(res, r[j])
			j++
		}
	}

	res = append(res, l[i:]...)
	res = append(res, r[j:]...)
	return res
}
