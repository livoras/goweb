package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func sayHeloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Print(r.URL.Path)
	fmt.Fprint(w, "Hello World")
}

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.html")
		t.Execute(w, nil)
	} else if r.Method == "POST" {
		r.ParseForm()
		fmt.Println(r.Form)
		strs := []string{}
		for k, v := range r.Form {
			field := k + ":" + v[0]
			strs = append(strs, field)
		}
		fmt.Println(r.Form["Jerry"], ">>>>>")
		fmt.Println(r.Form.Get("Jerry"), ">>>>>")
		str := strings.Join(strs, "\n")
		template.HTMLEscape(w, []byte(str))
	}
}

func main() {
	fmt.Println("Fuck")
	http.HandleFunc("/", sayHeloWorld)
	http.HandleFunc("/login", login)
	log.Fatal(http.ListenAndServe(":9090", nil))
}
