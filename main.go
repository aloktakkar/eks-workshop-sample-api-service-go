package main

import (
    "log"
    "net/http"
)

func main() {

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, r.URL.Path[1:])
    })

    http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hi")
    })
	log.Fatal(http.ListenAndServe(":8080", nil))
    http.Handle("/", http.FileServer(http.Dir("./static")))

    log.Fatal(http.ListenAndServe(":8080", nil))
  }
