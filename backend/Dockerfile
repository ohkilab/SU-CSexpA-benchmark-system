FROM go:1.20 AS builder

WORKDIR /app
COPY . .
RUN go build -mod=read-only -o backend-server

FROM ubuntu:22.04

WORKDIR /app
COPY --from=builder /app/backend-server .

ENTRYPOINT [ "/app/backend-server" ]
