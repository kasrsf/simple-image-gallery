package utils

import (
	"image"
	"image/color"
	"image/png"
	"io"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goregular"
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
		// Load the font
		f, err := truetype.Parse(goregular.TTF)
		if err != nil {
			return nil, err
		}

		// Create font context
		c := freetype.NewContext()
		c.SetDPI(72)
		c.SetFont(f)
		c.SetClip(img.Bounds())
		c.SetDst(img)
		c.SetSrc(image.NewUniform(color.Black))

		// Calculate font size based on image dimensions
		fontSize := float64(height) / 10
		c.SetFontSize(fontSize)

		// Get font metrics
		opts := truetype.Options{
			Size: fontSize,
			DPI: 72,
			Hinting: font.HintingFull,
		}
		face := truetype.NewFace(f, &opts)
		metrics := face.Metrics()

		// Calculate text bounds
		textWidth := font.MeasureString(face, text).Ceil()
		textHeight := metrics.Height.Ceil()

		// Calculate position to center the text
		x := (width - textWidth) / 2
		y := (height + textHeight) / 2

		// Draw text
		pt := freetype.Pt(x, y)
		_, err = c.DrawString(text, pt)
		if err != nil {
			return nil, err
		}
	}

	return img, nil
}

func SaveImageToPNG(img *image.RGBA, w io.Writer) error {
	return png.Encode(w, img)
}
