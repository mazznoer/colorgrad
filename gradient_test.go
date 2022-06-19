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

func TestSharpGradient(t *testing.T) {
	grad, _ := NewGradient().
		HtmlColors("#f00", "#0f0", "#00f").
		Build()

	// Sharp(0)
	grad0 := grad.Sharp(0, 0)
	testStr(t, grad0.At(0.0).Hex(), "#ff0000")
	testStr(t, grad0.At(0.5).Hex(), "#ff0000")
	testStr(t, grad0.At(1.0).Hex(), "#ff0000")

	// Sharp(1)
	grad1 := grad.Sharp(1, 0)
	testStr(t, grad1.At(0.0).Hex(), "#ff0000")
	testStr(t, grad1.At(0.5).Hex(), "#ff0000")
	testStr(t, grad1.At(1.0).Hex(), "#ff0000")

	// Sharp(3)
	grad3 := grad.Sharp(3, 0)
	testStr(t, grad3.At(0.0).Hex(), "#ff0000")
	testStr(t, grad3.At(0.2).Hex(), "#ff0000")

	testStr(t, grad3.At(0.4).Hex(), "#00ff00")
	testStr(t, grad3.At(0.5).Hex(), "#00ff00")
	testStr(t, grad3.At(0.6).Hex(), "#00ff00")

	testStr(t, grad3.At(0.9).Hex(), "#0000ff")
	testStr(t, grad3.At(1.0).Hex(), "#0000ff")

	testStr(t, grad3.At(-0.1).Hex(), "#ff0000")
	testStr(t, grad3.At(1.1).Hex(), "#0000ff")
	testStr(t, grad3.At(math.NaN()).Hex(), "#ff0000")

	// Sharp(2)
	grad, _ = NewGradient().
		HtmlColors("#f00", "#0f0", "#00f").
		Domain(-1, 1).
		Build()
	grad2 := grad.Sharp(2, 0)
	testStr(t, grad2.At(-1.0).Hex(), "#ff0000")
	testStr(t, grad2.At(-0.5).Hex(), "#ff0000")
	testStr(t, grad2.At(-0.1).Hex(), "#ff0000")

	testStr(t, grad2.At(0.1).Hex(), "#0000ff")
	testStr(t, grad2.At(0.5).Hex(), "#0000ff")
	testStr(t, grad2.At(1.0).Hex(), "#0000ff")
}

func TestSharpGradientX(t *testing.T) {
	grad, _ := NewGradient().
		HtmlColors("#f00", "#0f0", "#00f").
		Build()

	grad0 := grad.Sharp(0, 0.5)
	testStr(t, grad0.At(0.0).Hex(), "#ff0000")
	testStr(t, grad0.At(0.5).Hex(), "#ff0000")
	testStr(t, grad0.At(1.0).Hex(), "#ff0000")

	grad1 := grad.Sharp(1, 0.5)
	testStr(t, grad1.At(0.0).Hex(), "#ff0000")
	testStr(t, grad1.At(0.5).Hex(), "#ff0000")
	testStr(t, grad1.At(1.0).Hex(), "#ff0000")

	grad = grad.Sharp(3, 0.1)

	testStr(t, grad.At(0).Hex(), "#ff0000")
	testStr(t, grad.At(0.1).Hex(), "#ff0000")

	testStr(t, grad.At(1.0/3).Hex(), "#808000")

	testStr(t, grad.At(0.45).Hex(), "#00ff00")
	testStr(t, grad.At(0.55).Hex(), "#00ff00")

	testStr(t, grad.At(1.0/3*2).Hex(), "#008080")

	testStr(t, grad.At(0.9).Hex(), "#0000ff")
	testStr(t, grad.At(1).Hex(), "#0000ff")

	testStr(t, grad.At(-0.01).Hex(), "#ff0000")
	testStr(t, grad.At(1.01).Hex(), "#0000ff")
	testStr(t, grad.At(math.NaN()).Hex(), "#ff0000")
}

func TestGetColors(t *testing.T) {
	grad, _ := NewGradient().Build()
	colorsA := grad.ColorfulColors(5) // []colorful.Color
	colorsB := grad.Colors(5)         // []color.Color

	for i, c2 := range colorsB {
		var c1 color.Color = colorsA[i]
		if c1 != c2 {
			t.Errorf("%v != %v", c1, c2)
		}
	}

	colors0 := grad.ColorfulColors(0)
	if len(colors0) != 0 {
		t.Errorf("Error.")
	}
	colors1 := grad.ColorfulColors(1)
	testStr(t, colors1[0].Hex(), "#000000")

	colors2 := grad.ColorfulColors(2)
	testStr(t, colors2[0].Hex(), "#000000")
	testStr(t, colors2[1].Hex(), "#ffffff")

	colors3 := grad.ColorfulColors(3)
	testStr(t, colors3[0].Hex(), "#000000")
	testStr(t, colors3[1].Hex(), "#808080")
	testStr(t, colors3[2].Hex(), "#ffffff")

	grad, _ = NewGradient().
		HtmlColors("#f00", "#0f0", "#00f").
		Domain(-1, 1).
		Build()

	colors5 := grad.ColorfulColors(5)
	testStr(t, colors5[0].Hex(), "#ff0000")
	testStr(t, colors5[1].Hex(), "#808000")
	testStr(t, colors5[2].Hex(), "#00ff00")
	testStr(t, colors5[3].Hex(), "#008080")
	testStr(t, colors5[4].Hex(), "#0000ff")
}

func TestSpreadRepeat(t *testing.T) {
	grad, _ := NewGradient().
		HtmlColors("#000", "#fff").
		Build()

	testStr(t, grad.RepeatAt(-2.0).Hex(), "#000000")
	testStr(t, grad.RepeatAt(-1.9).Hex(), "#1a1a1a")
	testStr(t, grad.RepeatAt(-1.5).Hex(), "#808080")
	testStr(t, grad.RepeatAt(-1.1).Hex(), "#e5e5e5")

	testStr(t, grad.RepeatAt(-1.0).Hex(), "#000000")
	testStr(t, grad.RepeatAt(-0.9).Hex(), "#191919")
	testStr(t, grad.RepeatAt(-0.5).Hex(), "#808080")
	testStr(t, grad.RepeatAt(-0.1).Hex(), "#e6e6e6")

	testStr(t, grad.RepeatAt(0.0).Hex(), "#000000")
	testStr(t, grad.RepeatAt(0.1).Hex(), "#1a1a1a")
	testStr(t, grad.RepeatAt(0.5).Hex(), "#808080")
	testStr(t, grad.RepeatAt(0.9).Hex(), "#e5e5e5")

	testStr(t, grad.RepeatAt(1.0).Hex(), "#000000")
	testStr(t, grad.RepeatAt(1.1).Hex(), "#1a1a1a")
	testStr(t, grad.RepeatAt(1.5).Hex(), "#808080")
	testStr(t, grad.RepeatAt(1.9).Hex(), "#e5e5e5")

	testStr(t, grad.RepeatAt(2.0).Hex(), "#000000")
	testStr(t, grad.RepeatAt(2.1).Hex(), "#1a1a1a")
	testStr(t, grad.RepeatAt(2.5).Hex(), "#808080")
	testStr(t, grad.RepeatAt(2.9).Hex(), "#e5e5e5")
}

func TestSpreadReflect(t *testing.T) {
	grad, _ := NewGradient().
		HtmlColors("#000", "#fff").
		Build()

	testStr(t, grad.ReflectAt(-2.0).Hex(), "#000000")
	testStr(t, grad.ReflectAt(-1.9).Hex(), "#1a1a1a")
	testStr(t, grad.ReflectAt(-1.5).Hex(), "#808080")
	testStr(t, grad.ReflectAt(-1.1).Hex(), "#e5e5e5")

	testStr(t, grad.ReflectAt(-1.0).Hex(), "#ffffff")
	testStr(t, grad.ReflectAt(-0.9).Hex(), "#e5e5e5")
	testStr(t, grad.ReflectAt(-0.5).Hex(), "#808080")
	testStr(t, grad.ReflectAt(-0.1).Hex(), "#1a1a1a")

	testStr(t, grad.ReflectAt(0.0).Hex(), "#000000")
	testStr(t, grad.ReflectAt(0.1).Hex(), "#1a1a1a")
	testStr(t, grad.ReflectAt(0.5).Hex(), "#808080")
	testStr(t, grad.ReflectAt(0.9).Hex(), "#e5e5e5")

	testStr(t, grad.ReflectAt(1.0).Hex(), "#ffffff")
	testStr(t, grad.ReflectAt(1.1).Hex(), "#e5e5e5")
	testStr(t, grad.ReflectAt(1.5).Hex(), "#808080")
	testStr(t, grad.ReflectAt(1.9).Hex(), "#1a1a1a")

	testStr(t, grad.ReflectAt(2.0).Hex(), "#000000")
	testStr(t, grad.ReflectAt(2.1).Hex(), "#1a1a1a")
	testStr(t, grad.ReflectAt(2.5).Hex(), "#808080")
	testStr(t, grad.ReflectAt(2.9).Hex(), "#e5e5e5")
}

func testStr(t *testing.T, a, b string) {
	if a != b {
		t.Helper()
		t.Errorf("Left: %s, right: %s", a, b)
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
