package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
	
	"io"
	"time"
	"crypto/md5"
	"strconv"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!")
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		cruntime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(cruntime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, token)
	} else {
		r.ParseForm()
		token := r.Form.Get("token")
		if token != "" {
			// token の合法性を検証する
			fmt.Println("Exist Token: ", token)
		} else {
			// token が存在しなければエラーを出す
			fmt.Println("Not exist Token")
		}
		fmt.Println("username length: ", len(r.Form["username"][0]))
		fmt.Println("username: ", r.Form["username"])
		fmt.Println("password: ", r.Form["password"])
		template.HTMLEscape(w, []byte(r.Form.Get("username"))) //クライアントに出力
	}
}

func main() {
	http.HandleFunc("/", sayhelloName)
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}