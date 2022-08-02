package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("../../assets")))
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Failed to styart server!", err)
		return
	}
}
