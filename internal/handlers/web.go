package handlers

import (
	"net/http"
)

func HomePage(rw http.ResponseWriter, r *http.Request) {
	if r.RequestURI != "/" {
		rw.Header().Set("Content-Type", "text/html")
		rw.WriteHeader(http.StatusNotFound)
		rw.Write([]byte("<h4>Page not found</h4>"))
		return
	}
	rw.Header().Set("Content-Type", "text/html")
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte("<h2>Welcome to the home page</h2>"))
}
