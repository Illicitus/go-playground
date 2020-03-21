package core

import (
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

func ErrorHandler(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func createFormDataFileCopy(fileName string, source *multipart.FileHeader) (string, error) {
	// Create copy file
	filePath := fmt.Sprintf("./static/%s", fileName)
	newFile, err := os.Create(filePath)
	if err != nil {
		return "", err
	}

	// Read source file
	sourceFile, err := source.Open()
	if err != nil {
		return "", err
	}

	// Copy source to new file
	_, err = io.Copy(newFile, sourceFile)
	if err != nil {
		return "", err
	}

	// Close files
	err = newFile.Close()
	if err != nil {
		return "", err
	}

	err = sourceFile.Close()
	if err != nil {
		return "", err
	}

	return filePath, nil
}

// Take all files from form data input and copy selected files by their keys
func ParseValidateAndCopyFile(w http.ResponseWriter, r *http.Request, keys []string) (map[string]string, error) {
	result := make(map[string]string)

	for _, key := range keys {
		_, file, err := r.FormFile(key)
		if JsonErrorHandler500(w, err) {
			return result, err
		}

		filePath, err := createFormDataFileCopy(file.Filename, file)
		if JsonErrorHandler500(w, err) {
			return result, err
		}

		result[key] = filePath
	}

	return result, nil
}
