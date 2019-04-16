package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		hostname, _ := os.Hostname()
		fmt.Fprintf(w, "<h1>Hello from host %v</h1>\n\nLast update: %v", hostname, time.Now().Format("Mon, Jan _2 2006, 15:04:05"))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
