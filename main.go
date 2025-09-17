package main

import (
    "log"

    "github.com/gofiber/fiber/v2"
    "go-users-api/config"
    "go-users-api/routes"
)

func main() {
    config.ConnectDB()

    app := fiber.New()
    routes.SetupRoutes(app)

    log.Println("🚀 Servidor rodando em http://localhost:3000")
    log.Fatal(app.Listen(":3000"))
}
