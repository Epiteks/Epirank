FROM golang:alpine

ADD . /go/src/github.com/Shakarang/Epirank

RUN apk add --update alpine-sdk

RUN go get -u github.com/Sirupsen/logrus
RUN go get -u github.com/mattn/go-sqlite3
RUN go get -u github.com/gin-gonic/gin

RUN go install github.com/Shakarang/Epirank

ENTRYPOINT /go/bin/Epirank

EXPOSE 8080