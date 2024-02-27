package controller

import (
	"fmt"
	"log"

	"github.com/alejandrowaiz98/te-eme-backend-beta/models"
	"github.com/gofiber/fiber"
)

func (c *Controller) Register(ctx *fiber.Ctx) error {

	var data map[string]string

	if err := ctx.BodyParser(&data); err != nil {
		return err
	}

	var user models.User

	user.Username = data["Username"]
	user.Hash = data["Password"]
	user.Email = data["Email"]

	err := c.db.Register(user)

	if err != nil {
		return ctx.JSON(fmt.Sprintf("error: %v", err))
	}

	return ctx.JSON("Welcome!")

}

func (c *Controller) Login(ctx *fiber.Ctx) error {

	var data map[string]string

	if err := ctx.BodyParser(&data); err != nil {
		return err
	}

	var user models.User
	user.Hash = data["Hash"]
	user.Username = data["Username"]

	loggedUser, err := c.db.Login(user)

	if err != nil {
		return err
	}

	log.Println(loggedUser)
	return nil

}
