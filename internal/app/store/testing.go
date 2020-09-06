package store

import (
	"fmt"
	"strings"
	"testing"
)

// TestStore ...                                              // собираем таблицы
func TestStore(t *testing.T, databaseURL string) (*Store, func(...string)) {
	// Указывает тестам, что это тестовый метод и его не надо тестировать
	t.Helper()

	// Создаем конфиг
	config := NewConfig()
	config.DatabaseURL = databaseURL

	// Создаем хранилище
	// store
	s :=New(config)
	// Пробуем открыть
	if err := s.Open(); err != nil {
		t.Fatal(err)
	}

    // Если нет ошибок
	return s, func(tables ...string) {
		// И если таблицы вообще были
		if len(tables) > 0 {
			                                   // Очищаем таблицы от тестовых данных  // подставляем объединенные через запятую таблицы
			if _, err := s.db.Exec(fmt.Sprintf("TRUNCATE %s CASCADE", strings.Join(tables, ", "))); err != nil {
				t.Fatal(err)
			}
		}
        // Не забыть закрыть стор
		s.Close()
	}
}
