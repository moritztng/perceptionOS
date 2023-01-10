# syntax=docker/dockerfile:1

FROM golang:1.19

RUN apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y libopencv-dev && mv /usr/include/opencv4/opencv2 /usr/include/opencv2

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /app

CMD ["/app"]
