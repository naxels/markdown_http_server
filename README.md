## Markdown HTTP Server

Go HTTP server that serves Markdown (.md) files from the directory the server is run

To get the source run:
```
go get github.com/naxels/markdown_http_server
```

You can use images in the Markdown files, have them refer to a file in
`/assets/(full filename with extension)`

*Warning*: since we only needed a simple server, not many user errors / security concerns are checked


Makes use of `blackfriday` for Markdown to HTML conversion

To download blackfriday run:
```
go get github.com/russross/blackfriday
```

Also uses the css from: https://github.com/sindresorhus/github-markdown-css
