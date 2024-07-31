package colorgrad

import (
	"math"
	"testing"
)

func Test_CatmullRomGradient(t *testing.T) {
	grad, err := NewGradient().
		HtmlColors("#f00", "#0f0", "#00f").
		Mode(BlendRgb).
		Interpolation(InterpolationCatmullRom).
		Build()

	test(t, err, nil)
	test(t, grad.At(0.00).HexString(), "#ff0000")
	test(t, grad.At(0.25).HexString(), "#609f00")
	test(t, grad.At(0.50).HexString(), "#00ff00")
	test(t, grad.At(0.75).HexString(), "#009f60")
	test(t, grad.At(1.00).HexString(), "#0000ff")

	testSlice(t, colors2hex(grad.Colors(5)), []string{
		"#ff0000",
		"#609f00",
		"#00ff00",
		"#009f60",
		"#0000ff",
	})

	test(t, grad.At(-0.1).HexString(), "#ff0000")
	test(t, grad.At(1.11).HexString(), "#0000ff")
	test(t, grad.At(math.NaN()).HexString(), "#000000")
}
