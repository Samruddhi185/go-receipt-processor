FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY app/*.go ./

RUN go build -o /receipt-processor

EXPOSE 8000

CMD ["/receipt-processor"]