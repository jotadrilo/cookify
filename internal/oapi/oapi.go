package oapi

import "github.com/google/uuid"

func UUID(v string) *uuid.UUID {
	if v == "" {
		return nil
	}
	id := uuid.MustParse(v)
	return &id
}

func UUIDValue(v *uuid.UUID) string {
	if v == nil {
		return ""
	}
	return v.String()
}

func Float32(v float32) *float32 {
	return &v
}

func Float32Value(v *float32) float32 {
	if v == nil {
		return 0
	}
	return *v
}

func String(v string) *string {
	return &v
}

func StringValue(v *string) string {
	if v == nil {
		return ""
	}
	return *v
}

func StringSlice(v []string) *[]string {
	return &v
}

func StringSliceValue(v *[]string) []string {
	if v == nil {
		return nil
	}
	return *v
}
