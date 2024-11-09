package gin

import (
	openapi_types "github.com/oapi-codegen/runtime/types"

	"github.com/jotadrilo/cookify/app/api"
	"github.com/jotadrilo/cookify/app/core/domain"
	"github.com/jotadrilo/cookify/internal/oapi"
	"github.com/jotadrilo/cookify/internal/slices"
)

func DomainUserToAPIUser(v *domain.User) *api.User {
	if v == nil {
		return nil
	}

	return &api.User{
		Uuid:                     oapi.UUID(v.UUID),
		Name:                     v.Name,
		Email:                    v.Email,
		Birthday:                 openapi_types.Date{Time: v.BirthDate},
		Gender:                   v.Gender.String(),
		Height:                   v.Height,
		Weight:                   v.Weight,
		BmrMifflinStJeor:         oapi.Float32(v.BMRMifflinStJeor),
		BmrRevisedHarrisBenedict: oapi.Float32(v.BMRRevisedHarrisBenedict),
	}
}

func DomainUsersToAPIUsers(s []*domain.User) []*api.User {
	return slices.Map(s, DomainUserToAPIUser)
}
