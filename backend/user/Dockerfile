FROM golang:1.17.6-alpine

RUN apk update

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go get -u github.com/cespare/reflex 
CMD reflex -r '(\.go$|go\.mod)' -s go run *.go