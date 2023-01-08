# syntax=docker/dockerfile:1

FROM golang:1.19

RUN apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y libopencv-dev

WORKDIR /usr/local/bin/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/app ./...

CMD ["app"]
