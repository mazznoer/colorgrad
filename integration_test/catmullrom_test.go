package integration_test

import (
	"math"
	"testing"

	"github.com/mazznoer/colorgrad"
)

func Test_CatmullRom_Basic(t *testing.T) {
	g, err := colorgrad.NewGradient().
		HtmlColors("#f00", "#0f0", "#00f").
		Mode(colorgrad.BlendRgb).
		Interpolation(colorgrad.InterpolationCatmullRom).
		Build()

	test(t, err, nil)
	dmin, dmax := g.Domain()
	testF(t, dmin, 0.0)
	testF(t, dmax, 1.0)

	test(t, g.At(0.00).HexString(), "#ff0000")
	test(t, g.At(0.25).HexString(), "#609f00")
	test(t, g.At(0.50).HexString(), "#00ff00")
	test(t, g.At(0.75).HexString(), "#009f60")
	test(t, g.At(1.00).HexString(), "#0000ff")

	colors := []string{
		"#ff0000", "#e71800", "#cc3300", "#b04f00", "#926d00", "#758a00", "#59a600", "#3fc000",
		"#28d700", "#15ea00", "#08f700", "#01fe00", "#00fe01", "#00f708", "#00ea15", "#00d728",
		"#00c03f", "#00a659", "#008a75", "#006d92", "#004fb0", "#0033cc", "#0018e7", "#0000ff",
	}
	testSlice(t, colors2hex(g.Colors(24)), colors)

	test(t, g.At(-0.1).HexString(), "#ff0000")
	test(t, g.At(1.11).HexString(), "#0000ff")
	test(t, g.At(math.Inf(-1)).HexString(), "#ff0000")
	test(t, g.At(math.Inf(1)).HexString(), "#0000ff")
	test(t, g.At(math.NaN()).HexString(), "#000000")
}

func Test_CatmullRom_CustomDomain(t *testing.T) {
	g, err := colorgrad.NewGradient().
		HtmlColors("#00f", "#f00", "#fff", "#ff0").
		Domain(-0.75, 0.6).
		Mode(colorgrad.BlendRgb).
		Interpolation(colorgrad.InterpolationCatmullRom).
		Build()

	test(t, err, nil)
	dmin, dmax := g.Domain()
	testF(t, dmin, -0.75)
	testF(t, dmax, 0.6)

	colors := []string{
		"#0000ff", "#1b00e4", "#3900c6", "#5900a6", "#7a0085", "#9a0065", "#b80047", "#d2002d",
		"#e80017", "#f70008", "#fe0001", "#ff0303", "#ff1212", "#ff2a2a", "#ff4a4a", "#ff6d6d",
		"#ff9292", "#ffb5b5", "#ffd5d5", "#ffeded", "#fffcfc", "#fffffe", "#fffff7", "#ffffe8",
		"#ffffd2", "#ffffb8", "#ffff9a", "#ffff7a", "#ffff59", "#ffff39", "#ffff1b", "#ffff00",
	}
	testSlice(t, colors2hex(g.Colors(32)), colors)

	test(t, g.At(-0.8).HexString(), "#0000ff")
	test(t, g.At(0.63).HexString(), "#ffff00")
}

func Test_CatmullRom_CustomPosition(t *testing.T) {
	g, err := colorgrad.NewGradient().
		HtmlColors("#2ec149", "#efd037", "#ed6212", "#db1ed5", "#544fdd").
		Domain(20.0, 25.0, 37.0, 59.0, 87.0).
		Mode(colorgrad.BlendRgb).
		Interpolation(colorgrad.InterpolationCatmullRom).
		Build()

	test(t, err, nil)
	dmin, dmax := g.Domain()
	testF(t, dmin, 20.0)
	testF(t, dmax, 87.0)

	colors := []string{
		"#2ec149", "#7fc743", "#cfce3d", "#efd037", "#efcb31", "#efbe2b", "#eead24", "#ee991d",
		"#ee8417", "#ed7013", "#ed6112", "#ec5b16", "#eb541f", "#ea4d2c", "#e9463d", "#e83f50",
		"#e73865", "#e5327a", "#e42c90", "#e227a4", "#e023b6", "#de20c5", "#dc1ed0", "#da1ed5",
		"#d41fd6", "#ce20d7", "#c822d7", "#c024d8", "#b827d8", "#b02ad9", "#a72dd9", "#9e31da",
		"#9535da", "#8b39db", "#823ddb", "#7841db", "#6f45dc", "#6548dc", "#5d4cdd", "#544fdd",
	}
	testSlice(t, colors2hex(g.Colors(40)), colors)

	test(t, g.At(19.0).HexString(), "#2ec149")
	test(t, g.At(88.0).HexString(), "#544fdd")
}
