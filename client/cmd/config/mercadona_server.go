package config

import "github.com/spf13/cobra"

type MercadonaServer struct {
	Address                    string
	IgnoreCategoriesWithPrefix []string
	TargetProductSlugs         []string
	ZipCode                    uint32
	ProductsFile               string
	ReloadProducts             bool
}

var mercadonaNonEatableCategories = []string{
	"Cuidado facial y corporal",
	"Fitoterapia y parafarmacia",
	"Cuidado del cabello",
	"Limpieza y hogar",
	"Maquillaje",
	"Mascotas",
	"Bebé",
}

func BindMercadonaServerFlags(cmd *cobra.Command, cfg *MercadonaServer) {
	cmd.Flags().StringVar(&cfg.Address, "mercadona-server-api-address", "https://tienda.mercadona.es/api", "server API address")
	cmd.Flags().StringSliceVar(&cfg.IgnoreCategoriesWithPrefix, "mercadona-ignore-categories-with-prefix", mercadonaNonEatableCategories, "ignore mercadona categories starting with this prefix")
	cmd.Flags().StringSliceVar(&cfg.TargetProductSlugs, "mercadona-target-product-slugs", []string{}, "mercadona target product slugs")
	cmd.Flags().Uint32Var(&cfg.ZipCode, "mercadona-zip-code", 0, "zip code to locate the nearby mercadona warehouse")
	cmd.Flags().StringVar(&cfg.ProductsFile, "mercadona-products-file", "products.json", "load/store mercadona products from/in a JSON file")
	cmd.Flags().BoolVar(&cfg.ReloadProducts, "mercadona-reload-products", false, "reload products instad of loading them from the JSON file")
}
