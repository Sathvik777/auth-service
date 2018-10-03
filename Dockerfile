FROM golang:1.10 as builder
WORKDIR $GOPATH/src/github.com/Sathvik777/auth-service
COPY ./ .
RUN GOOS=linux GOARCH=386 go build -ldflags="-w -s" -v
RUN cp /auth-service /

FROM alpine:latest
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
COPY --from=builder //auth-service /
CMD //auth-service
