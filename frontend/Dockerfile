# Note: build from repository root
#  e.g.: docker build -t exp-a-frontend -f ./frontend/Dockerfile .
# build
FROM node:18.17-bookworm-slim AS builder

RUN npm i -g pnpm@8.15.9

WORKDIR /app/proto-gen/web
COPY ./proto-gen/web/package.json ./proto-gen/web/pnpm-lock.yaml ./
RUN pnpm install --frozen-lockfile

WORKDIR /app/frontend
COPY ./frontend/package.json ./frontend/pnpm-lock.yaml ./
RUN pnpm install --frozen-lockfile

COPY ./proto-gen /app/proto-gen
COPY ./frontend /app/frontend
RUN pnpm build

# serve in nginx
FROM nginx:1.24

COPY frontend/nginx.conf /etc/nginx/conf.d/default.conf
COPY --from=builder /app/frontend/dist /usr/share/nginx/html
