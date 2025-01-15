FROM golang:1.20

WORKDIR /app

COPY . /app

RUN go build -o receipt-processor

EXPOSE 8080

CMD ["./receipt-processor"]
