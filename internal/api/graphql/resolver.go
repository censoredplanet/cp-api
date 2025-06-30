package graph

import (
	service "github.com/censoredplanet/cp-api/internal/services"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Service service.ServicePort
}
