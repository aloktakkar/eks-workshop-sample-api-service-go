package main

import (
	"net/http"
	"log"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)
  
	log.Println("Listening...")
	http.ListenAndServe(":3000", nil)
  }
  
 type response struct {
	Message string   `json:"message"`
	EnvVars []string `json:"env"`
	Fib     []int    `json:"fib"`
}

func fib() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}
