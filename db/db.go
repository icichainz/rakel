package db

import (
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

var DB *gorm.DB

// Paste represents the structure of a paste in the database
type Paste struct {
    ID      uint   `json:"id" gorm:"primaryKey"`
    Title   string `json:"title"`
    Content string `json:"content"`
}

// InitDB initializes the SQLite database and runs migrations
func InitDB() {
    var err error
    DB, err = gorm.Open(sqlite.Open("pastebin.db"), &gorm.Config{})
    if err != nil {
        panic("Failed to connect to database")
    }

    // Auto-migrate the Paste model
    DB.AutoMigrate(&Paste{})
}