package main

import (
	"net/http"
	"log"
	"flag"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
  
	log.Println("Listening...")
	log.Fatal(http.ListenAndServe(":8080", nil))
  }
