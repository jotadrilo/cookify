package gin

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/jotadrilo/cookify/app/api"
	"github.com/jotadrilo/cookify/internal/errorutils"
	"github.com/jotadrilo/cookify/internal/logger"
	"github.com/jotadrilo/cookify/internal/oapi"
)

func (x *Controller) GetUsersParamUserIDRecipes(c *gin.Context, _ api.ParamUserID) {
	var ctx = c.Request.Context()

	items, err := x.Recipes.ListRecipes(ctx)
	if err != nil {
		RestErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusOK, DomainRecipesToAPIRecipes(items))
}

func (x *Controller) PostUsersParamUserIDRecipes(c *gin.Context, userID api.ParamUserID) {
	var (
		ctx  = c.Request.Context()
		body api.PostUsersParamUserIDRecipesJSONRequestBody
	)

	if err := c.BindJSON(&body); err != nil {
		RestErrorHandler(c, err)
		return
	}

	logger.Infof("Received %+v", body)

	var product = RecipePostRequestToDomainRecipe(&body)

	if err := product.Validate(); err != nil {
		RestErrorHandler(c, errorutils.NewErrBadRequest(err.Error()))
		return
	}

	item, err := x.Recipes.CreateUserRecipe(ctx, userID.String(), product)
	if err != nil {
		RestErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusCreated, &api.RecipeID{Uuid: oapi.UUID(item.UUID)})
}

func (x *Controller) GetUsersParamUserIDRecipesParamRecipeID(c *gin.Context, _ api.ParamUserID, recipeID api.ParamRecipeID) {
	var ctx = c.Request.Context()

	v, err := x.Recipes.GetRecipeByUUID(ctx, recipeID.String())
	if err != nil {
		RestErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusOK, DomainRecipeToAPIRecipe(v))
}

func (x *Controller) PatchUsersParamUserIDRecipesParamRecipeID(c *gin.Context, userID api.ParamUserID, recipeID api.ParamRecipeID) {
	var (
		ctx  = c.Request.Context()
		body api.PatchUsersParamUserIDRecipesParamRecipeIDJSONRequestBody
	)

	if err := c.BindJSON(&body); err != nil {
		RestErrorHandler(c, err)
		return
	}

	logger.Infof("Received %+v", body)

	var newRecipe = RecipePatchRequestToDomainRecipe(&body)

	curRecipe, err := x.Recipes.GetUserRecipeByUUID(ctx, userID.String(), recipeID.String())
	if err != nil {
		RestErrorHandler(c, err)
		return
	}

	var update bool

	if n := len(newRecipe.Ingredients); n > 0 && n != len(curRecipe.Ingredients) {
		curRecipe.Ingredients = newRecipe.Ingredients
		update = true
	}

	if n := newRecipe.Name; n != "" && n != curRecipe.Name {
		curRecipe.Name = newRecipe.Name
		update = true
	}

	if !update {
		c.AbortWithStatus(http.StatusNotModified)
		return
	}

	if err := curRecipe.Validate(); err != nil {
		RestErrorHandler(c, errorutils.NewErrBadRequest(err.Error()))
		return
	}

	if _, err := x.Recipes.UpdateUserRecipeByUUID(ctx, userID.String(), recipeID.String(), curRecipe); err != nil {
		RestErrorHandler(c, err)
		return
	}

	c.AbortWithStatus(http.StatusNoContent)
}
