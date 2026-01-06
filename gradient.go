package colorgrad

import (
	"image/color"
	"math"

	"github.com/mazznoer/csscolorparser"
)

type BlendMode int

const (
	BlendRgb BlendMode = iota
	BlendLinearRgb
	BlendLab
	BlendOklab
)

func (b BlendMode) String() string {
	switch b {
	case BlendRgb:
		return "BlendRgb"
	case BlendLinearRgb:
		return "BlendLinearRgb"
	case BlendLab:
		return "BlendLab"
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
var Lab = csscolorparser.FromLab
var Lch = csscolorparser.FromLch
var Oklab = csscolorparser.FromOklab
var Oklch = csscolorparser.FromOklch

func Rgb(r, g, b, a float64) Color {
	return Color{R: r, G: g, B: b, A: a}
}

func Rgb8(r, g, b, a uint8) Color {
	return Color{R: float64(r) / 255, G: float64(g) / 255, B: float64(b) / 255, A: float64(a) / 255}
}

func GoColor(col color.Color) Color {
	r, g, b, a := col.RGBA()
	if a == 0 {
		return csscolorparser.Color{}
	}
	r *= 0xffff
	r /= a
	g *= 0xffff
	g /= a
	b *= 0xffff
	b /= a
	return csscolorparser.Color{R: float64(r) / 65535.0, G: float64(g) / 65535.0, B: float64(b) / 65535.0, A: float64(a) / 65535.0}
}

type GradientCore interface {
	// Get color at certain position
	At(float64) Color
}

type Gradient struct {
	Core GradientCore
	Min  float64
	Max  float64
}

// Get color at certain position
func (g Gradient) At(t float64) Color {
	return g.Core.At(t)
}

// Get color at certain position
func (g Gradient) RepeatAt(t float64) Color {
	t = norm(t, g.Min, g.Max)
	return g.Core.At(g.Min + modulo(t, 1)*(g.Max-g.Min))
}

// Get color at certain position
func (g Gradient) ReflectAt(t float64) Color {
	t = norm(t, g.Min, g.Max)
	return g.Core.At(g.Min + math.Abs(modulo(1+t, 2)-1)*(g.Max-g.Min))
}

// Get n colors evenly spaced across gradient
func (g Gradient) Colors(count uint) []Color {
	d := g.Max - g.Min
	l := float64(count) - 1
	colors := make([]Color, count)
	for i := range colors {
		colors[i] = g.Core.At(g.Min + (float64(i)*d)/l).Clamp()
	}
	return colors
}

// Get the gradient domain min and max
func (g Gradient) Domain() (float64, float64) {
	return g.Min, g.Max
}

// Return a new hard-edge gradient
func (g Gradient) Sharp(segment uint, smoothness float64) Gradient {
	colors := []Color{}
	if segment >= 2 {
		colors = g.Colors(segment)
	} else {
		colors = append(colors, g.At(g.Min))
		colors = append(colors, g.At(g.Min))
	}
	return newSharpGradient(colors, g.Min, g.Max, smoothness)
}

type zeroGradient struct {
}

func (zg zeroGradient) At(t float64) Color {
	return Color{R: 0, G: 0, B: 0, A: 0}
}
