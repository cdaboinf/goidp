package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
	//"web" 
	//"os"
	//"golang.org/x/oauth2"
	//"golang.org/x/oauth2/google"
)

func mainHandler(w http.ResponseWriter, r *http.Request) {
	mainTemplate, err := template.ParseFiles("web/templates/main.html")
	if err != nil {
		panic(err)
	}

	mainTemplate.Execute(w, mainTemplate)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	// redirect to google
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	// get token
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

// add utility validation
// add storage for client registration

func exchangeCodeForToken(code string) (token string) {
	return ""
}

func main() {
	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/login", loginHandler)
	//web.Run()
	log.Fatal(http.ListenAndServe(":8888", nil))
}
