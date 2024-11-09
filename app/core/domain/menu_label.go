package domain

import (
	"encoding/json"
	"fmt"
)

type MenuLabel uint

const (
	MenuLabel_UNKNOWN MenuLabel = iota
	MenuLabel_BREAKFAST
	MenuLabel_MORNING_SNACK
	MenuLabel_LUNCH
	MenuLabel_AFTERNOON_SNACK
	MenuLabel_DINNER
)

var (
	MenuLabel_names = map[MenuLabel]string{
		MenuLabel_UNKNOWN:         "UNKNOWN",
		MenuLabel_BREAKFAST:       "BREAKFAST",
		MenuLabel_MORNING_SNACK:   "MORNING_SNACK",
		MenuLabel_LUNCH:           "LUNCH",
		MenuLabel_AFTERNOON_SNACK: "AFTERNOON_SNACK",
		MenuLabel_DINNER:          "DINNER",
	}
	MenuLabel_values = map[string]MenuLabel{
		"UNKNOWN":         MenuLabel_UNKNOWN,
		"BREAKFAST":       MenuLabel_BREAKFAST,
		"MORNING_SNACK":   MenuLabel_MORNING_SNACK,
		"LUNCH":           MenuLabel_LUNCH,
		"AFTERNOON_SNACK": MenuLabel_AFTERNOON_SNACK,
		"DINNER":          MenuLabel_DINNER,
	}
)

func ParseMenuLabel(s string) MenuLabel {
	if v, ok := MenuLabel_values[s]; ok {
		return v
	}
	return MenuLabel_UNKNOWN
}

func (x MenuLabel) String() string {
	if v, ok := MenuLabel_names[x]; ok {
		return v
	}
	return MenuLabel_UNKNOWN.String()
}

func (x MenuLabel) MarshalJSON() ([]byte, error) {
	if s, ok := MenuLabel_names[x]; ok {
		return json.Marshal(s)
	}
	return nil, fmt.Errorf("unknown menu label %d", x)
}

func (x *MenuLabel) UnmarshalJSON(text []byte) error {
	var s string
	if err := json.Unmarshal(text, &s); err != nil {
		return err
	}

	v := ParseMenuLabel(s)
	if v == MenuLabel_UNKNOWN {
		return fmt.Errorf("unknown menu label %s", s)
	}

	*x = v
	return nil
}

func (x MenuLabel) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

func (x *MenuLabel) UnmarshalText(text []byte) error {
	s := string(text)

	v := ParseMenuLabel(s)
	if v == MenuLabel_UNKNOWN {
		return fmt.Errorf("unknown menu label %s", s)
	}

	*x = v
	return nil
}
