package config

import "github.com/spf13/cobra"

type CookifyUser struct {
	UUID string
}

func BindCookifyUserFlags(cmd *cobra.Command, cfg *CookifyUser) {
	cmd.Flags().StringVar(&cfg.UUID, "cookify-user-uuid", "ecc5fe1d-f8fa-4f20-b353-85b57b4c3d28", "cookify user UUID")
}
