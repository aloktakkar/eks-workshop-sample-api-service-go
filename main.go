package main

import (
    "log"
    "net/http"
    "fmt"
)

func main() {

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        http.Handle("/", http.FileServer(http.Dir("./environment/eks-workshop-sample-api-service-go/static")))
    })

    http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hi")
    })
    log.Fatal(http.ListenAndServe(":8080", nil))
  }
