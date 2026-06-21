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

`make up` は公開用の sample V2026 bundle を read-only mount して起動します。sample は動作確認用であり、公式ベンチマークケースではありません。

## V2026 official data

公式の `v2026.json` とタグ txt 群は public repository と public Docker image に含めません。公式データは `expA-admin-tools` で bundle として生成し、サーバ上の非公開ディレクトリから実行時に read-only mount してください。

bundle の形は次の通りです。

```text
v2026/
  v2026.json
  manifest.json
  tags/
    random.txt
    1.txt
    ...
    10.txt
```

公式データで起動する場合:

```shell
$ V2026_BUNDLE_DIR=/secure/path/v2026 make up-official
```

`V2026_BUNDLE_DIR` は `.env` に書いておくこともできます。一時的に別の bundle を使う場合は、上のようにコマンドの環境変数で上書きできます。

起動後は管理画面の「コンテスト作成」で validator に `V2026` を選んで contest を作成します。V2026 ではタグ本文は入力せず、bundle 内の既存タグファイルを使います。`slug` は URL/API 上の contest slug、`tag slug` は `/app/storage/tags/{tag slug}` のディレクトリ名です。

通常の V2026 tag slug は `v2026` です。別のタグセットを使う場合は、mount 先のタグディレクトリ名を管理画面の `tag slug` に合わせます。

```shell
$ V2026_BUNDLE_DIR=/secure/path/v2026 V2026_CONTEST_SLUG=exp-a-2026-final-tags make up-official
```

公式データを配置してから Docker image を build しないでください。benchmark-service の image は `v2026.json` を含まず、`/app/data/v2026.json` を volume mount して読み込みます。

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
