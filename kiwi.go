package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

// Page defines how page data will be stored in memory.
type Page struct {
	Title string
	Body []byte
}

// This method saves the Page's Body to a text file and return
// an error if it fails to do so. For simplicity, it use the
// Title as the file name.
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
    filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func main() {
    p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
    p1.save()
	p2, err := loadPage("TestPage")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(p2.Body))
}
