package handlers

import (
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	//abs, err := filepath.Abs("./main.html")
	//if err == nil {
	//	fmt.Println("Absolute:", abs)
	//}

	absPath, _ := filepath.Abs("../src/goidp/web/templates/main.html")
	mainTemplate, err := template.ParseFiles(absPath)
	if err != nil {
		panic(err)
	}

	log.Println("new connection...")

	mainTemplate.Execute(w, mainTemplate)
}
