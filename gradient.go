package colorgrad

import (
	"image/color"
	"math"

	"github.com/lucasb-eyer/go-colorful"
)

type BlendMode int

const (
	BlendHcl BlendMode = iota
	BlendHsv
	BlendLab
	BlendLinearRgb
	BlendLuv
	BlendRgb
	BlendOklab
)

func (b BlendMode) String() string {
	switch b {
	case BlendHcl:
		return "BlendHcl"
	case BlendHsv:
		return "BlendHsv"
	case BlendLab:
		return "BlendLab"
	case BlendLinearRgb:
		return "BlendLinearRgb"
	case BlendLuv:
		return "BlendLuv"
	case BlendRgb:
		return "BlendRgb"
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

type gradientBase interface {
	// Get color at certain position
	At(float64) colorful.Color
}

type Gradient struct {
	grad gradientBase
	dmin float64
	dmax float64
}

// Get color at certain position
func (g Gradient) At(t float64) colorful.Color {
	return g.grad.At(t)
}

// Get color at certain position
func (g Gradient) RepeatAt(t float64) colorful.Color {
	t = norm(t, g.dmin, g.dmax)
	return g.grad.At(g.dmin + modulo(t, 1)*(g.dmax-g.dmin))
}

// Get color at certain position
func (g Gradient) ReflectAt(t float64) colorful.Color {
	t = norm(t, g.dmin, g.dmax)
	return g.grad.At(g.dmin + math.Abs(modulo(1+t, 2)-1)*(g.dmax-g.dmin))
}

// Get n colors evenly spaced across gradient
func (g Gradient) ColorfulColors(count uint) []colorful.Color {
	d := g.dmax - g.dmin
	l := float64(count) - 1
	colors := make([]colorful.Color, count)
	for i := range colors {
		colors[i] = g.grad.At(g.dmin + (float64(i)*d)/l).Clamped()
	}
	return colors
}

// Get n colors evenly spaced across gradient
func (g Gradient) Colors(count uint) []color.Color {
	colors := make([]color.Color, count)
	for i, col := range g.ColorfulColors(count) {
		colors[i] = col
	}
	return colors
}

// Get the gradient domain min and max
func (g Gradient) Domain() (float64, float64) {
	return g.dmin, g.dmax
}

// Return a new hard-edge gradient
func (g Gradient) Sharp(segment uint, smoothness float64) Gradient {
	colors := []colorful.Color{}
	if segment >= 2 {
		colors = g.ColorfulColors(segment)
	} else {
		colors = append(colors, g.At(g.dmin))
		colors = append(colors, g.At(g.dmin))
	}
	return newSharpGradient(colors, g.dmin, g.dmax, smoothness)
}

type zeroGradient struct {
	color colorful.Color
}

func (zg zeroGradient) At(t float64) colorful.Color {
	return zg.color
}
