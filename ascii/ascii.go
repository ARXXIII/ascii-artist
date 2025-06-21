package ascii

import (
	"fmt"
	"image"
	"image/color"
)

func grayScale(c color.Color) int {
	r, g, b, _ := c.RGBA()
	return int(0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b))
}

func avgPixel(img image.Image, x, y, w, h int) int {
	cnt, sum, max := 0, 0, img.Bounds().Max
	for i := x; i < x+w && i < max.X; i++ {
		for j := y; j < y+h && j < max.Y; j++ {
			sum += grayScale(img.At(i, j))
			cnt++
		}
	}
	return sum / cnt
}

func PrintASCII(img image.Image, asciiWidth int) {
	ramp := "@#+=. "
	max := img.Bounds().Max
	if max.X < asciiWidth {
		asciiWidth = max.X
	}
	scaleY := max.X / asciiWidth
	if scaleY < 1 {
		scaleY = 1
	}
	scaleX := scaleY * 2
	for y := 0; y < max.Y; y += scaleX {
		for x := 0; x < max.X; x += scaleY {
			c := avgPixel(img, x, y, scaleX, scaleY)
			fmt.Print(string(ramp[len(ramp)*c/65536]))
		}
		fmt.Println()
	}
}
