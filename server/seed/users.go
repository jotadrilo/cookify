package seed

import (
	"time"

	"github.com/jotadrilo/cookify/app/core/domain"
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
	cookify = &domain.User{
		UUID:      "5eb63182-29a0-4924-a3bc-4b5e552670f6",
		Name:      "cookify",
		Email:     "cookify@cookify.com",
		Gender:    domain.Gender_BOT,
		BirthDate: mustParseDate("2024-11-30"),
		Weight:    7,
		Height:    1,
	}

	users = []*domain.User{
		jotadrilo,
		cookify,
	}
)
