package colorgrad

import (
	"math"
	"testing"
)

func TestSplineGradient(t *testing.T) {
	grad, _ := NewGradient().
		HtmlColors("#f00", "#0f0", "#00f").
		Interpolation(InterpolationCatmullRom).
		Build()
	testStr(t, grad.At(0).Hex(), "#ff0000")
	testStr(t, grad.At(0.5).Hex(), "#00ff00")
	testStr(t, grad.At(1).Hex(), "#0000ff")
	testStr(t, grad.At(math.NaN()).Hex(), "#000000")

	grad, _ = NewGradient().
		HtmlColors("#f00", "#0f0", "#00f").
		Mode(BlendLinearRgb).
		Interpolation(InterpolationCatmullRom).
		Build()
	testStr(t, grad.At(0).Hex(), "#ff0000")
	testStr(t, grad.At(0.5).Hex(), "#00ff00")
	testStr(t, grad.At(1).Hex(), "#0000ff")
	testStr(t, grad.At(math.NaN()).Hex(), "#000000")

	grad, _ = NewGradient().
		HtmlColors("#f00", "#0f0", "#00f").
		Mode(BlendLab).
		Interpolation(InterpolationCatmullRom).
		Build()
	testStr(t, grad.At(0).Hex(), "#ff0000")
	testStr(t, grad.At(0.5).Hex(), "#00ff00")
	testStr(t, grad.At(1).Hex(), "#0000ff")
	testStr(t, grad.At(math.NaN()).Hex(), "#000000")

	grad, _ = NewGradient().
		HtmlColors("#f00", "#0f0", "#00f").
		Mode(BlendLuv).
		Interpolation(InterpolationCatmullRom).
		Build()
	testStr(t, grad.At(0).Hex(), "#ff0000")
	testStr(t, grad.At(0.5).Hex(), "#00ff00")
	testStr(t, grad.At(1).Hex(), "#0000ff")
	testStr(t, grad.At(math.NaN()).Hex(), "#000000")

	grad, _ = NewGradient().
		HtmlColors("#f00", "#0f0", "#00f").
		Mode(BlendHcl).
		Interpolation(InterpolationCatmullRom).
		Build()
	testStr(t, grad.At(0).Hex(), "#ff0000")
	testStr(t, grad.At(0.5).Hex(), "#00ff00")
	testStr(t, grad.At(1).Hex(), "#0000ff")
	testStr(t, grad.At(math.NaN()).Hex(), "#000000")

	grad, _ = NewGradient().
		HtmlColors("#f00", "#0f0", "#00f").
		Mode(BlendHsv).
		Interpolation(InterpolationCatmullRom).
		Build()
	testStr(t, grad.At(0).Hex(), "#ff0000")
	testStr(t, grad.At(0.5).Hex(), "#00ff00")
	testStr(t, grad.At(1).Hex(), "#0000ff")
	testStr(t, grad.At(math.NaN()).Hex(), "#000000")

	grad, _ = NewGradient().
		HtmlColors("#f00", "#0f0", "#00f").
		Mode(BlendOklab).
		Interpolation(InterpolationCatmullRom).
		Build()
	testStr(t, grad.At(0).Hex(), "#ff0000")
	testStr(t, grad.At(0.5).Hex(), "#00ff00")
	testStr(t, grad.At(1).Hex(), "#0000ff")
	testStr(t, grad.At(math.NaN()).Hex(), "#000000")

	grad, _ = NewGradient().
		HtmlColors("#f00", "#0f0", "#00f").
		Interpolation(InterpolationBasis).
		Build()
	testStr(t, grad.At(0).Hex(), "#ff0000")
	testStr(t, grad.At(1).Hex(), "#0000ff")
	testStr(t, grad.At(math.NaN()).Hex(), "#000000")
}
