package colorgrad

import (
	"testing"
)

func TestDomain(t *testing.T) {
	grad, _ := NewGradient().
		HtmlColors("yellow", "blue", "lime").
		Domain(0, 100).
		Build()

	dmin, dmax := grad.Domain()
	test(t, dmin, 0.0)
	test(t, dmax, 100.0)

	test(t, grad.At(0).HexString(), "#ffff00")
	test(t, grad.At(50).HexString(), "#0000ff")
	test(t, grad.At(100).HexString(), "#00ff00")
	// outside domain
	test(t, grad.At(-10).HexString(), "#ffff00")
	test(t, grad.At(150).HexString(), "#00ff00")

	grad, _ = NewGradient().
		HtmlColors("yellow", "blue", "lime").
		Domain(-1, 1).
		Build()

	dmin, dmax = grad.Domain()
	test(t, dmin, -1.0)
	test(t, dmax, 1.0)

	test(t, grad.At(-1).HexString(), "#ffff00")
	test(t, grad.At(0).HexString(), "#0000ff")
	test(t, grad.At(1).HexString(), "#00ff00")
	// outside domain
	test(t, grad.At(-1.5).HexString(), "#ffff00")
	test(t, grad.At(1.5).HexString(), "#00ff00")

	grad, _ = NewGradient().
		HtmlColors("#00ff00", "#ff0000", "#ffff00").
		Domain(0, 0.75, 1).
		Build()

	dmin, dmax = grad.Domain()
	test(t, dmin, 0.0)
	test(t, dmax, 1.0)

	test(t, grad.At(0).HexString(), "#00ff00")
	test(t, grad.At(0.75).HexString(), "#ff0000")
	test(t, grad.At(1).HexString(), "#ffff00")
	// outside domain
	test(t, grad.At(-1).HexString(), "#00ff00")
	test(t, grad.At(1.5).HexString(), "#ffff00")

	grad, _ = NewGradient().
		HtmlColors("#00ff00", "#ff0000", "#0000ff", "#ffff00").
		Domain(15, 25, 29, 63).
		Build()

	dmin, dmax = grad.Domain()
	test(t, dmin, 15.0)
	test(t, dmax, 63.0)

	test(t, grad.At(15).HexString(), "#00ff00")
	test(t, grad.At(25).HexString(), "#ff0000")
	test(t, grad.At(29).HexString(), "#0000ff")
	test(t, grad.At(63).HexString(), "#ffff00")
	// outside domain
	test(t, grad.At(10).HexString(), "#00ff00")
	test(t, grad.At(67).HexString(), "#ffff00")
}

func TestGetColors(t *testing.T) {
	grad, _ := NewGradient().Build()
	test(t, len(grad.Colors(0)), 0)
	test(t, grad.Colors(1)[0].HexString(), "#000000")
	testSlice(t, colors2hex(grad.Colors(2)), []string{
		"#000000",
		"#ffffff",
	})
	testSlice(t, colors2hex(grad.Colors(3)), []string{
		"#000000",
		"#808080",
		"#ffffff",
	})

	grad, _ = NewGradient().
		HtmlColors("#f00", "#0f0", "#00f").
		Domain(-1, 1).
		Build()

	testSlice(t, colors2hex(grad.Colors(5)), []string{
		"#ff0000",
		"#808000",
		"#00ff00",
		"#008080",
		"#0000ff",
	})
}

func TestSpreadRepeat(t *testing.T) {
	grad, _ := NewGradient().
		HtmlColors("#000", "#fff").
		Build()

	test(t, grad.RepeatAt(-2.0).HexString(), "#000000")
	test(t, grad.RepeatAt(-1.9).HexString(), "#1a1a1a")
	test(t, grad.RepeatAt(-1.5).HexString(), "#808080")
	test(t, grad.RepeatAt(-1.1).HexString(), "#e5e5e5")

	test(t, grad.RepeatAt(-1.0).HexString(), "#000000")
	test(t, grad.RepeatAt(-0.9).HexString(), "#191919")
	test(t, grad.RepeatAt(-0.5).HexString(), "#808080")
	test(t, grad.RepeatAt(-0.1).HexString(), "#e6e6e6")

	test(t, grad.RepeatAt(0.0).HexString(), "#000000")
	test(t, grad.RepeatAt(0.1).HexString(), "#1a1a1a")
	test(t, grad.RepeatAt(0.5).HexString(), "#808080")
	test(t, grad.RepeatAt(0.9).HexString(), "#e5e5e5")

	test(t, grad.RepeatAt(1.0).HexString(), "#000000")
	test(t, grad.RepeatAt(1.1).HexString(), "#1a1a1a")
	test(t, grad.RepeatAt(1.5).HexString(), "#808080")
	test(t, grad.RepeatAt(1.9).HexString(), "#e5e5e5")

	test(t, grad.RepeatAt(2.0).HexString(), "#000000")
	test(t, grad.RepeatAt(2.1).HexString(), "#1a1a1a")
	test(t, grad.RepeatAt(2.5).HexString(), "#808080")
	test(t, grad.RepeatAt(2.9).HexString(), "#e5e5e5")
}

func TestSpreadReflect(t *testing.T) {
	grad, _ := NewGradient().
		HtmlColors("#000", "#fff").
		Build()

	test(t, grad.ReflectAt(-2.0).HexString(), "#000000")
	test(t, grad.ReflectAt(-1.9).HexString(), "#1a1a1a")
	test(t, grad.ReflectAt(-1.5).HexString(), "#808080")
	test(t, grad.ReflectAt(-1.1).HexString(), "#e5e5e5")

	test(t, grad.ReflectAt(-1.0).HexString(), "#ffffff")
	test(t, grad.ReflectAt(-0.9).HexString(), "#e5e5e5")
	test(t, grad.ReflectAt(-0.5).HexString(), "#808080")
	test(t, grad.ReflectAt(-0.1).HexString(), "#1a1a1a")

	test(t, grad.ReflectAt(0.0).HexString(), "#000000")
	test(t, grad.ReflectAt(0.1).HexString(), "#1a1a1a")
	test(t, grad.ReflectAt(0.5).HexString(), "#808080")
	test(t, grad.ReflectAt(0.9).HexString(), "#e5e5e5")

	test(t, grad.ReflectAt(1.0).HexString(), "#ffffff")
	test(t, grad.ReflectAt(1.1).HexString(), "#e5e5e5")
	test(t, grad.ReflectAt(1.5).HexString(), "#808080")
	test(t, grad.ReflectAt(1.9).HexString(), "#1a1a1a")

	test(t, grad.ReflectAt(2.0).HexString(), "#000000")
	test(t, grad.ReflectAt(2.1).HexString(), "#1a1a1a")
	test(t, grad.ReflectAt(2.5).HexString(), "#808080")
	test(t, grad.ReflectAt(2.9).HexString(), "#e5e5e5")
}
