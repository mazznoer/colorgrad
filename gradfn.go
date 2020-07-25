package colorgrad

import (
	"math"

	"github.com/lucasb-eyer/go-colorful"
)

// Algorithms adapted from: https://github.com/d3/d3-scale-chromatic

const pi_1_3 = math.Pi / 3
const pi_2_3 = math.Pi * 2 / 3

// Sinebow

type sinebowGradient struct{}

func Sinebow() Gradient {
	return sinebowGradient{}
}

func (self sinebowGradient) At(t float64) colorful.Color {
	t = (0.5 - t) * math.Pi
	return colorful.Color{
		R: math.Pow(math.Sin(t), 2),
		G: math.Pow(math.Sin(t+pi_1_3), 2),
		B: math.Pow(math.Sin(t+pi_2_3), 2),
	}
}

func (self sinebowGradient) Colors(count uint) []colorful.Color {
	l := float64(count - 1)
	colors := make([]colorful.Color, count)
	for i := range colors {
		colors[i] = self.At(float64(i) / l)
	}
	return colors
}

// Turbo

type turboGradient struct{}

func Turbo() Gradient {
	return turboGradient{}
}

func (self turboGradient) At(t float64) colorful.Color {
	t = math.Max(0, math.Min(1, t))
	r := math.Round(34.61 + t*(1172.33-t*(10793.56-t*(33300.12-t*(38394.49-t*14825.05)))))
	g := math.Round(23.31 + t*(557.33+t*(1225.33-t*(3574.96-t*(1073.77+t*707.56)))))
	b := math.Round(27.2 + t*(3211.1-t*(15327.97-t*(27814-t*(22569.18-t*6838.66)))))
	return colorful.Color{
		R: clamp01(r / 255),
		G: clamp01(g / 255),
		B: clamp01(b / 255),
	}
}

func (self turboGradient) Colors(count uint) []colorful.Color {
	l := float64(count - 1)
	colors := make([]colorful.Color, count)
	for i := range colors {
		colors[i] = self.At(float64(i) / l)
	}
	return colors
}

// Cividis

type cividisGradient struct{}

func Cividis() Gradient {
	return cividisGradient{}
}

func (self cividisGradient) At(t float64) colorful.Color {
	t = math.Max(0, math.Min(1, t))
	r := math.Round(-4.54 - t*(35.34-t*(2381.73-t*(6402.7-t*(7024.72-t*2710.57)))))
	g := math.Round(32.49 + t*(170.73+t*(52.82-t*(131.46-t*(176.58-t*67.37)))))
	b := math.Round(81.24 + t*(442.36-t*(2482.43-t*(6167.24-t*(6614.94-t*2475.67)))))
	return colorful.Color{
		R: clamp01(r / 255),
		G: clamp01(g / 255),
		B: clamp01(b / 255),
	}
}

func (self cividisGradient) Colors(count uint) []colorful.Color {
	l := float64(count - 1)
	colors := make([]colorful.Color, count)
	for i := range colors {
		colors[i] = self.At(float64(i) / l)
	}
	return colors
}

func clamp01(t float64) float64 {
	return math.Max(0, math.Min(1, t))
}
