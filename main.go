package main

import (
	"fmt"
	"net/http"
)

func fibHandler(w http.ResponseWriter, r *http.Request) {
	number := r.URL.Query().Get("=")
	fmt.Fprintf(w, number)

	if number != "" {
		fmt.Fprintf(w, number)
	} else {
		http.Error(w, "Missing number in query", http.StatusBadRequest)
	}
}

func main() {
	http.HandleFunc("/fib", fibHandler)

	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}

