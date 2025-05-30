package middleware

import (
	"log"
	"net/http"
	"strings"
)

func IsAuthed(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")
		if authorization == "" {
			// next.ServeHTTP(w, r)
			log.Println("Я вас не звал, идите на хуй!")
			return
		}
		// token := strings.Split(authorization, " ")[1]
		token := strings.TrimPrefix(authorization, "Bearer ")
		log.Println("Token = " + token)

		next.ServeHTTP(w, r)
	})
}
