FROM golang:1.10-alpine
RUN apk update
RUN apk add --no-cache ca-certificates netcat-openbsd
RUN apk add --no-cache gcc musl-dev

ADD https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64 /usr/bin/dep
RUN chmod +x /usr/bin/dep

