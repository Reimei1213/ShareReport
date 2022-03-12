FROM golang:1.17.6-alpine

RUN apk update && \
    apk upgrade && \
    apk add git

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

# RUN go get -u github.com/cespare/reflex
RUN go build -o main bff/main.go