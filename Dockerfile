FROM golang:latest

RUN mkdir /app
COPY ./src /app

RUN apt -y update && apt -y install git
RUN go get github.com/gin-gonic/gin
RUN go get github.com/lib/pq

WORKDIR /app
