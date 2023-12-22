package constraints

type SignedIntegral interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}
type UnsignedIntegral interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}
type Integral interface {
	SignedIntegral | UnsignedIntegral
}
type Floating interface {
	~float32 | ~float64
}

type Signed interface {
	SignedIntegral | Floating
}

type Numeric interface {
	Integral | Floating
}
