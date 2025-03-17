package main

import (
    "rakel/db"
    "rakel/handlers"

    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/template/html/v2"
)

func main() {
    // Initialize the database
    db.InitDB()

    // Create a new Fiber app with HTML templating
    engine := html.New("./templates", ".html")
    app := fiber.New(fiber.Config{
        Views: engine,
    })

    // Serve static files
    app.Static("/static", "./static")

    // Routes
    app.Get("/", func(c *fiber.Ctx) error {
        return c.Render("create", fiber.Map{}) // Render the create.html template
    })

    app.Post("/pastes", handlers.CreatePaste)

    // Start the server
    app.Listen(":3000")
}