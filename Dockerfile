FROM golang:latest

WORKDIR /var/lib/app

COPY . .
RUN go mod download
