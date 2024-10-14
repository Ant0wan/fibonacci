package main

import (
	"fmt"
	"log"
	"math/big"
	"net/http"

	"fibonacci/lib"
)

func fibHandler(w http.ResponseWriter, r *http.Request) {
	// Only allow GET requests
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Parse the query parameters
	nStr := r.URL.Query().Get("n")
	if nStr == "" {
		http.Error(w, "Missing 'n' parameter", http.StatusBadRequest)
		return
	}

	// THis needs to be computed - maybe at start up depending on resource - to fix an abuse limit
	//if len(nStr) > 1000 {
    	//	http.Error(w, "Input is too large", http.StatusBadRequest)
    	//	return
	//}

	// Use math/big to handle very large integers
	n := new(big.Int)
	_, success := n.SetString(nStr, 10) // Parse base-10 string into big.Int
	if !success {
		http.Error(w, "Invalid 'n' parameter. It must be a valid number.", http.StatusBadRequest)
		return
	}
	// Estimate the time it will take to compute fibonacci on the number given before computing in order to reject it
	// If everything is fine, return the value of 'n'

	// Compute the nth Fibonacci number using the fiblib library
	fib := fiblib.FibonacciMatrix(n)

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

