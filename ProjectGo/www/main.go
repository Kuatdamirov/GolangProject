package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	// "log"
	// _ "github.com/lib/pq"
)

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")

	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	t.ExecuteTemplate(w, "index", nil)
}

func create(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/create.html", "templates/header.html", "templates/footer.html")

	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	t.ExecuteTemplate(w, "create", nil)
}

func saveArticle(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	anons := r.FormValue("anons")
	fullText := r.FormValue("fullText")
	if title == "" || anons == "" || fullText == "" {
		fmt.Fprintf(w, "Не все поля данные заполнены!")
	}
	connStr := "user=postgres dbname=mydb password=1234 host=5432 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	insert, err := db.Query(fmt.Sprintf("INSERT INTO articles (title, anons, fullText) VALUES('%s', '%s', '%s')", title, anons, fullText))

	if err != nil {
		panic(err)
	}
	defer insert.Close()

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func handleFunc() {
	http.HandleFunc("/", index)
	http.HandleFunc("/create", create)
	http.HandleFunc("/saveArticle", saveArticle)
	http.ListenAndServe(":8080", nil)
}

func main() {
	handleFunc()
}
