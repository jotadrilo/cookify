package gin

import (
	openapi_types "github.com/oapi-codegen/runtime/types"

	"github.com/jotadrilo/cookify/api"
	"github.com/jotadrilo/cookify/core/domain"
	"github.com/jotadrilo/cookify/internal/slices"
)

func Float32(v float32) *float32 {
	return &v
}

func String(v string) *string {
	return &v
}

func AnyPtr(v any) *any {
	return &v
}

func DomainUserToAPIUser(v *domain.User) api.User {
	if v == nil {
		return api.User{}
	}
	return api.User{
		Uuid:                     AnyPtr(v.UUID),
		Name:                     v.Name,
		Email:                    v.Email,
		Birthday:                 openapi_types.Date{Time: v.BirthDate},
		Gender:                   v.Gender.String(),
		Height:                   v.Height,
		Weight:                   v.Weight,
		BmrMifflinStJeor:         Float32(v.BMRMifflinStJeor),
		BmrRevisedHarrisBenedict: Float32(v.BMRRevisedHarrisBenedict),
	}
}

func DomainUsersToAPIUsers(v []*domain.User) api.Users {
	return slices.Map(v, DomainUserToAPIUser)
}

func DomainProductToAPIProduct(v *domain.Product) api.Product {
	if v == nil {
		return api.Product{}
	}
	return api.Product{
		LangEnUs: v.LangEnUS,
		LangEsEs: v.LangEsES,
		Name:     v.Name,
		Unit:     DomainUnitToAPIUnit(v.Unit),
		Uuid:     AnyPtr(v.UUID),
		Vendor:   v.Vendor,
	}
}

func DomainProductToAPIProductDetailed(v *domain.Product) api.ProductDetailed {
	if v == nil {
		return api.ProductDetailed{}
	}
	return api.ProductDetailed{
		LangEnUs:          v.LangEnUS,
		LangEsEs:          v.LangEsES,
		Name:              v.Name,
		Unit:              DomainUnitToAPIUnit(v.Unit),
		Uuid:              AnyPtr(v.UUID),
		Vendor:            v.Vendor,
		NutritionFacts100: AnyPtr(DomainNutritionFactsToAPINutritionFacts(v.NutritionFacts)),
	}
}

func DomainUnitToAPIUnit(v domain.Unit) any {
	return v.String()
}

func DomainMenuLabelToAPIMenuLabel(v domain.MenuLabel) any {
	return v.String()
}

func DomainProductsToAPIProducts(v []*domain.Product) api.Products {
	return slices.Map(v, DomainProductToAPIProduct)
}

func DomainNutritionFactsToAPINutritionFacts(v *domain.NutritionFacts) api.NutritionFacts {
	if v == nil {
		return api.NutritionFacts{}
	}
	return api.NutritionFacts{
		Cal:                v.Cal,
		Caffeine:           Float32(v.Caffeine),
		Calcium:            Float32(v.Calcium),
		CarbohydrateSugar:  Float32(v.CarbohydrateSugar),
		CarbohydrateTotal:  Float32(v.CarbohydrateTotal),
		Cholesterol:        Float32(v.Cholesterol),
		FatMonounsaturated: Float32(v.FatMonounsaturated),
		FatPolyunsaturated: Float32(v.FatPolyunsaturated),
		FatSaturated:       Float32(v.FatSaturated),
		FatTotal:           Float32(v.FatTotal),
		Fiber:              Float32(v.Fiber),
		Iron:               Float32(v.Iron),
		Potassium:          Float32(v.Potassium),
		Protein:            Float32(v.Protein),
		Salt:               Float32(v.Salt),
		Sodium:             Float32(v.Sodium),
		VitaminA:           Float32(v.VitaminA),
		VitaminB1:          Float32(v.VitaminB1),
		VitaminB12:         Float32(v.VitaminB12),
		VitaminB2:          Float32(v.VitaminB2),
		VitaminB3:          Float32(v.VitaminB3),
		VitaminB4:          Float32(v.VitaminB4),
		VitaminB5:          Float32(v.VitaminB5),
		VitaminB6:          Float32(v.VitaminB6),
		VitaminC:           Float32(v.VitaminC),
		VitaminD:           Float32(v.VitaminD),
		VitaminE:           Float32(v.VitaminE),
		VitaminK:           Float32(v.VitaminK),
	}
}

func DomainIngredientToAPIIngredient(v *domain.Ingredient) api.Ingredient {
	if v == nil {
		return api.Ingredient{}
	}
	return api.Ingredient{
		Product:  DomainProductToAPIProduct(v.Product),
		Quantity: v.Quantity,
	}
}

func DomainIngredientsToAPIIngredients(v []*domain.Ingredient) api.Ingredients {
	return slices.Map(v, DomainIngredientToAPIIngredient)
}

func DomainIngredientToAPIIngredientDetailed(v *domain.Ingredient) api.IngredientDetailed {
	if v == nil {
		return api.IngredientDetailed{}
	}
	return api.IngredientDetailed{
		Product:  DomainProductToAPIProductDetailed(v.Product),
		Quantity: v.Quantity,
	}
}

func DomainIngredientsToAPIIngredientsDetailed(v []*domain.Ingredient) api.IngredientsDetailed {
	return slices.Map(v, DomainIngredientToAPIIngredientDetailed)
}

func DomainRecipeToAPIRecipe(v *domain.Recipe) api.Recipe {
	if v == nil {
		return api.Recipe{}
	}
	return api.Recipe{
		Name: v.Name,
		Uuid: AnyPtr(v.UUID),
	}
}

func DomainRecipesToAPIRecipes(v []*domain.Recipe) api.Recipes {
	return slices.Map(v, DomainRecipeToAPIRecipe)
}

func DomainRecipeToAPIRecipeDetailed(v *domain.Recipe) api.RecipeDetailed {
	if v == nil {
		return api.RecipeDetailed{}
	}

	return api.RecipeDetailed{
		Uuid:                AnyPtr(v.UUID),
		Name:                v.Name,
		Ingredients:         AnyPtr(DomainIngredientsToAPIIngredientsDetailed(v.Ingredients)),
		NutritionFacts100:   AnyPtr(DomainNutritionFactsToAPINutritionFacts(v.NutritionFacts)),
		NutritionFactsTotal: AnyPtr(DomainNutritionFactsToAPINutritionFacts(v.NutritionFactsTotal)),
		Quantity:            Float32(v.Quantity),
	}
}

func DomainRecipesToAPIRecipesDetailed(v []*domain.Recipe) api.RecipesDetailed {
	return slices.Map(v, DomainRecipeToAPIRecipeDetailed)
}

func DomainMenuToAPIMenu(v *domain.Menu) api.Menu {
	if v == nil {
		return api.Menu{}
	}
	return api.Menu{
		Label: DomainMenuLabelToAPIMenuLabel(v.Label),
		Name:  v.Name,
		Uuid:  v.UUID,
	}
}

func DomainMenusToAPIMenus(v []*domain.Menu) api.Menus {
	return slices.Map(v, DomainMenuToAPIMenu)
}

func DomainMenuToAPIMenuDetailed(v *domain.Menu) api.MenuDetailed {
	if v == nil {
		return api.MenuDetailed{}
	}
	return api.MenuDetailed{
		Uuid:                v.UUID,
		Name:                v.Name,
		Label:               DomainMenuLabelToAPIMenuLabel(v.Label),
		Ingredients:         AnyPtr(DomainIngredientsToAPIIngredientsDetailed(v.Ingredients)),
		Recipes:             AnyPtr(DomainRecipesToAPIRecipesDetailed(v.Recipes)),
		NutritionFactsTotal: AnyPtr(DomainNutritionFactsToAPINutritionFacts(v.NutritionFactsTotal)),
	}
}

func DomainMenusToAPIMenusDetailed(v []*domain.Menu) api.MenusDetailed {
	return slices.Map(v, DomainMenuToAPIMenuDetailed)
}

func DomainDailyMenuToAPIDailyMenu(v *domain.DailyMenu) api.DailyMenu {
	if v == nil {
		return api.DailyMenu{}
	}
	return api.DailyMenu{
		Name: v.Name,
		Uuid: v.UUID,
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
		Uuid:                v.UUID,
		Name:                v.Name,
		Menus:               AnyPtr(DomainMenusToAPIMenusDetailed(v.Menus)),
		NutritionFactsTotal: AnyPtr(DomainNutritionFactsToAPINutritionFacts(v.NutritionFactsTotal)),
	}
}

func DomainDailyMenusToAPIDailyMenusDetailed(v []*domain.DailyMenu) api.DailyMenusDetailed {
	return slices.Map(v, DomainDailyMenuToAPIDailyMenuDetailed)
}
