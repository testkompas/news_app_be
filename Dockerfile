FROM golang:1.17 AS builder
ENV APP_HOME /go/src/github.com/test_kompas/news_app
RUN mkdir -p $APP_HOME
WORKDIR $APP_HOME
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app ./cmd/

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /go/src/github.com/test_kompas/news_app/app ./
ENTRYPOINT [ "./app", "seed" ]
