// package main

// import (
// 	"fmt"
// 	"net/http"
// 	"html/template"
// )

// type User struct {
// 	Name                 string
// 	Age                  uint16
// 	Money                int16
// 	AvgGrades, happyness float64
// 	Hobbies []string
// }

// func (u User) getAllInfo() string {
// 	return fmt.Sprintf("User name is %s. He is %d and he "+
// 		"has money equal: %d", u.Name, u.Age, u.Money)
// }

// func (u *User) setNewName(newName string){
// 	u.Name=newName
// }

// func homePage(w http.ResponseWriter, r *http.Request) {
// 	Kuat := User{"Kuat", 19, 1000, 4.2, 0.8, []string{"Football", "Skate", "Dance"}}
// 	// fmt.Fprintf(w, `<h1>Main text</h1>
// 	// <b>Main text</b>`)
// 	tmpl, _:= template.ParseFiles("templates/homePage.html")
// 	tmpl.Execute(w, Kuat)
// }
// func contactsPage(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "Contacts page!")
// }
// func handleRequest() {
// 	http.HandleFunc("/", homePage)
// 	http.HandleFunc("/contacts/", contactsPage)
// 	http.ListenAndServe(":8080", nil)
// }
// func main() {

// 	handleRequest()
// }
