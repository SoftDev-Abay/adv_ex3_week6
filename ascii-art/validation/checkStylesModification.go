package validation

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"io"
	"os"
)

func CheckStyleModification(styleName string) error {
	originalHash := GetStyleHash(styleName)
	pathToFile := "ascii-art/styles/" + styleName + ".txt"

	currentHash, err := HashFile(pathToFile)

	if err != nil {
		return err
	}

	if originalHash != currentHash {
		err := errors.New("File has been modified.")
		return err
	}

	return nil
}

func HashFile(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}

func GetStyleHash(styleName string) string {
	filemap := make(map[string]string)
	filemap["standard"] = "e194f1033442617ab8a78e1ca63a2061f5cc07a3f05ac226ed32eb9dfd22a6bf"
	filemap["thinkertoy"] = "64285e4960d199f4819323c4dc6319ba34f1f0dd9da14d07111345f5d76c3fa3"
	filemap["shadow"] = "26b94d0b134b77e9fd23e0360bfd81740f80fb7f6541d1d8c5d85e73ee550f73"
	return filemap[styleName]
}
