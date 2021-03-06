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
	imgFunc := func(x, y int) uint8 {
		//or return uint8(x*y)
		//or return uint8((x+y) / 2)
		return uint8(x ^ y)
	}
	v := imgFunc(x, y)
	return color.RGBA{v, v, 255, 255}
}

func main() {
	m := Image{256, 64}
	pic.ShowImage(m)
}
