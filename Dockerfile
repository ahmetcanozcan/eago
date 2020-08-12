
FROM golang:1.13-alpine AS build


ENV GOOS=linux
ENV GO111MODULE=on
ENV EAGO_BUILD_VERSION=0.0.2

WORKDIR /go/src/github.com/ahmetcanozcan/eago

COPY . /go/src/github.com/ahmetcanozcan/eago

RUN apk update && \
  go get github.com/magefile/mage

# build eago
RUN mage eago

