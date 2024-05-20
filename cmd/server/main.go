package main

import (
	"os"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	dist, set := os.LookupEnv("KUSSUMA_WWW")
	
	if !set {
		dist = "dist"
	}

	app.Static("/", dist)

	app.Listen("127.0.0.1:3000")
}
