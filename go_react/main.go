package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/poboisvert/go_projects/go_react/bootstrap"
	"github.com/poboisvert/go_projects/go_react/repository"
)

type Repository repository.Repository

func main() {
	app := fiber.New()
	bootstrap.InitializeApp(app)
}