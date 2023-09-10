# Note: build from repository root
#  e.g.: docker build -t exp-a-frontend -f ./frontend/Dockerfile .
# build
FROM node:18.17-alpine3.18 AS builder

RUN npm i -g pnpm

# WORKDIR /app/proto-gen
COPY ./proto-gen /app/proto-gen

WORKDIR /app/proto-gen/web
RUN pnpm i --prod

# WORKDIR /app/frontend
COPY ./frontend /app/frontend

WORKDIR /app/frontend
RUN pnpm i

COPY . .
RUN pnpm build

# serve in nginx
FROM nginx:1.24

COPY frontend/nginx.conf /etc/nginx/conf.d/default.conf
COPY --from=builder /app/frontend/dist /usr/share/nginx/html
