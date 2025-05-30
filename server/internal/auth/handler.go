package auth

import (
	"fmt"
	"net/http"

	"go/adv-demo/configs"
	"go/adv-demo/pkg/request"
	"go/adv-demo/pkg/response"
)

type (
	AuthHandlerDeps struct {
		*configs.AuthConfig
		*AuthService
	}
	AuthHandler struct {
		*configs.AuthConfig
		*AuthService
	}
)

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		AuthConfig:  deps.AuthConfig,
		AuthService: deps.AuthService,
	}
	router.HandleFunc("POST /auth/login", handler.Login)
	router.HandleFunc("POST /auth/register", handler.Register)
}

func (handler *AuthHandler) Login(w http.ResponseWriter, req *http.Request) {
	body, err := request.HandleBody[LoginRequest](w, req)
	if err != nil {
		return
	}
	fmt.Println(body)

	data := &LoginResponse{
		Token: handler.Secret,
	}
	response.Json(w, data, http.StatusOK)
}

func (handler *AuthHandler) Register(w http.ResponseWriter, req *http.Request) {
	body, err := request.HandleBody[RegisterRequest](w, req)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	email, err := handler.AuthService.Register(body.Email, body.Password, body.Name)
	if err != nil {
		response.Json(w, err.Error(), http.StatusConflict)
		// fmt.Println(err.Error())
		return
	}

	data := &RegisterResponse{
		Token: email,
		// Token: handler.Secret,
	}
	response.Json(w, data, http.StatusOK)
}

// fmt.Println("Register")            // В консоль
// fmt.Fprint(w, "<h1>Register</h1>") // В браузер
// Проверка через Regexp
// if payload.Email == "" || payload.Password == "" {
// 	res.Json(w, "Чего то не хватило", http.StatusUnauthorized)
// 	return
// }

// reg, _ := regexp.Compile(`[A-Za-z0-9\._%+\-]+@[A-Za-z0-9\.\-]+\.[A-Za-z]{2,}`)
// if !reg.MatchString(payload.Email) {
// 	res.Json(w, "Битый Email", http.StatusUnauthorized)
// 	return
// }
// mailAddress, err := mail.ParseAddress(payload.Email) // Встроенный парсер Email

// match, _ := regexp.MatchString(`[A-Za-z0-9\._%+\-]+@[A-Za-z0-9\.\-]+\.[A-Za-z]{2,}`, payload.Email)
// if !match {
// 	res.Json(w, "Битый Email", http.StatusUnauthorized)
// 	return
// }

// fmt.Println(payload)
