package model

import (
	"context"

	"github.com/uptrace/bun"

	"github.com/jotadrilo/cookify/app/core/domain"
)

type Ingredient struct {
	bun.BaseModel `bun:"table:ingredients"`

	ID int64 `bun:"id,pk,autoincrement"`

	RecipeUUID  string   `bun:"recipe_uuid,nullzero"`
	MenuUUID    string   `bun:"menu_uuid,nullzero"`
	ProductUUID string   `bun:"product_uuid,pk"`
	Product     *Product `bun:"rel:belongs-to,join:product_uuid=uuid"`
	Quantity    float32  `bun:"quantity,notnull"`
}

var _ bun.BeforeAppendModelHook = (*Ingredient)(nil)

func (r *Ingredient) BeforeAppendModel(_ context.Context, query bun.Query) error {
	if r.ProductUUID == "" && r.Product != nil {
		r.ProductUUID = r.Product.UUID
	}

	return nil
}

func IngredientToDomainIngredient(x *Ingredient) *domain.Ingredient {
	if x == nil {
		return nil
	}

	var ingredient = &domain.Ingredient{
		Product:  ProductToDomainProduct(x.Product),
		Quantity: x.Quantity,
	}

	return ingredient
}

func DomainIngredientToIngredient(x *domain.Ingredient) *Ingredient {
	if x == nil {
		return nil
	}

	var ingredient = &Ingredient{
		Product:  DomainProductToProduct(x.Product),
		Quantity: x.Quantity,
	}

	if x.Product != nil {
		ingredient.ProductUUID = x.Product.UUID
	}

	return ingredient
}
