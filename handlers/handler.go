package handlers

import (
	"net/http"

	"github.com/bshome19/programming-languages-with-db/models"
	"github.com/bshome19/programming-languages-with-db/repositories"
	"github.com/gofiber/fiber/v2"
)

type Handler struct{}

var singleton *Handler

func HandlerInstance() *Handler {
	if singleton == nil {
		singleton = &Handler{}
	}
	return singleton
}


func (_ *Handler) GetAllLanguages(c *fiber.Ctx) error {

	result, err := repositories.ReposInstance().GetAllLanguages()
	if err != nil {
		return err
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{"data": result})
}

func (_ *Handler) GetLanguagesById(c *fiber.Ctx) error {
	id := c.Params("id")
	
	result, err := repositories.ReposInstance().GetLanguagesById(id)
	if err != nil {
		return err
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{"data": result})
}

func (_ *Handler) CreateNewLanguage(c *fiber.Ctx) error {
	var language models.ProgrammingLanguage

	if err := c.BodyParser(&language); err != nil {
		c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	result, err := repositories.ReposInstance().CreateNewLanguage(language)
	if err != nil {
		return err
	}
	return c.Status(http.StatusCreated).JSON(fiber.Map{"data": result})
}

func (_ *Handler) DeleteLanguageById(c *fiber.Ctx) error {
	id := c.Params("id")

	result, err := repositories.ReposInstance().DeleteLanguageById(id)
	if err != nil {
		return err
	}
	return c.Status(http.StatusNoContent).JSON(fiber.Map{"data": result})

}

func (_ *Handler) UpdateLanguageDataById(c *fiber.Ctx) error {
	id := c.Params("id")
	var updatedlanguage models.ProgrammingLanguage
	updatedlanguage.Id = id

	if err := c.BodyParser(&updatedlanguage); err != nil {
		c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	result, err := repositories.ReposInstance().UpdateLanguageDataById(updatedlanguage)
	if err != nil {
		return err
	}
	return c.Status(http.StatusCreated).JSON(fiber.Map{"data": result})
}