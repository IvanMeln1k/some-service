# some-service

Проект для освоения работы с кодогенератором oapi-codegen

Для запуска oapi-codegen:
```
oapi-codegen -generate echo,strict-server,spec,types -package handler openapi.yaml > internal/handler/http.gen.go
```

Установить все зависимости:
```
go mod tidy
```