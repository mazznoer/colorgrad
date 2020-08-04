package colorgrad

import (
	"image/color"
	"math"

	"github.com/lucasb-eyer/go-colorful"
)

type sharpGradient struct {
	colors []colorful.Color
	pos    []float64
	n      int
}

func SharpGradient(grad Gradient, count uint) Gradient {
	return sharpGradient{
		colors: grad.Colors(count),
		pos:    linspace(0, 1, count+1),
		n:      int(count),
	}
}

func (sg sharpGradient) At(t float64) colorful.Color {
	if math.IsNaN(t) || t < 0 {
		return sg.colors[0]
	}

	if t > 1 {
		return sg.colors[sg.n-1]
	}

	for i := 0; i < sg.n; i++ {
		if (sg.pos[i] <= t) && (t <= sg.pos[i+1]) {
			return sg.colors[i]
		}
	}
	return sg.colors[sg.n-1]
}

func (sg sharpGradient) Colors(count uint) []colorful.Color {
	l := float64(count - 1)
	colors := make([]colorful.Color, count)
	for i := range colors {
		colors[i] = sg.At(float64(i) / l)
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

// IntoColors convert []colorful.Color to []color.Color
func IntoColors(colors []colorful.Color) []color.Color {
	res := make([]color.Color, len(colors))
	for i, col := range colors {
		res[i] = col
	}
	return res
}
