FROM golang:1.25 AS builder

WORKDIR /app

COPY . .

RUN go build -o app .

FROM debian:stable-slim

WORKDIR /app

COPY --from=builder /app/app .

EXPOSE 8084

CMD ["./app"]
