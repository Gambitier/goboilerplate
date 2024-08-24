package imageProcessor

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/gambitier/gocomm/storage/localfile"
	"github.com/h2non/bimg"
)

type ImageProcessor struct {
	OutputDir   string // Directory where processed images will be saved
	FileStorage localfile.LocalFileStorageImpl
}

func NewImageProcessor(
	outputDir string,
	fileStorage localfile.LocalFileStorageImpl,
) *ImageProcessor {
	return &ImageProcessor{
		outputDir,
		fileStorage,
	}
}

func (p *ImageProcessor) GeneratePreviewImageFromPath(inputFilePath string) (string, error) {
	// Read input image file using the file storage
	inputImage, err := p.FileStorage.ReadFile(inputFilePath)
	if err != nil {
		return "", err
	}

	// Generate preview image
	previewImage, err := p.GeneratePreviewImage(inputImage)
	if err != nil {
		return "", err
	}

	// Prepare output file path
	outputFileName := fmt.Sprintf("preview_%s", filepath.Base(inputFilePath))
	outputFilePath := filepath.Join(p.OutputDir, outputFileName)

	// Write preview image to file using the file storage
	err = p.FileStorage.WriteFile(outputFilePath, previewImage)
	if err != nil {
		return "", err
	}

	return outputFilePath, nil
}

func (p *ImageProcessor) GeneratePreviewImage(inputImage []byte) ([]byte, error) {
	_, err := bimg.Size(inputImage)
	if err != nil {
		return nil, err
	}

	// Check if the image format is supported by bimg
	imageType := bimg.DetermineImageTypeName(inputImage)
	log.Printf("image format: %v", imageType)
	if imageType == "unknown" {
		return nil, fmt.Errorf("unsupported image format")
	}

	// Resize the image to generate a preview image (e.g., 300x300 pixels)
	resizedImage, err := bimg.Resize(inputImage, bimg.Options{
		Width:  300,
		Height: 300,
	})
	if err != nil {
		return nil, err
	}

	return resizedImage, nil
}
