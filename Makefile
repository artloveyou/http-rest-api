# Билд приложения
.PHONY: build
build:
	go build -v ./cmd/apiserver
# > make build

# Тесты в всех каталогах внутри ./internal/... с таймаутом 30 сек
.PHONY: test
test:
	go test -v -race -timeout 30s ./internal/...

# > make test

# По умолчанию
# > make.exe
# выполнит
# > make build
.DEFAULT_GOAL := build
