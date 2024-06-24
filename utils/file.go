package utils

import (
	"io"
	"mime/multipart"
	"os"
	"path"
	"path/filepath"
	"runtime"

	"github.com/rs/zerolog/log"
)

func GetProjectDirectoryPath() string {
	// See https://stackoverflow.com/questions/31873396/is-it-possible-to-get-the-current-root-of-package-structure-as-a-string-in-golan
	_, f, _, _ := runtime.Caller(0)
	// Return the project root directory path.
	return filepath.Dir(filepath.Dir(f))
}

func CreateDirectoryIfNotExist(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		// If the directory does not exist, create it
		err = os.MkdirAll(dir, 0o755)
		if err != nil {
			log.Printf("Failed to create directory: %v", err)
			return err
		}
	}
	return nil
}

func StoreMultipartFile(file multipart.File, directoryPath, fileName string) (storedFilePath string, err error) {
	// Create the full path for the file
	filePath := directoryPath
	if !filepath.IsAbs(directoryPath) {
		filePath = filepath.Join(GetProjectDirectoryPath(), directoryPath)
	}

	fullPath := path.Join(filePath, fileName)
	tempFile, err := os.OpenFile(fullPath, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0o600)
	if err != nil {
		log.Error().Err(err).Msg("Cannot create temporary file")
		return "", err
	}

	_, err = io.Copy(tempFile, file)
	if err != nil {
		log.Error().Err(err).Msg("Failed to write to temporary file")
		return "", err
	}

	err = tempFile.Close()
	if err != nil {
		log.Error().Err(err).Msg("Failed to close temporary file")
		return "", err
	}

	return tempFile.Name(), nil
}
