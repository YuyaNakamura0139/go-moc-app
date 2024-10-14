# https://hub.docker.com/_/golang
FROM golang:1.23-alpine

ENV ROOT=/go/src/work
WORKDIR ${ROOT}

RUN apk update \
    && apk add --no-cache git \
    # グローバルインストール,$GOPATH/binに配置
    # ホットリロードツール(コードの変更を検知して自動でビルドする)
    && go install github.com/air-verse/air@latest

CMD ["air", "-c", ".air.toml"]