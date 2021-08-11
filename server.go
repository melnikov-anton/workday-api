package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/melnikov-anton/workday-api/internal/router"
)

func main() {
	portNumb := flag.String("port", "8080", "port number for server to listen")

	flag.Parse()

	*portNumb = fmt.Sprintf(":%s", *portNumb)
	router := router.NewRouter()

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Printf("Starting server on port %s ...\n", *portNumb)

	log.Fatal(http.ListenAndServe(*portNumb, router))
}
