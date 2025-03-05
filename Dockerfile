FROM golang:latest

WORKDIR /app

COPY . .

RUN go build -o server main.go

CMD ["./server", "--host", "0.0.0.0"]