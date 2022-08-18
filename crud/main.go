package main

import (
	"fmt"
	"log"
	"net/http"
)

func hiHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/online" {
		http.Error(w, "404 Error", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "404 Error", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Yes, online!")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w,"ParseForm err: %s", err)
		return
	}
	fmt.Fprintf(w, "POST req success")
	login := r.FormValue("login")
	password := r.FormValue("password")
	fmt.Fprintf(w, "login: %s \n", login)
	fmt.Fprintf(w, "password: %s \n", password)
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/online", hiHandler)

	fmt.Printf("Starting Server")
	if err := http.ListenAndServe(":8080", nil); err !=nil {
		log.Fatal(err)
	}
}