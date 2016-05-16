package pages

import (
  "net/http"
)

// Page is a generalized definition of a code representation of an html view. 
type Page interface {
  Path() string
  Serve(http.ResponseWriter, *http.Request)
}