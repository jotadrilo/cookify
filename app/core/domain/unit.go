package domain

import (
	"encoding/json"
	"fmt"
)

type Unit uint

const (
	Unit_UNKNOWN Unit = iota
	Unit_CALORIE
	Unit_GRAM
	Unit_LITER
)

var (
	Unit_names = map[Unit]string{
		Unit_UNKNOWN: "n/a",
		Unit_CALORIE: "cal",
		Unit_GRAM:    "g",
		Unit_LITER:   "l",
	}
	Unit_values = map[string]Unit{
		"n/a": Unit_UNKNOWN,
		"cal": Unit_CALORIE,
		"g":   Unit_GRAM,
		"G":   Unit_GRAM,
		"l":   Unit_LITER,
		"L":   Unit_LITER,
	}
)

func ParseUnit(s string) Unit {
	if v, ok := Unit_values[s]; ok {
		return v
	}
	return Unit_UNKNOWN
}

func (x Unit) String() string {
	if v, ok := Unit_names[x]; ok {
		return v
	}
	return Unit_UNKNOWN.String()
}

func (x Unit) MarshalJSON() ([]byte, error) {
	if s, ok := Unit_names[x]; ok {
		return json.Marshal(s)
	}
	return nil, fmt.Errorf("unknown unit %d", x)
}

func (x *Unit) UnmarshalJSON(text []byte) error {
	var s string
	if err := json.Unmarshal(text, &s); err != nil {
		return err
	}

	v := ParseUnit(s)
	if v == Unit_UNKNOWN {
		return fmt.Errorf("unknown unit %s", s)
	}

	*x = v
	return nil
}

func (x Unit) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

func (x *Unit) UnmarshalText(text []byte) error {
	s := string(text)

	v := ParseUnit(s)
	if v == Unit_UNKNOWN {
		return fmt.Errorf("unknown unit %s", s)
	}

	*x = v
	return nil
}
