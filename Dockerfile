# https://hub.docker.com/_/golang
FROM golang:1.22.8-alpine

ENV ROOT=/go/src/work
WORKDIR ${ROOT}

RUN apk update \
    && apk add --no-cache git