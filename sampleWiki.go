package main

import (
  "fmt"
  "io/ioutil"
  "net/http"
)

type Page struct {
  Title string
  Body []byte
}

func (p *Page) save() error {
  fileName := p.Title + ".txt"
  return ioutil.WriteFile(fileName, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
  fileName := title + ".txt"
  body, err := ioutil.ReadFile(fileName)
 
  if err != nil {
    return nil, err
  }
 
  return &Page{Title: title, Body: body}, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
  title := r.URL.Path[len("/view/"):]
  p,_ := loadPage(title)
  fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func main() {

  http.HandleFunc("/view/", handler)
  http.ListenAndServe(":8080", nil)
}
