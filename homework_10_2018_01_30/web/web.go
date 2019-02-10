package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"
	"time"

	"gopkg.in/mgo.v2/bson"
)

type (
	// User struct
	User struct {
		ID          bson.ObjectId `bson:"_id"`
		Login       string        `bson:"login"`
		Password    string        `bson:"password"`
		IsLogin     bool          `bson:"is_login"`
		Description `bson:"description"`
	}

	// Description struct
	Description struct {
		Name string `bson:"name"`
		Age  uint8  `bson:"age"`
	}

	// Cookie struct
	Cookie struct {
		Login    string
		Password string
		Expires  time.Time
	}
)

var database Databaser

func init() {
	database = InitCacheDB()
}
func main() {
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))
	http.HandleFunc("/", Home)
	http.HandleFunc("/register", Registration)
	http.HandleFunc("/login", LogIn)
	http.HandleFunc("/exit", LogOff)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Home handler
func Home(w http.ResponseWriter, r *http.Request) {
	if login, err := r.Cookie("login"); err == nil {
		t, err := template.ParseFiles(
			"./templates/main-page.html",
			"./templates/auth/header.html",
			"./templates/auth/main.html",
		)
		if err != nil {
			fmt.Fprint(w, err.Error())
			return
		}
		fmt.Println(login)
		fmt.Println(login.Value)
		t.ExecuteTemplate(w, "index", login.Value)
	} else {
		t, err := template.ParseFiles(
			"./templates/main-page.html",
			"./templates/no-auth/header.html",
			"./templates/no-auth/main.html",
		)
		if err != nil {
			fmt.Fprint(w, err.Error())
			return
		}
		t.ExecuteTemplate(w, "index", nil)
	}
}

// Registration handler
func Registration(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles(
		"./templates/register-page.html",
		"./templates/no-auth/header.html",
		"./templates/no-auth/register.html",
	)
	r.ParseForm()
	var login, password, name, age string
	if r.Form.Get("login") != "" {
		for k, v := range r.Form {
			switch k {
			case "login":
				login = v[0]
			case "password":
				password = v[0]
			case "name":
				name = v[0]
			case "age":
				age = v[0]
			}
		}
		userAge, _ := strconv.Atoi(age)
		if len([]rune(login)) >= 3 && len([]rune(password)) >= 3 && len([]rune(name)) >= 3 && userAge >= 16 && userAge <= 99 {
			encPass := encryptString(password)
			if err := database.AddUser(User{bson.NewObjectId(), login, encPass, true, Description{name, uint8(userAge)}}); err != nil {
				fmt.Fprintf(w, err.Error())
			} else {
				expiration := time.Now().Add(30 * time.Minute)
				cookie := http.Cookie{Name: "login", Value: login, Expires: expiration}
				http.SetCookie(w, &cookie)
				cookie = http.Cookie{Name: "password", Value: encPass, Expires: expiration}
				http.SetCookie(w, &cookie)
				http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			}
		} else {
			fmt.Fprintf(w, "You put incorrect values")
		}
	} else {
		t.ExecuteTemplate(w, "index", nil)
	}
}

// LogIn handler
func LogIn(w http.ResponseWriter, r *http.Request) {
	if _, err := r.Cookie("login"); err == nil {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	} else {
		t, _ := template.ParseFiles(
			"./templates/login-page.html",
			"./templates/no-auth/header.html",
			"./templates/no-auth/login.html",
		)
		r.ParseForm()
		var login, password string
		if r.Form.Get("login") != "" {
			for k, v := range r.Form {
				switch k {
				case "login":
					login = v[0]
				case "password":
					password = v[0]
				}
			}
			if login != "" || password != "" {
				encPass := encryptString(password)
				if err := database.LogIn(login, encPass); err != nil {
					fmt.Fprintf(w, err.Error())
				} else {
					expiration := time.Now().Add(30 * time.Second)
					cookie := http.Cookie{Name: "login", Value: login, Expires: expiration}
					http.SetCookie(w, &cookie)
					cookie = http.Cookie{Name: "password", Value: encPass, Expires: expiration}
					http.SetCookie(w, &cookie)
					http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
				}
			}
		} else {
			t.ExecuteTemplate(w, "index", nil)
		}
	}
}

// LogOff handler
func LogOff(w http.ResponseWriter, r *http.Request) {
	if login, err := r.Cookie("login"); err == nil {
		cookie := http.Cookie{Name: "login", Value: "", Expires: time.Unix(0, 0)}
		http.SetCookie(w, &cookie)
		cookie = http.Cookie{Name: "password", Value: "", Expires: time.Unix(0, 0)}
		http.SetCookie(w, &cookie)
		if err := database.LogOff(login.Value); err != nil {
			fmt.Fprintf(w, err.Error())
		}
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	}
}
