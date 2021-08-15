package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

//  {{/* a comment */}}	Defines a comment
/*
{{.}}	Renders the root element
{{.Name}}	Renders the “Name”-field in a nested element
{{if .Done}} {{else}} {{end}}	Defines an if/else-Statement
{{range .List}} {{.}} {{end}}	Loops over all “List” field and renders each using {{.}}
*/

var tpl *template.Template
var name = "Parth"

func main() {
	/* Added html files */
	tpl, _ = tpl.ParseGlob("templates/*.html")

	/* Added css Directory */
	fs := http.FileServer(http.Dir("templates/css/"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))

	/* Added js Directory */
	js := http.FileServer(http.Dir("templates/js/"))
	http.Handle("/js/", http.StripPrefix("/js/", js))

	/* Added images Directory */
	images := http.FileServer(http.Dir("templates/images/"))
	http.Handle("/images/", http.StripPrefix("/images/", images))

	http.HandleFunc("/", handle)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func handle(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	// fmt.Fprint(w, "Hello world!")
	fmt.Println("Server Running!")
	tpl.ExecuteTemplate(w, "index.html", name)
}
