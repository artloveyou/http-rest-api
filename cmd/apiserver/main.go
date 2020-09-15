package main

import (
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/artloveyou/http-rest-api/internal/app/apiserver"
	"log"
)

var (
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

	flag.Parse()

	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	if err := apiserver.Start(config); err != nil {
		log.Fatal(err)
	}
}
