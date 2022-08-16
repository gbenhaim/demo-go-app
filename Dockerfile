FROM golang:1.17 as builder

WORKDIR /usr/src/app
COPY * ./
RUN go build -v -o bin/app

FROM registry.access.redhat.com/ubi8/ubi:latest

COPY --from=builder /usr/src/app/bin/app /usr/bin

CMD ["/usr/bin/app"]