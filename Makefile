.PHONY: all

all: clean build 

build:
	GOOS=linux go build .

clean:
	rm -f markdown_http_server
