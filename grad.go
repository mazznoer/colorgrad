package colorgrad

import (
	"fmt"
	"image/color"
	"math"

	"github.com/lucasb-eyer/go-colorful"
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

type Gradient interface {
	// Get color at certain position
	At(float64) colorful.Color

	// Get n colors evenly spaced across gradient
	Colors(uint) []colorful.Color
}

type GradientBuilder struct {
	colors []colorful.Color
	pos    []float64
	mode   BlendMode
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

func (gb *GradientBuilder) HexColors(hexColors ...string) *GradientBuilder {
	colors := []colorful.Color{}

	for _, v := range hexColors {
		c, err := colorful.Hex(v)
		if err != nil {
			continue
		}
		colors = append(colors, c)
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
	if len(gb.colors) == 0 {
		// Default colors
		gb.colors = []colorful.Color{
			colorful.Hsv(0, 0, 0), // black
			colorful.Hsv(0, 0, 1), // white
		}
	} else if len(gb.colors) == 1 {
		gb.colors = append(gb.colors, gb.colors[0])
	}

	if len(gb.pos) > 0 && len(gb.pos) != len(gb.colors) {
		return nil, fmt.Errorf("Domain count not equal colors count")
	}

	if len(gb.pos) == 0 {
		w := 1.0 / float64(len(gb.colors)-1)
		gb.pos = make([]float64, len(gb.colors))

		for i := range self.pos {
			gb.pos[i] = float64(i) * w
		}
	}

	for i := 0; i < len(gb.pos)-1; i++ {
		if gb.pos[i] > gb.pos[i+1] {
			return nil, fmt.Errorf("Domain is wrong")
		}
	}

	//fmt.Printf("Pos: %v Colors: %v Mode: %v\n", gb.pos, gb.colors, gb.mode)

	return gradientX{
		colors: gb.colors,
		pos:    gb.pos,
		min:    gb.pos[0],
		max:    gb.pos[len(gb.pos)-1],
		count:  len(gb.colors),
		mode:   gb.mode,
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

	if t > gx.max {
		return gx.colors[gx.count-1]
	}

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

func (gx gradientX) Colors(count uint) []colorful.Color {
	d := gx.max - gx.min
	l := float64(count - 1)
	colors := make([]colorful.Color, count)

	for i := range colors {
		colors[i] = gx.At(gx.min + (float64(i)*d)/l)
	}
	return colors
}

// Algorithm taken from: https://github.com/gka/chroma.js

func blendLrgb(a, b colorful.Color, t float64) colorful.Color {
	return colorful.Color{
		R: math.Sqrt(math.Pow(a.R, 2)*(1-t) + math.Pow(b.R, 2)*t),
		G: math.Sqrt(math.Pow(a.G, 2)*(1-t) + math.Pow(b.G, 2)*t),
		B: math.Sqrt(math.Pow(a.B, 2)*(1-t) + math.Pow(b.B, 2)*t),
	}
}
