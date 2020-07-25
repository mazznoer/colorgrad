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

func (self *GradientBuilder) Colors(colors ...color.Color) *GradientBuilder {
	self.colors = make([]colorful.Color, len(colors))

	for i, v := range colors {
		c, _ := colorful.MakeColor(v)
		self.colors[i] = c
	}
	return self
}

func (self *GradientBuilder) HexColors(hex_colors ...string) *GradientBuilder {
	colors := []colorful.Color{}

	for _, v := range hex_colors {
		c, err := colorful.Hex(v)
		if err != nil {
			continue
		}
		colors = append(colors, c)
	}
	self.colors = colors
	return self
}

func (self *GradientBuilder) Domain(domain ...float64) *GradientBuilder {
	self.pos = domain
	return self
}

func (self *GradientBuilder) Mode(mode BlendMode) *GradientBuilder {
	self.mode = mode
	return self
}

func (self *GradientBuilder) Build() (Gradient, error) {
	if len(self.colors) == 0 {
		// Default colors
		self.colors = []colorful.Color{
			colorful.Hsv(0, 0, 0), // black
			colorful.Hsv(0, 0, 1), // white
		}
	} else if len(self.colors) == 1 {
		self.colors = append(self.colors, self.colors[0])
	}

	if len(self.pos) > 0 && len(self.pos) != len(self.colors) {
		return nil, fmt.Errorf("Domain count not equal colors count")
	}

	if len(self.pos) == 0 {
		w := 1.0 / float64(len(self.colors)-1)
		self.pos = make([]float64, len(self.colors))

		for i := range self.pos {
			self.pos[i] = float64(i) * w
		}
	}

	for i := 0; i < len(self.pos)-1; i++ {
		if self.pos[i] > self.pos[i+1] {
			return nil, fmt.Errorf("Domain is wrong")
		}
	}

	//fmt.Printf("Pos: %v Colors: %v Mode: %v\n", self.pos, self.colors, self.mode)

	return gradientX{
		colors: self.colors,
		pos:    self.pos,
		min:    self.pos[0],
		max:    self.pos[len(self.pos)-1],
		count:  len(self.colors),
		mode:   self.mode,
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

func (self gradientX) At(t float64) colorful.Color {
	if math.IsNaN(t) || t < self.min {
		return self.colors[0]
	}

	if t > self.max {
		return self.colors[self.count-1]
	}

	for i := 0; i < self.count-1; i++ {
		p1 := self.pos[i]
		p2 := self.pos[i+1]

		if (p1 <= t) && (t <= p2) {
			t := (t - p1) / (p2 - p1)
			a := self.colors[i]
			b := self.colors[i+1]

			switch self.mode {
			case HCL:
				return a.BlendHcl(b, t)
			case HSV:
				return a.BlendHsv(b, t)
			case LAB:
				return a.BlendLab(b, t)
			case LRGB:
				return blend_lrgb(a, b, t)
			case LUV:
				return a.BlendLuv(b, t)
			case RGB:
				return a.BlendRgb(b, t)
			}
		}
	}
	return self.colors[self.count-1]
}

func (self gradientX) Colors(count uint) []colorful.Color {
	d := self.max - self.min
	l := float64(count - 1)
	colors := make([]colorful.Color, count)

	for i := range colors {
		colors[i] = self.At(self.min + (float64(i)*d)/l)
	}
	return colors
}

// Algorithm taken from: https://github.com/gka/chroma.js

func blend_lrgb(a, b colorful.Color, t float64) colorful.Color {
	return colorful.Color{
		R: math.Sqrt(math.Pow(a.R, 2)*(1-t) + math.Pow(b.R, 2)*t),
		G: math.Sqrt(math.Pow(a.G, 2)*(1-t) + math.Pow(b.G, 2)*t),
		B: math.Sqrt(math.Pow(a.B, 2)*(1-t) + math.Pow(b.B, 2)*t),
	}
}
