package colorgrad

import (
	"math"

	"github.com/lucasb-eyer/go-colorful"
)

func linspace(min, max float64, n uint) []float64 {
	d := max - min
	l := float64(n - 1)
	res := make([]float64, n)
	for i := range res {
		res[i] = (min + (float64(i)*d)/l)
	}
	return res
}

// Algorithm taken from: https://github.com/gka/chroma.js

func blendLrgb(a, b colorful.Color, t float64) colorful.Color {
	return colorful.Color{
		R: math.Sqrt(math.Pow(a.R, 2)*(1-t) + math.Pow(b.R, 2)*t),
		G: math.Sqrt(math.Pow(a.G, 2)*(1-t) + math.Pow(b.G, 2)*t),
		B: math.Sqrt(math.Pow(a.B, 2)*(1-t) + math.Pow(b.B, 2)*t),
	}
}
