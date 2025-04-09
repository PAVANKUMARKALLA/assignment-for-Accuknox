package main

import (
    "fmt"
    "net/http"
    "time"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        currentTime := time.Now().Format(time.RFC1123)
        fmt.Fprintf(w, "Current date and time: %s", currentTime)
    })

    http.ListenAndServe(":8080", nil)
}

