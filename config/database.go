package config

import (
    "log"

    "github.com/glebarez/sqlite" // Driver sem Cgo
    "gorm.io/gorm"
    "go-users-api/models"
)

var DB *gorm.DB

func ConnectDB() {
    database, err := gorm.Open(sqlite.Open("users.db"), &gorm.Config{})
    if err != nil {
        log.Fatal("❌ Erro ao conectar ao banco:", err)
    }

    database.AutoMigrate(&models.User{})
    DB = database

    log.Println("📦 Banco conectado e migrado com sucesso")
}
