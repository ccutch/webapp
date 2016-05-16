package providers

import (
	"net/http"
)

// aliases for http interfaces
type handler http.Handler

// Provider defines the generalized definition of a service provider
type Provider interface {
	Mount() http.Handler
}
