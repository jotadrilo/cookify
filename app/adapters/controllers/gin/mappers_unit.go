package gin

import (
	"github.com/jotadrilo/cookify/app/api"
	"github.com/jotadrilo/cookify/app/core/domain"
)

func DomainUnitToAPIUnit(v domain.Unit) api.Unit {
	return api.Unit(v.String())
}

func UnitToDomainUnit(v api.Unit) domain.Unit {
	return domain.ParseUnit(string(v))
}
