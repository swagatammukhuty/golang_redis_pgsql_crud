package main

import (
	"golang_redis_pgsql/internal/handlers"
	"golang_redis_pgsql/pkg/database"
	"golang_redis_pgsql/pkg/redis"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

const serverPort = ":8080"

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}
	database.InitPgDB()
	redis.InitRedis()
	app := fiber.New()
	app.Get("/getAllUsers", func(c *fiber.Ctx) error {
		return handlers.GetAllUsers(c)
	})
	log.Printf("Server is running on the port %v", serverPort)
	err := app.Listen(serverPort)
	if err != nil {
		log.Fatal("Error while listening the port")
	}
}
