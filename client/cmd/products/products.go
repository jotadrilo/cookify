package products

import (
	"github.com/spf13/cobra"
)

func NewProductsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "products",
		Short: "Products operations",
	}

	cmd.AddCommand(
		newImportCmd(),
	)

	return cmd
}
