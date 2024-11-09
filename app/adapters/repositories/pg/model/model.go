package model

import (
	"context"

	"github.com/uptrace/bun"
)

////go:embed migrations/*.sql
//var migrations embed.FS

func Register(ctx context.Context, db *bun.DB) error {
	db.RegisterModel(
		(*MenuRecipe)(nil),
		(*DailyMenuMenu)(nil),
	)

	var tables = []any{
		(*NutritionFacts)(nil),
		(*Product)(nil),
		(*Ingredient)(nil),
		(*Recipe)(nil),
		(*Menu)(nil),
		(*MenuRecipe)(nil),
		(*DailyMenu)(nil),
		(*DailyMenuMenu)(nil),
	}

	for _, t := range tables {
		_, err := db.NewCreateTable().
			Model(t).
			IfNotExists().
			Exec(ctx)
		if err != nil {
			return err
		}
	}

	type indexValue struct {
		Model   any
		Columns []string
		Unique  bool
	}

	var indexes = map[string]indexValue{
		// A user cannot add the same product name and vendor more than once
		"idx_products_name_vendor": {
			Model:   (*Product)(nil),
			Columns: []string{"name", "vendor"},
			Unique:  true,
		},
		// For efficient lookups
		"idx_ingredients_recipe_uuid": {
			Model:   (*Ingredient)(nil),
			Columns: []string{"recipe_uuid"},
		},
		"idx_ingredients_menu_uuid": {
			Model:   (*Ingredient)(nil),
			Columns: []string{"menu_uuid"},
		},
		"idx_ingredients_product_uuid": {
			Model:   (*Ingredient)(nil),
			Columns: []string{"product_uuid"},
		},
		// A user cannot add the same ingredient to a recipe more than once
		"idx_ingredients_product_uuid_recipe_uuid": {
			Model:   (*Ingredient)(nil),
			Columns: []string{"product_uuid", "recipe_uuid"},
			Unique:  true,
		},
		// For efficient lookups
		"idx_ingredients_product_uuid_menu_uuid": {
			Model:   (*Ingredient)(nil),
			Columns: []string{"product_uuid", "menu_uuid"},
		},
	}

	for idx, v := range indexes {
		create := db.NewCreateIndex().
			Model(v.Model).
			Index(idx).
			Column(v.Columns...).
			IfNotExists()

		if v.Unique {
			create.Unique()
		}

		if _, err := create.Exec(ctx); err != nil {
			return err
		}
	}

	//m := migrate.NewMigrations()
	//if err := m.Discover(migrations); err != nil {
	//	return err
	//}
	//
	//// Initialize migration tables
	//migrator := migrate.NewMigrator(db, m)
	//
	//if err := migrator.Init(ctx); err != nil {
	//	return err
	//}
	//
	//// Run migrations
	//mg, err := migrator.Migrate(ctx)
	//if err != nil {
	//	return err
	//}
	//
	//if !mg.IsZero() {
	//	logger.Infof("Migrations applied successfully: %s", mg.String())
	//}

	return nil
}
