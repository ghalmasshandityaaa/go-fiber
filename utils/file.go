package utils

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"strings"
	"time"
)

func UploadFile(ctx *fiber.Ctx) error {
	/** Handle upload file */
	file, errFile := ctx.FormFile("file")
	if errFile != nil {
		log.Println("Error File => ", errFile)
	}

	var filename *string
	if file != nil {
		file.Filename = strings.ReplaceAll("GO"+time.Now().Format("20060102150405")+file.Filename, " ", "_")
		filename = &file.Filename

		errSaveFile := ctx.SaveFile(file, fmt.Sprintf("./public/images/books/cover/%s", *filename))
		if errSaveFile != nil {
			log.Println("Error Save File => ", errSaveFile)
		}
	} else {
		log.Println("Nothing file to uploading.")
	}

	if filename != nil {
		ctx.Locals("filename", *filename)
	} else {
		ctx.Locals("filename", nil)
	}

	return ctx.Next()
}
