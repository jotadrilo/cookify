package seed

import (
	"context"
	"fmt"

	ginctl "github.com/jotadrilo/cookify/app/adapters/controllers/gin"
)

func Run(ctx context.Context, ctl *ginctl.Controller) error {
	var (
		usersUC = ctl.Users
		//productsUC   = ctl.ProductsController.Products
		//recipesUC    = ctl.RecipesController.Recipes
		//menusUC      = ctl.MenusController.Menus
		//dailyMenusUC = ctl.DailyMenusController.DailyMenus
	)

	//for ix, product := range products {
	//	p, err := productsUC.CreateProduct(ctx, product)
	//	if err != nil {
	//		return fmt.Errorf("error creating product: %v", err)
	//	}
	//	products[ix].UUID = p.UUID
	//}

	for ix, user := range users {
		p, err := usersUC.CreateUser(ctx, user)
		if err != nil {
			return fmt.Errorf("error creating user: %v", err)
		}
		users[ix].UUID = p.UUID
	}

	//for ix, recipe := range recipes {
	//	r, err := recipesUC.CreateUserRecipe(ctx, recipe.UserUUID, recipe)
	//	if err != nil {
	//		return fmt.Errorf("error creating recipe: %v", err)
	//	}
	//	recipes[ix].UUID = r.UUID
	//}
	//
	//for ix, menu := range menus {
	//	m, err := menusUC.CreateUserMenu(ctx, menu.UserUUID, menu)
	//	if err != nil {
	//		return fmt.Errorf("error creating menu: %v", err)
	//	}
	//	menus[ix].UUID = m.UUID
	//}
	//
	//for ix, dailyMenu := range dailyMenus {
	//	m, err := dailyMenusUC.CreateUserDailyMenu(ctx, dailyMenu.UserUUID, dailyMenu)
	//	if err != nil {
	//		return fmt.Errorf("error creating menu: %v", err)
	//	}
	//	dailyMenus[ix].UUID = m.UUID
	//}

	return nil
}
