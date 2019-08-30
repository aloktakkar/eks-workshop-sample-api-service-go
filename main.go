package main

import (
    "log"
    "net/http"
    "fmt"
)

func main() {

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "This is an example of and OLDER version of the docker image")
    //http.Handle("/", http.FileServer(http.Dir("/static")))
    })

    http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hi")
    })
    log.Fatal(http.ListenAndServe(":9080", nil))
  }
