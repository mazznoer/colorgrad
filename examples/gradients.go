// +build ignore

package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"

	"github.com/lucasb-eyer/go-colorful"
	"github.com/mazznoer/colorgrad"
)

type data struct {
	gradient colorgrad.Gradient
	name     string
}

func main() {
	presetGradients := []data{
		{colorgrad.CubehelixDefault(), "CubehelixDefault"},
		{colorgrad.Warm(), "Warm"},
		{colorgrad.Cool(), "Cool"},
		{colorgrad.Rainbow(), "Rainbow"},
		{colorgrad.Cividis(), "Cividis"},
		{colorgrad.Sinebow(), "Sinebow"},
		{colorgrad.Turbo(), "Turbo"},
		{colorgrad.Viridis(), "Viridis"},
		{colorgrad.Plasma(), "Plasma"},
		{colorgrad.Magma(), "Magma"},
		{colorgrad.Inferno(), "Inferno"},
		{colorgrad.BrBG(), "BrBG"},
		{colorgrad.PRGn(), "PRGn"},
		{colorgrad.PiYG(), "PiYG"},
		{colorgrad.PuOr(), "PuOr"},
		{colorgrad.RdBu(), "RdBu"},
		{colorgrad.RdGy(), "RdGy"},
		{colorgrad.RdYlBu(), "RdYlBu"},
		{colorgrad.RdYlGn(), "RdYlGn"},
		{colorgrad.Spectral(), "Spectral"},
		{colorgrad.Blues(), "Blues"},
		{colorgrad.Greens(), "Greens"},
		{colorgrad.Greys(), "Greys"},
		{colorgrad.Oranges(), "Oranges"},
		{colorgrad.Purples(), "Purples"},
		{colorgrad.Reds(), "Reds"},
		{colorgrad.BuGn(), "BuGn"},
		{colorgrad.BuPu(), "BuPu"},
		{colorgrad.GnBu(), "GnBu"},
		{colorgrad.OrRd(), "OrRd"},
		{colorgrad.PuBuGn(), "PuBuGn"},
		{colorgrad.PuBu(), "PuBu"},
		{colorgrad.PuRd(), "PuRd"},
		{colorgrad.RdPu(), "RdPu"},
		{colorgrad.YlGnBu(), "YlGnBu"},
		{colorgrad.YlGn(), "YlGn"},
		{colorgrad.YlOrBr(), "YlOrBr"},
		{colorgrad.YlOrRd(), "YlOrRd"},
	}

	// Custom gradients

	grad1, _ := colorgrad.NewGradient().Build()

	grad2, _ := colorgrad.NewGradient().
		Colors(
			color.RGBA{0, 206, 209, 255},
			color.RGBA{255, 105, 180, 255},
			colorful.Color{R: 0.274, G: 0.5, B: 0.7},
			colorful.Hsv(50, 1, 1),
			colorful.Hsv(348, 0.9, 0.8),
		).
		Build()

	grad3, _ := colorgrad.NewGradient().
		HtmlColors("#C41189", "#00BFFF", "#FFD700").
		Build()

	grad4, _ := colorgrad.NewGradient().
		HtmlColors("gold", "hotpink", "darkturquoise").
		Build()

	grad5, _ := colorgrad.NewGradient().
		HtmlColors(
			"rgb(125,110,221)",
			"rgb(90%,45%,97%)",
			"hsl(229,79%,85%)",
		).
		Build()

	// Domain & color position

	domain1, _ := colorgrad.NewGradient().
		HtmlColors("deeppink", "gold", "seagreen").
		Build()

	domain2, _ := colorgrad.NewGradient().
		HtmlColors("deeppink", "gold", "seagreen").
		Domain(0, 100).
		Build()

	domain3, _ := colorgrad.NewGradient().
		HtmlColors("deeppink", "gold", "seagreen").
		Domain(-1, 1).
		Build()

	colorPos1, _ := colorgrad.NewGradient().
		HtmlColors("deeppink", "gold", "seagreen").
		Domain(0, 0.7, 1).
		Build()

	colorPos2, _ := colorgrad.NewGradient().
		HtmlColors("deeppink", "gold", "seagreen").
		Domain(15, 30, 80).
		Build()

	colorPos3, _ := colorgrad.NewGradient().
		HtmlColors("deeppink", "#6d27a1", "#ff0", "#1185e4").
		Domain(0, 0.7, 0.7, 1).
		Build()

	invalidRgbTest, _ := colorgrad.NewGradient().
		HtmlColors("#DC143C", "#FFD700", "#4682B4").
		Mode(colorgrad.BlendHcl).
		Build()

	// Blending modes

	colors := []string{"#fff", "#00f"}

	blendHcl, _ := colorgrad.NewGradient().
		HtmlColors(colors...).
		Mode(colorgrad.BlendHcl).
		Build()

	blendHsv, _ := colorgrad.NewGradient().
		HtmlColors(colors...).
		Mode(colorgrad.BlendHsv).
		Build()

	blendLab, _ := colorgrad.NewGradient().
		HtmlColors(colors...).
		Mode(colorgrad.BlendLab).
		Build()

	blendLinearRgb, _ := colorgrad.NewGradient().
		HtmlColors(colors...).
		Mode(colorgrad.BlendLinearRgb).
		Build()

	blendLuv, _ := colorgrad.NewGradient().
		HtmlColors(colors...).
		Mode(colorgrad.BlendLuv).
		Build()

	blendRgb, _ := colorgrad.NewGradient().
		HtmlColors(colors...).
		Mode(colorgrad.BlendRgb).
		Build()

	blendOklab, _ := colorgrad.NewGradient().
		HtmlColors(colors...).
		Mode(colorgrad.BlendOklab).
		Build()

	customGradients := []data{
		{grad1, "custom-default"},
		{grad2, "custom-colors"},
		{grad3, "custom-hex-colors"},
		{grad4, "custom-named-colors"},
		{grad5, "custom-css-colors"},
		{domain1, "domain-default"},
		{domain2, "domain-0-100"},
		{domain3, "domain-neg1-1"},
		{colorPos1, "color-position-1"},
		{colorPos2, "color-position-2"},
		{colorPos3, "color-position-3"},
		{invalidRgbTest, "invalid-rgb-test"},
		{blendRgb, "blend-rgb"},
		{blendLinearRgb, "blend-linear-rgb"},
		{blendOklab, "blend-oklab"},
		{blendHcl, "blend-hcl"},
		{blendHsv, "blend-hsv"},
		{blendLab, "blend-lab"},
		{blendLuv, "blend-luv"},
	}

	// Sharp gradients

	grad := colorgrad.Rainbow()
	var segments uint = 11

	sharpGradients := []data{
		{grad.Sharp(segments, 0.0), "0.0"},
		{grad.Sharp(segments, 0.1), "0.1"},
		{grad.Sharp(segments, 0.2), "0.2"},
		{grad.Sharp(segments, 0.3), "0.3"},
		{grad.Sharp(segments, 0.4), "0.4"},
		{grad.Sharp(segments, 0.5), "0.5"},
		{grad.Sharp(segments, 0.6), "0.6"},
		{grad.Sharp(segments, 0.7), "0.7"},
		{grad.Sharp(segments, 0.8), "0.8"},
		{grad.Sharp(segments, 0.9), "0.9"},
		{grad.Sharp(segments, 1.0), "1.0"},
	}

	width := 1500
	height := 70

	for _, d := range presetGradients {
		img := gradientImage(d.gradient, width, height)
		filepath := fmt.Sprintf("output/preset-%s.png", d.name)
		fmt.Println(filepath)
		savePNG(img, filepath)
	}

	for _, d := range customGradients {
		img := gradientImage(d.gradient, width, height)
		filepath := fmt.Sprintf("output/%s.png", d.name)
		fmt.Println(filepath)
		savePNG(img, filepath)
	}

	for _, d := range sharpGradients {
		img := gradientImage(d.gradient, width, height)
		filepath := fmt.Sprintf("output/sharp-smoothness-%s.png", d.name)
		fmt.Println(filepath)
		savePNG(img, filepath)
	}
}

func gradientImage(gradient colorgrad.Gradient, width, height int) image.Image {
	fw := float64(width)
	dmin, dmax := gradient.Domain()
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for x := 0; x < width; x++ {
		col := gradient.At(remap(float64(x), 0, fw, dmin, dmax))
		for y := 0; y < height; y++ {
			img.Set(x, y, col)
		}
	}
	return img
}

// Map t which is in range [a, b] to range [c, d]
func remap(t, a, b, c, d float64) float64 {
	return (t-a)*((d-c)/(b-a)) + c
}

func savePNG(img image.Image, filepath string) {
	file, err := os.Create(filepath)
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()
	png.Encode(file, img)
}
