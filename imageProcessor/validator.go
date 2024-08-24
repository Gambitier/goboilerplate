package imageProcessor

import (
	"fmt"
	"mime/multipart"

	"github.com/h2non/bimg"
)

func ValidateMimeFile(file multipart.File) error {
	buf := make([]byte, 512) // Read a small portion to determine the file type
	_, err := file.Read(buf)
	if err != nil {
		return fmt.Errorf("failed to read file: %v", err)
	}

	err = ValidateImageBuffer(buf)
	if err != nil {
		return err
	}

	// Reset the file pointer to the beginning for subsequent operations
	_, err = file.Seek(0, 0)
	if err != nil {
		return fmt.Errorf("failed to reset file pointer")
	}

	return nil
}

func ValidateImageBuffer(buf []byte) error {
	imageType := bimg.DetermineImageType(buf)
	if !bimg.IsImageTypeSupportedByVips(imageType).Load {
		return fmt.Errorf("image type is not supported. Only JPEG, PNG or WEBP are supported")
	}
	return nil
}
