package main

import (
	"fmt"
	"io/ioutil"
)

// Page defines how page data will be stored in memory.
type Page struct {
	Title string
	Body []byte
}

// This method saves the Page's Body to a text file and return
// an error if it fails to do so. For simplicity, it use the
// Title as the file name.
func savePage(p *Page) error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
    filename := title + ".txt"
    if body, err := ioutil.ReadFile(filename), err == nil {
		return &Page{Title: title, Body: body}, nil
	}
	else {
		nil, err
	}
}

