package main

import "github.com/gofiber/fiber"

func setupRoutes() {
	app.Get(GetLeads)
	app.Post(NewLead)
	app.Delete(DeleteLead)
}

func main() {
	app := fiber.New()
}
