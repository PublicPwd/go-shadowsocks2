FROM golang:1.18.3 AS builder

WORKDIR /app
USER root
ARG GO111MODULE=on
ARG GOPROXY=https://goproxy.cn

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -a -o shadowsocks

FROM debian:buster-slim

ENV TZ=Asia/Shanghai

RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime \
    && echo $TZ > /etc/timezone

COPY --from=builder /app/shadowsocks /usr/bin/shadowsocks

ENTRYPOINT ["shadowsocks"]
