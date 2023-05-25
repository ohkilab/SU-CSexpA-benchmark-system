FROM node:18 AS builder

RUN npm i -g pnpm serve

WORKDIR /app/proto-gen

COPY proto-gen .

WORKDIR /app/proto-gen/web

RUN pnpm i

WORKDIR /app/frontend

COPY frontend/package.json frontend/pnpm-lock.yaml ./
RUN pnpm i

COPY frontend .

RUN pnpm build

FROM nginx:1.24

COPY frontend/nginx.conf /etc/nginx/conf.d/default.conf

COPY --from=builder /app/frontend/dist /usr/share/nginx/html
