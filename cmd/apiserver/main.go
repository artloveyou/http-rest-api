package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/artloveyou/http-rest-api/internal/app/apiserver"
	"log"
)

var (
	// переменная для хранения флага
	configPath string
)

func init() {
	// создаем флаг              // название         // где хранится значение              // описание
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

func main() {
	// парсим флаги
	flag.Parse()

	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	// Создаем s сервер
	s := apiserver.New(config)
	// Пробуем запустить // Если ошибка
	if err := s.Start(); err != nil {
		log.Fatal(err) // пишем лог
	}
}
