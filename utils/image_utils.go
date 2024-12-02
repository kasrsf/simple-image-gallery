package utils

import (
	"image"
	"image/color"
	"image/png"
	"io"
)

func CreateImage(width, height int, text string) (*image.RGBA, error) {
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	// Fill background with white
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			img.Set(x, y, color.White)
		}
	}

	// Add text
	if text != "" {
		// TODO
	}

	return img, nil
}

func SaveImageToPNG(img *image.RGBA, w io.Writer) error {
	return png.Encode(w, img)
}
