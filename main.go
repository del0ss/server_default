package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type Note struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Text    string `json:"text"`
}

var Notes = []Note{}

func main() {
	//json.Marshal()
	http.HandleFunc("/", mainPage)
	http.HandleFunc("/notes", notesPage)
	http.HandleFunc("/create_notes", saveNotePage)

	err := http.ListenAndServe(":3000", nil)
	fmt.Println("Server is starting on port 3000")
	if err != nil {
		log.Fatal("ListenAndServer", err)
	}
}

func mainPage(w http.ResponseWriter, req *http.Request) {
	tmpl, err := template.ParseFiles("static/index.html")

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
}

func notesPage(w http.ResponseWriter, req *http.Request) {
	//notes := []Note{Note{"NAME", "SURNAME", "TEXT"}, Note{"NAME2", "SURNAME2", "TEXT2"}}
	//js, _ := json.Marshal(note)
	tmpl, err := template.ParseFiles("static/notes.html")
	if err != nil {
		http.Error(w, err.Error(), 400)
	}

	if err := tmpl.Execute(w, Notes); err != nil {
		http.Error(w, err.Error(), 400)
	}

}

func saveNotePage(w http.ResponseWriter, req *http.Request) {
	tmpl, err := template.ParseFiles("static/create.html")
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	name := req.FormValue("name")
	surname := req.FormValue("surname")
	text := req.FormValue("text")
	if name != "" && surname != "" && text != "" {
		Notes = append(Notes, Note{name, surname, text})
	}
	if err := tmpl.Execute(w, Notes); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	fmt.Println("FROM CREATE", name, surname, text)

}
