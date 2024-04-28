package main

import (
    "github.com/gofiber/fiber/v2"
	"fmt"
)

func main() {
    // Crie uma nova instância do Fiber
    app := fiber.New()

    // Defina uma rota simples
    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Olá, mundo!")
    })

    clientes := ObterClientes()
    fmt.Println(clientes)

    db, err := db.ConnectDB()
    if err != nil {
        fmt.Println("Erro ao conectar ao banco de dados:", err)
        return
    }
    defer db.Close()


    // Inicie o servidor na porta 3000
    app.Listen(":3005")
}

