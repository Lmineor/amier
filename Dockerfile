FROM golang:alpine AS builder

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.io,direct

WORKDIR /go/src/ziyue
COPY . .

RUN go env && go build -o server .


FROM alpine:latest

WORKDIR /go/src/ziyue

COPY --from=builder /go/src/ziyue ./

EXPOSE 8888

ENTRYPOINT ./server
