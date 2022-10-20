//go:build ignore
// +build ignore

package main

import (
	"image"
	"image/png"
	"os"

	"github.com/mazznoer/colorgrad"
)

func main() {
	grad, _ := colorgrad.NewGradient().
		HtmlColors("#c41189", "#00BFFF", "#FFD700").
		Build()

	w := 1500
	h := 70
	fw := float64(w)

	img := image.NewRGBA(image.Rect(0, 0, w, h))

	for x := 0; x < w; x++ {
		col := grad.At(float64(x) / fw)
		for y := 0; y < h; y++ {
			img.Set(x, y, col)
		}
	}

	file, err := os.Create("gradient.png")
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()
	png.Encode(file, img)
}
