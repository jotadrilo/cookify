package recipes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"

	"github.com/AlecAivazis/survey/v2"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/spf13/cobra"

	"github.com/jotadrilo/cookify/app/api"
	"github.com/jotadrilo/cookify/client/cmd/config"
	"github.com/jotadrilo/cookify/internal/logger"
	"github.com/jotadrilo/cookify/internal/oapi"
	"github.com/jotadrilo/cookify/internal/prompt"
	"github.com/jotadrilo/cookify/internal/slices"
)

func newUpdateCmd() *cobra.Command {
	type cmdConfig struct {
		Cookify config.CookifyServer
		User    config.CookifyUser
	}

	cfg := &cmdConfig{}

	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update a new recipe",
		RunE: func(cmd *cobra.Command, args []string) error {
			var ctx = cmd.Context()

			apiCli, err := api.NewClientWithResponses(cfg.Cookify.Address)
			if err != nil {
				return err
			}

			var userID api.ID

			if rsp, err := apiCli.GetUsersParamUserIDWithResponse(ctx, *oapi.UUID(cfg.User.UUID)); err != nil {
				return err
			} else {
				switch rsp.StatusCode() {
				case http.StatusOK:
					userID = *rsp.JSON200.Uuid
				default:
					var apiError api.Error
					if err := json.Unmarshal(rsp.Body, &apiError); err != nil {
						return err
					}
					return fmt.Errorf("user %q cannot be obtained: (%d - %s - %s)", cfg.User.UUID, rsp.StatusCode(), rsp.Status(), apiError.Title)
				}
			}

			var recipes []api.Recipe

			if rsp, err := apiCli.GetUsersParamUserIDRecipesWithResponse(ctx, userID); err != nil {
				return err
			} else {
				switch rsp.StatusCode() {
				case http.StatusOK:
					recipes = *rsp.JSON200
				default:
					var apiError api.Error
					if err := json.Unmarshal(rsp.Body, &apiError); err != nil {
						return err
					}
					return fmt.Errorf("recipes cannot be obtained: (%d - %s - %s)", rsp.StatusCode(), rsp.Status(), apiError.Title)
				}
			}

			var (
				recipeSlugs   []string
				recipesBySlug = make(map[string]*api.Recipe, len(recipes))
			)

			for _, recipe := range recipes {
				var slug string
				slug = fmt.Sprintf("%s (%s)", text.FgGreen.Sprint(recipe.Name), text.FgYellow.Sprint(recipe.Uuid.String()))
				recipeSlugs = append(recipeSlugs, slug)
				recipesBySlug[slug] = &recipe
			}

			sort.Strings(recipeSlugs)

			var selectedRecipeSlug string

			prompt.MustAsk(&selectedRecipeSlug, prompt.WithAskPrompt(&survey.Select{
				Message:  "Select the recipe to update:",
				Options:  recipeSlugs,
				PageSize: 50,
			}), prompt.WithAskKeepFilter())

			if selectedRecipeSlug == "" {
				logger.Warnf("You didn't select a recipe. Exiting now!")
				return nil
			}

			var selectedRecipe = recipesBySlug[selectedRecipeSlug]

			var (
				curRecipe *api.Recipe
				recipeID  = *oapi.UUID(selectedRecipe.Uuid.String())
			)

			if rsp, err := apiCli.GetUsersParamUserIDRecipesParamRecipeIDWithResponse(ctx, userID, recipeID); err != nil {

				return err
			} else {
				switch rsp.StatusCode() {
				case http.StatusOK:
					curRecipe = rsp.JSON200
				default:
					var apiError api.Error
					if err := json.Unmarshal(rsp.Body, &apiError); err != nil {
						return err
					}
					return fmt.Errorf("recipe %q cannot be obtained: (%d - %s - %s)", selectedRecipe.Uuid.String(), rsp.StatusCode(), rsp.Status(), apiError.Title)
				}
			}

			var (
				recipeNewName        string
				recipeNewIngredients []api.Ingredient
			)

			if prompt.MustAskConfirm(prompt.WithAskConfirmMessage("Do you want to update the name?")) {
				prompt.MustAsk(&recipeNewName, prompt.WithAskPrompt(&survey.Input{
					Message: "Enter the new recipe name or leave it intact:",
					Default: curRecipe.Name,
				}))
			}

			if prompt.MustAskConfirm(prompt.WithAskConfirmMessage("Do you want to update the ingredients?")) {
				var currentProductSlugs = slices.Map(curRecipe.Ingredients, func(v api.Ingredient) string {
					return v.Product.Slug
				})

				if prompt.MustAskConfirm(prompt.WithAskConfirmMessage("Do you want to remove ingredients?")) {

				}

				if prompt.MustAskConfirm(prompt.WithAskConfirmMessage("Do you want to add ingredients?")) {
					recipeNewIngredients = curRecipe.Ingredients

					var products []api.Product

					if rsp, err := apiCli.GetProductsWithResponse(ctx); err != nil {
						return err
					} else {
						switch rsp.StatusCode() {
						case http.StatusOK:
							products = *rsp.JSON200
						default:
							var apiError api.Error
							if err := json.Unmarshal(rsp.Body, &apiError); err != nil {
								return err
							}
							return fmt.Errorf("products cannot be obtained: (%d - %s - %s)", rsp.StatusCode(), rsp.Status(), apiError.Title)
						}
					}

					var (
						productSlugs   []string
						productsBySlug = make(map[string]*api.Product, len(products))
					)

					for _, prod := range products {
						if slices.Contains(currentProductSlugs, prod.Slug) {
							continue
						}

						var slug string
						if prod.DisplayNameLangEsEs != nil {
							slug = fmt.Sprintf("%s (%s - %s)", text.FgGreen.Sprint(prod.Slug), text.FgYellow.Sprint(prod.Uuid), *prod.DisplayNameLangEsEs)
						} else {
							slug = fmt.Sprintf("%s (%s)", text.FgGreen.Sprint(prod.Slug), text.FgYellow.Sprint(prod.Uuid))
						}

						productSlugs = append(productSlugs, slug)
						productsBySlug[slug] = &prod
					}

					sort.Strings(productSlugs)

					var selectedProductSlugs []string

					prompt.MustAsk(&selectedProductSlugs, prompt.WithAskPrompt(&survey.MultiSelect{
						Message:  "Select the recipe ingredients:",
						Options:  productSlugs,
						PageSize: 50,
					}), prompt.WithAskKeepFilter())

					if len(selectedProductSlugs) == 0 {
						return fmt.Errorf("recipe ingredients cannot be empty")
					}

					var selectedProducts []*api.Product

					for _, slug := range selectedProductSlugs {
						selectedProducts = append(selectedProducts, productsBySlug[slug])
					}

					for _, prod := range selectedProducts {
						var p *api.Product

						rsp, err := apiCli.GetProductsParamProductIDWithResponse(ctx, *prod.Uuid)
						if err != nil {
							return err
						} else {
							switch rsp.StatusCode() {
							case http.StatusOK:
								p = rsp.JSON200
							default:
								var apiError api.Error
								if err := json.Unmarshal(rsp.Body, &apiError); err != nil {
									return err
								}
								return fmt.Errorf("product %q cannot be obtained: (%d - %s - %s)", *prod.Uuid, rsp.StatusCode(), rsp.Status(), apiError.Title)
							}
						}

						var ingredient = api.Ingredient{Product: *p}

						var q float32
						prompt.MustAskFloat32(&q, prompt.WithAskPrompt(&survey.Input{
							Message: fmt.Sprintf("Enter the quantity for %q product (%s):", p.Slug, *p.Uuid),
							Default: "0.0",
						}))

						if q == 0 {
							return fmt.Errorf("product quantity cannot be zero")
						}

						ingredient.Quantity = float32(q)

						recipeNewIngredients = append(recipeNewIngredients, ingredient)
					}
				}
			}

			var recipe api.PatchUsersParamUserIDRecipesParamRecipeIDJSONRequestBody

			if recipeNewName != "" {
				recipe.Name = oapi.String(recipeNewName)
			}

			if len(recipeNewIngredients) > 0 {
				recipe.Ingredients = &recipeNewIngredients
			}

			rsp, err := apiCli.PatchUsersParamUserIDRecipesParamRecipeIDWithResponse(ctx, userID, recipeID, recipe)
			if err != nil {
				return err
			}

			switch rsp.StatusCode() {
			case http.StatusNoContent:
				logger.Infof("Recipe %q was updated: %s", curRecipe.Name, recipeID.String())
			case http.StatusNotModified:
				logger.Infof("Recipe %q was not updated: %s", curRecipe.Name, recipeID.String())
			default:
				var apiError api.Error
				if err := json.Unmarshal(rsp.Body, &apiError); err != nil {
					return err
				}
				return fmt.Errorf("recipe %q was not updated: (%d - %s - %s)", curRecipe.Name, rsp.StatusCode(), rsp.Status(), apiError.Title)
			}

			return nil
		},
	}

	config.BindCookifyServerFlags(cmd, &cfg.Cookify)
	config.BindCookifyUserFlags(cmd, &cfg.User)

	return cmd
}
