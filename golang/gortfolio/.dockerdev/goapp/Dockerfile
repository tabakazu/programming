FROM golang:1.13-alpine

ENV GOPATH=/go
ENV GO111MODULE=on

RUN apk add --update --no-cache bash ca-certificates git mysql-client

# アプリ用のディレクトリ作成
RUN mkdir -p /go/src/github.com/tabakazu/gortfolio
WORKDIR /go/src/github.com/tabakazu/gortfolio
