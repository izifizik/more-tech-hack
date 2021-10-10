# More.tech 3.0

## Стек бэкенда

* Go 1.17
* Postgres
* Keycloak

## Пример .env файла

```
PORT=8080
COOKIE_DATAHUB=""
DATAHUB_URL="http://datahub.yc.pbd.ai:9002/api/graphql"
KC_CLIENT_PATH=""
KC_ADMIN_USERNAME=""
KC_ADMIN_PASSWORD=""
KC_REALM="dima"
KC_SECRET=""
KC_CLIENT="golang-app"

PORT_DB=5432
USER_DB="vtb"
NAME_DB="vtb"
PASSWD_DB=""
HOST_DB=""
```

## Запуск проекта

`go run cmd/main.go`
