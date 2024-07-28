package colorgrad

import (
	"fmt"

	"github.com/mazznoer/csscolorparser"
)

type GradientBuilder struct {
	colors            []Color
	positions         []float64
	mode              BlendMode
	interpolation     Interpolation
	invalidHtmlColors []string
	clean             bool
}

func NewGradient() *GradientBuilder {
	return &GradientBuilder{
		mode:          BlendRgb,
		interpolation: InterpolationLinear,
		clean:         false,
	}
}

func (gb *GradientBuilder) Colors(colors ...Color) *GradientBuilder {
	for _, col := range colors {
		gb.colors = append(gb.colors, col)
	}
	gb.clean = false
	return gb
}

func (gb *GradientBuilder) HtmlColors(htmlColors ...string) *GradientBuilder {
	for _, s := range htmlColors {
		c, err := csscolorparser.Parse(s)
		if err != nil {
			gb.invalidHtmlColors = append(gb.invalidHtmlColors, s)
			continue
		}
		gb.colors = append(gb.colors, c)
	}
	gb.clean = false
	return gb
}

func (gb *GradientBuilder) Domain(positions ...float64) *GradientBuilder {
	gb.positions = positions
	gb.clean = false
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

func (gb *GradientBuilder) prepareBuild() error {
	if gb.clean {
		return nil
	}

	if gb.invalidHtmlColors != nil {
		return fmt.Errorf("invalid HTML colors: %q", gb.invalidHtmlColors)
	}

	if len(gb.colors) == 0 {
		// Default colors
		gb.colors = []Color{
			{R: 0, G: 0, B: 0, A: 1}, // black
			{R: 1, G: 1, B: 1, A: 1}, // white
		}
	} else if len(gb.colors) == 1 {
		gb.colors = append(gb.colors, gb.colors[0])
	}

	var positions []float64

	if len(gb.positions) == 0 {
		positions = linspace(0, 1, uint(len(gb.colors)))
	} else if len(gb.positions) == len(gb.colors) {
		for i := 0; i < len(gb.positions)-1; i++ {
			if gb.positions[i] > gb.positions[i+1] {
				return fmt.Errorf("invalid domain")
			}
		}
		positions = gb.positions
	} else if len(gb.positions) == 2 {
		if gb.positions[0] >= gb.positions[1] {
			return fmt.Errorf("invalid domain")
		}
		positions = linspace(gb.positions[0], gb.positions[1], uint(len(gb.colors)))
	} else {
		return fmt.Errorf("invalid domain")
	}

	gb.positions = positions
	gb.clean = true
	return nil
}

func (gb *GradientBuilder) Build() (Gradient, error) {
	if err := gb.prepareBuild(); err != nil {
		return Gradient{
			grad: zeroGradient{},
			dmin: 0,
			dmax: 1,
		}, err
	}

	if gb.interpolation == InterpolationLinear {
		return newLinearGradient(gb.colors, gb.positions, gb.mode), nil
	}

	if gb.interpolation == InterpolationBasis {
		return newBasisGradient(gb.colors, gb.positions, gb.mode), nil
	}

	return newCatmullRomGradient(gb.colors, gb.positions, gb.mode), nil
}

// For testing purposes
func (gb *GradientBuilder) GetColors() *[]Color {
	return &gb.colors
}

// For testing purposes
func (gb *GradientBuilder) GetPositions() *[]float64 {
	return &gb.positions
}
