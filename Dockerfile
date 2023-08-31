FROM golang:1.20-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o /app/api /cmd/main.go

FROM alpine:latest

WORKDIR /app/
COPY --from=builder /app/api /app/.env ./

CMD [ "/app/api" ]