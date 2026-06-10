package handlers

import (
	"context"
	"encoding/json"
	"golang_redis_pgsql/model"
	"golang_redis_pgsql/pkg/database"
	"golang_redis_pgsql/pkg/redis"
	"time"

	"github.com/gofiber/fiber/v2"
	goredis "github.com/redis/go-redis/v9"
)

var (
	ctx = context.Background()
)

func GetAllUsers(c *fiber.Ctx) error {
	val, err := redis.RedisDB.Get(ctx, "all_users").Result()
	if err == goredis.Nil {
		var users []model.User
		if err := database.DB.Find(&users).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Postgres DB Error"})
		}
		data, _ := json.Marshal(users)
		redis.RedisDB.Set(ctx, "all_users", data, 10*time.Minute)
		return c.Status(fiber.StatusOK).JSON(users)
	} else if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Redis DB Error"})
	}
	var users []model.User
	if err := json.Unmarshal([]byte(val), &users); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error while Unmarshal"})
	}
	return c.Status(fiber.StatusOK).JSON(users)
}

func CreateUser(c *fiber.Ctx) error {
	users := new(model.User)
	if err := c.BodyParser(users); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Request"})
	}
	var lastUser model.User
	if err := database.DB.Order("id desc").First(&lastUser).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Database Error"})
	}
	users.ID = lastUser.ID + 1
	err := database.DB.Create(&users).Error
	data, _ := json.Marshal(users)
	redis.RedisDB.Del(ctx, "all_users")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Database Error"})
	}
	var u []model.User
	if err := database.DB.Find(&u).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Postgres DB Error"})
	}
	data, _ = json.Marshal(u)
	redis.RedisDB.Set(ctx, "all_users", data, 10*time.Minute)
	return c.Status(fiber.StatusCreated).JSON(users)
}
