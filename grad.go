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
	HCL BlendMode = iota
	HSV
	LAB
	LRGB
	LUV
	RGB
)

type gradientBase interface {
	// Get color at certain position
	At(float64) colorful.Color
}

type Gradient interface {
	// Get color at certain position
	At(float64) colorful.Color

	// Get n colors evenly spaced across gradient
	ColorfulColors(uint) []colorful.Color

	// Get n colors evenly spaced across gradient
	Colors(uint) []color.Color

	// Get the gradient domain min and max
	Domain() (float64, float64)

	// Return a new hard-edge gradient
	Sharp(uint) Gradient
}

type gradient struct {
	grad gradientBase
	min  float64
	max  float64
}

func (g gradient) At(t float64) colorful.Color {
	return g.grad.At(t)
}

func (g gradient) ColorfulColors(count uint) []colorful.Color {
	d := g.max - g.min
	l := float64(count - 1)
	colors := make([]colorful.Color, count)
	for i := range colors {
		colors[i] = g.grad.At(g.min + (float64(i)*d)/l).Clamped()
	}
	return colors
}

func (g gradient) Colors(count uint) []color.Color {
	colors := make([]color.Color, count)
	for i, col := range g.ColorfulColors(count) {
		colors[i] = col
	}
	return colors
}

func (g gradient) Domain() (float64, float64) {
	return g.min, g.max
}

func (g gradient) Sharp(n uint) Gradient {
	gradbase := sharpGradient{
		colors: g.ColorfulColors(n),
		pos:    linspace(g.min, g.max, n+1),
		n:      int(n),
		min:    g.min,
		max:    g.max,
	}
	return gradient{
		grad: gradbase,
		min:  g.min,
		max:  g.max,
	}
}

type zeroGradient struct{}

func (zg zeroGradient) At(t float64) colorful.Color {
	return colorful.Color{R: 0, G: 0, B: 0}
}

type GradientBuilder struct {
	colors            []colorful.Color
	pos               []float64
	mode              BlendMode
	invalidHtmlColors []string
}

func NewGradient() *GradientBuilder {
	return &GradientBuilder{
		mode: RGB,
	}
}

func (gb *GradientBuilder) Colors(colors ...color.Color) *GradientBuilder {
	gb.colors = make([]colorful.Color, len(colors))
	for i, v := range colors {
		c, _ := colorful.MakeColor(v)
		gb.colors[i] = c
	}
	return gb
}

func (gb *GradientBuilder) HtmlColors(htmlColors ...string) *GradientBuilder {
	colors := []colorful.Color{}
	invalidColors := []string{}

	for _, s := range htmlColors {
		c, err := csscolorparser.Parse(s)
		if err != nil {
			invalidColors = append(invalidColors, s)
			continue
		}
		col, _ := colorful.MakeColor(c)
		colors = append(colors, col)
	}

	if len(invalidColors) > 0 {
		gb.invalidHtmlColors = invalidColors
	}

	gb.colors = colors
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

func (gb *GradientBuilder) Build() (Gradient, error) {
	zgrad := gradient{
		grad: zeroGradient{},
		min:  0,
		max:  1,
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
		pos = gb.pos
	} else if len(gb.pos) == 2 {
		pos = linspace(gb.pos[0], gb.pos[1], uint(len(gb.colors)))
	} else {
		return zgrad, fmt.Errorf("Wrong domain.")
	}
	gb.pos = pos

	for i := 0; i < len(gb.pos)-1; i++ {
		if gb.pos[i] > gb.pos[i+1] {
			return zgrad, fmt.Errorf("Domain number %v (%v) is bigger than the next domain (%v)", i+1, gb.pos[i], gb.pos[i+1])
		}
	}

	gradbase := gradientX{
		colors: gb.colors,
		pos:    gb.pos,
		min:    gb.pos[0],
		max:    gb.pos[len(gb.pos)-1],
		count:  len(gb.colors),
		mode:   gb.mode,
	}

	return gradient{
		grad: gradbase,
		min:  gb.pos[0],
		max:  gb.pos[len(gb.pos)-1],
	}, nil
}

type gradientX struct {
	colors []colorful.Color
	pos    []float64
	min    float64
	max    float64
	count  int
	mode   BlendMode
}

func (gx gradientX) At(t float64) colorful.Color {
	if math.IsNaN(t) || t < gx.min {
		return gx.colors[0]
	}

	//if t > gx.max {
	//	return gx.colors[gx.count-1]
	//}

	for i := 0; i < gx.count-1; i++ {
		p1 := gx.pos[i]
		p2 := gx.pos[i+1]

		if (p1 <= t) && (t <= p2) {
			t := (t - p1) / (p2 - p1)
			a := gx.colors[i]
			b := gx.colors[i+1]

			switch gx.mode {
			case HCL:
				return a.BlendHcl(b, t)
			case HSV:
				return a.BlendHsv(b, t)
			case LAB:
				return a.BlendLab(b, t)
			case LRGB:
				return blendLrgb(a, b, t)
			case LUV:
				return a.BlendLuv(b, t)
			case RGB:
				return a.BlendRgb(b, t)
			}
		}
	}
	return gx.colors[gx.count-1]
}

type sharpGradient struct {
	colors []colorful.Color
	pos    []float64
	n      int
	min    float64
	max    float64
}

func (sg sharpGradient) At(t float64) colorful.Color {
	if math.IsNaN(t) || t < sg.min {
		return sg.colors[0]
	}

	//if t > sg.max {
	//	return sg.colors[sg.n-1]
	//}

	for i := 0; i < sg.n; i++ {
		if (sg.pos[i] <= t) && (t <= sg.pos[i+1]) {
			return sg.colors[i]
		}
	}
	return sg.colors[sg.n-1]
}
