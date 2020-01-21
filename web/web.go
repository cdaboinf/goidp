package web

import(
	"net/http"
	"text/template"
)

func mainHandler(w http.ResponseWriter, r *http.Request) {
	mainTemplate, err := template.ParseFiles("web/templates/main.html")
	if err != nil {
		panic(err)
	}

	mainTemplate.Execute(w, mainTemplate)
}

func Run() {
	http.HandleFunc("/", mainHandler)
	//http.HandleFunc("/login", loginHandler)

	//log.Fatal(http.ListenAndServe(":8888", nil))
}