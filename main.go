package main

import (
	"net/http"
	"log"
)

func main() {
	directory := flag.String("d", "/static", "the directory of static file to host")
	http.Handle("/", http.FileServer(http.Dir(*static)))
  
	log.Println("Listening...")
	log.Fatal(http.ListenAndServe(":8080", nil))
  }
