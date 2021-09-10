# syntax=docker/dockerfile:1
FROM golang:1.16-alpine

WORKDIR /app

RUN apk add --update make

COPY ./ ./

RUN make install

RUN make build

CMD ["./jsonops"]