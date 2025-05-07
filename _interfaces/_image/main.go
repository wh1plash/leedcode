package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	w, h int
}

// ColorModel returns the color model (required by image.Image)
func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

// Bounds returns the image size (required by image.Image)
func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.w, img.h)
}

// At returns the color at a given (x, y) (required by image.Image)
func (img Image) At(x, y int) color.Color {
	v := uint8(x ^ y) // example pattern
	return color.RGBA{v, v, 255, 255}
}

func main() {
	m := Image{w: 256, h: 256}
	pic.ShowImage(m)
}
