FROM golang:latest

RUN mkdir /app
COPY ./src /app

RUN apt-get -y update && apt-get -y install git
RUN go get github.com/gin-gonic/gin
RUN go get github.com/lib/pq

WORKDIR /app
