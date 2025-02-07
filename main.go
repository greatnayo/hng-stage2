package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func ClassifyNumberHandler(w http.ResponseWriter, r *http.Request) {
	// Set Response headers
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Retrieve the number from URL
	vars := mux.Vars(r)
	numStr, ok := vars["number"]

	// If not found in path, check query parameter
	if !ok {
		numStr = r.URL.Query().Get("number")
	}

	// Convert string to integer
	num, err := strconv.Atoi(numStr)
	if err != nil {
		http.Error(w, `{
		"number": "alphabet",
    	"error": true
		}`, http.StatusBadRequest)
		return
	}

	// Classification and properties
	classification := classifyNumber(num)
	armstrong := isArmstrong(num)

	// Create JSON response
	response := map[string]interface{}{
		"number":     num,
		"is_prime":   isPrime(num),
		"is_perfect": isPerfect(num),
		"properties": getProperties(armstrong, classification),
		"digit_sum":  findDigitSum(num), // sum of its digits
		"fun_fact":   getFunfact(num),
	}

	// Encode response as JSON
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// Function to classify a number
func classifyNumber(n int) string {
	if n == 0 {
		return "zero"
	} else if n%2 == 0 {
		return "even"
	} else {
		return "odd"
	}
}

// Function to check if a number is Armstrong
func isArmstrong(number int) bool {
	temp := number
	result := 0

	for temp > 0 {
		remainder := temp % 10
		result += remainder * remainder * remainder
		temp /= 10
	}

	return result == number
}

// Function to check if a number is Prime
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// Function to check if a number is Perfect
func isPerfect(num int) bool {
	sum := 0
	for count := 1; count < num; count++ {
		if num%count == 0 {
			sum += count
		}
	}
	return num == sum
}

// Function to calculate the sum of digits of a number
func findDigitSum(num int) int {
	res := 0
	for num > 0 {
		res += num % 10
		num /= 10
	}
	return res
}

// Function to fetch a fun fact from numbers API
func getFunfact(number int) string {
	url := fmt.Sprintf("http://numbersapi.com/%d/math", number)

	// Make HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return "Error fetching fact"
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return "Error reading fact"
	}

	// Save response into a variable
	return string(body)
}

// Function to get properties of a number
func getProperties(armstrong bool, classification string) []string {
	if armstrong {
		return []string{"armstrong", classification}
	}
	return []string{classification}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/classify-number/{number}", ClassifyNumberHandler).Methods(http.MethodGet)
	r.HandleFunc("/api/classify-number", ClassifyNumberHandler).Methods(http.MethodGet)

	fmt.Println("Server running on port 8080...")
	http.ListenAndServe(":8080", r)
}
