package main

import (
	"learngo/users"
	"net/http"
	"text/template"
)

func signInUser(w http.ResponseWriter, r *http.Request) {
	newUser := getUser(r)
	ok := users.DefaultUserService.VerifyUser(newUser)
	fileName := "sign-in.html"
	t, _ := template.ParseFiles(fileName)
	if !ok {
		t.ExecuteTemplate(w, fileName, "New User Sign-in Failure")
		return;
	}
	t.ExecuteTemplate(w, fileName, "New User Sign-in Success")
}

func signUpUser(w http.ResponseWriter, r *http.Request) {
	newUser := getUser(r)
	err := users.DefaultUserService.CreateUser(newUser)
	fileName := "sign-up.html"
	t, _ := template.ParseFiles(fileName)
	if err != nil {
		t.ExecuteTemplate(w, fileName, "New User Sign-up Failure")
		return;
	}

	data := map[string]interface{}{
		"Message":           "New User Sign-up",
		"ShowSignInButton":  true,
	}
	t.ExecuteTemplate(w, fileName, data)

}

func getUser( r *http.Request) users.User {
	email := r.FormValue("email")
	password := r.FormValue("password")
	return users.User {
		Email: email,
		Password: password,
	}
}


func userHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/sign-in":
		signInUser(w, r)
	case "/sign-up":
		signUpUser(w, r)
	case "/sign-in-form":
		getSignInPage(w, r)
	case "/sign-up-form":
		getSignUpPage(w, r)
	}
}

func getSignInPage(w http.ResponseWriter, r *http.Request) {
	templating(w, "sign-in.html", nil)
}

func getSignUpPage(w http.ResponseWriter, r *http.Request) {
	templating(w, "sign-up.html", nil)
}

func templating(w http.ResponseWriter, fileName string, data interface{}) {
	t, _ := template.ParseFiles(fileName)
	t.ExecuteTemplate(w, fileName, data)
}

func main() {
	http.HandleFunc("/", userHandler)
	http.ListenAndServe(":8080", nil)
}