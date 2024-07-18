package colorgrad

import (
	"math"
	"testing"
)

func TestSplineGradient(t *testing.T) {
	grad, _ := NewGradient().
		HtmlColors("#f00", "#0f0", "#00f").
		Mode(BlendRgb).
		Interpolation(InterpolationCatmullRom).
		Build()
	testStr(t, grad.At(0).HexString(), "#ff0000")
	testStr(t, grad.At(0.5).HexString(), "#00ff00")
	testStr(t, grad.At(1).HexString(), "#0000ff")
	testStr(t, grad.At(math.NaN()).HexString(), "#00000000")

	grad, _ = NewGradient().
		HtmlColors("#f00", "#0f0", "#00f").
		Mode(BlendLinearRgb).
		Interpolation(InterpolationCatmullRom).
		Build()
	testStr(t, grad.At(0).HexString(), "#ff0000")
	testStr(t, grad.At(0.5).HexString(), "#00ff00")
	testStr(t, grad.At(1).HexString(), "#0000ff")
	testStr(t, grad.At(math.NaN()).HexString(), "#00000000")

	grad, _ = NewGradient().
		HtmlColors("#f00", "#0f0", "#00f").
		Mode(BlendOklab).
		Interpolation(InterpolationCatmullRom).
		Build()
	testStr(t, grad.At(0).HexString(), "#ff0000")
	//testStr(t, grad.At(0.5).HexString(), "#00ff00")
	testStr(t, grad.At(1).HexString(), "#0000ff")
	testStr(t, grad.At(math.NaN()).HexString(), "#00000000")

	grad, _ = NewGradient().
		HtmlColors("#f00", "#0f0", "#00f").
		Interpolation(InterpolationBasis).
		Build()
	testStr(t, grad.At(0).HexString(), "#ff0000")
	testStr(t, grad.At(1).HexString(), "#0000ff")
	testStr(t, grad.At(math.NaN()).HexString(), "#00000000")
}
