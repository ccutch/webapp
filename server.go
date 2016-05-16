package webapp

import (
	"net/http"
)

// Server type, fulfilled from `http` package.
// Server should accept network requests, provide requests to providers, and return responses from
// providers to requesting clients.
type Server interface {
  Handle(string, http.Handler)
  ListenAndServe(string, http.Handler)
}