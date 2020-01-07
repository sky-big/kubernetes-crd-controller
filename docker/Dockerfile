FROM alpine:3.6

RUN apk add --no-cache ca-certificates

ADD kubernetes-crd-controller /usr/local/bin/pulsar-operator

RUN adduser -D kubernetes-crd-controller
USER kubernetes-crd-controller
