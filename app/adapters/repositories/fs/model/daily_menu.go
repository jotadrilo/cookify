package model

import "github.com/jotadrilo/cookify/app/core/domain"

type DailyMenu struct {
	UUID     string  `json:"uuid"`
	UserUUID string  `json:"user_uuid"`
	Name     string  `json:"name"`
	Menus    []*Menu `json:"menus"`
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
		UUID:     x.UUID,
		UserUUID: x.UserUUID,
		Name:     x.Name,
		Menus:    menus,
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
		UUID:     x.UUID,
		UserUUID: x.UserUUID,
		Name:     x.Name,
		Menus:    menus,
	}
}
