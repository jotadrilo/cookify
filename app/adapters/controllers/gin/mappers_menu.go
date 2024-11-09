package gin

import (
	"github.com/jotadrilo/cookify/app/api"
	"github.com/jotadrilo/cookify/app/core/domain"
	"github.com/jotadrilo/cookify/internal/oapi"
	"github.com/jotadrilo/cookify/internal/slices"
)

func DomainMenuToAPIMenu(v *domain.Menu) api.Menu {
	if v == nil {
		return api.Menu{}
	}
	return api.Menu{
		Uuid:  oapi.UUID(v.UUID),
		Label: DomainMenuLabelToAPIMenuLabel(v.Label),
		Name:  v.Name,
	}
}

func DomainMenusToAPIMenus(v []*domain.Menu) api.Menus {
	return slices.Map(v, DomainMenuToAPIMenu)
}

func DomainMenuToAPIMenuDetailed(v *domain.Menu) api.MenuDetailed {
	if v == nil {
		return api.MenuDetailed{}
	}

	var (
		ingredients = DomainIngredientsToAPIIngredients(v.Ingredients)
		recipes     = DomainRecipesToAPIRecipes(v.Recipes)
	)

	return api.MenuDetailed{
		Uuid:                oapi.UUID(v.UUID),
		Name:                v.Name,
		Label:               DomainMenuLabelToAPIMenuLabel(v.Label),
		Ingredients:         &ingredients,
		Recipes:             &recipes,
		NutritionFactsTotal: DomainNutritionFactsToAPINutritionFacts(v.NutritionFactsTotal),
	}
}

func DomainMenusToAPIMenusDetailed(v []*domain.Menu) api.MenusDetailed {
	return slices.Map(v, DomainMenuToAPIMenuDetailed)
}
