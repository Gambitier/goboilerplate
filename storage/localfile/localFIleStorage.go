package localfile

import (
	"fmt"
	"io"
	"mime/multipart"
	"path/filepath"

	"github.com/spf13/afero"
)

// LocalFileStorage wraps an Afero file system to provide simplified file operations.
type LocalFileStorage struct {
	FileSystem afero.Fs // File system to use for file operations
}

// LocalFileStorageImpl defines the interface for local file storage operations.
type LocalFileStorageImpl interface {
	ReadFile(filePath string) ([]byte, error)
	WriteFile(filePath string, data []byte) error
	SaveMultipartFile(file multipart.File, savePath string) error
}

// NewLocalFileStorage creates a new instance of LocalFileStorage.
func NewLocalFileStorage(fs afero.Fs) *LocalFileStorage {
	return &LocalFileStorage{FileSystem: fs}
}

// ReadFile reads a file from the file system.
func (lfs *LocalFileStorage) ReadFile(filePath string) ([]byte, error) {
	return afero.ReadFile(lfs.FileSystem, filePath)
}

// WriteFile writes data to a file on the file system.
func (lfs *LocalFileStorage) WriteFile(filePath string, data []byte) error {
	// Ensure the directory exists
	err := lfs.FileSystem.MkdirAll(filepath.Dir(filePath), 0755)
	if err != nil {
		return fmt.Errorf("failed to create directories: %v", err)
	}

	// Write the file
	return afero.WriteFile(lfs.FileSystem, filePath, data, 0644)
}

// SaveMultipartFile saves a multipart file to the file system.
func (lfs *LocalFileStorage) SaveMultipartFile(file multipart.File, savePath string) error {
	// Ensure the directory exists
	err := lfs.FileSystem.MkdirAll(filepath.Dir(savePath), 0755)
	if err != nil {
		return fmt.Errorf("failed to create directories: %v", err)
	}

	// Create the file on the file system
	outFile, err := lfs.FileSystem.Create(savePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer outFile.Close()

	// Copy the uploaded file to the destination file
	_, err = io.Copy(outFile, file)
	if err != nil {
		return fmt.Errorf("failed to write file: %v", err)
	}

	return nil
}
