FROM golang:1.13 AS builder

ADD . /go/src/go-con-client

WORKDIR /go/src/go-con-client
ENV GO111MODULE=on

RUN go build -mod=vendor -o main

FROM debian:9.9-slim

ENV LANG C.UTF-8
ENV GOPATH /go

COPY --from=builder /go/src/go-con-client/main /app/main
WORKDIR /app

RUN chmod +x main

ENTRYPOINT ["./main"]
EXPOSE 8081