package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Debugging: Print if files exist
	templates := []string{
		"../frontend/templates/index.html",
		"../frontend/templates/about.html",
		"../frontend/templates/projects.html",
		"../frontend/templates/contact.html",
	}

	for _, path := range templates {
		if _, err := os.Stat(path); os.IsNotExist(err) {
			log.Fatalf("ERROR: File %s does not exist", path)
		} else {
			fmt.Println("File exists:", path)
		}
	}

	// Serve static files
	app.Static("/static", "../frontend/static")
	app.Static("/css", "../frontend/css")

	// Serve HTML pages
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendFile("../frontend/templates/index.html")
	})
	app.Get("/about", func(c *fiber.Ctx) error {
		return c.SendFile("../frontend/templates/about.html")
	})
	app.Get("/projects", func(c *fiber.Ctx) error {
		return c.SendFile("../frontend/templates/projects.html")
	})
	app.Get("/contact", func(c *fiber.Ctx) error {
		return c.SendFile("../frontend/templates/contact.html")
	})

	// Start server on port 3001
	log.Fatal(app.Listen(":3001"))
}
