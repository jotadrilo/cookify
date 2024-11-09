package domain

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// UnitHelper defines a set of functions to ease managing units
type UnitHelper interface {
	Unit() Unit
	Value() float32
}

var quantityStrRe = regexp.MustCompile(`([\d\\.]+)\s*(\w+)`)

func ParseQuantityString(s string) (UnitHelper, error) {
	matches := quantityStrRe.FindStringSubmatch(s)

	if matches[1] == "" || matches[2] == "" {
		return nil, fmt.Errorf("invalid quantity string: %s", s)
	}

	q, err := strconv.ParseFloat(matches[1], 32)
	if err != nil {
		return nil, err
	}

	return ParseUnitHelper(strings.ToLower(matches[2]), float32(q)), nil
}

func ParseUnitHelper(s string, q float32) UnitHelper {
	switch strings.ToLower(s) {
	case "cal":
		return Cal(q)
	case "kcal":
		return Kcal(q)
	case "mcg":
		return MicroGram(q)
	case "mg":
		return MilliGram(q)
	case "g":
		return Gram(q)
	case "kg":
		return KiloGram(q)
	case "ml":
		return MilliLiter(q)
	case "l":
		return Liter(q)
	default:
		return Unknown(q)
	}
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
