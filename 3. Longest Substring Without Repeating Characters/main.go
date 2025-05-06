package main

import "fmt"

func main() {
	s := "pwwkew"
	fmt.Println("final result:", lengthOfLongestSubstring(s))
}

func lengthOfLongestSubstring(s string) int {
	max := 0

	tmp := []string{}
	for i := 0; i < len(s); i++ {
		sub := s[0 : len(s)-i]
		rev := s[i:]
		tmp = append(tmp, sub)

		res := validate(tmp)
		if res >= max {
			max = res
		}
		tmp = nil
		tmp = append(tmp, rev)
		res = validate(tmp)
		if res >= max {
			max = res
		}

		tmp = nil
	}
	return max
}

func validate(s []string) int {
	res := 0
	for _, str := range s {
		m := make(map[rune]bool)
		for _, ch := range str {
			if _, exists := m[ch]; !exists {
				res++
				m[ch] = true
			} else {
				// return when char appiers in map. We can don't return then we get non repeated character map
				m = nil
				return res
			}
		}
	}

	return res
}
