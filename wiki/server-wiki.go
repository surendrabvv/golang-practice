package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "html/template"
)

type Page struct {
    Title string
    Body []byte
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/view/"):]
    page, _ := load(title)
    fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", page.Title, page.Body)    
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
    
}

func editHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/edit/"):]
    page, err := load(title)
    if err != nil {
        page := &Page{Title: title}
    }
    template := template.parseFiles("wiki-edit.html")
    template.execute(w, page) 
}

func load (title string) (*Page, error) {
    filename := title + ".txt"
    body, err := ioutil.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    return &Page{Title: title, Body: body}, nil
}

func main() {
    http.HandleFunc("/view/", viewHandler)
    http.HandleFunc("/save/", saveHanlder)
    http.HandleFunc("/edit/", editHandler)
    http.ListenAndServe(":3000", nil)
}