package integration_test

import (
	"math"
	"testing"

	"github.com/mazznoer/colorgrad"
)

func Test_LinearGradient_Basic(t *testing.T) {
	g, err := colorgrad.NewGradient().
		HtmlColors("#f00", "#0f0", "#00f").
		Mode(colorgrad.BlendRgb).
		Interpolation(colorgrad.InterpolationLinear).
		Build()

	test(t, err, nil)
	dmin, dmax := g.Domain()
	testF(t, dmin, 0.0)
	testF(t, dmax, 1.0)

	test(t, g.At(0.00).HexString(), "#ff0000")
	test(t, g.At(0.25).HexString(), "#808000")
	test(t, g.At(0.50).HexString(), "#00ff00")
	test(t, g.At(0.75).HexString(), "#008080")
	test(t, g.At(1.00).HexString(), "#0000ff")

	colors := []string{
		"#ff0000", "#e91600", "#d32c00", "#bc4300", "#a65900", "#906f00", "#7a8500", "#649b00",
		"#4eb100", "#37c800", "#21de00", "#0bf400", "#00f40b", "#00de21", "#00c837", "#00b14e",
		"#009b64", "#00857a", "#006f90", "#0059a6", "#0043bc", "#002cd3", "#0016e9", "#0000ff",
	}
	testSlice(t, colors2hex(g.Colors(24)), colors)

	test(t, g.At(-0.1).HexString(), "#ff0000")
	test(t, g.At(1.11).HexString(), "#0000ff")
	test(t, g.At(math.Inf(-1)).HexString(), "#ff0000")
	test(t, g.At(math.Inf(1)).HexString(), "#0000ff")
	test(t, g.At(math.NaN()).HexString(), "#000000")
}

func Test_LinearGradient_CustomDomain(t *testing.T) {
	g, err := colorgrad.NewGradient().
		HtmlColors("#00f", "#f00", "#fff", "#ff0").
		Domain(-0.75, 0.6).
		Mode(colorgrad.BlendRgb).
		Interpolation(colorgrad.InterpolationLinear).
		Build()

	test(t, err, nil)
	dmin, dmax := g.Domain()
	testF(t, dmin, -0.75)
	testF(t, dmax, 0.6)

	colors := []string{
		"#0000ff", "#1900e6", "#3100ce", "#4a00b5", "#63009c", "#7b0084", "#94006b", "#ad0052",
		"#c5003a", "#de0021", "#f70008", "#ff1010", "#ff2929", "#ff4242", "#ff5a5a", "#ff7373",
		"#ff8c8c", "#ffa5a5", "#ffbdbd", "#ffd6d6", "#ffefef", "#fffff7", "#ffffde", "#ffffc5",
		"#ffffad", "#ffff94", "#ffff7b", "#ffff63", "#ffff4a", "#ffff31", "#ffff19", "#ffff00",
	}
	testSlice(t, colors2hex(g.Colors(32)), colors)

	test(t, g.At(-0.8).HexString(), "#0000ff")
	test(t, g.At(0.63).HexString(), "#ffff00")
}

func Test_LinearGradient_CustomPosition(t *testing.T) {
	g, err := colorgrad.NewGradient().
		HtmlColors("#2ec149", "#efd037", "#ed6212", "#db1ed5", "#544fdd").
		Domain(20.0, 25.0, 37.0, 59.0, 87.0).
		Mode(colorgrad.BlendRgb).
		Interpolation(colorgrad.InterpolationLinear).
		Build()

	test(t, err, nil)
	dmin, dmax := g.Domain()
	testF(t, dmin, 20.0)
	testF(t, dmax, 87.0)

	colors := []string{
		"#2ec149", "#70c643", "#b3cb3d", "#efcf37", "#efbf31", "#eeaf2c", "#ee9f27", "#ee9021",
		"#ee801c", "#ed7017", "#ed6114", "#eb5c23", "#ea5732", "#e95241", "#e74c51", "#e64760",
		"#e4426f", "#e33c7e", "#e2378d", "#e0329d", "#df2cac", "#dd27bb", "#dc22ca", "#d91fd5",
		"#d022d6", "#c825d6", "#c028d7", "#b72bd7", "#af2ed8", "#a731d8", "#9f34d9", "#9637d9",
		"#8e3ada", "#863dda", "#7d40db", "#7543db", "#6d46dc", "#6549dc", "#5c4cdd", "#544fdd",
	}
	testSlice(t, colors2hex(g.Colors(40)), colors)

	test(t, g.At(19.0).HexString(), "#2ec149")
	test(t, g.At(88.0).HexString(), "#544fdd")
}
