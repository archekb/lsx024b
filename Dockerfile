FROM golang:1.17-buster

WORKDIR /app
COPY go.mod go.mod
RUN go mod download
