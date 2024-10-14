package main

import (
	"fmt"
	"log"
	"math/big"
	"net/http"

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

	// This needs to be set for production purpose - to fix an abuse limit.
	// It could dynamically be set depending on the hardware it runs on.
	// Comment this for test purposes, if you do not mind wait for compute.
	// Estimate the time it will take to compute fibonacci on the number given before computing in order to reject it.
	if len(nStr) > 7 {
    		http.Error(w, "Input is too large to be computed quickly.", http.StatusBadRequest)
    		return
	}

	n := new(big.Int)
	_, success := n.SetString(nStr, 10)
	if !success {
		http.Error(w, "Invalid 'n' parameter. It must be a valid number.", http.StatusBadRequest)
		return
	}
	// If everything is fine, return the value of 'n'

	fib := fibonacci.FibonacciMatrix(n)

	// Return the Fibonacci number as a string
	fmt.Fprintf(w, "%s", fib.String())
}

func main() {
	// Register the handler for /fib
	http.HandleFunc("/fib", fibHandler)

	// Start the server on port 8000
	log.Println("Starting server on :8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatalf("Server failed: %s", err)
	}
}

