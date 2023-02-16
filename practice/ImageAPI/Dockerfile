# syntax=docker/dockerfile:1

FROM golang:1.15.7-alpine

RUN mkdir /app
ADD . /app
WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /app

EXPOSE 4444 

CMD [ "./main"]