package colorgrad

import (
	"image/color"
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
	gb := NewGradient()
	grad, err = gb.HtmlColors("#00f", "#00ffff").
		Colors(
			Rgb8(255, 255, 0, 255),
			Hwb(320, 0.1, 0.3, 1),
			GoColor(color.RGBA{R: 127, G: 0, B: 0, A: 127}),
			GoColor(color.Gray{185}),
		).
		HtmlColors("gold", "hwb(320, 10%, 30%)").
		Domain(10, 50).
		Mode(BlendRgb).
		Interpolation(InterpolationLinear).
		Build()
	test(t, err, nil)
	test(t, domain(grad.Domain()), [2]float64{10, 50})
	testSlice(t, colors2hex(grad.Colors(8)), []string{
		"#0000ff",
		"#00ffff",
		"#ffff00",
		"#b31980", // xxx
		"#ff00007f",
		"#b9b9b9",
		"#ffd700",
		"#b31a80",
	})
	testSlice(t, colors2hex(*gb.GetColors()), []string{
		"#0000ff",
		"#00ffff",
		"#ffff00",
		"#b31a80",
		"#ff00007f",
		"#b9b9b9",
		"#ffd700",
		"#b31a80",
	})

	// Filter stops
	gb = NewGradient()
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

func Test_CssGradient(t *testing.T) {
	testData := []struct {
		s        string
		position []float64
		colors   []string
	}{
		{
			"blue",
			[]float64{0, 1},
			[]string{"#0000ff", "#0000ff"},
		},
		{
			"red, blue",
			[]float64{0, 1},
			[]string{"#ff0000", "#0000ff"},
		},
		{
			"red, lime, blue",
			[]float64{0, 0.5, 1},
			[]string{"#ff0000", "#00ff00", "#0000ff"},
		},
		{
			"red, lime 75%, blue",
			[]float64{0, 0.75, 1},
			[]string{"#ff0000", "#00ff00", "#0000ff"},
		},
		{
			"red 70%, lime, blue",
			[]float64{0, 0.7, 0.85, 1},
			[]string{"#ff0000", "#ff0000", "#00ff00", "#0000ff"},
		},
		{
			"#00f, #ff0 10% 35%, #f00",
			[]float64{0, 0.1, 0.35, 1},
			[]string{"#0000ff", "#ffff00", "#ffff00", "#ff0000"},
		},
		{
			"red, 75%, #ff0",
			[]float64{0, 0.75, 1},
			[]string{"#ff0000", "#ff8000", "#ffff00"},
		},
	}

	for _, d := range testData {
		gb := NewGradient()
		gb.Css(d.s)
		grad, err := gb.Build()

		test(t, err, nil)
		testSlice(t, d.position, *gb.GetPositions())
		testSlice(t, d.colors, colors2hex(*gb.GetColors()))

		dmin, dmax := grad.Domain()
		test(t, 0.0, dmin)
		test(t, 1.0, dmax)
	}

	// Invalid format
	invalid := []string{
		"",
		" ",
		"reds, blue",
		"0, red, lime",
		"red, lime, 100%",
		"deeppink, 0.4, 0.9, pink",
		"50%",
		"0%, 100%",
		"æ",
		"red â 15%, blue",
		"red, ä, blue",
	}
	for _, s := range invalid {
		_, err := NewGradient().Css(s).Build()
		testTrue(t, err != nil)
		test(t, err.Error(), "invalid CSS gradient")
	}
}
