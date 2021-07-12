FROM golang:latest as builder

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
WORKDIR /go/src/github.com/xxarupakaxx/sample1-bot
COPY . .
RUN go github.com/line/line-bot-sdk-go/linebot
RUN go build main.go

FROM alpine
COPY --from=buidler /go/src/github.com/xxarupakaxx/sample1-bot /app

CMD /app/main $PORT