package test

import (
	"fmt"
	"net/http"
)

// Ping func ...
func Ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong")
}
