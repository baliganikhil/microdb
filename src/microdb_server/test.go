package main

import (
   "fmt"
   "net/http"
)

func main() {
    fmt.Println("Hello")

    http.HandleFunc("/", handleIndex);
    http.ListenAndServe(":8993", nil)
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello");
}
