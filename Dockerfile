FROM golang:1.13 as builder
WORKDIR $GOPATH/src/github.com/Sathvik777/go-api-skeleton
COPY ./ .
RUN GOOS=linux GOARCH=386 go build -ldflags="-w -s" -v
RUN cp /go-api-skeleton /

FROM alpine:latest
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
COPY --from=builder //go-api-skeleton /
CMD //go-api-skeleton
