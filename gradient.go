package colorgrad

import (
	"math"

	"github.com/mazznoer/csscolorparser"
)

type BlendMode int

const (
	BlendRgb BlendMode = iota
	BlendLinearRgb
	BlendOklab
)

func (b BlendMode) String() string {
	switch b {
	case BlendRgb:
		return "BlendRgb"
	case BlendLinearRgb:
		return "BlendLinearRgb"
	case BlendOklab:
		return "BlendOklab"
	}
	return ""
}

type Interpolation int

const (
	InterpolationLinear Interpolation = iota
	InterpolationCatmullRom
	InterpolationBasis
)

func (i Interpolation) String() string {
	switch i {
	case InterpolationLinear:
		return "InterpolationLinear"
	case InterpolationCatmullRom:
		return "InterpolationCatmullRom"
	case InterpolationBasis:
		return "InterpolationBasis"
	}
	return ""
}

type Color = csscolorparser.Color

var Hwb = csscolorparser.FromHwb
var Hsv = csscolorparser.FromHsv
var Hsl = csscolorparser.FromHsl
var LinearRgb = csscolorparser.FromLinearRGB
var Oklab = csscolorparser.FromOklab
var Oklch = csscolorparser.FromOklch

func Rgb(r, g, b, a float64) Color {
	return Color{R: r, G: g, B: b, A: a}
}

func Rgb8(r, g, b, a uint8) Color {
	return Color{R: float64(r) / 255, G: float64(g) / 255, B: float64(b) / 255, A: float64(a) / 255}
}

type gradientBase interface {
	// Get color at certain position
	At(float64) Color
}

type Gradient struct {
	grad gradientBase
	dmin float64
	dmax float64
}

// Get color at certain position
func (g Gradient) At(t float64) Color {
	return g.grad.At(t)
}

// Get color at certain position
func (g Gradient) RepeatAt(t float64) Color {
	t = norm(t, g.dmin, g.dmax)
	return g.grad.At(g.dmin + modulo(t, 1)*(g.dmax-g.dmin))
}

// Get color at certain position
func (g Gradient) ReflectAt(t float64) Color {
	t = norm(t, g.dmin, g.dmax)
	return g.grad.At(g.dmin + math.Abs(modulo(1+t, 2)-1)*(g.dmax-g.dmin))
}

// Get n colors evenly spaced across gradient
func (g Gradient) Colors(count uint) []Color {
	d := g.dmax - g.dmin
	l := float64(count) - 1
	colors := make([]Color, count)
	for i := range colors {
		colors[i] = g.grad.At(g.dmin + (float64(i)*d)/l).Clamp()
	}
	return colors
}

// Get the gradient domain min and max
func (g Gradient) Domain() (float64, float64) {
	return g.dmin, g.dmax
}

// Return a new hard-edge gradient
func (g Gradient) Sharp(segment uint, smoothness float64) Gradient {
	colors := []Color{}
	if segment >= 2 {
		colors = g.Colors(segment)
	} else {
		colors = append(colors, g.At(g.dmin))
		colors = append(colors, g.At(g.dmin))
	}
	return newSharpGradient(colors, g.dmin, g.dmax, smoothness)
}

type zeroGradient struct {
	color Color
}

func (zg zeroGradient) At(t float64) Color {
	return zg.color
}
