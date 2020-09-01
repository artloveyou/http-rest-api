package apiserver

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

//   Запуск
// > go test ./internal/app/apiserver
func TestAPIServer_HandleHello(t *testing.T) {
	// Создаем наш сервер который принимает NewConfig()
	s := New(NewConfig())
	// Рекордер
	rec := httptest.NewRecorder()
	// Реквест
	                              // методом гет  // по адресу   // передаем ничего (просто вызов)
	req, _ := http.NewRequest(http.MethodGet, "/hello", nil)
	s.handleHello().ServeHTTP(rec, req)
	// Проверим что боди возвращает строку равную "Hello"
	assert.Equal(t, rec.Body.String(), "Hello")
}