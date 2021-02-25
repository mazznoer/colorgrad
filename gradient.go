package colorgrad

import (
	"fmt"
	"image/color"

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
	if segment < 2 {
		return Gradient{
			grad: zeroGradient{color: g.grad.At(g.dmin)},
			dmin: g.dmin,
			dmax: g.dmax,
		}
	}
	if smoothness > 0 {
		return newSharpGradientX(g, segment, smoothness)
	}
	return newSharpGradient(g, segment)
}

type zeroGradient struct {
	color colorful.Color
}

func (zg zeroGradient) At(t float64) colorful.Color {
	return zg.color
}

type linearGradient struct {
	colors []colorful.Color
	pos    []float64
	dmin   float64
	dmax   float64
	count  int
	mode   BlendMode
}

func (lg linearGradient) At(t float64) colorful.Color {
	if t < lg.dmin {
		return lg.colors[0]
	}
	if t > lg.dmax {
		return lg.colors[lg.count]
	}
	for i := 0; i < lg.count; i++ {
		p1 := lg.pos[i]
		p2 := lg.pos[i+1]
		if (p1 <= t) && (t <= p2) {
			t := (t - p1) / (p2 - p1)
			a := lg.colors[i]
			b := lg.colors[i+1]
			switch lg.mode {
			case BlendHcl:
				return a.BlendHcl(b, t).Clamped()
			case BlendHsv:
				return a.BlendHsv(b, t)
			case BlendLab:
				return a.BlendLab(b, t).Clamped()
			case BlendLinearRgb:
				return blendLrgb(a, b, t)
			case BlendLuv:
				return a.BlendLuv(b, t).Clamped()
			case BlendRgb:
				return a.BlendRgb(b, t)
			case BlendOklab:
				return blendOklab(a, b, t).Clamped()
			}
		}
	}
	return lg.colors[0]
}

type sharpGradient struct {
	colors []colorful.Color
	pos    []float64
	n      int
	dmin   float64
	dmax   float64
}

func (sg sharpGradient) At(t float64) colorful.Color {
	if t < sg.dmin {
		return sg.colors[0]
	}
	if t > sg.dmax {
		return sg.colors[sg.n-1]
	}
	for i := 0; i < sg.n; i++ {
		if (sg.pos[i] <= t) && (t <= sg.pos[i+1]) {
			return sg.colors[i]
		}
	}
	return sg.colors[0]
}

func newSharpGradient(grad Gradient, segment uint) Gradient {
	dmin, dmax := grad.Domain()
	gradbase := sharpGradient{
		colors: grad.ColorfulColors(segment),
		pos:    linspace(dmin, dmax, segment+1),
		n:      int(segment),
		dmin:   dmin,
		dmax:   dmax,
	}
	return Gradient{
		grad: gradbase,
		dmin: dmin,
		dmax: dmax,
	}
}

type sharpGradientX struct {
	colors []colorful.Color
	pos    []float64
	last   int
	dmin   float64
	dmax   float64
}

func (sg sharpGradientX) At(t float64) colorful.Color {
	if t < sg.dmin {
		return sg.colors[0]
	}
	if t > sg.dmax {
		return sg.colors[sg.last]
	}
	for i := 0; i < sg.last; i++ {
		p1 := sg.pos[i]
		p2 := sg.pos[i+1]
		if (p1 <= t) && (t <= p2) {
			if i%2 == 0 {
				return sg.colors[i]
			}
			t := (t - p1) / (p2 - p1)
			a := sg.colors[i]
			b := sg.colors[i+1]
			return a.BlendRgb(b, t)
		}
	}
	return sg.colors[0]
}

func newSharpGradientX(grad Gradient, segment uint, smoothness float64) Gradient {
	colors := make([]colorful.Color, segment*2)
	i := 0
	for _, c := range grad.ColorfulColors(segment) {
		colors[i] = c
		i++
		colors[i] = c
		i++
	}
	dmin, dmax := grad.Domain()
	t := clamp01(smoothness) * (dmax - dmin) / float64(segment) / 4
	p := linspace(dmin, dmax, segment+1)
	pos := make([]float64, segment*2)
	i = 0
	j := 0
	for x := 0; x < int(segment); x++ {
		pos[i] = p[j]
		if i > 0 {
			pos[i] += t
		}
		i++
		j++
		pos[i] = p[j]
		if i < len(colors)-1 {
			pos[i] -= t
		}
		i++
	}
	gradbase := sharpGradientX{
		colors: colors,
		pos:    pos,
		last:   int(segment*2 - 1),
		dmin:   dmin,
		dmax:   dmax,
	}
	return Gradient{
		grad: gradbase,
		dmin: dmin,
		dmax: dmax,
	}
}

type GradientBuilder struct {
	colors            []colorful.Color
	pos               []float64
	mode              BlendMode
	invalidHtmlColors []string
}

func NewGradient() *GradientBuilder {
	return &GradientBuilder{
		mode: BlendRgb,
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

	gradbase := linearGradient{
		colors: gb.colors,
		pos:    pos,
		dmin:   pos[0],
		dmax:   pos[len(pos)-1],
		count:  len(gb.colors) - 1,
		mode:   gb.mode,
	}

	return Gradient{
		grad: gradbase,
		dmin: pos[0],
		dmax: pos[len(pos)-1],
	}, nil
}
