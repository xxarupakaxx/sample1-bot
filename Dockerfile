FROM golang:alpine as builder

RUN apk update \
    && apk add --no-cache git curl make gcc g++
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
WORKDIR /go/src/github.com/xxarupakaxx/sample1-bot
COPY . .
RUN go mod download \
    && go build main.go

FROM alpine
COPY --from=builder /go/src/github.com/xxarupakaxx/sample1-bot /app

CMD /app/main $PORT