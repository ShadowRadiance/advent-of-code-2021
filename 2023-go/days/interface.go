package days

type DayInterface interface {
	Part01(string) string
	Part02(string) string
}

type Test struct {
	input    string
	expected string
}
