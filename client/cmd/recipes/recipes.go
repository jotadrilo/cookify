package recipes

import (
	"github.com/spf13/cobra"
)

func NewRecipesCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "recipes",
		Short: "Recipes operations",
	}

	cmd.AddCommand(
		newAddCmd(),
		newUpdateCmd(),
	)

	return cmd
}
