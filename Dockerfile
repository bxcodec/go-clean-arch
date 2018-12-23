# Builder
FROM golang:1.11.4-alpine3.8 as builder

RUN apk update && apk upgrade && \
    apk --update add git gcc make && \
    go get -u github.com/golang/dep/cmd/dep

WORKDIR /go/src/github.com/bxcodec/go-clean-arch

COPY . .

RUN make engine

# Distribution
FROM alpine:latest

RUN apk update && apk upgrade && \
    apk --update --no-cache add tzdata && \
    mkdir /app 

WORKDIR /app 

EXPOSE 9090

COPY --from=builder /go/src/github.com/bxcodec/go-clean-arch/engine /app

CMD /app/engine
