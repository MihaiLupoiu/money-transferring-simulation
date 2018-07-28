FROM golang:1.10-alpine
RUN apk update && apk upgrade
RUN apk add --no-cache ca-certificates netcat-openbsd
RUN apk add --no-cache gcc musl-dev
RUN apk add --no-cache bash git openssh

ADD https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64 /usr/bin/dep
RUN chmod +x /usr/bin/dep
