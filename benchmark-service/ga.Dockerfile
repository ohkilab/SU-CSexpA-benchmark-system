FROM golang:1.21 AS builder

WORKDIR /app
COPY . .

WORKDIR /app/benchmark-service
RUN go mod vendor \
  && go build -mod=vendor -o benchmark-server ./cmd/server

FROM ubuntu:22.04

RUN apt-get update && apt-get install -y --no-install-recommends tzdata  \
  && rm -rf /var/lib/apt/lists/*
ENV TZ Asia/Tokyo

WORKDIR /app
COPY --from=builder /app/benchmark-service/benchmark-server .
COPY --from=builder /app/benchmark-service/data ./data

ENTRYPOINT [ "/app/benchmark-server" ]
