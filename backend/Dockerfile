FROM golang:1.22 AS builder

WORKDIR /src
COPY . .
RUN go build \
    -tags timetzdata \
    -mod=vendor \
    -o backend-server \
    ./cmd/server

FROM ubuntu:22.04

RUN apt-get update && apt-get install -y --no-install-recommends tzdata \
    && rm -rf /var/lib/apt/lists/*
ENV TZ Asia/Tokyo

WORKDIR /app
COPY --from=builder /src/backend-server .

ENTRYPOINT [ "/app/backend-server" ]
