FROM golang:1.17-buster

COPY go.mod /root
RUN cd /root && go mod download
WORKDIR /app