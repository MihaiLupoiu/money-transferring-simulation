FROM alpine
RUN apk update && apk add ca-certificates netcat-openbsd

COPY users.alpine users
    
ENTRYPOINT ["./users"]