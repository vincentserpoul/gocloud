FROM alpine:3.4
RUN apk add --no-cache ca-certificates apache2-utils
COPY ./testdeploy /testdeploy
EXPOSE 8086
ENTRYPOINT ["/testdeploy"]