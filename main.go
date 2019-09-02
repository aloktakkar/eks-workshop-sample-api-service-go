apackage main

import (
    "log"
    "net/http"
    "fmt"
)

func main() {

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "This is an example of a NEW DOCKER...you are now load balancing. Please upgrade to a new one shortly")
    //http.Handle("/", http.FileServer(http.Dir("/static")))
    })

    http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hi")
    })
    log.Fatal(http.ListenAndServe(":9080", nil))
  }
