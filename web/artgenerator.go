package web

import (
	"log"
)

// generateAsciiArt generates ASCII art for the given text and banner
func generateAsciiArt(text, banner string) (string, error) {
	bannerFile, err := ReadBannerFile("./banners/" + banner + ".txt")
	if err != nil {
		log.Printf("Error reading banner file '%s': %v", banner, err)
		return "", err
	}

	runeAsciiArtMap, err := MapCreator(bannerFile)
	if err != nil {
		log.Printf("Error creating printable ASCII map for '%s': %v\n", text, err)
		return "", err
	}

	artText, err := ArtRetriever(text, runeAsciiArtMap)
	if err != nil {
		// Log the error with additional context
		log.Printf("Error retrieving ASCII art for text '%s': %v\n", text, err)
		return "", err
	}

	return artText, nil
}
