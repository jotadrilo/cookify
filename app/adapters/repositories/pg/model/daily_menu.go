package model

import (
	"context"

	"github.com/google/uuid"
	"github.com/uptrace/bun"

	"github.com/jotadrilo/cookify/app/core/domain"
)

type DailyMenu struct {
	bun.BaseModel `bun:"table:daily_menus"`

	ID int64 `bun:"id,pk,autoincrement"`

	UUID  string  `bun:"uuid,pk"`
	Name  string  `bun:"name"`
	Menus []*Menu `bun:"m2m:daily_menu_menus,join:DailyMenu=Menu"`
}

type DailyMenuMenu struct {
	ID int64 `bun:"id,pk,autoincrement"`

	DailyMenuUUID string     `bun:"daily_menu_uuid,notnull"`
	DailyMenu     *DailyMenu `bun:"rel:belongs-to,join:daily_menu_uuid=uuid"`
	MenuUUID      string     `bun:"menu_uuid,notnull"`
	Menu          *Menu      `bun:"rel:belongs-to,join:menu_uuid=uuid"`
}

var _ bun.BeforeAppendModelHook = (*DailyMenu)(nil)

func (r *DailyMenu) BeforeAppendModel(_ context.Context, query bun.Query) error {
	if r.UUID == "" {
		r.UUID = uuid.NewString()
	}

	return nil
}

func DailyMenuToDomainDailyMenu(x *DailyMenu) *domain.DailyMenu {
	if x == nil {
		return nil
	}

	var menus []*domain.Menu

	for _, menu := range x.Menus {
		menus = append(menus, MenuToDomainMenu(menu))
	}

	var dm = &domain.DailyMenu{
		UUID:  x.UUID,
		Name:  x.Name,
		Menus: menus,
	}

	dm.Fixup()

	return dm
}

func DomainDailyMenuToDailyMenu(x *domain.DailyMenu) *DailyMenu {
	if x == nil {
		return nil
	}

	var menus []*Menu

	for _, menu := range x.Menus {
		menus = append(menus, DomainMenuToMenu(menu))
	}

	return &DailyMenu{
		UUID:  x.UUID,
		Name:  x.Name,
		Menus: menus,
	}
}
