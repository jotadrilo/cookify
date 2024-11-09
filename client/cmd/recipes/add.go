package recipes

import (
	"encoding/json"
	"fmt"
	"net/http"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/google/uuid"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/spf13/cobra"

	"github.com/jotadrilo/cookify/app/api"
	"github.com/jotadrilo/cookify/client/cmd/config"
	"github.com/jotadrilo/cookify/client/ui/recipes/add"
	"github.com/jotadrilo/cookify/internal/logger"
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

			m, err := tea.NewProgram(add.NewModel(
				add.WithModelProducts(products...),
			)).Run()
			if err != nil {
				return err
			}

			var recipe = api.PostUsersParamUserIDRecipesJSONRequestBody{}

			if v, ok := m.(*add.Model); ok {
				recipe.Name = v.GetName()

				var (
					selectedProducts   = v.GetSelectedProducts()
					selectedQuantities = v.GetSelectedQuantities()
				)

				for ix, product := range selectedProducts {
					var quantity = selectedQuantities[ix]

					recipe.Ingredients = append(recipe.Ingredients, api.Ingredient{
						Product:  *product,
						Quantity: quantity,
					})
				}
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
