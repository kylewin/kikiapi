package main

import (
	"fmt"
	"log"
	"net/http"
)

var (
	buildVersion string
	buildDate    string
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>KiKi Api</h1>")
	fmt.Fprintf(w, "Build Version: %s</br>", buildVersion)
	fmt.Fprintf(w, "Build Date: %s</br>", buildDate)
}

func main() {
	http.HandleFunc("/", hello)
	log.Println("Listening on port 8088...")
	log.Fatal(http.ListenAndServe(":8088", nil))
}
