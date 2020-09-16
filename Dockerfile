######################################################
# Go build container
######################################################
FROM golang:latest as build
WORKDIR /go/src/github.com/ticketmaster/lbapi
ADD . .
RUN make build-linux
######################################################
# Move to alpine containera
######################################################
FROM alpine:3.11
EXPOSE 8080
EXPOSE 8443
RUN ulimit -n 15000
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
RUN mkdir /usr/local/share/ca-certificates/exth

WORKDIR /app
COPY --from=build /go/src/github.com/ticketmaster/lbapi/etc etc
COPY --from=build /go/src/github.com/ticketmaster/lbapi/target/linux/lbapi lbapi
ENTRYPOINT ["/app/lbapi"]
