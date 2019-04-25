package main

import (
    "net/http"
    "log"
)

func main() {
    addr := ":8080"
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello, kubernetes!"))
    })
    log.Println("listen on", addr)
    http.ListenAndServe(addr, nil)
}

