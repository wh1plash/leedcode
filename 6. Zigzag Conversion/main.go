package main

import "fmt"

func main() {
	str := "AB"
	s := convert(str, 1)
	fmt.Println(s)
}

func convert(s string, numRows int) string {
	if numRows == 1 || numRows >= len(s) {
		return s
	}

	res := make([][]string, numRows)

	for i := range res {
		res[i] = []string{}
	}

	row := 0
	goingDown := false

	for i := range s {
		res[row] = append(res[row], string(s[i]))
		if row == 0 || row == numRows-1 {
			goingDown = !goingDown
		}
		if goingDown {
			row++

		} else {
			row--
		}
	}

	str := ""
	for n := range res {
		for j := range res[n] {
			str = str + res[n][j]
		}
	}
	return str
}
