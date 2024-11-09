package model

type ProductToUnit struct {
	ProductID int64    `bun:",pk"`
	Product   *Product `bun:"rel:belongs-to,join:product_id=id"`
	UnitID    int64    `bun:",pk"`
	Unit      *Unit    `bun:"rel:belongs-to,join:unit_id=id"`
}
