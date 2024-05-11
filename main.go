package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

var (
	buildVersion string
	buildDate    string
)

func hello(w http.ResponseWriter, r *http.Request) {
	// Get the HTTP method from the request
	method := r.Method

	// 10% return 500
	randomNumber := rand.Intn(100)
	var statusCode int
	if randomNumber < 90 {
		statusCode = http.StatusOK
	} else {
		statusCode = http.StatusInternalServerError
	}

	w.WriteHeader(statusCode)
	strStatusCode := strconv.Itoa(statusCode)

	currentTime := time.Now().Format("2006/01/02 15:04:05")

	fmt.Fprintf(w, "<h1>KiKi Api</h1>")
	fmt.Fprintf(w, "Build Version: %s</br>", buildVersion)
	fmt.Fprintf(w, "Build Date: %s</br>", buildDate)

	fmt.Fprintf(w, "%s %s request received</br>", currentTime, method)
	fmt.Fprintf(w, "%s %d response returned</br>", currentTime, statusCode)

	// Return the HTTP method as the response
	log.Println(method + " request received")
	log.Println(strStatusCode + " response returned")
}

func main() {
	http.HandleFunc("/", hello)
	log.Println("Listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
