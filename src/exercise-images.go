package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct {
	width  int
	height int
}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.width, img.height)
}

func (img Image) At(x, y int) color.Color {
	return color.RGBA{R: uint8(float64(x) + float64(y)/float64(img.width)*255), G: 255, B: 255, A: 255}
}

func main() {
	m := Image{255, 255}
	pic.ShowImage(m)
}
