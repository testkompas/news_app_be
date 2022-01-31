FROM golang:1.17
ENV APP_HOME /go/src/github.com/test_kompas/news_app
RUN mkdir -p $APP_HOME
WORKDIR $APP_HOME
COPY . .
RUN GOOS=linux go build -o app ./cmd/

FROM alpine:latest
WORKDIR /root/
COPY --from=0 /go/src/github.com/test_kompas/news_app/app ./
ENTRYPOINT [ "./app", "seed" ]
