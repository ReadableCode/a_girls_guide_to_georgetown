package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

var templatesDir = "../frontend/templates"
var articlesDir = "../frontend/templates/articles"

type Article struct {
	Title string `json:"title"`
	Link  string `json:"link"`
}

func main() {
	app := fiber.New()

	// Enable logging
	app.Use(logger.New())

	// Serve static files
	app.Static("/static", "../frontend/static")
	app.Static("/css", "../frontend/css")

	// API endpoint to list all articles dynamically
	app.Get("/api/articles", func(c *fiber.Ctx) error {
		articles := []Article{}

		// Scan the articles directory for .html files
		if err := filepath.WalkDir(articlesDir, func(path string, d os.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if !d.IsDir() && filepath.Ext(path) == ".html" {
				filename := filepath.Base(path)
				title := strings.TrimSuffix(filename, ".html")
				title = strings.ReplaceAll(title, "-", " ") // Format title nicely

				articles = append(articles, Article{
					Title: strings.Title(title), // Capitalize first letters
					Link:  "/articles/" + strings.TrimSuffix(filename, ".html"),
				})
			}
			return nil
		}); err != nil {
			return err
		}

		return c.JSON(articles) // Return the list of articles as JSON
	})

	// Serve all pages in templates/ and templates/articles/
	if err := filepath.WalkDir(templatesDir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && filepath.Ext(path) == ".html" {
			// Generate the correct route based on the relative path
			relPath, _ := filepath.Rel(templatesDir, path)
			relPath = filepath.ToSlash(relPath) // Ensure cross-platform compatibility
			route := "/" + strings.TrimSuffix(relPath, filepath.Ext(relPath))

			// Ensure that index.html is served at the root (/) path
			if route == "/index" {
				route = "/"
			}

			fmt.Println("✅ Registering route:", route, "->", path) // Debugging log

			app.Get(route, func(c *fiber.Ctx) error {
				return c.SendFile(path) // Serve the page file
			})
		}
		return nil
	}); err != nil {
		log.Fatalf("Error walking templates directory: %v", err)
	}

	log.Fatal(app.Listen(":8504"))
}
