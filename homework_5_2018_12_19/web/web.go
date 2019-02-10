package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

// User struct
type User struct {
	login    string
	password string
	isLogin  bool
}

var database Cache

func main() {
	database = InitCache()
	http.HandleFunc("/", Home)
	http.HandleFunc("/logoff", LogOff)
	http.HandleFunc("/login", Login)
	http.HandleFunc("/registation", Registration)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Home handler
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `
	<h1>Hi there!</h1>
	<h3>URL pathes is:</h3>
	<div>/registation with keys name=your login, password=your password</br>
	/login with keys name=your login, password=your password</br>
	/logoff with keys name=your login</br>
	</div>
	`)
}

// Registration handler
func Registration(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	login := ""
	password := ""
	for k, v := range r.Form {
		switch k {
		case "name":
			login = strings.Join(v, "")
		case "password":
			password = strings.Join(v, "")
		}
	}
	if login != "" || password != "" {
		database.AddUser(login, password)
		fmt.Fprintf(w, "Registration complete, please login %s", login)
	} else {
		fmt.Fprintf(w, "You incorrect login or password")
	}
}

// Login handler
func Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	login := ""
	password := ""
	for k, v := range r.Form {
		switch k {
		case "name":
			login = strings.Join(v, "")
		case "password":
			password = strings.Join(v, "")
		}
	}
	if err := database.Login(login, password); err != nil {
		fmt.Fprintf(w, err.Error())
	} else {
		fmt.Fprintf(w, "You logged in.")
	}
}

// LogOff handler
func LogOff(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	login := ""
	for k, v := range r.Form {
		switch k {
		case "name":
			login = strings.Join(v, "")
		}
	}
	if err := database.Logoff(login); err != nil {
		fmt.Fprintf(w, err.Error())
	} else {
		fmt.Fprintf(w, "You logged off in.")
	}
}
