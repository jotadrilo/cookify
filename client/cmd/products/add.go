package products

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"

	"github.com/jotadrilo/cookify/app/adapters/controllers/gin"
	"github.com/jotadrilo/cookify/app/api"
	"github.com/jotadrilo/cookify/app/core/domain"
	"github.com/jotadrilo/cookify/client/cmd/config"
	"github.com/jotadrilo/cookify/internal/logger"
	"github.com/jotadrilo/cookify/internal/oapi"
	"github.com/jotadrilo/cookify/internal/prompt"
)

func newAddCmd() *cobra.Command {
	type cmdConfig struct {
		Cookify config.CookifyServer
	}

	cfg := &cmdConfig{}

	cmd := &cobra.Command{
		Use:   "add",
		Short: "Add a new product",
		RunE: func(cmd *cobra.Command, args []string) error {
			var ctx = cmd.Context()

			apiCli, err := api.NewClientWithResponses(cfg.Cookify.Address)
			if err != nil {
				return err
			}

			var (
				displayNameEsES        string
				displayNameEnEN        string
				price                  float32
				priceQuantityReference string
				brand                  string
			)

			prompt.MustAsk(&displayNameEsES, prompt.WithAskPrompt(&survey.Input{
				Message: "Enter the display name (es-ES):",
				Default: "",
			}))

			prompt.MustAsk(&displayNameEnEN, prompt.WithAskPrompt(&survey.Input{
				Message: "Enter the display name (en-EN):",
				Default: "",
			}))

			prompt.MustAskFloat32(&price, prompt.WithAskPrompt(&survey.Input{
				Message: "Enter the price (ex. 1.50):",
				Default: "0.0",
			}))

			prompt.MustAsk(&priceQuantityReference, prompt.WithAskPrompt(&survey.Input{
				Message: "Enter the price quantity reference (ex. 1 kg):",
				Default: "1 kg",
			}))

			prompt.MustAsk(&brand, prompt.WithAskPrompt(&survey.Input{
				Message: "Enter the branch:",
				Default: "",
			}))

			uh, err := domain.ParseQuantityString(priceQuantityReference)
			if err != nil {
				return err
			}

			var body = api.PostProductsJSONRequestBody{
				DisplayNameLangEsEs: oapi.String(displayNameEsES),
				DisplayNameLangEnUs: oapi.String(displayNameEnEN),
				Unit:                gin.DomainUnitToAPIUnit(uh.Unit()),
				CurrentPrice: &api.Price{
					Price:    price,
					Quantity: uh.Value(),
				},
				Vendors: oapi.StringSlice([]string{brand}),
			}

			body.Slug = strings.ReplaceAll(strings.ToLower(displayNameEsES), " ", "-")

			bindNutritionValues(&body)

			rsp, err := apiCli.PostProductsWithResponse(ctx, body)
			if err != nil {
				return err
			}

			switch rsp.StatusCode() {
			case http.StatusCreated:
				var id api.ProductID
				if err := json.Unmarshal(rsp.Body, &id); err != nil {
					return err
				}
				logger.Infof("Product %q created: %s", body.Slug, id.Uuid.String())
			case http.StatusConflict:
				var id api.ProductID
				if err := json.Unmarshal(rsp.Body, &id); err == nil {
					logger.Infof("Product %q already exists: %s", body.Slug, id.Uuid.String())
					break
				}

				var apiError api.Error
				if err := json.Unmarshal(rsp.Body, &apiError); err != nil {
					return err
				}
				logger.Warnf("Product %q was not created: (%d - %s - %s)", body.Slug, rsp.StatusCode(), rsp.Status(), apiError.Title)
			default:
				var apiError api.Error
				if err := json.Unmarshal(rsp.Body, &apiError); err != nil {
					return err
				}
				return fmt.Errorf("product %q was not created: (%d - %s - %s)", body.Slug, rsp.StatusCode(), rsp.Status(), apiError.Title)
			}

			return nil
		},
	}

	config.BindCookifyServerFlags(cmd, &cfg.Cookify)

	return cmd
}
