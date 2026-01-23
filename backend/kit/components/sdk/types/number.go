package types

type Number interface {
	int64 | int32 | int16 | int8 | int | float64 | float32
}

type Float interface {
	float64 | float32
}

type Int interface {
	int64 | int32 | int16 | int8 | int
}
