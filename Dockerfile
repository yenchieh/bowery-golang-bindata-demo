FROM alpine:3.10

ADD build/binary /binary

ENTRYPOINT ["/binary"]
