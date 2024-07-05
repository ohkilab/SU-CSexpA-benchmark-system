# SU-CSexpA-benchmark-system

## requirements

- [docker](https://docs.docker.com/engine/install/)
- [go](https://go.dev/doc/install)

## setup

put `benchmark-service/csv/tag.csv` and `benchmark-service/csv/geotag.csv`

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

## known issues

### cannot connect to the frontend via http://localhost:80 on docker desktop for mac

In docker desktop for mac, the `host` network mode is not available, so you cannot connect to the frontend in default.  
To launch services, comment out all lines which contain `network_mode=host` in `compose.yaml`. After this operation, services are launched with the `bridge` network mode.


```diff
 backend:
     build:
       context: backend
       dockerfile: Dockerfile
     env_file:
       - .env
     restart: always
-    network_mode: host
     depends_on:
       - db
       - benchmark-service
     extra_hosts:
       - "host.docker.internal:host-gateway"
```

### Frontend image has to be build from repository root
Since frontend project is dependent on files outside of its directory, the image has to be built from the project root. The following command can be used to build the frontend image.

```sh
docker build -t exp-a-frontend -f ./frontend/Dockerfile .
```
