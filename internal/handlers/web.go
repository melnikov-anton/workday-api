package handlers

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"time"
)

type TemplateData struct {
	IsWorkday bool
}

// HomePage shows home page of the app
func HomePage(rw http.ResponseWriter, r *http.Request) {
	if r.RequestURI != "/" {
		rw.Header().Set(HEADER_CONTENT_TYPE, "text/html")
		rw.WriteHeader(http.StatusNotFound)
		rw.Write([]byte("<h4>Page not found</h4>"))
		return
	}

	ts, err := template.ParseFiles(filepath.Join(appConfig.AppRootDir, "templates", "home.page.tmpl"))
	if err != nil {
		log.Println(err.Error())
		http.Error(rw, STATUS_INTERNAL_SERVER_ERROR, 500)
		return
	}

	todayString := time.Now().Format("2006.01.02")
	isWorkday, err := IsDateWorkday(todayString, "ru")
	if err != nil {
		log.Println(err.Error())
		http.Error(rw, STATUS_INTERNAL_SERVER_ERROR, 500)
		return
	}
	tmplData := TemplateData{
		IsWorkday: isWorkday,
	}

	err = ts.Execute(rw, tmplData)
	if err != nil {
		log.Println(err.Error())
		http.Error(rw, STATUS_INTERNAL_SERVER_ERROR, 500)
	}
}
