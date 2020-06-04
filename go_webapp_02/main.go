package main

import (
	"html/template"
	"log"
	"net/http"
)

func htmlHandler0(w http.ResponseWriter, r *http.Request) {
	// テンプレートをparse
	t := template.Must(template.ParseFiles("templates/template000.html.tpl"))
	str := "Sample Message"

	// テンプレートを描画
	if err := t.ExecuteTemplate(w, "template000.html.tpl", str);err != nil {
		log.Fatal(err)
	}
}

func htmlHandler1(w http.ResponseWriter, r *http.Request) {
	// テンプレートをparse
	t := template.Must(template.ParseFiles("templates/template001.html.tpl"))
	m := map[string]int{
		"key1": 101,
		"key2": 202,
		"key3": 303,
		"key4": -404,
	}

	// テンプレートを描画
	if err := t.ExecuteTemplate(w, "template001.html.tpl", m); err != nil {
		log.Fatal(err)
	}
}

func htmlHandler2(w http.ResponseWriter, r *http.Request) {
	// テンプレートをparse
	t := template.Must(template.ParseFiles("templates/template002.html.tpl"))
	type SampleData struct {
		Name string
		Age int
	}
	data := SampleData{Name: "Taro", Age: 25}

	// テンプレートを描画
	if err := t.ExecuteTemplate(w, "template002.html.tpl", data); err != nil {
		log.Fatal(err)
	}
}

func main() {
	log.Println("Launching...")
	http.HandleFunc("/page0", htmlHandler0)
	http.HandleFunc("/page1", htmlHandler1)
	http.HandleFunc("/page2", htmlHandler2)

	// サーバーを起動
	http.ListenAndServe(":8080", nil)
}