package main

import (
    "fmt"
    "io/ioutil"
)

type Page struct {
    Title string
    Body []byte
}

func save (p *Page) error {
    filename := p.Title + ".txt"
    return ioutil.WriteFile(filename, p.Body, 0600)
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
    save(&Page{Title: "TestPage", Body: []byte("This is a test string")})
    content, error := load ("TestPage")
    if error != nil {
        fmt.Println(error)
    } else {
        fmt.Println(string(content.Body))   
    }    
}