package model

import (
	"time"

	"github.com/jotadrilo/cookify/app/core/domain"
)

type User struct {
	UUID      string    `json:"uuid"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Gender    string    `json:"gender"`
	BirthDate time.Time `json:"birth_date"`
	Weight    float32   `json:"weight"`
	Height    float32   `json:"height"`
}

func (x User) OwnsRecipe(v *Recipe) bool {
	return x.UUID == v.UserUUID
}

func (x User) OwnsMenu(v *Menu) bool {
	return x.UUID == v.UserUUID
}

func (x User) OwnsDailyMenu(v *DailyMenu) bool {
	return x.UUID == v.UserUUID
}

func UserToDomainUser(x *User) *domain.User {
	if x == nil {
		return nil
	}

	var user = &domain.User{
		UUID:      x.UUID,
		Name:      x.Name,
		Email:     x.Email,
		Gender:    domain.ParseGender(x.Gender),
		BirthDate: x.BirthDate,
		Weight:    x.Weight,
		Height:    x.Height,
	}

	user.Fixup()

	return user
}

func DomainUserToUser(x *domain.User) *User {
	if x == nil {
		return nil
	}

	var user = &User{
		UUID:      x.UUID,
		Name:      x.Name,
		Email:     x.Email,
		Gender:    x.Gender.String(),
		BirthDate: x.BirthDate,
		Weight:    x.Weight,
		Height:    x.Height,
	}

	return user
}
