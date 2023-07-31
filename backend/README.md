# SU-CSexpA-benchmark-system backend

## Generate

### all

```shell
$ make generate
```

### ent(db ORM)

```shell
$ make ent/generate
```

### tbls(db schema document)

Before generate, make sure that DB(mysql) is launched on localhost:3306.

```shell
$ make tbls/generate
```

## Test

**e2e**

```shell
$ make e2e_test
```
