package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"interview.heeko/login/database"
	"interview.heeko/login/model"
)

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	switch r.Method {
	case "GET":
		t, _ := template.ParseFiles("./views/login.html")
		t.Execute(w, nil)
	case "POST":
		r.ParseForm()
		username := r.Form["username"][0]
		password := r.Form["password"][0]
		fmt.Println("username:", username)
		fmt.Println("password:", password)
		user, err := database.GetUser(username, password)
		w.Header().Set("Content-Type", "application/json")
		if err != nil {
			log.Println(fmt.Sprintf("User not found: %s", err.Error()))
			w.WriteHeader(http.StatusBadRequest)
			responseError := model.ErrorResponse{
				Response: err.Error()}
			json.NewEncoder(w).Encode(responseError)
		}
		json.NewEncoder(w).Encode(user)
	default:
		fmt.Fprintf(w, "Unsupported function")
	}
}

func register(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		t, _ := template.ParseFiles("./views/register.html")
		t.Execute(w, nil)
	case "POST":
		r.ParseForm()
		username := r.Form["username"][0]
		password := r.Form["password"][0]
		fmt.Println("username:", username)
		fmt.Println("password:", password)
		user, err := database.CreateUser(username, password)
		w.Header().Set("Content-Type", "application/json")
		if err != nil {
			log.Println(fmt.Sprintf("User not created due to: %s", err.Error()))
			w.WriteHeader(http.StatusBadRequest)
			responseError := model.ErrorResponse{
				Response: "User already exist"}
			json.NewEncoder(w).Encode(responseError)
		}
		json.NewEncoder(w).Encode(user)
	default:
		fmt.Fprintf(w, "Unsupported function")
	}
}

func loadEnvFile() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err.Error())
	}
}

func main() {
	loadEnvFile()
	database.Database()

	// serve css/html files
	http.Handle("/views/", http.StripPrefix("/views/", http.FileServer(http.Dir("./views"))))

	http.HandleFunc("/", login)
	http.HandleFunc("/login", login)
	http.HandleFunc("/register", register)

	err := http.ListenAndServe(os.ExpandEnv(":${SERVER_PORT}"), nil) // setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
