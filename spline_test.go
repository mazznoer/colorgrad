package colorgrad

import (
	"math"
	"testing"
)

func TestSplineGradient(t *testing.T) {
	grad, err := NewGradient().
		HtmlColors("#f00", "#0f0", "#00f").
		Mode(BlendRgb).
		Interpolation(InterpolationCatmullRom).
		Build()
	test(t, err, nil)
	test(t, grad.At(0).HexString(), "#ff0000")
	test(t, grad.At(0.5).HexString(), "#00ff00")
	test(t, grad.At(1).HexString(), "#0000ff")
	test(t, grad.At(math.NaN()).HexString(), "#000000")

	grad, err = NewGradient().
		HtmlColors("#f00", "#0f0", "#00f").
		Mode(BlendLinearRgb).
		Interpolation(InterpolationCatmullRom).
		Build()
	test(t, err, nil)
	test(t, grad.At(0).HexString(), "#ff0000")
	test(t, grad.At(0.5).HexString(), "#00ff00")
	test(t, grad.At(1).HexString(), "#0000ff")
	test(t, grad.At(math.NaN()).HexString(), "#000000")

	grad, err = NewGradient().
		HtmlColors("#f00", "#0f0", "#00f").
		Mode(BlendOklab).
		Interpolation(InterpolationCatmullRom).
		Build()
	test(t, err, nil)
	test(t, grad.At(0).HexString(), "#ff0000")
	//test(t, grad.At(0.5).HexString(), "#00ff00")
	test(t, grad.At(1).HexString(), "#0000ff")
	test(t, grad.At(math.NaN()).HexString(), "#000000")

	grad, err = NewGradient().
		HtmlColors("#f00", "#0f0", "#00f").
		Interpolation(InterpolationBasis).
		Build()
	test(t, err, nil)
	test(t, grad.At(0).HexString(), "#ff0000")
	test(t, grad.At(1).HexString(), "#0000ff")
	test(t, grad.At(math.NaN()).HexString(), "#000000")
}
