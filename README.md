# SU-CSexpA-benchmark-system

## requirements

- [docker](https://docs.docker.com/engine/install/)
- [go](https://go.dev/doc/install)

## launch

### 1. create `.env`

```shell
$ cp .env.sample .env
```

### 2. run

You can access to the web application on http://localhost:80

```shell
$ make up
```

### 3. if you want to test, generate seeds

You can login with the user(id: `ohkilab`, password: `ohkilab`)

```shell
$ cd backend
$ make batch/seed
```
