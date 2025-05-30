package request

import (
	"fmt"
	"regexp"

	"github.com/go-playground/validator/v10"
)

func IsValid[T any](payload T) error {
	validate := validator.New()

	// Регистрация кастомного валидатора с тегом "password"
	validate.RegisterValidation("password", validatePassword)

	err := validate.Struct(payload)
	if err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			for _, valErr := range validationErrors {
				// fmt.Println(valErr.Field(), valErr.Tag())
				// valErr.Field() — имя поля
				// valErr.Tag() — название проверки, которую поле не прошло
				fmt.Printf("Поле %s не прошло проверку %s\n", valErr.Field(), valErr.Tag())
			}
		}
		return err
	}
	// if err != nil {
	// 	// response.Json(w, err.Error(), http.StatusBadRequest)
	// 	return err
	// }
	return nil
}

// validatePassword — кастомная функция валидации пароля
func validatePassword(fl validator.FieldLevel) bool {
	password := fl.Field().String()

	// Проверка на наличие хотя бы одного спецсимвола
	matched, _ := regexp.MatchString(`[!@#~$%^&*()_+{}\[\]:;<>,.?\\/-]`, password)
	return matched
}
