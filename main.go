package main

import (
    "log"
    "net/http"
    "fmt"
)

func main() {

    log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("./static"))))
    
  }
