package colorgrad

import (
	"math"
	"testing"
)

func Test_BasisGradient(t *testing.T) {
	grad, err := NewGradient().
		HtmlColors("#f00", "#0f0", "#00f").
		Mode(BlendRgb).
		Interpolation(InterpolationBasis).
		Build()

	test(t, err, nil)
	test(t, grad.At(0.00).HexString(), "#ff0000")
	test(t, grad.At(0.25).HexString(), "#857505")
	test(t, grad.At(0.50).HexString(), "#2baa2b")
	test(t, grad.At(0.75).HexString(), "#057585")
	test(t, grad.At(1.00).HexString(), "#0000ff")

	testSlice(t, colors2hex(grad.Colors(5)), []string{
		"#ff0000",
		"#857505",
		"#2baa2b",
		"#057585",
		"#0000ff",
	})

	test(t, grad.At(-0.1).HexString(), "#ff0000")
	test(t, grad.At(1.11).HexString(), "#0000ff")
	test(t, grad.At(math.NaN()).HexString(), "#000000")
}
