package main

import (
	"fmt"
	"image" //defines interface Image
)

/*
type Image interface {
    ColorModel() color.Model //also interface
    Bounds() Rectangle
    At(x, y int) color.Color //also interface
}
*/
func main() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
}
