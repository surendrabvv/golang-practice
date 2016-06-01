//Declaring package main

package main

//Importing the necessary packages
//like fmt (format) and net/http (http package)

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi every one, %s are coming", r.URL.Path[1:])
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":3000", nil)
}