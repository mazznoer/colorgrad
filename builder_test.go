package colorgrad

import (
	"testing"
)

func domain(min, max float64) [2]float64 {
	return [2]float64{min, max}
}

func Test_Builder(t *testing.T) {
	var grad Gradient
	var err error

	// Default colors
	grad, err = NewGradient().Build()
	test(t, err, nil)
	test(t, domain(grad.Domain()), [2]float64{0, 1})
	test(t, grad.At(0).HexString(), "#000000")
	test(t, grad.At(1).HexString(), "#ffffff")

	// Single color
	grad, err = NewGradient().
		Colors(Rgb8(0, 255, 0, 255)).
		Build()
	test(t, err, nil)
	test(t, domain(grad.Domain()), [2]float64{0, 1})
	test(t, grad.At(0).HexString(), "#00ff00")
	test(t, grad.At(1).HexString(), "#00ff00")

	// Default domain
	grad, err = NewGradient().
		HtmlColors("red", "lime", "blue").
		Build()
	test(t, err, nil)
	test(t, domain(grad.Domain()), [2]float64{0, 1})
	test(t, grad.At(0.0).HexString(), "#ff0000")
	test(t, grad.At(0.5).HexString(), "#00ff00")
	test(t, grad.At(1.0).HexString(), "#0000ff")

	// Custom domain
	grad, err = NewGradient().
		HtmlColors("red", "lime", "blue").
		Domain(-100, 100).
		Build()
	test(t, err, nil)
	test(t, domain(grad.Domain()), [2]float64{-100, 100})
	test(t, grad.At(-100).HexString(), "#ff0000")
	test(t, grad.At(0).HexString(), "#00ff00")
	test(t, grad.At(100).HexString(), "#0000ff")

	// Color position
	grad, err = NewGradient().
		HtmlColors("red", "lime", "blue").
		Domain(13, 27.3, 90).
		Build()
	test(t, err, nil)
	test(t, domain(grad.Domain()), [2]float64{13, 90})
	test(t, grad.At(13).HexString(), "#ff0000")
	test(t, grad.At(27.3).HexString(), "#00ff00")
	test(t, grad.At(90).HexString(), "#0000ff")

	// Multiple colors, custom domain
	grad, err = NewGradient().
		HtmlColors("#00f", "#00ffff").
		Colors(Rgb8(255, 255, 0, 255), Hwb(0, 0, 0, 1)).
		HtmlColors("gold").
		Domain(10, 50).
		Build()
	test(t, err, nil)
	test(t, domain(grad.Domain()), [2]float64{10, 50})
	testSlice(t, colors2hex(grad.Colors(5)), []string{
		"#0000ff",
		"#00ffff",
		"#ffff00",
		"#ff0000",
		"#ffd700",
	})

	// Filter stops
	gb := NewGradient()
	gb.HtmlColors("gold", "red", "blue", "yellow", "black", "white", "plum")
	gb.Domain(0, 0, 0.5, 0.5, 0.5, 1, 1)
	_, err = gb.Build()
	test(t, err, nil)
	testSlice(t, *gb.GetPositions(), []float64{0, 0.5, 0.5, 1})
	testSlice(t, colors2hex(*gb.GetColors()), []string{
		"#ff0000",
		"#0000ff",
		"#000000",
		"#ffffff",
	})

	// --- Builder Error

	// Invalid HTML colors
	grad, err = NewGradient().
		HtmlColors("#777", "bloodred", "#bbb", "#zzz").
		Build()
	testTrue(t, err != nil)
	testTrue(t, isZeroGradient(grad))

	// Invalid domain
	grad, err = NewGradient().
		HtmlColors("#777", "#fff", "#ccc", "#222").
		Domain(0, 0.5, 1).
		Build()
	testTrue(t, err != nil)
	testTrue(t, isZeroGradient(grad))

	// Invalid domain
	grad, err = NewGradient().
		HtmlColors("#777", "#fff", "#ccc", "#222").
		Domain(0, 0.71, 0.70, 1).
		Build()
	testTrue(t, err != nil)
	testTrue(t, isZeroGradient(grad))

	// Invalid domain
	grad, err = NewGradient().
		HtmlColors("#f00", "#0f0").
		Domain(1, 1).
		Build()
	testTrue(t, err != nil)
	testTrue(t, isZeroGradient(grad))

	// Invalid domain
	grad, err = NewGradient().
		HtmlColors("#777", "#fff", "#ccc", "#222").
		Domain(1, 0).
		Build()
	testTrue(t, err != nil)
	testTrue(t, isZeroGradient(grad))
}
