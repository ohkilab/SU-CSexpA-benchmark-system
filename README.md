# SU-CSexpA-benchmark-system

## how to work

1. create `.env`

```shell
$ cp .env.sample .env
```

2. run

```shell
$ make up
```

3. if you want to test, generate seeds

```shell
$ cd backend
$ make seed
```

## test

- e2e test

```shell
$ make e2e_test
```
