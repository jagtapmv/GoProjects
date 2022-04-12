package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileserver)
	http.HandleFunc("/form", formHander)
	http.HandleFunc("/hello", helloHandeler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func formHander(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/form" {
		http.Error(res, "404 not found", http.StatusNotFound)
	}
	if err := req.ParseForm(); err != nil {
		fmt.Fprintf(res, "Parseform() err %v\n", err)
	}
	name := req.FormValue("name")
	mobile := req.FormValue("Mobile")

	fmt.Fprintf(res, "Name : %s\n", name)
	fmt.Fprintf(res, "Mobile Number : %s\n", mobile)
}

func helloHandeler(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/hello" {
		http.Error(res, "404 not found", http.StatusNotFound)
	}
	if req.Method != "GET" {
		http.Error(res, "Method not supported!", http.StatusNotFound)
	}
	fmt.Fprintf(res, "Hello world!")
}
