package integration_test

import (
	"testing"

	"github.com/mazznoer/colorgrad"
)

func TestLinearize(t *testing.T) {
	// Simple (2 colors)
	{
		g0, err := colorgrad.NewGradient().
			HtmlColors("#00f", "#f00").
			Mode(colorgrad.BlendRgb).
			Interpolation(colorgrad.InterpolationLinear).
			Build()

		testTrue(t, err == nil)

		g1 := g0.Linearize(0.005)
		lg := g1.Core.(colorgrad.LinearGradient)

		test(t, lg.Mode(), colorgrad.BlendRgb)
		test(t, len(lg.Colors()), 2)
		test(t, len(lg.Positions()), 2)
	}

	// Simple (4 colors)
	{
		g0, err := colorgrad.NewGradient().
			HtmlColors("#00f", "#f00", "#fff", "#ff0").
			Mode(colorgrad.BlendRgb).
			Interpolation(colorgrad.InterpolationLinear).
			Build()

		testTrue(t, err == nil)

		g1 := g0.Linearize(0.005)
		lg := g1.Core.(colorgrad.LinearGradient)

		test(t, len(lg.Colors()), 4)
		test(t, len(lg.Positions()), 4)
	}

	// Complex
	{
		g0, err := colorgrad.NewGradient().
			HtmlColors("#00f", "#f00", "#fff", "#ff0").
			Mode(colorgrad.BlendOklab).
			Interpolation(colorgrad.InterpolationBasis).
			Build()

		testTrue(t, err == nil)

		g1 := g0.Linearize(0.005)
		lg := g1.Core.(colorgrad.LinearGradient)

		test(t, len(lg.Colors()), 21)
		test(t, len(lg.Positions()), 21)

		expected := []string{
			"#0000ff", "#242cf0", "#3c3de0", "#5146d2", "#654fc3", "#7752b7",
			"#8955aa", "#9a589e", "#ab5b92", "#b95e8a", "#c86082", "#d5647c",
			"#e06a7a", "#ea7179", "#f27b7c", "#f98581", "#fd9088", "#ff9c90",
			"#ffa898", "#ffb3a1", "#ffbfaa", "#ffcbb3", "#ffd5b9", "#ffdebe",
			"#ffe6c0", "#ffedc0", "#fff2bb", "#fff6b5", "#fffaaf", "#fffca3",
			"#fffd96", "#feff89", "#feff76", "#feff62", "#ffff45", "#ffff00",
		}
		testSlice(t, expected, colors2hex(g1.Colors(36)))
	}
}
