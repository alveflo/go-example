package entities

import "io/ioutil"

type Page struct {
	Title string
	Body  []byte // byte slices: https://blog.golang.org/go-slices-usage-and-internals
}

func (p *Page) Save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}
