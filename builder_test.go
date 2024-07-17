//go:build ignore
package colorgrad

import (
	"image/color"
	"math"
	"testing"
)

func TestBasic1(t *testing.T) {
	// Single color
	grad, _ := NewGradient().
		Colors(color.RGBA{0, 255, 0, 255}).
		Build()
	dmin, dmax := grad.Domain()
	if dmin != 0 || dmax != 1 {
		t.Errorf("Domain got: (%v, %v), expected: (0, 1)", dmin, dmax)
	}
	testStr(t, grad.At(0).Hex(), "#00ff00")
	testStr(t, grad.At(1).Hex(), "#00ff00")

	// Named colors
	grad, _ = NewGradient().
		HtmlColors("tomato", "skyblue", "gold", "springgreen").
		Build()
	colors := grad.ColorfulColors(4)
	testStr(t, colors[0].Hex(), "#ff6347")
	testStr(t, colors[1].Hex(), "#87ceeb")
	testStr(t, colors[2].Hex(), "#ffd700")
	testStr(t, colors[3].Hex(), "#00ff7f")

	// Blend mode
	grad, _ = NewGradient().
		HtmlColors("#333", "#bbb").
		Mode(BlendHcl).
		Build()
	testStr(t, grad.At(0).Hex(), "#333333")
	testStr(t, grad.At(1).Hex(), "#bbbbbb")

	grad, _ = NewGradient().
		HtmlColors("#333", "#bbb").
		Mode(BlendHsv).
		Build()
	testStr(t, grad.At(0).Hex(), "#333333")
	testStr(t, grad.At(1).Hex(), "#bbbbbb")

	grad, _ = NewGradient().
		HtmlColors("#333", "#bbb").
		Mode(BlendLab).
		Build()
	testStr(t, grad.At(0).Hex(), "#333333")
	testStr(t, grad.At(1).Hex(), "#bbbbbb")

	grad, _ = NewGradient().
		HtmlColors("#333", "#bbb").
		Mode(BlendLinearRgb).
		Build()
	testStr(t, grad.At(0).Hex(), "#333333")
	testStr(t, grad.At(1).Hex(), "#bbbbbb")

	grad, _ = NewGradient().
		HtmlColors("#333", "#bbb").
		Mode(BlendLuv).
		Build()
	testStr(t, grad.At(0).Hex(), "#333333")
	testStr(t, grad.At(1).Hex(), "#bbbbbb")

	grad, _ = NewGradient().
		HtmlColors("#333", "#bbb").
		Mode(BlendRgb).
		Build()
	testStr(t, grad.At(0).Hex(), "#333333")
	testStr(t, grad.At(1).Hex(), "#bbbbbb")

	grad, _ = NewGradient().
		HtmlColors("#333", "#bbb").
		Mode(BlendOklab).
		Build()
	testStr(t, grad.At(0).Hex(), "#333333")
	testStr(t, grad.At(1).Hex(), "#bbbbbb")
}

func TestBasic2(t *testing.T) {
	// Custom gradient default
	grad, _ := NewGradient().Build()
	colors := grad.ColorfulColors(2)

	if len(colors) != 2 {
		t.Errorf("Expected 2, got %v", len(colors))
	}
	testStr(t, colors[0].Hex(), "#000000")
	testStr(t, colors[1].Hex(), "#ffffff")

	testStr(t, grad.At(math.NaN()).Hex(), "#000000")

	// Custom colors
	grad, _ = NewGradient().
		Colors(
			color.RGBA{255, 0, 0, 255},
			color.RGBA{255, 255, 0, 255},
			color.RGBA{0, 0, 255, 255},
		).
		Build()
	colors = grad.ColorfulColors(3)

	if len(colors) != 3 {
		t.Errorf("Expected 3, got %v", len(colors))
	}
	testStr(t, colors[0].Hex(), "#ff0000")
	testStr(t, colors[1].Hex(), "#ffff00")
	testStr(t, colors[2].Hex(), "#0000ff")

	testStr(t, grad.At(math.NaN()).Hex(), "#000000")

	// Custom colors #2
	grad, _ = NewGradient().
		HtmlColors("#00f", "#00ffff").
		Colors(color.RGBA{255, 255, 0, 255}).
		HtmlColors("lime").
		Build()
	colors = grad.ColorfulColors(4)
	testStr(t, colors[0].Hex(), "#0000ff")
	testStr(t, colors[1].Hex(), "#00ffff")
	testStr(t, colors[2].Hex(), "#ffff00")
	testStr(t, colors[3].Hex(), "#00ff00")
}

func TestError(t *testing.T) {
	// Invalid HTML colors
	grad, err := NewGradient().
		HtmlColors("#777", "bloodred", "#bbb", "#zzz").
		Build()
	if err == nil {
		t.Errorf("It should error")
	}
	if !isZeroGradient(grad) {
		t.Errorf("It should zeroGradient")
	}

	// Wrong domain 1
	grad, err = NewGradient().
		HtmlColors("#777", "#fff", "#ccc", "#222").
		Domain(0, 0.5, 1).
		Build()
	if err == nil {
		t.Errorf("It should error")
	}
	if !isZeroGradient(grad) {
		t.Errorf("It should zeroGradient")
	}

	// Wrong domain 2
	grad, err = NewGradient().
		HtmlColors("#777", "#fff", "#ccc", "#222").
		Domain(0, 0.71, 0.70, 1).
		Build()
	if err == nil {
		t.Errorf("It should error")
	}
	if !isZeroGradient(grad) {
		t.Errorf("It should zeroGradient")
	}

	// Wrong domain 3
	grad, err = NewGradient().
		HtmlColors("#777", "#fff", "#ccc", "#222").
		Domain(1, 0).
		Build()
	if err == nil {
		t.Errorf("It should error")
	}
	if !isZeroGradient(grad) {
		t.Errorf("It should zeroGradient")
	}
}
