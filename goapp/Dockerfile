FROM golang:1.16.5-alpine

ADD . /home

WORKDIR /home

RUN \
  apk add --no-cache bash git openssh && \
  go get -u go.mongodb.org/mongo-driver/mongo && \
  go get -u go.mongodb.org/mongo-driver/mongo && \
  go get -u go.mongodb.org/mongo-driver/mongo/options && \
  go get -u go.mongodb.org/mongo-driver/mongo/readpref"
