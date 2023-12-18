package util

func Filter[T any](ss []T, test func(T) bool) (ret []T) {
	for _, item := range ss {
		if test(item) {
			ret = append(ret, item)
		}
	}
	return
}

func Accumulate[T any](values []T, operation func(a T, b T) T) T {
	if len(values) < 1 {
		panic("util.Accumulate: empty list")
	}

	var total = values[0]
	for _, value := range values[1:] {
		total = operation(total, value)
	}
	return total
}

func Transform[T any, U any](values []T, operation func(item T) U) []U {
	// example: Transform(blocks, func(s string) int { return s(item) })
	result := make([]U, 0)
	for _, value := range values {
		result = append(result, operation(value))
	}
	return result
}

func All[T any](values []T, predicate func(value T) bool) bool {
	for _, value := range values {
		if !predicate(value) {
			return false
		}
	}
	return true
}

func Any[T any](values []T, predicate func(value T) bool) bool {
	for _, value := range values {
		if predicate(value) {
			return true
		}
	}
	return false
}

func None[T any](values []T, predicate func(value T) bool) bool {
	for _, value := range values {
		if predicate(value) {
			return false
		}
	}
	return true
}
