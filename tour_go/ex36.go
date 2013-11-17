package main

import "code.google.com/p/go-tour/pic"

func Pic(dx, dy int) [][]uint8 {
    pic := make([][]uint8, dy)
    for i := range(pic) {
      pic[i] =  make([]uint8, dx)
    }
    for x, row := range(pic) {
        for y := range(row) {
          pic[x][y] = uint8(x * y)
        }
    }
    return pic
}

func main() {
    pic.Show(Pic)
}
