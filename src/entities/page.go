package entities

import (
	"io/ioutil"
	"os"
)

type Page struct {
	Title string
	Body  []byte // byte slices: https://blog.golang.org/go-slices-usage-and-internals
}

var path = "./data/"

// Save Page
// Saves page into /data/{title}
func (p *Page) Save() error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0600)
	}

	filename := path + p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

// LoadPage loads page from /data/{title}.
// If file not exists, return nil
func LoadPage(title string) (*Page, error) {
	filename := path + title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return &Page{Title: title, Body: body}, nil
}
