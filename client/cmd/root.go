package cmd

import (
	"flag"

	"github.com/spf13/cobra"

	"github.com/jotadrilo/cookify/client/cmd/products"
	"github.com/jotadrilo/cookify/client/cmd/recipes"
)

func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "cookify",
		Short:             "Cookify Client",
		Long:              `cookify - Controls the Cookify service`,
		SilenceUsage:      true,
		DisableAutoGenTag: true,
	}

	cmd.PersistentFlags().AddGoFlagSet(flag.CommandLine)

	cmd.AddCommand(
		products.NewProductsCommand(),
		recipes.NewRecipesCommand(),
	)

	return cmd
}
