package middleware

import "net/http"

type WrapperWriter struct {
	http.ResponseWriter
	StatusCode int
}

func (w *WrapperWriter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
	w.StatusCode = statusCode
}

// Как корректно "пробросить" значения (например, userId после проверки авторизации)
// из одной middleware в другую или до хендлера, если стандартные параметры (w, r) никак не меняются?
// одсказка: речь про передачу данных без сторонних пакетов.
