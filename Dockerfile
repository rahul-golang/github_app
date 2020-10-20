FROM golang:1.14-alpine

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

RUN mkdir -p /go/src/github.com/rahul-golang/github_app
ADD .     /go/src/github.com/rahul-golang/github_app

WORKDIR /go/src/github.com/rahul-golang/github_app
RUN pwd
RUN ls
RUN go build -o github-app .


EXPOSE 8080

ENTRYPOINT [ "./github-app" ]

