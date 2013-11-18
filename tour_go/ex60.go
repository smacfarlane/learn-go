package main

import (
    "code.google.com/p/go-tour/pic"
    "image"
    "image/color"
)

type Image struct{W, H int}

func (Image) ColorModel() color.Model {
  return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
    return image.Rect(0,0,i.W,i.H)
}

func (i Image) At(x, y int) color.Color {
    z := uint8(x * y) //Use the same value from ex36
    return color.RGBA{z, z, 255, 255}
}

func main() {
    m := Image{256, 256}
    pic.ShowImage(m)
}
