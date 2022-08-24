package main

import (
	"fmt"
	"log"
	"net/http"
)

// Request
// Reponse
// fun(req,res){}

func Welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the QTSolution intern!")
	fmt.Println("Endpoint Hit: homePage")
}

func Register(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Register intern!")
	fmt.Println("Endpoint Hit: homePage")
}

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Login Page!")
	fmt.Println("Endpoint Hit: homePage")
}

func ForgetPassword(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Forget Password!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/", Welcome)
	http.HandleFunc("/register", Register)
	http.HandleFunc("/login", Login)
	http.HandleFunc("/password", ForgetPassword)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRequests()
}
