package days

import (
	"testing"
)

type test struct {
	input    string
	expected string
}

func TestDay01_Part01(t *testing.T) {
	tests := []test{
		{input: "1abc2", expected: "12"},
		{input: "pqr3stu8vwx", expected: "38"},
		{input: "a1b2c3d4e5f", expected: "15"},
		{input: "treb7uchet", expected: "77"},
		{input: "1abc2\npqr3stu8vwx\na1b2c3d4e5f\ntreb7uchet", expected: "142"},
	}

	for _, test := range tests {
		actual := Day01{}.Part01(test.input)
		if test.expected != actual {
			t.Errorf("expected: %v, got: %v", test.expected, actual)
		}
	}
}
