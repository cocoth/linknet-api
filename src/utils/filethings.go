package utils

import (
	"bufio"
	"fmt"
	"os"

	"io"
	"mime/multipart"
	"path/filepath"
	"strings"
)

func ReadFileLines(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func SaveMultipartFile(fileHandler *multipart.FileHeader) (string, error) {
	// Get the upload directory from environment variable
	uploadDir := os.Getenv("APP_UPLOAD_DIR")
	if uploadDir == "" {
		return "", fmt.Errorf("APP_UPLOAD_DIR environment variable is not set")
	}

	// Open the uploaded file
	src, err := fileHandler.Open()
	if err != nil {
		return "", fmt.Errorf("failed to open file: %v", err)
	}
	defer src.Close()

	// Determine the file type (extension)
	fileExt := strings.ToLower(filepath.Ext(fileHandler.Filename))
	if len(fileExt) > 0 {
		fileExt = fileExt[1:] // Remove the leading dot
	}

	// Create subdirectory based on file type
	subDir := filepath.Join(uploadDir, fileExt)
	if err := os.MkdirAll(subDir, os.ModePerm); err != nil {
		return "", fmt.Errorf("failed to create subdirectory: %v", err)
	}

	// Create the destination file
	destPath := filepath.Join(subDir, fileHandler.Filename)
	dest, err := os.Create(destPath)
	if err != nil {
		return "", fmt.Errorf("failed to create destination file: %v", err)
	}
	defer dest.Close()

	// Copy the file content
	if _, err := io.Copy(dest, src); err != nil {
		return "", fmt.Errorf("failed to save file: %v", err)
	}

	return destPath, nil
}
