package handlers

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go-fiber-api/database"
	"go-fiber-api/models/entity"
	"go-fiber-api/models/request"
	"go-fiber-api/utils"
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

func DeleteBookById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	var book entity.Book
	err := database.DB.Debug().First(&book, "id = ?", id).Error

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"OK":      false,
			"message": "Book not found",
		})
	}

	/** Delete book file from disk */
	errRemove := utils.RemoveFile(book.Cover, "books/cover/")
	if errRemove != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"OK":      false,
			"message": errRemove.Error(),
		})
	} else {
		fmt.Println("Success remove file")
	}

	errDelete := database.DB.Debug().Delete(&book, "id = ?", id).Error
	if errDelete != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"OK":      false,
			"message": errDelete,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"OK":      true,
		"message": "Success delete book",
	})
}
