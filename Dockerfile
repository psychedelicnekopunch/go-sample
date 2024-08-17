# https://hub.docker.com/_/golang
FROM golang:1.21rc3

WORKDIR /var/www/html

RUN go install github.com/cosmtrek/air@latest

COPY ./app /var/www/html
