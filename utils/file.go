package utils

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"mime/multipart"
	"os"
	"strings"
	"time"
)

const DefaultPathAssetImages string = "./public/images/"

func UploadFile(ctx *fiber.Ctx) error {
	/** Handle upload file */
	file, errFile := ctx.FormFile("file")
	if errFile != nil {
		log.Println("Error File => ", errFile)
	}

	var filename *string
	if file != nil {
		//contentType := file.Header.Get("Content-Type")
		allowedTypes := []string{"image/jpg", "image/jpeg", "image/png"}
		errCheck := checkAllowFileType(file, allowedTypes)
		if errCheck != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"OK":      false,
				"message": errCheck.Error(),
			})
		}

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

func RemoveFile(filename string, path ...string) error {
	var err error
	if len(path) > 0 {
		err = os.Remove(DefaultPathAssetImages + path[0] + filename)
	} else {
		err = os.Remove(DefaultPathAssetImages + filename)
	}

	if err != nil {
		log.Println("Failed remove file => ", err)
		return err
	}

	return nil
}

func checkAllowFileType(file *multipart.FileHeader, contentTypes []string) error {
	if len(contentTypes) > 0 {
		for _, content := range contentTypes {
			contentType := file.Header.Get("Content-Type")

			if contentType == content {
				return nil
			}
		}

		return errors.New("Not allowed file type")
	}

	return nil
}
