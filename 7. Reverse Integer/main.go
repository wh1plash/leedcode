package main

import (
	"fmt"
	"strconv"
)

func main() {
	i := reverse(1534236469)
	fmt.Println(i)
}

func reverse(i int) int {
	const (
		MinInt32 = -2147483648
		MaxInt32 = 2147483647
	)
	d := i
	if i < 0 {
		d = -d
	}
	s := strconv.Itoa(int(d))
	b := []byte(s)
	for r, j := 0, len(b)-1; r < j; r, j = r+1, j-1 {
		b[r], b[j] = b[j], b[r]
	}
	res, err := strconv.Atoi(string(b))
	if err != nil {
		panic(err)
	}
	if res < MinInt32 || res > MaxInt32 {
		return 0
	}
	if i < 0 {

		res = -res
	}

	return res
}
