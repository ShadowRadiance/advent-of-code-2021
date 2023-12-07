package days

import (
	"testing"
)

func TestDay01_Part01(t *testing.T) {
	tests := []Test{
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

func TestDay01_Part02(t *testing.T) {
	tests := []Test{
		{input: "two1nine", expected: "29"},
		{input: "eightwothree", expected: "83"},
		{input: "abcone2threexyz", expected: "13"},
		{input: "xtwone3four", expected: "24"},
		{input: "4nineeightseven2", expected: "42"},
		{input: "zoneight234", expected: "14"},
		{input: "7pqrstsixteen", expected: "76"},
		{input: "two1nine\neightwothree\nabcone2threexyz\nxtwone3four\n4nineeightseven2\nzoneight234\n7pqrstsixteen", expected: "281"},
	}

	for _, test := range tests {
		actual := Day01{}.Part02(test.input)
		if test.expected != actual {
			t.Errorf("expected: %v, got: %v", test.expected, actual)
		}
	}
}
