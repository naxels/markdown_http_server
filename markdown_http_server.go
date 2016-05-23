package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/russross/blackfriday"
)

//Page struct for HTML pages
type Page struct {
	Title string
	Body  []byte
}

var (
	httpPort = flag.String("p", "8080", "web port to listen to")
	fileDir  = flag.String("f", ".", "base directory to start server from")
)

func main() {
	flag.Parse()

	if *httpPort == "" {
		fmt.Fprintln(os.Stderr, "require a web port")
		flag.Usage()
		os.Exit(1)
	}

	if *fileDir == "" {
		fmt.Fprintln(os.Stderr, "require a folder")
		flag.Usage()
		os.Exit(1)
	}

	http.HandleFunc("/", serveIndex)
	http.HandleFunc("/file/show", serveMarkdown)
	http.HandleFunc("/assets/", serveAssets)
	http.ListenAndServe(":"+*httpPort, nil)
}

func serveIndex(w http.ResponseWriter, r *http.Request) {
	log.Println("index request")

	files, _ := filepath.Glob(filepath.Join(*fileDir, "*.md"))

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

	p, _ := loadMarkdownFile(filename)

	fmt.Fprintf(w, "<head><link href='%s' rel='stylesheet' type='text/css'></head>", "/assets/github-markdown.css")
	//fmt.Fprintf(w, "<head><>%s<></head>", "assets/github-markdown.css")

	fmt.Fprintf(w, "<a href='/'>Back to Index</a><h1>%s</h1><div class='markdown-body'>%s</div>", p.Title, blackfriday.MarkdownCommon(p.Body))
}

func serveAssets(w http.ResponseWriter, r *http.Request) {
	log.Println("serveAssets request")

	asset := r.URL.Path[1:]

	log.Println("Asset:", asset, "requested")

	http.ServeFile(w, r, asset)
}

func loadMarkdownFile(filename string) (*Page, error) {
	//filename := file
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: filename, Body: body}, nil
}
