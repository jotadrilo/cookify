package gin

import (
	"github.com/jotadrilo/cookify/app/api"
	"github.com/jotadrilo/cookify/app/core/domain"
)

func DomainMenuLabelToAPIMenuLabel(v domain.MenuLabel) api.MenuLabel {
	return api.MenuLabel(v.String())
}
