// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oapi-codegen/runtime"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (GET /admin/users)
	GetAdminUsers(c *gin.Context)

	// (GET /products)
	GetProducts(c *gin.Context)

	// (POST /products)
	PostProducts(c *gin.Context)

	// (GET /products/{ParamProductID})
	GetProductsParamProductID(c *gin.Context, paramProductID ParamProductID)

	// (GET /products/{ParamProductID}/equivalents)
	GetProductsParamProductIDEquivalents(c *gin.Context, paramProductID ParamProductID)

	// (GET /users/{ParamUserID})
	GetUsersParamUserID(c *gin.Context, paramUserID ParamUserID)

	// (GET /users/{ParamUserID}/daily-menus)
	GetUsersParamUserIDDailyMenus(c *gin.Context, paramUserID ParamUserID)

	// (GET /users/{ParamUserID}/daily-menus/{ParamDailyMenuID})
	GetUsersParamUserIDDailyMenusParamDailyMenuID(c *gin.Context, paramUserID ParamUserID, paramDailyMenuID ParamDailyMenuID)

	// (GET /users/{ParamUserID}/menus)
	GetUsersParamUserIDMenus(c *gin.Context, paramUserID ParamUserID)

	// (GET /users/{ParamUserID}/menus/{ParamMenuID})
	GetUsersParamUserIDMenusParamMenuID(c *gin.Context, paramUserID ParamUserID, paramMenuID ParamMenuID)

	// (GET /users/{ParamUserID}/recipes)
	GetUsersParamUserIDRecipes(c *gin.Context, paramUserID ParamUserID)

	// (POST /users/{ParamUserID}/recipes)
	PostUsersParamUserIDRecipes(c *gin.Context, paramUserID ParamUserID)

	// (GET /users/{ParamUserID}/recipes/{ParamRecipeID})
	GetUsersParamUserIDRecipesParamRecipeID(c *gin.Context, paramUserID ParamUserID, paramRecipeID ParamRecipeID)

	// (PATCH /users/{ParamUserID}/recipes/{ParamRecipeID})
	PatchUsersParamUserIDRecipesParamRecipeID(c *gin.Context, paramUserID ParamUserID, paramRecipeID ParamRecipeID)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandler       func(*gin.Context, error, int)
}

type MiddlewareFunc func(c *gin.Context)

// GetAdminUsers operation middleware
func (siw *ServerInterfaceWrapper) GetAdminUsers(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetAdminUsers(c)
}

// GetProducts operation middleware
func (siw *ServerInterfaceWrapper) GetProducts(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetProducts(c)
}

// PostProducts operation middleware
func (siw *ServerInterfaceWrapper) PostProducts(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PostProducts(c)
}

// GetProductsParamProductID operation middleware
func (siw *ServerInterfaceWrapper) GetProductsParamProductID(c *gin.Context) {

	var err error

	// ------------- Path parameter "ParamProductID" -------------
	var paramProductID ParamProductID

	err = runtime.BindStyledParameterWithOptions("simple", "ParamProductID", c.Param("ParamProductID"), &paramProductID, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter ParamProductID: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetProductsParamProductID(c, paramProductID)
}

// GetProductsParamProductIDEquivalents operation middleware
func (siw *ServerInterfaceWrapper) GetProductsParamProductIDEquivalents(c *gin.Context) {

	var err error

	// ------------- Path parameter "ParamProductID" -------------
	var paramProductID ParamProductID

	err = runtime.BindStyledParameterWithOptions("simple", "ParamProductID", c.Param("ParamProductID"), &paramProductID, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter ParamProductID: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetProductsParamProductIDEquivalents(c, paramProductID)
}

// GetUsersParamUserID operation middleware
func (siw *ServerInterfaceWrapper) GetUsersParamUserID(c *gin.Context) {

	var err error

	// ------------- Path parameter "ParamUserID" -------------
	var paramUserID ParamUserID

	err = runtime.BindStyledParameterWithOptions("simple", "ParamUserID", c.Param("ParamUserID"), &paramUserID, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter ParamUserID: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetUsersParamUserID(c, paramUserID)
}

// GetUsersParamUserIDDailyMenus operation middleware
func (siw *ServerInterfaceWrapper) GetUsersParamUserIDDailyMenus(c *gin.Context) {

	var err error

	// ------------- Path parameter "ParamUserID" -------------
	var paramUserID ParamUserID

	err = runtime.BindStyledParameterWithOptions("simple", "ParamUserID", c.Param("ParamUserID"), &paramUserID, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter ParamUserID: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetUsersParamUserIDDailyMenus(c, paramUserID)
}

// GetUsersParamUserIDDailyMenusParamDailyMenuID operation middleware
func (siw *ServerInterfaceWrapper) GetUsersParamUserIDDailyMenusParamDailyMenuID(c *gin.Context) {

	var err error

	// ------------- Path parameter "ParamUserID" -------------
	var paramUserID ParamUserID

	err = runtime.BindStyledParameterWithOptions("simple", "ParamUserID", c.Param("ParamUserID"), &paramUserID, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter ParamUserID: %w", err), http.StatusBadRequest)
		return
	}

	// ------------- Path parameter "ParamDailyMenuID" -------------
	var paramDailyMenuID ParamDailyMenuID

	err = runtime.BindStyledParameterWithOptions("simple", "ParamDailyMenuID", c.Param("ParamDailyMenuID"), &paramDailyMenuID, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter ParamDailyMenuID: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetUsersParamUserIDDailyMenusParamDailyMenuID(c, paramUserID, paramDailyMenuID)
}

// GetUsersParamUserIDMenus operation middleware
func (siw *ServerInterfaceWrapper) GetUsersParamUserIDMenus(c *gin.Context) {

	var err error

	// ------------- Path parameter "ParamUserID" -------------
	var paramUserID ParamUserID

	err = runtime.BindStyledParameterWithOptions("simple", "ParamUserID", c.Param("ParamUserID"), &paramUserID, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter ParamUserID: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetUsersParamUserIDMenus(c, paramUserID)
}

// GetUsersParamUserIDMenusParamMenuID operation middleware
func (siw *ServerInterfaceWrapper) GetUsersParamUserIDMenusParamMenuID(c *gin.Context) {

	var err error

	// ------------- Path parameter "ParamUserID" -------------
	var paramUserID ParamUserID

	err = runtime.BindStyledParameterWithOptions("simple", "ParamUserID", c.Param("ParamUserID"), &paramUserID, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter ParamUserID: %w", err), http.StatusBadRequest)
		return
	}

	// ------------- Path parameter "ParamMenuID" -------------
	var paramMenuID ParamMenuID

	err = runtime.BindStyledParameterWithOptions("simple", "ParamMenuID", c.Param("ParamMenuID"), &paramMenuID, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter ParamMenuID: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetUsersParamUserIDMenusParamMenuID(c, paramUserID, paramMenuID)
}

// GetUsersParamUserIDRecipes operation middleware
func (siw *ServerInterfaceWrapper) GetUsersParamUserIDRecipes(c *gin.Context) {

	var err error

	// ------------- Path parameter "ParamUserID" -------------
	var paramUserID ParamUserID

	err = runtime.BindStyledParameterWithOptions("simple", "ParamUserID", c.Param("ParamUserID"), &paramUserID, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter ParamUserID: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetUsersParamUserIDRecipes(c, paramUserID)
}

// PostUsersParamUserIDRecipes operation middleware
func (siw *ServerInterfaceWrapper) PostUsersParamUserIDRecipes(c *gin.Context) {

	var err error

	// ------------- Path parameter "ParamUserID" -------------
	var paramUserID ParamUserID

	err = runtime.BindStyledParameterWithOptions("simple", "ParamUserID", c.Param("ParamUserID"), &paramUserID, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter ParamUserID: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PostUsersParamUserIDRecipes(c, paramUserID)
}

// GetUsersParamUserIDRecipesParamRecipeID operation middleware
func (siw *ServerInterfaceWrapper) GetUsersParamUserIDRecipesParamRecipeID(c *gin.Context) {

	var err error

	// ------------- Path parameter "ParamUserID" -------------
	var paramUserID ParamUserID

	err = runtime.BindStyledParameterWithOptions("simple", "ParamUserID", c.Param("ParamUserID"), &paramUserID, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter ParamUserID: %w", err), http.StatusBadRequest)
		return
	}

	// ------------- Path parameter "ParamRecipeID" -------------
	var paramRecipeID ParamRecipeID

	err = runtime.BindStyledParameterWithOptions("simple", "ParamRecipeID", c.Param("ParamRecipeID"), &paramRecipeID, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter ParamRecipeID: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetUsersParamUserIDRecipesParamRecipeID(c, paramUserID, paramRecipeID)
}

// PatchUsersParamUserIDRecipesParamRecipeID operation middleware
func (siw *ServerInterfaceWrapper) PatchUsersParamUserIDRecipesParamRecipeID(c *gin.Context) {

	var err error

	// ------------- Path parameter "ParamUserID" -------------
	var paramUserID ParamUserID

	err = runtime.BindStyledParameterWithOptions("simple", "ParamUserID", c.Param("ParamUserID"), &paramUserID, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter ParamUserID: %w", err), http.StatusBadRequest)
		return
	}

	// ------------- Path parameter "ParamRecipeID" -------------
	var paramRecipeID ParamRecipeID

	err = runtime.BindStyledParameterWithOptions("simple", "ParamRecipeID", c.Param("ParamRecipeID"), &paramRecipeID, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter ParamRecipeID: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PatchUsersParamUserIDRecipesParamRecipeID(c, paramUserID, paramRecipeID)
}

// GinServerOptions provides options for the Gin server.
type GinServerOptions struct {
	BaseURL      string
	Middlewares  []MiddlewareFunc
	ErrorHandler func(*gin.Context, error, int)
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router gin.IRouter, si ServerInterface) {
	RegisterHandlersWithOptions(router, si, GinServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router gin.IRouter, si ServerInterface, options GinServerOptions) {
	errorHandler := options.ErrorHandler
	if errorHandler == nil {
		errorHandler = func(c *gin.Context, err error, statusCode int) {
			c.JSON(statusCode, gin.H{"msg": err.Error()})
		}
	}

	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandler:       errorHandler,
	}

	router.GET(options.BaseURL+"/admin/users", wrapper.GetAdminUsers)
	router.GET(options.BaseURL+"/products", wrapper.GetProducts)
	router.POST(options.BaseURL+"/products", wrapper.PostProducts)
	router.GET(options.BaseURL+"/products/:ParamProductID", wrapper.GetProductsParamProductID)
	router.GET(options.BaseURL+"/products/:ParamProductID/equivalents", wrapper.GetProductsParamProductIDEquivalents)
	router.GET(options.BaseURL+"/users/:ParamUserID", wrapper.GetUsersParamUserID)
	router.GET(options.BaseURL+"/users/:ParamUserID/daily-menus", wrapper.GetUsersParamUserIDDailyMenus)
	router.GET(options.BaseURL+"/users/:ParamUserID/daily-menus/:ParamDailyMenuID", wrapper.GetUsersParamUserIDDailyMenusParamDailyMenuID)
	router.GET(options.BaseURL+"/users/:ParamUserID/menus", wrapper.GetUsersParamUserIDMenus)
	router.GET(options.BaseURL+"/users/:ParamUserID/menus/:ParamMenuID", wrapper.GetUsersParamUserIDMenusParamMenuID)
	router.GET(options.BaseURL+"/users/:ParamUserID/recipes", wrapper.GetUsersParamUserIDRecipes)
	router.POST(options.BaseURL+"/users/:ParamUserID/recipes", wrapper.PostUsersParamUserIDRecipes)
	router.GET(options.BaseURL+"/users/:ParamUserID/recipes/:ParamRecipeID", wrapper.GetUsersParamUserIDRecipesParamRecipeID)
	router.PATCH(options.BaseURL+"/users/:ParamUserID/recipes/:ParamRecipeID", wrapper.PatchUsersParamUserIDRecipesParamRecipeID)
}
