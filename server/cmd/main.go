package main

import (
	"fmt"
	"net/http"

	"go/adv-demo/configs"
	"go/adv-demo/internal/auth"
	"go/adv-demo/internal/link"
	"go/adv-demo/internal/user"
	"go/adv-demo/pkg/db"
	"go/adv-demo/pkg/middleware"
)

func main() {
	conf := configs.LoadConfig()
	db := db.NewDb(conf)
	router := http.NewServeMux()

	// Repositories
	LinkRepository := link.NewLinkRepository(db)
	UserRepository := user.NewUserRepository(db)

	// Services
	authService := auth.NewAuthService(UserRepository)

	// Handlers
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		AuthConfig:  &conf.Auth,
		AuthService: authService,
	})
	// fmt.Println("Main Auth after")
	link.NewLinkHandler(router, link.LinkHandlerDeps{
		LinkRepository: LinkRepository,
	})
	// fmt.Println("Main Link after")

	// Middlewares
	stack := middleware.Chain(
		middleware.CORS,
		middleware.Logging,
		// middleware.IsAuthed,
	)

	server := http.Server{
		Addr: ":8081",
		// Handler: router,
		Handler: stack(router),
	}

	fmt.Println("Server is listening on port 8081")
	server.ListenAndServe()
}
