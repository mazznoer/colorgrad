package colorgrad

import (
	"image/color"
	"math"

	"github.com/lucasb-eyer/go-colorful"
)

type classesGrad struct {
	colors []colorful.Color
	pos    []float64
	n      int
}

func Classes(grad Gradient, count uint) Gradient {
	return classesGrad{
		colors: grad.Colors(count),
		pos:    linspace(0, 1, count+1),
		n:      int(count),
	}
}

func (cg classesGrad) At(t float64) colorful.Color {
	if math.IsNaN(t) || t < 0 {
		return cg.colors[0]
	}

	if t > 1 {
		return cg.colors[cg.n-1]
	}

	for i := 0; i < cg.n; i++ {
		if (cg.pos[i] <= t) && (t <= cg.pos[i+1]) {
			return cg.colors[i]
		}
	}
	return cg.colors[cg.n-1]
}

func (cg classesGrad) Colors(count uint) []colorful.Color {
	l := float64(count - 1)
	colors := make([]colorful.Color, count)
	for i := range colors {
		colors[i] = cg.At(float64(i) / l)
	}
	return colors
}

func linspace(min, max float64, n uint) []float64 {
	d := max - min
	l := float64(n - 1)
	res := make([]float64, n)

	for i := range res {
		res[i] = (min + (float64(i)*d)/l)
	}
	return res
}

// Convert []colorful.Color to []color.Color
func IntoColors(colors []colorful.Color) []color.Color {
	res := make([]color.Color, len(colors))

	for i, col := range colors {
		res[i] = col
	}
	return res
}
