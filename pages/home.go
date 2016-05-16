package views

import (
	"fmt"
	"net/http"
)

type response http.ResponseWriter
type request http.Request

// HomePage is a code representation for home.html
type HomePage struct {
	cache map[string]string
}

// Path fulfils Page interface returning route path for page
func (h HomePage) Path() string {
	return "/"
}

// Serve fulfils interface for page and writes page data to response writer
func (h HomePage) Serve(w response, r *request) {
	// view data should be loaded from file and cached.
	fmt.Fprint(w, `
<html>
  <head></head>
  <body>
    <h1>Home Page</h1>
  </body>
</html>  
  `)
}
