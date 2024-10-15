package main

import (
	"fmt"
	"log"
	"math/big"
	"net/http"
	"os"

	"fibonacci/lib"
)

func fibHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	nStr := r.URL.Query().Get("n")
	if nStr == "" {
		http.Error(w, "Missing 'n' parameter", http.StatusBadRequest)
		return
	}

	n := new(big.Int)
	if _, success := n.SetString(nStr, 10); !success {
		http.Error(w, "Invalid 'n' parameter. It must be a valid number.", http.StatusBadRequest)
		return
	}

	if len(nStr) > 7 {
		http.Error(w, "Input is too large to be computed quickly.", http.StatusBadRequest)
		return
	}

	fib := fibonacci.FibonacciMatrix(n)

	fmt.Fprintf(w, "%s", fib.String())
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	http.HandleFunc("/fib", fibHandler)

	log.Printf("Starting server on :%s", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatalf("Server failed: %s", err)
	}
}
