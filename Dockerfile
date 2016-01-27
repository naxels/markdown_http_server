FROM centos:latest

MAINTAINER Patrick van de Glind

ADD markdown_http_server /home/markdown_http_server

WORKDIR /home

RUN chmod +x /home/markdown_http_server

CMD ["/home/markdown_http_server"]

EXPOSE 8080
