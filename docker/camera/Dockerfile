# syntax=docker/dockerfile:1

FROM golang:1.19

RUN apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y libopencv-dev

WORKDIR /usr/local/bin/app

COPY go.mod ./
RUN go mod download && go mod verify

RUN mv /usr/include/opencv4/opencv2 /usr/include/opencv2

COPY . .
RUN go build -v -o /app

CMD ["/app"]
