package handlers

import (
	"fmt"
	"net/http"
)

func HomePage(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "Welcome to the home page")
}
