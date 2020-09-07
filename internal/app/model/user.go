package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
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
	return validation.ValidateStruct(u)
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