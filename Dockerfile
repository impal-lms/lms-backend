FROM alpine:3.7
WORKDIR /server
RUN apk update \
        && apk upgrade \
        && apk add --no-cache \
        ca-certificates \
        && update-ca-certificates 2>/dev/null || true
COPY ./bin/app .
EXPOSE 9000
CMD ["/server/app", "server"]

