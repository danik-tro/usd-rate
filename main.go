package main

import (
	"github.com/danik-tro/usd-rate/pkg/core"
	"github.com/danik-tro/usd-rate/pkg/handlers"
	storage "github.com/danik-tro/usd-rate/pkg/storages"
	"github.com/danik-tro/usd-rate/pkg/utils"
	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/robfig/cron/v3"
)

func main() {
	app := fiber.New()
	c := core.LoadConfig()
	s := storage.NewStorage(c)
	cache := storage.NewCache(c.RedisAddr)
	rh := handlers.NewRateHandler(&c, &s, cache)

	app.Get("/rate", rh.Rate)
	app.Post("/subscribe", rh.Subscribe)
	app.Post("/send-emails", rh.SendEmails)

	cfg := swagger.Config{
		BasePath: "/",
		FilePath: "./docs/swagger.json",
		Path:     "docs",
		Title:    "Swagger API Docs",
	}
	app.Use(swagger.New(cfg))

	cron := cron.New(cron.WithSeconds())

	cron.AddFunc("0 0 12 * * *", func() {
		utils.CronAction()
	})

	cron.Start()
	defer cron.Stop()
	app.Listen("0.0.0.0:8000")
}
