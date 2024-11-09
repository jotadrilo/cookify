package domain

type Error struct {
	Error string `json:"error,omitempty"`
	Code  int    `json:"code,omitempty"`
}
