package main

import (
    "fmt"
    "net/http"
	"common/day1"
)

func main() {
    http.HandleFunc("/hello", HelloServer)
	http.HandleFunc("/day1", Day1Server)
    http.ListenAndServe(":8080", nil)
}

func Day1Server(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, day1.Calculate())
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello!")
}