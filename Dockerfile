FROM golang:1.16-buster

COPY go.mod /root
RUN cd /root && go mod download
WORKDIR /app