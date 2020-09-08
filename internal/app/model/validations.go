package model

import validation "github.com/go-ozzo/ozzo-validation"
               // поле EncryptedPassword не пустое, значит юзер уже создан и Password не надо валидировать
func requiredIf(cond bool) validation.RuleFunc{

	return func(value interface{}) error {
		// создаем юзера, пароль валидируем
		if cond {
			return validation.Validate(value, validation.Required)
		}
		// если обновляем юзера - пароль не валидируем
		return nil
	}
}
