package handlers

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/gambitier/gocomm/imageProcessor"
	queuesubscribers "github.com/gambitier/gocomm/modules/users/queueSubscribers"
	"github.com/gofiber/fiber/v2"
)

const MaxUploadSize = 5 * 1024 * 1024 // 5 MB

func (h *UserHandler) UploadAvatar(c *fiber.Ctx) error {
	file, err := c.FormFile("image")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Failed to read image")
	}

	if file.Size > MaxUploadSize {
		return c.Status(fiber.StatusBadRequest).SendString("File size exceeds limit")
	}

	fileBytes, err := file.Open()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to open image")
	}
	defer fileBytes.Close()

	err = imageProcessor.ValidateMimeFile(fileBytes)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(fmt.Sprintf("%v", err))
	}

	// Prepare file path
	savePath := filepath.Join("avatars", file.Filename)

	// Write file to the file system
	err = h.FileStorage.SaveMultipartFile(fileBytes, savePath)
	if err != nil {
		log.Print(err)
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to save image")
	}

	// Publish a message to the message queue
	err = h.MessageQueue.Publish(queuesubscribers.UploadAvatarChannel, []byte(savePath))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to publish upload event")
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Upload successful",
		"data": map[string]string{
			"filename": file.Filename,
			"filepath": savePath,
		},
	})
}
