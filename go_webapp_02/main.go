package main

import (
	"html/template"
	"log"
	"net/http"
	"time"
	"fmt"
	"strings"
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

func htmlHandler3(w http.ResponseWriter, r *http.Request) {
	funcMap := template.FuncMap{
		"samplefunc1": func() string { return time.Now().String() },
		"samplefunc2": func(src string) string { return fmt.Sprintf("[$$ %s $$]", src)},
		"samplefunc3": strings.ToUpper,
	}
	// テンプレートをparse
	t := template.Must(template.New("t").Funcs(funcMap).ParseFiles("templates/template003.html.tpl"))
	m := map[string]string {
		"msg1": "golang is programming language.",
		"msg2": "hello template app.",
	}

	// テンプレートを描画
	if err := t.ExecuteTemplate(w, "template003.html.tpl", m); err != nil {
		log.Fatal(err)
	}
}

func htmlHandler4(w http.ResponseWriter, r *http.Request) {
	// テンプレートをparse
	t := template.Must(template.ParseFiles("templates/template004.html.tpl"))
	strArray := []string{"aa", "bbb", "cccc", "ddddd"}

	// テンプレートを描画
	if err := t.ExecuteTemplate(w, "template004.html.tpl", strArray); err != nil {
		log.Fatal(err)
	}
}

func htmlHandler5(w http.ResponseWriter, r *http.Request) {
	funcMap := template.FuncMap{
		"safehtml": func(text string) template.HTML { return template.HTML(text) },
	}
	// テンプレートをparse
	t := template.Must(template.New("T").Funcs(funcMap).ParseFiles("templates/template005.html.tpl"))
	st := struct {
		Param1 string
		Param2 string
	}{
		Param1: r.FormValue("param1"),
		Param2: r.FormValue("param2"),
	}

	// テンプレートを描画
	if err := t.ExecuteTemplate(w, "template005.html.tpl", st); err != nil {
		log.Fatal(err)
	}
}

func main() {
	log.Println("Launching...")
	http.HandleFunc("/page0", htmlHandler0)
	http.HandleFunc("/page1", htmlHandler1)
	http.HandleFunc("/page2", htmlHandler2)
	http.HandleFunc("/page3", htmlHandler3)
	http.HandleFunc("/page4", htmlHandler4)
	http.HandleFunc("/page5", htmlHandler5)

	// サーバーを起動
	http.ListenAndServe(":8080", nil)
}