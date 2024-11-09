package domain

import (
	"encoding/json"
	"fmt"
)

type Gender uint

const (
	Gender_UNKNOWN Gender = iota
	Gender_MALE
	Gender_FEMALE
	Gender_BOT
)

var (
	Gender_names = map[Gender]string{
		Gender_UNKNOWN: "n/a",
		Gender_MALE:    "male",
		Gender_FEMALE:  "female",
		Gender_BOT:     "bot",
	}
	Gender_values = map[string]Gender{
		"n/a":    Gender_UNKNOWN,
		"male":   Gender_MALE,
		"female": Gender_FEMALE,
		"bot":    Gender_BOT,
	}
)

func ParseGender(s string) Gender {
	if v, ok := Gender_values[s]; ok {
		return v
	}
	return Gender_UNKNOWN
}

func (x Gender) String() string {
	if v, ok := Gender_names[x]; ok {
		return v
	}
	return Gender_UNKNOWN.String()
}

func (x Gender) MarshalJSON() ([]byte, error) {
	if s, ok := Gender_names[x]; ok {
		return json.Marshal(s)
	}
	return nil, fmt.Errorf("unknown gender %d", x)
}

func (x *Gender) UnmarshalJSON(text []byte) error {
	var s string
	if err := json.Unmarshal(text, &s); err != nil {
		return err
	}

	v := ParseGender(s)
	if v == Gender_UNKNOWN {
		return fmt.Errorf("unknown gender %s", s)
	}

	*x = v
	return nil
}

func (x Gender) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

func (x *Gender) UnmarshalText(text []byte) error {
	s := string(text)

	v := ParseGender(s)
	if v == Gender_UNKNOWN {
		return fmt.Errorf("unknown gender %s", s)
	}

	*x = v
	return nil
}
