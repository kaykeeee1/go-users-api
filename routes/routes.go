package routes

import (
    "strconv"
    "strings"

    "github.com/gofiber/fiber/v2"
    "go-users-api/config"
    "go-users-api/models"
)

func SetupRoutes(app *fiber.App) {
    // Rota raiz
    app.Get("/", func(c *fiber.Ctx) error {
        return c.JSON(fiber.Map{"message": "API em Go funcionando! 🚀"})
    })

    // Listar todos os usuários
    app.Get("/users", func(c *fiber.Ctx) error {
        var users []models.User
        config.DB.Find(&users)
        return c.JSON(users)
    })

    // Criar novo usuário
    app.Post("/users", func(c *fiber.Ctx) error {
        user := new(models.User)
        if err := c.BodyParser(user); err != nil {
            return c.Status(400).JSON(fiber.Map{"error": "Dados inválidos"})
        }
        config.DB.Create(&user)
        return c.Status(201).JSON(user)
    })

    // Atualizar usuário
    app.Put("/users/:id", func(c *fiber.Ctx) error {
        idParam := strings.TrimSpace(c.Params("id"))
        id, err := strconv.Atoi(idParam)
        if err != nil {
            return c.Status(400).JSON(fiber.Map{"error": "ID inválido"})
        }

        var user models.User
        if err := config.DB.First(&user, id).Error; err != nil {
            return c.Status(404).JSON(fiber.Map{"error": "Usuário não encontrado"})
        }

        if err := c.BodyParser(&user); err != nil {
            return c.Status(400).JSON(fiber.Map{"error": "Dados inválidos"})
        }

        config.DB.Save(&user)
        return c.JSON(user)
    })

    // Deletar usuário
    app.Delete("/users/:id", func(c *fiber.Ctx) error {
        idParam := strings.TrimSpace(c.Params("id"))
        id, err := strconv.Atoi(idParam)
        if err != nil {
            return c.Status(400).JSON(fiber.Map{"error": "ID inválido"})
        }

        if err := config.DB.Delete(&models.User{}, id).Error; err != nil {
            return c.Status(404).JSON(fiber.Map{"error": "Usuário não encontrado"})
        }

        return c.SendStatus(204)
    })
}
