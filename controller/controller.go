package controller

import (
	"github.com/alejandrowaiz98/te-eme-backend-beta/database"
	"github.com/gofiber/fiber"
)

type Controller struct {
	db database.FirestoreImplementation
}

type ControllerImplementation interface {
	Register(ctx *fiber.Ctx)
	Login(ctx *fiber.Ctx)
}
