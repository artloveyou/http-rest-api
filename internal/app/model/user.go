package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"golang.org/x/crypto/bcrypt"
)

// User ...
type User struct {
	ID                int
	Email             string
	Password          string
	EncryptedPassword string
}

// Validate
func (u *User) Validate() error {
	return validation.ValidateStruct(
		u,
		validation.Field(&u.Email, validation.Required, is.Email),
		                                            // если создаем юзера - валидируем пароль
		validation.Field(&u.Password, validation.By(requiredIf(u.EncryptedPassword == "")), validation.Length(6, 100)),
		)
}

// BeforeCreate
func (u *User) BeforeCreate() error {
	if len(u.Password) > 0 {
		// Шифруем пароль
		enc, err := encryptString(u.Password)
		if err != nil {
			return err
		}
		// Записываем зашифрованную строку
		u.EncryptedPassword = enc
	}

	return nil
}

// принимает // возвращает
func encryptString(s string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		// пустая строка и ошибка
		return "", err
	}
	// зашифрованная строка // нулевая ошибка
	return string(b), nil
}
