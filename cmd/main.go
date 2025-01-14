package main

import (
	"fmt"
	"log"
	"minishop/internal/config"
	"minishop/pkg/di"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName:      config.Viper.App.Name,
		ReadTimeout:  time.Minute * time.Duration(config.Viper.Server.Timeout),
		WriteTimeout: time.Minute * time.Duration(config.Viper.Server.Timeout),
		IdleTimeout:  time.Minute * time.Duration(config.Viper.Server.Timeout),
		Prefork:      true,
	})
	app.Use(logger.New())
	di.DependencyInjection(app)
	log.Printf("version %s", config.Viper.App.Version)
	log.Printf("environtment %s", config.Viper.Server.Environment)
	app.Listen(":" + fmt.Sprint(config.Viper.Server.Port))
}
