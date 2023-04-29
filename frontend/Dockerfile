FROM node:18 AS builder

RUN npm i -g pnpm serve

WORKDIR /app/proto-gen

COPY proto-gen .

WORKDIR /app/frontend

COPY frontend/package.json frontend/pnpm-lock.yaml ./
RUN pnpm i

COPY frontend .

RUN pnpm build

FROM nginx:1.24

COPY --from=builder /app/frontend/dist /usr/share/nginx/html