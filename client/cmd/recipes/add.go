package recipes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"

	"github.com/AlecAivazis/survey/v2"
	"github.com/google/uuid"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/spf13/cobra"

	"github.com/jotadrilo/cookify/app/api"
	"github.com/jotadrilo/cookify/client/cmd/config"
	"github.com/jotadrilo/cookify/internal/logger"
	"github.com/jotadrilo/cookify/internal/prompt"
)

func newAddCmd() *cobra.Command {
	type cmdConfig struct {
		Cookify config.CookifyServer
		User    config.CookifyUser
	}

	cfg := &cmdConfig{}

	cmd := &cobra.Command{
		Use:   "add",
		Short: "Add a new recipe",
		RunE: func(cmd *cobra.Command, args []string) error {
			var ctx = cmd.Context()

			apiCli, err := api.NewClientWithResponses(cfg.Cookify.Address)
			if err != nil {
				return err
			}

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

			var recipe = api.PostUsersParamUserIDRecipesJSONRequestBody{}

			prompt.MustAsk(&recipe.Name, prompt.WithAskPrompt(&survey.Input{
				Message: "Enter the recipe name:",
				Default: "",
			}))

			if recipe.Name == "" {
				return fmt.Errorf("recipe name cannot be empty")
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

				prompt.MustAskFloat32(&ingredient.Quantity, prompt.WithAskPrompt(&survey.Input{
					Message: fmt.Sprintf("Enter the quantity for %q product (%s):", p.Slug, *p.Uuid),
					Default: "0.0",
				}))

				if ingredient.Quantity == 0 {
					return fmt.Errorf("product quantity cannot be zero")
				}

				recipe.Ingredients = append(recipe.Ingredients, ingredient)
			}

			rsp, err := apiCli.PostUsersParamUserIDRecipesWithResponse(ctx, uuid.MustParse(cfg.User.UUID), recipe)
			if err != nil {
				return err
			}

			switch rsp.StatusCode() {
			case http.StatusCreated:
				var id api.ProductID
				if err := json.Unmarshal(rsp.Body, &id); err != nil {
					return err
				}
				logger.Infof("Recipe %q created: %s", recipe.Name, id.Uuid.String())
			case http.StatusConflict:
				var id api.ProductID
				if err := json.Unmarshal(rsp.Body, &id); err == nil {
					logger.Infof("Recipe %q already exists: %s", recipe.Name, id.Uuid.String())
					break
				}

				var apiError api.Error
				if err := json.Unmarshal(rsp.Body, &apiError); err != nil {
					return err
				}
				logger.Warnf("Recipe %q was not created: (%d - %s - %s)", recipe.Name, rsp.StatusCode(), rsp.Status(), apiError.Title)
			default:
				var apiError api.Error
				if err := json.Unmarshal(rsp.Body, &apiError); err != nil {
					return err
				}
				return fmt.Errorf("recipe %q was not created: (%d - %s - %s)", recipe.Name, rsp.StatusCode(), rsp.Status(), apiError.Title)
			}

			return nil
		},
	}

	config.BindCookifyServerFlags(cmd, &cfg.Cookify)
	config.BindCookifyUserFlags(cmd, &cfg.User)

	return cmd
}
