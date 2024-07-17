//go:build ignore
package colorgrad

import (
	"image/color"
	"testing"

	"github.com/lucasb-eyer/go-colorful"
)

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
