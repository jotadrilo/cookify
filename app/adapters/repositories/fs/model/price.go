package model

import "github.com/jotadrilo/cookify/app/core/domain"

type Price struct {
	UUID     string  `json:"uuid"`
	Price    float32 `json:"price"`
	Quantity float32 `json:"quantity"`
}

func PriceToDomainPrice(x *Price) *domain.Price {
	if x == nil {
		return nil
	}

	var price = &domain.Price{
		UUID:     x.UUID,
		Price:    x.Price,
		Quantity: x.Quantity,
	}

	return price
}

func DomainPriceToPrice(x *domain.Price) *Price {
	if x == nil {
		return nil
	}

	var price = &Price{
		UUID:     x.UUID,
		Price:    x.Price,
		Quantity: x.Quantity,
	}

	return price
}
