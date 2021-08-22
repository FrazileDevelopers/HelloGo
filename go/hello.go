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
var name = "Madhav"

type User struct {
	Fname    string `json:"fname"`
	Lname    string `json:"lname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

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

	/// MySQL Implementation
	// fmt.Println("Go Mysql =>")

	// db, dberr := sql.Open("mysql", "root:FrazileGo@0511@34.133.196.208/users")

	// if dberr != nil {
	// 	panic(dberr.Error())
	// }

	// defer db.Close()

	// fmt.Println("Successfully Connected to MySQL Database")

	// insert, err := db.Query("INSERT INTO users SET name='Parth Aggarwal'")

	// if err != nil {
	// 	panic(err.Error())
	// }

	// defer insert.Close()

	// fmt.Println("Successfully Inserted to USERS Table into FRAZILE Database")

	// Initialize the first connection to the database, to see if everything works correctly.
	// Make sure to check the error.
	// err := db.Ping()

	// results, err := db.Query("SELECT * FROM users")

	// if err != nil {
	// 	panic(err.Error())
	// }

	// for results.Next() {
	// 	var user User

	// 	err = results.Scan(&user.Fname)

	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	fmt.Println(user.Fname)
	// }
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
