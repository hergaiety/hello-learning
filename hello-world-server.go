package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", homepage)
	fmt.Println("Server is running on http://localhost:8081")
	http.ListenAndServe(":8081", nil)
}

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello World</h1>")
}
