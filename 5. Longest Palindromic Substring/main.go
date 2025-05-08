package main

import (
	"fmt"
	"time"
)

func main() {
	s := "abababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababa"

	fmt.Println(longestPalindrome(s))
}

func longestPalindrome(s string) string {
	//subs := []string{}
	//subs := allstr(s)
	//sort(subs)
	subs := allPalindromicSubstrings(s)
	str := longest(subs)
	return str
}

func longest(s []string) string {
	start := time.Now()
	lng := 0
	indx := 0
	for i := range s {
		if len(s[i]) > lng {
			lng = len(s[i])
			indx = i
		}
	}

	fmt.Println("find longest", time.Since(start))
	return s[indx]
}

// func sort(s []string) {
// 	start := time.Now()
// 	for i := 1; i < len(s); i++ {
// 		key := s[i]
// 		j := i - 1

// 		for j >= 0 && len(s[j]) < len(key) {
// 			s[j+1] = s[j]
// 			j--
// 		}
// 		s[j+1] = key
// 	}
// 	fmt.Println("sort", time.Since(start))
// }

func allPalindromicSubstrings(s string) []string {
	start := time.Now()
	n := len(s)
	res := make([]string, 0, n)
	for center := 0; center < 2*n-1; center++ {
		left := center / 2
		right := left + center%2
		for left >= 0 && right < n && s[left] == s[right] {
			res = append(res, s[left:right+1])
			left--
			right++
		}
	}
	fmt.Println("find allPalindromicSubstrings", time.Since(start))
	return res
}

func allstr(s string) []string {
	start := time.Now()
	strs := []string{}
	for i := range len(s) {
		for j := i + 1; j <= len(s); j++ {
			ins := s[i:j]
			if isPalindrom(ins) {
				strs = append(strs, ins)

			}
		}
	}
	fmt.Println("all strings", time.Since(start))
	return strs
}

func isPalindrom(s string) bool {
	b := []byte(s)
	for r, j := 0, len(b)-1; r < j; r, j = r+1, j-1 {
		b[r], b[j] = b[j], b[r]
	}

	return s == string(b)
}
