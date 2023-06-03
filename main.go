package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("./public/*.html"))
}

func main() {
	// Serve static files from the "public" directory
	fs := http.FileServer(http.Dir("public"))

	http.HandleFunc("/home", home)

	// http.HandleFunc("/signup", nil)
	// http.HandleFunc("/login", nil)

	// http.HandleFunc("/post", nil)
	// http.HandleFunc("/post", nil)
	// http.HandleFunc("/post", nil)
	// http.HandleFunc("/post/most_liked", nil)
	// http.HandleFunc("/post/categ", nil)
	// http.HandleFunc("/post/like", nil)
	// http.HandleFunc("/post", nil)

	// Handle requests using the file server
	http.Handle("/", fs)

	// Start the server
	fmt.Println("http://localhost:8080/home")
	http.ListenAndServe(":8080", nil)
}
func home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("log")
	if err := tmpl.ExecuteTemplate(w, "home.html", nil); err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
	}
}
