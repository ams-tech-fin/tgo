package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

type CacheController struct {
	cache *redis.Client
}

func NewCacheController(cache *redis.Client) *CacheController {
	return &CacheController{cache: cache}
}

func (c *CacheController) CacheTest(ctx *fiber.Ctx) error {
	type Request struct {
		Data string `json:"data"`
	}

	var req Request
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}

	// Cache example
	v := context.Background()
	fmt.Println("cacheado", c.cache.Get(v, "bla"))

	valueRaw := fiber.Map{"data": req.Data}
	value, _ := json.Marshal(valueRaw)
	c.cache.Set(v, "bla", value, time.Second*10)

	return ctx.Status(fiber.StatusOK).JSON(valueRaw)
}
