package domain

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"

	"github.com/jotadrilo/cookify/app/core/biz"
)

type User struct {
	UUID      string
	Name      string
	Email     string
	Gender    Gender
	BirthDate time.Time
	Weight    float32
	Height    float32

	// Computed
	BMRMifflinStJeor         float32
	BMRRevisedHarrisBenedict float32
}

func (x *User) Validate() error {
	return validation.ValidateStruct(x,
		validation.Field(&x.Name, validation.Required),
		validation.Field(&x.Email, validation.Required),
		validation.Field(&x.Gender, validation.Required),
		validation.Field(&x.BirthDate, validation.Required),
		validation.Field(&x.Weight, validation.Required),
		validation.Field(&x.Height, validation.Required),
	)
}

func (x *User) Fixup() *User {
	var (
		now = time.Now()
		age = time.Now().Year() - x.BirthDate.Year()
	)

	if now.Month() < x.BirthDate.Month() || (now.Month() == x.BirthDate.Month() && now.Day() < x.BirthDate.Day()) {
		age--
	}

	if x.Gender == Gender_MALE {
		x.BMRMifflinStJeor = biz.GetBasalMetabolismMifflinStJeorMale(age, x.Height, x.Weight)
		x.BMRRevisedHarrisBenedict = biz.GetBasalMetabolismRevisedHarrisBenedictMale(age, x.Height, x.Weight)
	}

	if x.Gender == Gender_FEMALE {
		x.BMRMifflinStJeor = biz.GetBasalMetabolismMifflinStJeorFemale(age, x.Height, x.Weight)
		x.BMRRevisedHarrisBenedict = biz.GetBasalMetabolismRevisedHarrisBenedictFemale(age, x.Height, x.Weight)
	}

	return x
}
