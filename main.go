package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"

	"github.com/russross/blackfriday"
)

//Page struct for HTML pages
type Page struct {
	Title string
	Body  []byte
}

func main() {
	http.HandleFunc("/", serveIndex)
	http.HandleFunc("/file/show", serveMarkdown)
	http.ListenAndServe(":3000", nil)
}

func serveIndex(w http.ResponseWriter, r *http.Request) {
	log.Println("index request")

	files, _ := filepath.Glob("*.md")

	fmt.Fprintf(w, "<h1>%s</h1>", "Index")

	fmt.Fprint(w, "<ul>")
	for _, f := range files {
		fmt.Fprintf(w, "<li><a href='/file/show?name=%s'>%s</a></li>", f, f)
	}
	fmt.Fprint(w, "</ul>")
}

func serveMarkdown(w http.ResponseWriter, r *http.Request) {
	log.Println("file/show request")

	filename := r.FormValue("name")

	if filename == "" {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	log.Println(filename, "requested")

	s, _ := loadStylesheet("assets/github-markdown.css")

	p, _ := loadMarkdownFile(filename)

	//fmt.Fprintf(w, "<head><link href='%s' rel='stylesheet' type='text/css'></head>", s)
	fmt.Fprintf(w, "<head><style type='text/css'>%s</style></head>", s)

	fmt.Fprintf(w, "<a href='/'>Back to Index</a><h1>%s</h1><div class='markdown-body'>%s</div>", p.Title, blackfriday.MarkdownCommon(p.Body))
}

func loadMarkdownFile(filename string) (*Page, error) {
	//filename := file
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: filename, Body: body}, nil
}

func loadStylesheet(filename string) ([]byte, error) {
	stylesheet, err := ioutil.ReadFile(filename)
	if err != nil {
		return []byte(""), err
	}
	return stylesheet, nil
}
