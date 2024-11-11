// main.go
package main

import (
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"strconv"
	"time"
)

// Function to calculate the nth Fibonacci number
func fibonacci(n int) *big.Int {
	if n <= 1 {
		return big.NewInt(int64(n))
	}
	a, b := big.NewInt(0), big.NewInt(1)
	for i := 2; i <= n; i++ {
		a, b = b, a.Add(a, b)
	}
	return b
}

// Handler function to serve JSON response
func handler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	// Parse the "n" parameter from the URL
	n, err := strconv.Atoi(r.URL.Query().Get("n"))
	if err != nil {
		n = 40 // Default to a large number to simulate CPU-bound work
	}

	// Calculate Fibonacci number
	result := fibonacci(n)

	// Prepare the JSON response
	response := map[string]interface{}{
		"number":  n,
		"result":  result.String(),
		"elapsed": time.Since(start).String(),
	}

	// Set headers and respond
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/fibonacci", handler)
	fmt.Println("Go server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
