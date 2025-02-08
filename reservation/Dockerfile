FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o /app/main .

FROM alpine:latest

COPY --from=builder /app/main /app/main

CMD ["/app/main"]