FROM golang:1.20 AS builder

WORKDIR /app
COPY . .
RUN go build -mod=vendor -o benchmark-server ./cmd/server

FROM ubuntu:22.04

WORKDIR /app
COPY --from=builder /app/benchmark-server .
COPY --from=builder /app/data ./data

ENTRYPOINT [ "/app/benchmark-server" ]
