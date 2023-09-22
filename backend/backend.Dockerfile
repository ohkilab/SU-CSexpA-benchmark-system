FROM golang:1.21 AS builder

WORKDIR /src
COPY . .
# RUN echo $(ls -la)

WORKDIR /src/benchmark-service
# RUN echo $(ls -la)
RUN go mod vendor

WORKDIR /src/backend
RUN go mod vendor
RUN go build \
  -tags timetzdata \
  -mod=vendor \
  -o backend-server \
  ./cmd/server

FROM ubuntu:22.04

RUN apt-get update && apt-get install -y tzdata \
  && rm -rf /var/lib/apt/lists/*
ENV TZ Asia/Tokyo

WORKDIR /app
COPY --from=builder /src/backend/backend-server .

ENTRYPOINT [ "/app/backend-server" ]
