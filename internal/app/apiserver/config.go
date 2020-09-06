package apiserver

import "github.com/artloveyou/http-rest-api/internal/app/store"

// Config
// Стракт конфигурации сервера
type Config struct {
	// Адрес по которому запускается сервер
	BindAddr string `toml:"bind_addr"`
    // Уровень логирования
	LogLevel string `toml:"log_level"`
	Store *store.Config
}

// NewConfig
// Функция для удобства отдает некий конфиг по умолчанию
// с дефолтными параметрами
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
		Store: store.NewConfig(),
	}
}