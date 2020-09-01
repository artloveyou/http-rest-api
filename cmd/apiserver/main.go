package main

import (
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/artloveyou/http-rest-api/internal/app/apiserver"
	"log"
)

var (
	// переменная для хранения флага
	configPath string
)

func init() {
	// создаем флаг              // название флага        // где хранится значение              // описание
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
	// > apiserver.exe -h                                выведет хелп в консоли
	// > apiserver.exe -config-path another/config.toml  подключит другой конфиг
}

func main() {
	fmt.Println("It's run")

	// парсим флаги
	flag.Parse()

	config := apiserver.NewConfig()

	                          // что декодим // куда записываем
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
