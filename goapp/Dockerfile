FROM golang:1.16.5-alpine

ADD . /go/src

WORKDIR /go/src

RUN go get -d -v ./...

RUN go install -v ./...
