### Go http-rest-api tutorial from https://www.youtube.com/watch?v=LxJLuW5aUDQ

Сервер написан на windows поэтому для деплоя надо переключить среду 
компиляции под linux
```
// to linux
set GOARCH=amd64
set GOOS=linux

// проверка среды компиляции
go env 

// to windows
set GOARCH=amd64
set GOOS=windows

