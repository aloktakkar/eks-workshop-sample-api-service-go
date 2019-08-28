package main

import (
	"net/http"
	"log"
	"flag"
)

func main() {
	directory := flag.String("d", "/static", "the directory of static file to host")
	http.Handle("/", http.FileServer(http.Dir(*directory)))
  
	log.Println("Listening...")
	log.Fatal(http.ListenAndServe(":8080", nil))
  }
