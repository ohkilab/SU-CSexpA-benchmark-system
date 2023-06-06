FROM golang:1.20 AS builder

WORKDIR /src
COPY . .
RUN go build \
    -mod=vendor \
    -tags timetzdata \
    -o backend-server \
    ./cmd/server

FROM ubuntu:22.04

RUN apt-get update && apt-get install -y tzdata \
    && rm -rf /var/lib/apt/lists/*
ENV TZ Asia/Tokyo

WORKDIR /app
COPY --from=builder /src/backend-server .
COPY --from=builder /src/storage ./storage

ENTRYPOINT [ "/app/backend-server" ]
