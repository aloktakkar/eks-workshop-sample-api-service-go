package main

import (
    "log"
    "net/http"
    "fmt"
)

func main() {

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
fmt.Fprintf(w, "This is an example of a NEW DOCKER...you are now load balancing!!!!!!!!!!!!!")
    fs := http.FileServer(http.Dir("static/"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))
    })

    http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hi")
    })
    log.Fatal(http.ListenAndServe(":9080", nil))
  }
