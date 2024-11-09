package gin

import (
	"github.com/jotadrilo/cookify/app/api"
	"github.com/jotadrilo/cookify/app/core/domain"
	"github.com/jotadrilo/cookify/internal/oapi"
)

func DomainPriceToAPIPrice(v *domain.Price) *api.Price {
	return &api.Price{
		Uuid:     oapi.UUID(v.UUID),
		Price:    float32(v.Price),
		Quantity: float32(v.Quantity),
	}
}

func PriceToDomainPrice(v *api.Price) *domain.Price {
	if v == nil {
		return nil
	}
	return &domain.Price{
		UUID:     oapi.UUIDValue(v.Uuid),
		Price:    float32(v.Price),
		Quantity: float32(v.Quantity),
	}
}
