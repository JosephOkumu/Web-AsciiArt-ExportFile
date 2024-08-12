package web

import (
	"errors"
	"os"
)

// ReadBannerFile reads the content of a banner file and returns it as a string.
func ReadBannerFile(filename string) (string, error) {
	// Read file content
	bannerFile, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	if len(bannerFile) == 0 {
		return "", errors.New("banner file is empty")
	}
	return string(bannerFile), nil
}
