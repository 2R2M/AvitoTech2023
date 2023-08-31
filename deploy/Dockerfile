FROM golang:1.20.0-alpine3.16 as builder
RUN mkdir /build
COPY . /build/
WORKDIR /build
RUN CGO_ENABLED=0 GOOS=linux go build -a -o server cmd/main.go

# generate clean, final image for end users
FROM alpine:3.16.0
WORKDIR /app
COPY --from=builder /build/server /app
COPY --from=builder /build/.env /app
# executable
ENTRYPOINT [ "./server" ]
