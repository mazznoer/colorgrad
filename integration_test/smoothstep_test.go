package integration_test

import (
	"math"
	"testing"

	"github.com/mazznoer/colorgrad"
)

func Test_Smoothstep_Basic(t *testing.T) {
	g, err := colorgrad.NewGradient().
		HtmlColors("#f00", "#0f0", "#00f").
		Mode(colorgrad.BlendRgb).
		Interpolation(colorgrad.InterpolationSmoothstep).
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
		"#ff0000", "#fa0500", "#eb1400", "#d42b00", "#b84700", "#986700", "#778800", "#57a800",
		"#38c700", "#1fe000", "#0cf300", "#01fe00", "#00fe01", "#00f30c", "#00e01f", "#00c738",
		"#00a857", "#008877", "#006798", "#0047b8", "#002bd4", "#0014eb", "#0005fa", "#0000ff",
	}
	testSlice(t, colors2hex(g.Colors(24)), colors)

	test(t, g.At(-0.1).HexString(), "#ff0000")
	test(t, g.At(1.11).HexString(), "#0000ff")
	test(t, g.At(math.Inf(-1)).HexString(), "#ff0000")
	test(t, g.At(math.Inf(1)).HexString(), "#0000ff")
	test(t, g.At(math.NaN()).HexString(), "#000000")
}

func Test_Smoothstep_CustomDomain(t *testing.T) {
	g, err := colorgrad.NewGradient().
		HtmlColors("#00f", "#f00", "#fff", "#ff0").
		Domain(-0.75, 0.6).
		Mode(colorgrad.BlendRgb).
		Interpolation(colorgrad.InterpolationSmoothstep).
		Build()

	test(t, err, nil)
	dmin, dmax := g.Domain()
	testF(t, dmin, -0.75)
	testF(t, dmax, 0.6)

	colors := []string{
		"#0000ff", "#0700f8", "#1900e6", "#3400cb", "#5500aa", "#790086", "#9e0061", "#c1003e",
		"#de0021", "#f3000c", "#fe0001", "#ff0303", "#ff1212", "#ff2a2a", "#ff4a4a", "#ff6d6d",
		"#ff9292", "#ffb5b5", "#ffd5d5", "#ffeded", "#fffcfc", "#fffffe", "#fffff3", "#ffffde",
		"#ffffc1", "#ffff9e", "#ffff79", "#ffff55", "#ffff34", "#ffff19", "#ffff07", "#ffff00",
	}
	testSlice(t, colors2hex(g.Colors(32)), colors)

	test(t, g.At(-0.8).HexString(), "#0000ff")
	test(t, g.At(0.63).HexString(), "#ffff00")
}

func Test_Smoothstep_CustomPosition(t *testing.T) {
	g, err := colorgrad.NewGradient().
		HtmlColors("#2ec149", "#efd037", "#ed6212", "#db1ed5", "#544fdd").
		Domain(20.0, 25.0, 37.0, 59.0, 87.0).
		Mode(colorgrad.BlendRgb).
		Interpolation(colorgrad.InterpolationSmoothstep).
		Build()

	test(t, err, nil)
	dmin, dmax := g.Domain()
	testF(t, dmin, 20.0)
	testF(t, dmax, 87.0)

	colors := []string{
		"#2ec149", "#63c544", "#c2cd3b", "#efd037", "#efc935", "#efb82f", "#eea228", "#ee8b20",
		"#ed7619", "#ed6714", "#ed6212", "#ed6116", "#ec5d20", "#ea582f", "#e95241", "#e74a56",
		"#e5426d", "#e33a83", "#e03399", "#df2cae", "#dd26bf", "#dc21cc", "#db1fd3", "#db1ed5",
		"#d91fd5", "#d421d5", "#cd23d6", "#c426d6", "#b92ad7", "#ae2ed8", "#a233d8", "#9637d9",
		"#893cda", "#7d40db", "#7244db", "#6848dc", "#604bdc", "#5a4ddd", "#554edd", "#544fdd",
	}
	testSlice(t, colors2hex(g.Colors(40)), colors)

	test(t, g.At(19.0).HexString(), "#2ec149")
	test(t, g.At(88.0).HexString(), "#544fdd")
}
