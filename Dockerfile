FROM golang:1.19.1-alpine3.16

WORKDIR /

COPY . .

RUN go mod download && go mod verify
