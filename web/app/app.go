package app

import (
	"goidp/web/handlers"
	"log"
	"net/http"
)

func Run(port string) {
	http.HandleFunc("/", handlers.DefaultHandler)
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/callback", handlers.CodeExchangeHandler)

	log.Fatal(http.ListenAndServe(port, nil))
}
