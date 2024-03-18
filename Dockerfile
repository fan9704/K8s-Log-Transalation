FROM golang:latest

RUN mkdir -p /usr/local/go/src/ironman-2021
WORKDIR /usr/local/go/src/ironman-2021
ADD . /usr/local/go/src/ironman-2021

RUN go mod download
RUN go build ./main.go

EXPOSE 8080
CMD ["./main"]