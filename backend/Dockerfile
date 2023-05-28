FROM golang:1.20 AS builder

WORKDIR /src
COPY . .
RUN go build \
    -mod=vendor \
    -tags timetzdata \
    -o backend-server \
    ./cmd/server

FROM ubuntu:22.04

WORKDIR /app
COPY --from=builder /src/backend-server .
COPY --from=builder /src/storage ./storage

ENTRYPOINT [ "/app/backend-server" ]
