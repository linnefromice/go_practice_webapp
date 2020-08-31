package main

import (
	"fmt"
	"os"
	"text/template"
)

func main() {
	s1, _ := template.ParseFiles("header.tmpl", "content.tmpl", "footer.tmpl")
	fmt.Println("HEADER")
	s1.ExecuteTemplate(os.Stdout, "header", nil)
	fmt.Println("CONTENT")
	s1.ExecuteTemplate(os.Stdout, "content", nil)
	fmt.Println("FOOTER")
	s1.ExecuteTemplate(os.Stdout, "footer", nil)
	fmt.Println("-- Execute --")
	s1.Execute(os.Stdout, nil)
}