package main

import (
    "fmt"
	"log"
    "net/http"
)

func main() {

    fmt.Println("Kubernetes Secret Management Handbook - Chapter 1 - Hello World Example")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello from URL path: %s\n", r.URL.Path)
		fmt.Printf("User requested the URL path:%s\n", r.URL.Path)
    })
	fmt.Printf("Server running on http://localhost:8080\n")
    if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}