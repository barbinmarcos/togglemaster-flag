package main

import (
	"fmt"
	"net/http"
)

func health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "togglemaster-flag running")
}

func main() {

	http.HandleFunc("/health", health)

	fmt.Println("Server running on :8081")

	http.ListenAndServe(":8081", nil)
}
