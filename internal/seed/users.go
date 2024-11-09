package seed

import (
	"time"

	"github.com/jotadrilo/cookify/core/domain"
)

func mustParseDate(v string) time.Time {
	t, err := time.Parse(time.DateOnly, v)
	if err != nil {
		panic(err)
	}
	return t
}

var (
	jotadrilo = &domain.User{
		UUID:      "f0a443f6-0906-47ea-9c62-4c1fbec31942",
		Name:      "jotadrilo",
		Email:     "josriolop@gmail.com",
		Gender:    domain.Gender_MALE,
		BirthDate: mustParseDate("1992-11-12"),
		Weight:    100,
		Height:    175,
	}

	users = []*domain.User{
		jotadrilo,
	}
)
