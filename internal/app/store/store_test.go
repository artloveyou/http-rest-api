package store_test

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

// Вызывается один раз во всем пакете
// Тут можно делать разовые манипуляции перед тестами
func TestMain(m *testing.M) {
	databaseURL = os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "host=localhost user=postgres password=root dbname=restapi_test sslmode=disable"
	}
    // Не забываем выйти
	os.Exit(m.Run())
}
