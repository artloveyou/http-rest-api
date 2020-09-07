package store_test

import (
	"github.com/artloveyou/http-rest-api/internal/app/model"
	"github.com/artloveyou/http-rest-api/internal/app/store"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	// Тестим создание юзера
	u, err := s.User().Create(model.TestUser(t))
	// Не должно быть ошибок
	// Юзер не равен нулю
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	// Сначала тестим на отсутствие пользователя
	// по запрошенному емайлу
	email := "test@email.io"
	_, err := s.User().FindByEmail(email)
	// Ожидаем ошибку
	assert.Error(t, err)

	u := model.TestUser(t)
	u.Email = email
	// Создаем юзера
	s.User().Create(u)

	// u <-- помещаем юзера в переменную
	u, err = s.User().FindByEmail(email)

	// теперь юзер должен найтись
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
