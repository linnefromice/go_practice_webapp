package main

import (
	"html/template"
	"net/http"
	"path/filepath"
	"sync"
)

type templateHandler struct {
	once     sync.Once
	filename string
	tmpl     *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path, err := filepath.Abs("./templates/")
	if err != nil {
		panic(err)
	}
	t.once.Do(
		func() {
			t.tmpl = template.Must(template.ParseFiles(path + t.filename))
		})
	t.tmpl.Execute(w, nil)
}