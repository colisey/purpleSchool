package request

// import (
// 	"fmt"
// 	"regexp"

// 	"github.com/go-playground/locales/ru"
// 	ut "github.com/go-playground/universal-translator"
// 	"github.com/go-playground/validator/v10"
// 	rutrans "github.com/go-playground/validator/v10/translations/ru"
// )

// type RegisterRequest struct {
// 	Username string `validate:"required"`
// 	Password string `validate:"required,password"`
// }

// var (
// 	validate *validator.Validate
// 	trans    ut.Translator
// )

// func main() {
// 	// Создаем валидатор и переводчик
// 	validate = validator.New()

// 	russian := ru.New()
// 	uni := ut.New(russian, russian)
// 	trans, _ = uni.GetTranslator("ru")

// 	// Регистрируем стандартные русские переводы
// 	_ = rutrans.RegisterDefaultTranslations(validate, trans)

// 	// Регистрируем кастомный тег "password"
// 	validate.RegisterValidation("password", validatePassword)

// 	// Регистрируем перевод для "password"
// 	validate.RegisterTranslation("password", trans, func(ut ut.Translator) error {
// 		return ut.Add("password", "Пароль должен содержать хотя бы один спецсимвол", true)
// 	}, func(ut ut.Translator, fe validator.FieldError) string {
// 		t, _ := ut.T("password", fe.Field())
// 		return t
// 	})

// 	// Пример запроса
// 	req := RegisterRequest{
// 		Username: "admin",
// 		Password: "123456", // нет спецсимвола
// 	}

// 	// Валидируем
// 	err := validate.Struct(req)
// 	if err != nil {
// 		for _, e := range err.(validator.ValidationErrors) {
// 			fmt.Println(e.Translate(trans))
// 		}
// 	}
// }

// // validatePassword — кастомный валидатор
// func validatePassword(fl validator.FieldLevel) bool {
// 	password := fl.Field().String()
// 	matched, _ := regexp.MatchString(`[!@#\$%\^&\*\(\)_\+\-=\[\]{};':"\\|,.<>\/?]`, password)
// 	return matched
// }
