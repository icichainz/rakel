package handlers

import (
	"fmt"
	"net/http"
	"rakel/db"

	"github.com/gofiber/fiber/v2"
)

// CreatePaste handles the creation of a new paste
func CreatePaste(c *fiber.Ctx) error {
    // Parse form data
    title := c.FormValue("title")
    content := c.FormValue("content")

    // Validate input
    if content == "" {
        return c.Status(http.StatusBadRequest).SendString("Content is required")
    }

    // Create a new paste
    paste := db.Paste{
        Title:   title,
        Content: content,
    }

    // Save to the database
    result := db.DB.Create(&paste)
    if result.Error != nil {
        return c.Status(http.StatusInternalServerError).SendString("Failed to create paste")
    }

    // Redirect to the paste view page (we'll implement this later)
    return c.Redirect(fmt.Sprintf("/pastes/%d", paste.ID))
}