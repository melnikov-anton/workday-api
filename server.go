package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/melnikov-anton/workday-api/internal/config"
	"github.com/melnikov-anton/workday-api/internal/handlers"
	"github.com/melnikov-anton/workday-api/internal/router"
)

func main() {
	portNumb := flag.String("port", "8080", "port number for server to listen")
	flag.Parse()

	var app config.AppConfig
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	app.AppRootDir = dir
	handlers.InitHandlers(&app)

	router := router.NewRouter(&app)

	*portNumb = fmt.Sprintf(":%s", *portNumb)

	fmt.Printf("Starting server on port %s ...\n", *portNumb)

	log.Fatal(http.ListenAndServe(*portNumb, router))
}
