package main

import (
	"os"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	www, set := os.LookupEnv("KUSSUMA_WWW")
	
	if !set {
		www = "www"
	}

	app.Static("/", www)

	app.Listen("127.0.0.1:3000")
}
