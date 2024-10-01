package main

import (
	"log"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		log.Println("Usage: towebp <input_image_file> <output_webp_file>")
		return
	}

	srcPath := os.Args[1]
	webpPath := os.Args[2]

	imageData, err := os.ReadFile(srcPath)
	if err != nil {
		log.Fatal(err)
	}

	webpData, err := convertToWebp(imageData)
	if err != nil {
		log.Fatal(err)
	}

	os.WriteFile(webpPath, webpData, 0666)

	log.Printf("Successfully created %s file", webpPath)
}
