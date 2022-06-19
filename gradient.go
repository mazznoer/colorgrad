package colorgrad

import (
	"fmt"
	"image/color"
	"math"

	"github.com/lucasb-eyer/go-colorful"
	"github.com/mazznoer/csscolorparser"
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

type Interpolation = int

const (
	InterpolationLinear Interpolation = iota
	InterpolationCatmullRom
	InterpolationBasis
)

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

type GradientBuilder struct {
	colors            []colorful.Color
	pos               []float64
	mode              BlendMode
	interpolation     Interpolation
	invalidHtmlColors []string
}

func NewGradient() *GradientBuilder {
	return &GradientBuilder{
		mode:          BlendRgb,
		interpolation: InterpolationLinear,
	}
}

func (gb *GradientBuilder) Colors(colors ...color.Color) *GradientBuilder {
	for _, c := range colors {
		col, _ := colorful.MakeColor(c)
		gb.colors = append(gb.colors, col)
	}
	return gb
}

func (gb *GradientBuilder) HtmlColors(htmlColors ...string) *GradientBuilder {
	for _, s := range htmlColors {
		c, err := csscolorparser.Parse(s)
		if err != nil {
			gb.invalidHtmlColors = append(gb.invalidHtmlColors, s)
			continue
		}
		gb.colors = append(gb.colors, colorful.Color{R: c.R, G: c.G, B: c.B})
	}
	return gb
}

func (gb *GradientBuilder) Domain(domain ...float64) *GradientBuilder {
	gb.pos = domain
	return gb
}

func (gb *GradientBuilder) Mode(mode BlendMode) *GradientBuilder {
	gb.mode = mode
	return gb
}

func (gb *GradientBuilder) Interpolation(mode Interpolation) *GradientBuilder {
	gb.interpolation = mode
	return gb
}

func (gb *GradientBuilder) Build() (Gradient, error) {
	zgrad := Gradient{
		grad: zeroGradient{},
		dmin: 0,
		dmax: 1,
	}

	if gb.invalidHtmlColors != nil {
		return zgrad, fmt.Errorf("Invalid HTML colors: %q", gb.invalidHtmlColors)
	}

	if len(gb.colors) == 0 {
		// Default colors
		gb.colors = []colorful.Color{
			{R: 0, G: 0, B: 0}, // black
			{R: 1, G: 1, B: 1}, // white
		}
	} else if len(gb.colors) == 1 {
		gb.colors = append(gb.colors, gb.colors[0])
	}

	var pos []float64

	if len(gb.pos) == 0 {
		pos = linspace(0, 1, uint(len(gb.colors)))
	} else if len(gb.pos) == len(gb.colors) {
		for i := 0; i < len(gb.pos)-1; i++ {
			if gb.pos[i] > gb.pos[i+1] {
				return zgrad, fmt.Errorf("Domain number %v (%v) is bigger than the next domain (%v)", i+1, gb.pos[i], gb.pos[i+1])
			}
		}
		pos = gb.pos
	} else if len(gb.pos) == 2 {
		if gb.pos[0] >= gb.pos[1] {
			return zgrad, fmt.Errorf("Wrong domain.")
		}
		pos = linspace(gb.pos[0], gb.pos[1], uint(len(gb.colors)))
	} else {
		return zgrad, fmt.Errorf("Wrong domain.")
	}

	if gb.interpolation == InterpolationLinear {
		return newLinearGradient(gb.colors, pos, gb.mode), nil
	}
	return newSplineGradient(gb.colors, pos, gb.mode, gb.interpolation), nil
}
