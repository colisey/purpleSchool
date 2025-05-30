package main

import (
	"os"

	"go/adv-demo/internal/link"
	"go/adv-demo/internal/user"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// go run migrations/auto.go Команда миграции
func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(postgres.Open(os.Getenv("DB_DSN")), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(
		&link.Link{},
		&user.User{},
	)
}
