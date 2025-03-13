package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	// Ensure the log directory exists
	logDir := "../logs"
	logFile := logDir + "/access.log"
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		os.MkdirAll(logDir, 0755)
	}

	// Open log file
	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	defer file.Close()
	log.SetOutput(file) // Redirect all logs to this file

	fmt.Println("📝 Log file: ", logFile)

	// Enable Fiber's logging middleware (logs to the file)
	app.Use(logger.New(logger.Config{
		Output: file,
	}))

	// Log each request manually
	app.Use(func(c *fiber.Ctx) error {
		clientIP := c.IP()
		userAgent := c.Get("User-Agent")
		method := c.Method()
		path := c.Path()

		log.Printf("🔹 [%s] %s - %s | %s", method, path, clientIP, userAgent)
		return c.Next()
	})

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

	log.Fatal(app.Listen(":8504"))
}
