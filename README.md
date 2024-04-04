# some-service

Проект для освоения работы с кодогенератором oapi-codegen

Подробнее о oapi-codegen: https://github.com/deepmap/oapi-codegen

Установить oapi-codegen:
```
go install github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@latest
```

Для запуска oapi-codegen:
```
oapi-codegen -generate echo,strict-server,spec,types -package handler openapi.yaml > internal/handler/http.gen.go
```

Установить все зависимости:
```
go mod tidy
```