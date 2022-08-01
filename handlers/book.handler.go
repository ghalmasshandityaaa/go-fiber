package handlers

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go-fiber-api/database"
	"go-fiber-api/models/entity"
	"go-fiber-api/models/request"
)

func CreateBook(ctx *fiber.Ctx) error {
	book := new(request.CreateBookRequest)

	/** check body */
	if err := ctx.BodyParser(book); err != nil {
		fmt.Println("error => ", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"OK":      false,
			"message": err.Error(),
		})
	}

	/** Validate request body */
	validate := validator.New()
	errValidate := validate.Struct(book)

	if errValidate != nil {
		var errors []*ErrorResponse
		for _, err := range errValidate.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}

		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"OK":      false,
			"message": errors,
		})
	}

	filename := ctx.Locals("filename")
	fmt.Println("filename => ", filename)
	if filename == nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"OK":      false,
			"message": "Image cover is required",
		})
	}

	cover := fmt.Sprintf("%v", filename)

	/** Insert user into database */
	newBook := entity.Book{
		Title:  book.Title,
		Author: book.Author,
		Cover:  cover,
	}

	errInsert := database.DB.Create(&newBook).Error
	if errInsert != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"OK":      false,
			"message": errInsert,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"OK":      true,
		"message": "Success create book",
		"data":    newBook,
	})
}
