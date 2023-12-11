package day01

import (
	"testing"

	"github.com/shadowradiance/advent-of-code/2023-go/util"
)

func TestSolution_Part01(t *testing.T) {
	examples := []util.TestExample{
		{Input: "1abc2", Expected: "12"},
		{Input: "pqr3stu8vwx", Expected: "38"},
		{Input: "a1b2c3d4e5f", Expected: "15"},
		{Input: "treb7uchet", Expected: "77"},
		{Input: "1abc2\npqr3stu8vwx\na1b2c3d4e5f\ntreb7uchet", Expected: "142"},
	}

	for _, example := range examples {
		actual := Solution{}.Part01(example.Input)
		if example.Expected != actual {
			t.Errorf("expected: %v, got: %v", example.Expected, actual)
		}
	}
}

func TestSolution_Part02(t *testing.T) {
	examples := []util.TestExample{
		{Input: "two1nine", Expected: "29"},
		{Input: "eightwothree", Expected: "83"},
		{Input: "abcone2threexyz", Expected: "13"},
		{Input: "xtwone3four", Expected: "24"},
		{Input: "4nineeightseven2", Expected: "42"},
		{Input: "zoneight234", Expected: "14"},
		{Input: "7pqrstsixteen", Expected: "76"},
		{Input: "two1nine\neightwothree\nabcone2threexyz\nxtwone3four\n4nineeightseven2\nzoneight234\n7pqrstsixteen", Expected: "281"},
	}

	for _, example := range examples {
		actual := Solution{}.Part02(example.Input)
		if example.Expected != actual {
			t.Errorf("expected: %v, got: %v", example.Expected, actual)
		}
	}
}
