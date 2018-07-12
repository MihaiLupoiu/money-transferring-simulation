FROM alpine
RUN apk update && apk add ca-certificates netcat-openbsd

# ENV GIN_MODE=release

COPY users.alpine users
    
ENTRYPOINT ["./users"]