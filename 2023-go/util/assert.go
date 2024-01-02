package util

func Assert(b bool, s string) {
	if !b {
		panic(s)
	}
}
