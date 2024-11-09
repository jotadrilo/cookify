package products

import (
	"github.com/spf13/cobra"
)

func newImportCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "import",
		Short: "Import data from vendors",
	}

	cmd.AddCommand(
		newImportMercadonaCmd(),
	)

	return cmd
}
