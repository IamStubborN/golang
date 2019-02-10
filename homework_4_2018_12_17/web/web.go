package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

var name, lastname, age string = "New", "User", "16"

//HomeRouterHandler use for /
func HomeRouterHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	for k, v := range r.Form {
		fmt.Println("key:", k, "val:", strings.Join(v, ""))
		if k == "name" {
			name = strings.Join(v, "")
		}
		if k == "lastname" {
			lastname = strings.Join(v, "")
		}
	}
	fmt.Fprintf(w, "Hello %s %s", name, lastname)
}

// SomethingRouterHandler use for /something
func SomethingRouterHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	for k, v := range r.Form {
		fmt.Println("key:", k, "val:", strings.Join(v, ""))
		if k == "name" {
			name = strings.Join(v, "")
			fmt.Fprintf(w, "Привет %s \n", name)
		}
		if k == "lastname" {
			lastname = strings.Join(v, "")
			fmt.Fprintf(w, "Здравствуйте %s \n", lastname)
		}
		if k == "age" {
			age = strings.Join(v, "")
			fmt.Fprintf(w, "Тебе %s лет \n", age)
		}
	}
}

func main() {
	http.HandleFunc("/", HomeRouterHandler)
	http.HandleFunc("/something", SomethingRouterHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
