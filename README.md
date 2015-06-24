## Markdown HTTP Server

Go HTTP server that serves Markdown (.md) files from the directory the server is run

Makes use of `blackfriday` for Markdown to HTML conversion

To use run:
```
go get github.com/russross/blackfriday
```

*Warning*: since we only needed a simple server, not many user errors / security concerns are checked
