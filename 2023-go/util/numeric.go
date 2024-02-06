package util

import (
	"math"
	"slices"
	"strconv"

	"github.com/shadowradiance/advent-of-code/2023-go/util/constraints"
)

const NumberChars = "1234567890"

func ConvertNumeric(s string) (i int) {
	i, err := strconv.Atoi(s)
	PanicOnError(err)
	return
}

func MapStringsToIntegers(ss []string) []int {
	return Transform(ss, func(item string) int {
		return ConvertNumeric(item)
	})
}

func IntSliceContainsInt(slice []int, number int) bool {
	return slices.Contains(slice, number)
}

func ChunkIntSlice(slice []int, chunkSize int) [][]int {
	var chunks [][]int
	for {
		if len(slice) == 0 {
			break
		}

		// necessary check to avoid slicing beyond
		// slice capacity
		if len(slice) < chunkSize {
			chunkSize = len(slice)
		}

		chunks = append(chunks, slice[0:chunkSize])
		slice = slice[chunkSize:]
	}

	return chunks
}

func GreatestCommonDivisor(a, b interface{}) interface{} {
	switch a.(type) {
	case int, int64, int32, int16, int8:
		a64, _ := a.(int64)
		b64, _ := b.(int64)
		return GreatestCommonDivisorI(a64, b64)
	case float64, float32:
		a64, _ := a.(float64)
		b64, _ := b.(float64)
		return GreatestCommonDivisorF(a64, b64)
	default:
		panic("Cannot call GCD on a non numeric value")
	}
}

func GreatestCommonDivisorI[T constraints.Integral](a, b T) T {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}
func GreatestCommonDivisorF[T constraints.Floating](a, b T) T {
	for b != 0 {
		t := b
		b = T(math.Mod(float64(a), float64(b)))
		a = t
	}
	return a
}

func LowestCommonMultiple(a, b interface{}) interface{} {
	switch a.(type) {
	case int, int64, int32, int16, int8:
		a64, _ := a.(int64)
		b64, _ := b.(int64)
		return LowestCommonMultipleI(a64, b64)
	case float64, float32:
		a64, _ := a.(float64)
		b64, _ := b.(float64)
		return LowestCommonMultipleF(a64, b64)
	default:
		panic("Cannot call GCD on a non numeric value")
	}
}

func LowestCommonMultipleI[T constraints.Integral](a, b T) T {
	result := a * b / GreatestCommonDivisorI(a, b)

	return result
}

func LowestCommonMultipleF[T constraints.Floating](a, b T) T {
	result := a * b / GreatestCommonDivisorF(a, b)

	return result
}

func LowestCommonMultipleSlice(numbers []int) int {
	switch len(numbers) {
	case 0:
		return 0
	case 1:
		return numbers[0]
	case 2:
		return LowestCommonMultipleI(numbers[0], numbers[1])
	default:
		return LowestCommonMultipleI(numbers[0], LowestCommonMultipleSlice(numbers[1:]))
	}
}

func BoolInt(b bool) int {
	if b {
		return 1
	} else {
		return 0
	}
}
