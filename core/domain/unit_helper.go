package domain

// UnitHelper defines a set of functions to ease managing units
type UnitHelper interface {
	Unit() Unit
	Value() float32
}

type Unknown float32

func (x Unknown) Unit() Unit {
	return Unit_UNKNOWN
}

func (x Unknown) Value() float32 {
	return float32(x)
}

type Cal float32

func (x Cal) Unit() Unit {
	return Unit_CALORIE
}

func (x Cal) Value() float32 {
	return float32(x)
}

type Kcal float32

func (x Kcal) Unit() Unit {
	return Unit_CALORIE
}

func (x Kcal) Value() float32 {
	return float32(x) * 1000
}

type MilliGram float32

func (x MilliGram) Unit() Unit {
	return Unit_GRAM
}

func (x MilliGram) Value() float32 {
	return float32(x) / 1000
}

type MicroGram float32

func (x MicroGram) Unit() Unit {
	return Unit_GRAM
}

func (x MicroGram) Value() float32 {
	return float32(x) / 1000000
}

type Gram float32

func (x Gram) Unit() Unit {
	return Unit_GRAM
}

func (x Gram) Value() float32 {
	return float32(x)
}

type KiloGram float32

func (x KiloGram) Unit() Unit {
	return Unit_GRAM
}

func (x KiloGram) Value() float32 {
	return float32(x) * 1000
}

type MilliLiter float32

func (x MilliLiter) Unit() Unit {
	return Unit_GRAM
}

func (x MilliLiter) Value() float32 {
	return float32(x) / 1000
}

type Liter float32

func (x Liter) Unit() Unit {
	return Unit_LITER
}

func (x Liter) Value() float32 {
	return float32(x)
}
