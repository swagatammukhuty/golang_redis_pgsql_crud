package main

import (
	"golang_redis_pgsql/pkg/database"
	"golang_redis_pgsql/pkg/redis"
	"golang_redis_pgsql/routes"
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
	routes.SetUpRoutes(app)
	log.Printf("Server is running on the port %v", serverPort)
	err := app.Listen(serverPort)
	if err != nil {
		log.Fatal("Error while listening the port")
	}
}
