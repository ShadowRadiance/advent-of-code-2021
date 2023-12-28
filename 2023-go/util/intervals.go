package util

type Interval struct {
	Start int
	Final int
}

func (i Interval) Length() int {
	return i.Final - i.Start + 1
}

func (i Interval) Invalid() bool {
	return i.Final < i.Start
}
