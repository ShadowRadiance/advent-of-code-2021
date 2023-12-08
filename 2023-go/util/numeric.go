package util

import "strconv"

const NumberChars = "1234567890"

func ConvertNumeric(s string) (i int) {
	i, _ = strconv.Atoi(s)
	return
}

func MaxInt(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func MinInt(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
