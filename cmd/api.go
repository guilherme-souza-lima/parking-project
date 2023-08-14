package cmd

import (
	"ProjetoEstacionamento/infra"
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"time"
)

func StartHTTP(ctx context.Context, container *infra.Container) {
	app := fiber.New(fiber.Config{
		StrictRouting: true,
	})

	go func() {
		for {
			select {
			case <-ctx.Done():
				if err := app.Shutdown(); err != nil {
					panic(err)
				}
				return
			default:
				time.Sleep(1 * time.Second)
			}
		}
	}()

	app.Use(cors.New(cors.Config{
		AllowHeaders: "*",
	}))

	app.Get("/parking-info", container.ParkingHandler.Info)
	app.Post("/parking/occupy/:vehicle", container.ParkingHandler.Occupy)
	app.Post("/parking/release/:vehicle", container.ParkingHandler.Release)

	err := app.Listen(":8080")
	if err != nil {
		panic(err)
	}
}
