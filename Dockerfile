FROM golang:alpine

ENV DB_CONFIG=postgres://test:123456@127.0.0.1:5432/k8slog
ENV PORT=7780

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .
EXPOSE 7780

RUN go build -o main .

CMD ["./main"]