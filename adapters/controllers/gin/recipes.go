package gin

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/jotadrilo/cookify/api"
	"github.com/jotadrilo/cookify/core/domain"
	"github.com/jotadrilo/cookify/core/ports"
	"github.com/jotadrilo/cookify/internal/errorutils"
)

type RecipesController struct {
	Recipes ports.RecipesUseCase
}

func NewRecipesController(recipes ports.RecipesUseCase) *RecipesController {
	return &RecipesController{
		Recipes: recipes,
	}
}

func (x *RecipesController) CreateRecipe(c *gin.Context, userID api.UserID) {
	var (
		ctx    = c.Request.Context()
		Recipe domain.Recipe
	)

	if err := c.BindJSON(&Recipe); err != nil {
		RestErrorHandler(c, err)
		return
	}

	if err := Recipe.Validate(); err != nil {
		RestErrorHandler(c, errorutils.NewErrBadRequest(err.Error()))
		return
	}

	p, err := x.Recipes.CreateUserRecipe(ctx, userID.String(), &Recipe)
	if err != nil {
		RestErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusCreated, p)
}

func (x *RecipesController) GetUsersUserIdRecipes(c *gin.Context, _ api.UserID) {
	var ctx = c.Request.Context()

	items, err := x.Recipes.ListRecipes(ctx)
	if err != nil {
		RestErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusOK, DomainRecipesToAPIRecipes(items))
}

func (x *RecipesController) GetUsersUserIdRecipesRecipeId(c *gin.Context, _ api.UserID, recipeID api.RecipeID) {
	var ctx = c.Request.Context()

	v, err := x.Recipes.GetRecipeByUUID(ctx, recipeID.String())
	if err != nil {
		RestErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusOK, DomainRecipeToAPIRecipeDetailed(v))
}
