package main

import (
	"fmt"
	"net/http"

	"github.com/ccutch/webapp/providers"
)

func main() {
	http.Handle("/users", new(providers.UserProvider).Mount())
	fmt.Println("Server online 127.0.0.1:5000")
	if err := http.ListenAndServe(":5000", nil); err != nil {
		panic(err)
	}
}
