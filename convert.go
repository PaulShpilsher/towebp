package main

import (
	"bytes"
	"fmt"
	"image"

	"github.com/chai2010/webp"
	"github.com/disintegration/imaging"
	"github.com/rwcarlsen/goexif/exif"
	giftowebp "github.com/sizeofint/gif-to-webp"
)

// detectImageType detects the image type from the image bytes
func detectImageType(imageData []byte) (string, error) {
	if bytes.HasPrefix(imageData, []byte{0x47, 0x49, 0x46}) {
		return "gif", nil
	} else if bytes.HasPrefix(imageData, []byte{0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A}) {
		return "png", nil
	} else if bytes.HasPrefix(imageData, []byte{0xFF, 0xD8}) {
		return "jpeg", nil
	} else if bytes.HasPrefix(imageData, []byte{0x52, 0x49, 0x46, 0x46}) {
		return "webp", nil
	}
	return "", fmt.Errorf("unsupported image format")
}

func convertToWebp(imageData []byte) ([]byte, error) {

	imageType, err := detectImageType(imageData)
	if err != nil {
		return nil, err // unsupported format
	}

	if imageType == "webp" {
		// no conversion needed
		return imageData, nil
	}

	if imageType == "gif" {

		converter := giftowebp.NewConverter()
		converter.LoopCompatibility = false
		converter.WebPConfig.SetLossless(0)

		converter.WebPAnimEncoderOptions.SetKmin(9)
		converter.WebPAnimEncoderOptions.SetKmax(17)

		webpBin, err := converter.Convert(imageData)
		return webpBin, err
	}

	// png, jpg
	img, _, err := image.Decode(bytes.NewReader(imageData))
	if err != nil {
		return nil, err
	}

	// handle EXIF data
	exifData, err := exif.Decode(bytes.NewReader(imageData))
	if err == nil { // If there's no error, then EXIF data exists
		orientation, _ := exifData.Get(exif.Orientation)
		if orientation != nil {
			img = applyOrientation(img, orientation.String())
		}
	}

	// Resize the image to ensure the max-width or max-height is 800px
	// const maxSize = 800
	// bounds := img.Bounds()
	// if bounds.Dx() > maxSize || bounds.Dy() > maxSize {
	// 	img = imaging.Fit(img, maxSize, maxSize, imaging.Lanczos)
	// }

	var buf bytes.Buffer
	// Create a WEBP encoder
	err = webp.Encode(&buf, img, &webp.Options{Lossless: false, Quality: 85})
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func applyOrientation(img image.Image, orientation string) image.Image {
	switch orientation {
	case "2":
		return imaging.FlipH(img)
	case "3":
		return imaging.Rotate180(img)
	case "4":
		return imaging.FlipV(img)
	case "5":
		return imaging.Transpose(img)
	case "6":
		return imaging.Rotate270(img)
	case "7":
		return imaging.Transverse(img)
	case "8":
		return imaging.Rotate90(img)
	default:
		return img
	}
}
