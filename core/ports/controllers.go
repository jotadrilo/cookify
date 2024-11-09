package ports

import "github.com/jotadrilo/cookify/api"

// Controller is a marker to ensure that all Controller interfaces are unique
type Controller interface {
	controller()
}

type GinController interface {
	Controller
	api.ServerInterface
}
