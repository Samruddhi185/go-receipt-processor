FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY fetch-receipt-processor/*.go ./

RUN go build -o /receipt-processor

EXPOSE 8000

CMD ["/receipt-processor"]