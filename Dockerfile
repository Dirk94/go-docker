FROM golang:latest

WORKDIR /app

COPY ./ /app

RUN go mod download

ENTRYPOINT go run commands/runserver.go