package config

import "github.com/spf13/cobra"

type CookifyServer struct {
	Address string
}

func BindCookifyServerFlags(cmd *cobra.Command, cfg *CookifyServer) {
	cmd.Flags().StringVar(&cfg.Address, "cookify-server-api-address", "http://localhost:8080/api/v1", "cookify server API address")
}
