package colorgrad

import (
	"math"
	"testing"
)

func TestBasic1(t *testing.T) {
	// Single color
	grad, err := NewGradient().
		Colors(Color{R: 0, G: 1, B: 0, A: 1}).
		Build()
	dmin, dmax := grad.Domain()
	test(t, err, nil)
	test(t, dmin, 0.0)
	test(t, dmax, 1.0)
	test(t, grad.At(0).HexString(), "#00ff00")
	test(t, grad.At(1).HexString(), "#00ff00")

	// Named colors
	grad, err = NewGradient().
		HtmlColors("tomato", "skyblue", "gold", "springgreen").
		Build()
	test(t, err, nil)
	testSlice(t, colors2hex(grad.Colors(4)), []string{
		"#ff6347",
		"#87ceeb",
		"#ffd700",
		"#00ff7f",
	})

	// Blend mode
	grad, err = NewGradient().
		HtmlColors("#333", "#bbb").
		Mode(BlendRgb).
		Build()
	test(t, err, nil)
	test(t, grad.At(0).HexString(), "#333333")
	test(t, grad.At(1).HexString(), "#bbbbbb")

	grad, err = NewGradient().
		HtmlColors("#333", "#bbb").
		Mode(BlendLinearRgb).
		Build()
	test(t, err, nil)
	test(t, grad.At(0).HexString(), "#333333")
	test(t, grad.At(1).HexString(), "#bbbbbb")

	grad, err = NewGradient().
		HtmlColors("#333", "#bbb").
		Mode(BlendOklab).
		Build()
	test(t, err, nil)
	test(t, grad.At(0).HexString(), "#333333")
	test(t, grad.At(1).HexString(), "#bbbbbb")
}

func TestBasic2(t *testing.T) {
	// Custom gradient default
	grad, err := NewGradient().Build()
	test(t, err, nil)
	testSlice(t, colors2hex(grad.Colors(2)), []string{
		"#000000",
		"#ffffff",
	})
	test(t, grad.At(math.NaN()).HexString(), "#000000")

	// Custom colors
	grad, err = NewGradient().
		Colors(
			Rgb8(255, 0, 0, 255),
			Rgb8(255, 255, 0, 255),
			Rgb8(0, 0, 255, 255),
		).
		Build()
	test(t, err, nil)
	testSlice(t, colors2hex(grad.Colors(3)), []string{
		"#ff0000",
		"#ffff00",
		"#0000ff",
	})
	test(t, grad.At(math.NaN()).HexString(), "#000000")

	// Custom colors #2
	grad, err = NewGradient().
		HtmlColors("#00f", "#00ffff").
		Colors(Rgb8(255, 255, 0, 255)).
		HtmlColors("lime").
		Build()
	test(t, err, nil)
	testSlice(t, colors2hex(grad.Colors(4)), []string{
		"#0000ff",
		"#00ffff",
		"#ffff00",
		"#00ff00",
	})
}

func TestFilterStops(t *testing.T) {
	gb := NewGradient()
	gb.HtmlColors("gold", "red", "blue", "yellow", "black", "white", "plum")
	gb.Domain(0, 0, 0.5, 0.5, 0.5, 1, 1)
	_, err := gb.Build()
	test(t, err, nil)
	testSlice(t, *gb.GetPositions(), []float64{0, 0.5, 0.5, 1})
	testSlice(t, colors2hex(*gb.GetColors()), []string{
		"#ff0000",
		"#0000ff",
		"#000000",
		"#ffffff",
	})
}

func TestError(t *testing.T) {
	// Invalid HTML colors
	grad, err := NewGradient().
		HtmlColors("#777", "bloodred", "#bbb", "#zzz").
		Build()
	testTrue(t, err != nil)
	testTrue(t, isZeroGradient(grad))

	// Wrong domain 1
	grad, err = NewGradient().
		HtmlColors("#777", "#fff", "#ccc", "#222").
		Domain(0, 0.5, 1).
		Build()
	testTrue(t, err != nil)
	testTrue(t, isZeroGradient(grad))

	// Wrong domain 2
	grad, err = NewGradient().
		HtmlColors("#777", "#fff", "#ccc", "#222").
		Domain(0, 0.71, 0.70, 1).
		Build()
	testTrue(t, err != nil)
	testTrue(t, isZeroGradient(grad))

	// Wrong domain 3
	grad, err = NewGradient().
		HtmlColors("#777", "#fff", "#ccc", "#222").
		Domain(1, 0).
		Build()
	testTrue(t, err != nil)
	testTrue(t, isZeroGradient(grad))
}
