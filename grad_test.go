package colorgrad

import (
	"image/color"
	"math"
	"testing"

	"github.com/lucasb-eyer/go-colorful"
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
		Mode(HCL).
		Build()
	testStr(t, grad.At(0).Hex(), "#333333")
	testStr(t, grad.At(1).Hex(), "#bbbbbb")

	grad, _ = NewGradient().
		HtmlColors("#333", "#bbb").
		Mode(HSV).
		Build()
	testStr(t, grad.At(0).Hex(), "#333333")
	testStr(t, grad.At(1).Hex(), "#bbbbbb")

	grad, _ = NewGradient().
		HtmlColors("#333", "#bbb").
		Mode(LAB).
		Build()
	testStr(t, grad.At(0).Hex(), "#333333")
	testStr(t, grad.At(1).Hex(), "#bbbbbb")

	grad, _ = NewGradient().
		HtmlColors("#333", "#bbb").
		Mode(LRGB).
		Build()
	testStr(t, grad.At(0).Hex(), "#333333")
	testStr(t, grad.At(1).Hex(), "#bbbbbb")

	grad, _ = NewGradient().
		HtmlColors("#333", "#bbb").
		Mode(LUV).
		Build()
	testStr(t, grad.At(0).Hex(), "#333333")
	testStr(t, grad.At(1).Hex(), "#bbbbbb")

	grad, _ = NewGradient().
		HtmlColors("#333", "#bbb").
		Mode(RGB).
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

	testStr(t, grad.At(math.NaN()).Hex(), "#ff0000")

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
}

func TestDomain(t *testing.T) {
	grad, _ := NewGradient().
		HtmlColors("yellow", "blue", "lime").
		Domain(0, 100).
		Build()

	testStr(t, grad.At(0).Hex(), "#ffff00")
	testStr(t, grad.At(50).Hex(), "#0000ff")
	testStr(t, grad.At(100).Hex(), "#00ff00")
	// outside domain
	testStr(t, grad.At(-10).Hex(), "#ffff00")
	testStr(t, grad.At(150).Hex(), "#00ff00")

	grad, _ = NewGradient().
		HtmlColors("yellow", "blue", "lime").
		Domain(-1, 1).
		Build()

	testStr(t, grad.At(-1).Hex(), "#ffff00")
	testStr(t, grad.At(0).Hex(), "#0000ff")
	testStr(t, grad.At(1).Hex(), "#00ff00")
	// outside domain
	testStr(t, grad.At(-1.5).Hex(), "#ffff00")
	testStr(t, grad.At(1.5).Hex(), "#00ff00")

	grad, _ = NewGradient().
		HtmlColors("#00ff00", "#ff0000", "#ffff00").
		Domain(0, 0.75, 1).
		Build()

	testStr(t, grad.At(0).Hex(), "#00ff00")
	testStr(t, grad.At(0.75).Hex(), "#ff0000")
	testStr(t, grad.At(1).Hex(), "#ffff00")
	// outside domain
	testStr(t, grad.At(-1).Hex(), "#00ff00")
	testStr(t, grad.At(1.5).Hex(), "#ffff00")

	grad, _ = NewGradient().
		HtmlColors("#00ff00", "#ff0000", "#0000ff", "#ffff00").
		Domain(15, 25, 29, 63).
		Build()

	testStr(t, grad.At(15).Hex(), "#00ff00")
	testStr(t, grad.At(25).Hex(), "#ff0000")
	testStr(t, grad.At(29).Hex(), "#0000ff")
	testStr(t, grad.At(63).Hex(), "#ffff00")
	// outside domain
	testStr(t, grad.At(10).Hex(), "#00ff00")
	testStr(t, grad.At(67).Hex(), "#ffff00")
}

func TestSharp(t *testing.T) {
	grad, _ := NewGradient().Build()
	grad2 := grad.Sharp(7)
	testStr(t, grad2.At(0).Hex(), "#000000")
	testStr(t, grad2.At(1).Hex(), "#ffffff")

	testStr(t, grad2.At(math.NaN()).Hex(), "#000000")
	testStr(t, grad2.At(-0.01).Hex(), "#000000")
	testStr(t, grad2.At(1.01).Hex(), "#ffffff")

	colors := grad2.ColorfulColors(7)
	if len(colors) != 7 {
		t.Errorf("Expected 7, got %v", len(colors))
	}
	testStr(t, colors[0].Hex(), "#000000")
	testStr(t, colors[6].Hex(), "#ffffff")
}

func TestGetColors(t *testing.T) {
	grad, _ := NewGradient().Build()
	colors1 := grad.ColorfulColors(5) // []colorful.Color
	colors2 := grad.Colors(5)         // []color.Color

	for i, c2 := range colors2 {
		var c1 color.Color = colors1[i]
		if c1 != c2 {
			t.Errorf("%v != %v", c1, c2)
		}
	}
}

func testStr(t *testing.T, result, expected string) {
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func isZeroGradient(grad Gradient) bool {
	dmin, dmax := grad.Domain()
	if dmin != 0 || dmax != 1 {
		return false
	}
	colors := grad.ColorfulColors(13)
	black := colorful.Color{R: 0, G: 0, B: 0}
	for _, col := range colors {
		if col != black {
			return false
		}
	}
	return true
}
