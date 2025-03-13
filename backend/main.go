package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Define the frontend path
	frontendPath := "../frontend/templates/"

	// Debugging: Check if index.html exists
	indexPath := frontendPath + "index.html"
	if _, err := os.Stat(indexPath); os.IsNotExist(err) {
		log.Fatalf("ERROR: File %s does not exist", indexPath)
	} else {
		fmt.Println("File exists:", indexPath)
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
