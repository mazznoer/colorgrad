package colorgrad

import (
	"fmt"
	"image/color"

	"github.com/lucasb-eyer/go-colorful"
	"github.com/mazznoer/csscolorparser"
)

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

	if gb.interpolation == InterpolationBasis {
		return newBasisGradient(gb.colors, pos, gb.mode), nil
	}

	return newCatmullRomGradient(gb.colors, pos, gb.mode), nil
}
