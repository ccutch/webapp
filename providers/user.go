package providers

import (
	"encoding/json"
	"net/http"

	"github.com/ccutch/webapp/services"
	"github.com/gorilla/mux"
)

// Handler response and request are all aliases to http types
type Handler http.Handler

// UserProvider handles routes for users
type UserProvider struct{}

// Mount fulfills provider interface
func (u *UserProvider) Mount() Handler {
	r := mux.NewRouter()
	r.HandleFunc("/users", u.getUsers)
	return r
}

func (u *UserProvider) getUsers(w http.ResponseWriter, r *http.Request) {
	us := new(services.UserService).GetUsers(1, 100)
	err := json.NewEncoder(w).Encode(&us)
	if err != nil {
		http.Error(w, "Error encoding users", 500)
	}
}
