// +build ignore

package main

import (
	"html/template"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Page struct {
	Title string
	Body []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s", r.URL.Path[1:])
}

var templates = template.Must(template.ParseFiles("edit.html", "view.html"))
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	t, err := templates.ExecuteTemplate(w, tmpl = ".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusIntervalServerError)
		return
	}
	err = t.Execute(w, p)
	if err !=  nil {
		http.Error(w, err.Error(), http.StatusIntervalServerError)
	}
}

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")
func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
	m := validPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return "", errors.New("invalid Page Title")
	}
	return m[2], nil // The title is the second subexpression.
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title, err := getTitle(w, r)
	if err != nil {
		return
	}
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/" + title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title, err := getTitle(w, r)
	if err != nil {
		return
	}
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func savehandler(w http.ResponseWriter, r *http.Request) {
	title, err := getTitle(w, r)
	if err != nil {
		return
	}
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Errpr(w, err.Error(), http.StatusIntervalServerError)
		return
	}
	http.Request(w, r, "/view/" + title, http.StatusFound)
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/view/", viewHandler) // http://localhost:8080/view/TestPage
	http.HandleFunc("/edit/", editHandler) // http://localhost:8080/edit
	http.HandleFunc("/save/", savehandler) // http://localhost:8080/save
	log.Fatal(http.ListenAndServe(":8080", nil))
}