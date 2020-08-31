package apiserver

// Config
// Стракт конфигурации сервера
type Config struct {
	// Адрес по которому запускается сервер
	BindAddr string `toml:"bind_addr"`
}

// NewConfig
// Функция для удобства отдает некий конфиг по умолчанию
// с дефолтными параметрами
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
	}
}