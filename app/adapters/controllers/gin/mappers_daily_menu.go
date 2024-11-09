package gin

import (
	"github.com/jotadrilo/cookify/app/api"
	"github.com/jotadrilo/cookify/app/core/domain"
	"github.com/jotadrilo/cookify/internal/oapi"
	"github.com/jotadrilo/cookify/internal/slices"
)

func DomainDailyMenuToAPIDailyMenu(v *domain.DailyMenu) api.DailyMenu {
	if v == nil {
		return api.DailyMenu{}
	}
	return api.DailyMenu{
		Uuid: oapi.UUID(v.UUID),
		Name: v.Name,
	}
}

func DomainDailyMenusToAPIDailyMenus(v []*domain.DailyMenu) api.DailyMenus {
	return slices.Map(v, DomainDailyMenuToAPIDailyMenu)
}

func DomainDailyMenuToAPIDailyMenuDetailed(v *domain.DailyMenu) api.DailyMenuDetailed {
	if v == nil {
		return api.DailyMenuDetailed{}
	}
	return api.DailyMenuDetailed{
		Uuid:                oapi.UUID(v.UUID),
		Name:                v.Name,
		Menus:               DomainMenusToAPIMenusDetailed(v.Menus),
		NutritionFactsTotal: DomainNutritionFactsToAPINutritionFacts(v.NutritionFactsTotal),
	}
}

func DomainDailyMenusToAPIDailyMenusDetailed(v []*domain.DailyMenu) api.DailyMenusDetailed {
	return slices.Map(v, DomainDailyMenuToAPIDailyMenuDetailed)
}
