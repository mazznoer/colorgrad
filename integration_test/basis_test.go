package integration_test

import (
	"math"
	"testing"

	"github.com/mazznoer/colorgrad"
)

func Test_BasisGradient_Basic(t *testing.T) {
	g, err := colorgrad.NewGradient().
		HtmlColors("#f00", "#0f0", "#00f").
		Mode(colorgrad.BlendRgb).
		Interpolation(colorgrad.InterpolationBasis).
		Build()

	test(t, err, nil)
	dmin, dmax := g.Domain()
	testF(t, dmin, 0.0)
	testF(t, dmax, 1.0)

	test(t, g.At(0.00).HexString(), "#ff0000")
	test(t, g.At(0.25).HexString(), "#857505")
	test(t, g.At(0.50).HexString(), "#2baa2b")
	test(t, g.At(0.75).HexString(), "#057585")
	test(t, g.At(1.00).HexString(), "#0000ff")

	colors := []string{
		"#ff0000", "#e91600", "#d32c00", "#bd4101", "#a85502", "#946803", "#807906", "#6d880a",
		"#5c950e", "#4c9f14", "#3da61c", "#30aa25", "#25aa30", "#1ca63d", "#149f4c", "#0e955c",
		"#0a886d", "#067980", "#036894", "#0255a8", "#0141bd", "#002cd3", "#0016e9", "#0000ff",
	}
	testSlice(t, colors2hex(g.Colors(24)), colors)

	test(t, g.At(-0.1).HexString(), "#ff0000")
	test(t, g.At(1.11).HexString(), "#0000ff")
	test(t, g.At(math.Inf(-1)).HexString(), "#ff0000")
	test(t, g.At(math.Inf(1)).HexString(), "#0000ff")
	test(t, g.At(math.NaN()).HexString(), "#000000")
}

func Test_BasisGradient_CustomDomain(t *testing.T) {
	g, err := colorgrad.NewGradient().
		HtmlColors("#00f", "#f00", "#fff", "#ff0").
		Domain(-0.75, 0.6).
		Mode(colorgrad.BlendRgb).
		Interpolation(colorgrad.InterpolationBasis).
		Build()

	test(t, err, nil)
	dmin, dmax := g.Domain()
	testF(t, dmin, -0.75)
	testF(t, dmax, 0.6)

	colors := []string{
		"#0000ff", "#1900e6", "#3100ce", "#4901b7", "#6002a1", "#77058d", "#8c087c", "#a00d6d",
		"#b21461", "#c21c59", "#d02755", "#dc3356", "#e6425b", "#ee5263", "#f4646e", "#f87679",
		"#fb8986", "#fd9b91", "#fead9c", "#ffbda4", "#ffcca9", "#ffd8aa", "#ffe3a6", "#ffeb9e",
		"#fff292", "#fff783", "#fffa72", "#fffd5e", "#fffe48", "#ffff31", "#ffff19", "#ffff00",
	}
	testSlice(t, colors2hex(g.Colors(32)), colors)

	test(t, g.At(-0.8).HexString(), "#0000ff")
	test(t, g.At(0.63).HexString(), "#ffff00")
}

func Test_BasisGradient_CustomPosition(t *testing.T) {
	g, err := colorgrad.NewGradient().
		HtmlColors("#2ec149", "#efd037", "#ed6212", "#db1ed5", "#544fdd").
		Domain(20.0, 25.0, 37.0, 59.0, 87.0).
		Mode(colorgrad.BlendRgb).
		Interpolation(colorgrad.InterpolationBasis).
		Build()

	test(t, err, nil)
	dmin, dmax := g.Domain()
	testF(t, dmin, 20.0)
	testF(t, dmax, 87.0)

	colors := []string{
		"#2ec149", "#6fc543", "#a8c53c", "#d0bb33", "#dbb22f", "#e3a82c", "#e89c29", "#eb9029",
		"#ec822b", "#eb7530", "#ea6839", "#e96140", "#e85b48", "#e75552", "#e64f5c", "#e44a66",
		"#e24571", "#df407c", "#dd3c87", "#d93992", "#d5369d", "#d134a7", "#cc32b0", "#c631b8",
		"#c131bd", "#bc31c2", "#b632c7", "#b033cb", "#a934ce", "#a235d1", "#9b37d3", "#9439d5",
		"#8c3bd7", "#853ed9", "#7d41da", "#7543db", "#6d46db", "#6549dc", "#5c4cdd", "#544fdd",
	}
	testSlice(t, colors2hex(g.Colors(40)), colors)

	test(t, g.At(19.0).HexString(), "#2ec149")
	test(t, g.At(88.0).HexString(), "#544fdd")
}
