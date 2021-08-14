package handlers

import (
	"html/template"
	"log"
	"net/http"
)

func HomePage(rw http.ResponseWriter, r *http.Request) {
	if r.RequestURI != "/" {
		rw.Header().Set("Content-Type", "text/html")
		rw.WriteHeader(http.StatusNotFound)
		rw.Write([]byte("<h4>Page not found</h4>"))
		return
	}

	ts, err := template.ParseFiles("templates/home.page.tmpl")
	if err != nil {
		log.Println(err.Error())
		http.Error(rw, "Internal Server Error", 500)
		return
	}

	err = ts.Execute(rw, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(rw, "Internal Server Error", 500)
	}
}
