package colorgrad

import (
	"testing"

	"github.com/mazznoer/csscolorparser"
)

func Test_Utils(t *testing.T) {
	// linspace

	test(t, len(linspace(0, 1, 0)), 0)
	testTrue(t, linspace(0, 1, 1)[0] == 0.0)
	testSlice(t, linspace(0, 1, 2), []float64{0, 1})
	testSlice(t, linspace(0, 1, 3), []float64{0, 0.5, 1})
	testSlice(t, linspace(0, 100, 3), []float64{0, 50, 100})

	// norm

	test(t, norm(0.99, 0, 1), 0.99)
	test(t, norm(12, 0, 100), 0.12)
	test(t, norm(753, 0, 1000), 0.753)

	// modulo

	test(t, modulo(2.73, 1), 0.73)
	test(t, modulo(32, 25), 7.0)

	// clamp01

	test(t, clamp01(0), 0.0)
	test(t, clamp01(1), 1.0)
	test(t, clamp01(0.997), 0.997)
	test(t, clamp01(-0.51), 0.0)
	test(t, clamp01(1.0001), 1.0)

	// parseFloat

	validData := []struct {
		str string
		num float64
	}{
		{"0", 0},
		{"0.0", 0},
		{"1234", 1234},
		{"0.00027", 0.00027},
		{"-56.03", -56.03},
	}
	for _, dt := range validData {
		f, ok := parseFloat(dt.str)
		testTrue(t, ok)
		test(t, f, dt.num)
	}

	invalidData := []string{
		"",
		" ",
		"25.0x",
		"1.0d7",
		"x10",
		"o",
	}
	for _, s := range invalidData {
		_, ok := parseFloat(s)
		testTrue(t, !ok)
	}

	// convertColors

	colors := []Color{
		Rgb(1, 0.7, 0.1, 0.5),
		//Rgb8(10, 255, 125, 0), //
		LinearRgb(0.1, 0.9, 1, 1),
		Hwb(0, 0, 0, 1),
		Hwb(320, 0.1, 0.3, 1),
		Hsv(120, 0.3, 0.2, 0.1),
		Hsl(120, 0.3, 0.2, 1),
	}

	for i, arr := range convertColors(colors, BlendRgb) {
		col := Rgb(spreadF64(arr))
		test(t, colors[i].HexString(), col.HexString())
	}

	for i, arr := range convertColors(colors, BlendLinearRgb) {
		col := LinearRgb(spreadF64(arr))
		test(t, colors[i].HexString(), col.HexString())
	}

	/*for i, arr := range convertColors(colors, BlendOklab) {
		col := Oklab(spreadF64(arr))
		test(t, colors[i].HexString(), col.HexString())
	}*/

	for _, c := range colors {
		col, err := csscolorparser.Parse(c.HexString())
		test(t, err, nil)

		x := Oklab(spreadF64(col2oklab(col)))
		test(t, x.HexString(), c.HexString())

		y := LinearRgb(spreadF64(col2linearRgb(col)))
		test(t, y.HexString(), c.HexString())
	}

	hexColors := []string{
		"#000000",
		"#ffffff",
		"#999999",
		"#135cdf",
		"#ff0000",
		"#00ff7f",
		//"#0aff7d", //
		//"#09ff7d", //
		"#abc5679b",
	}
	for _, s := range hexColors {
		col, err := csscolorparser.Parse(s)
		test(t, err, nil)
		test(t, col.HexString(), s)

		x := Oklab(spreadF64(col2oklab(col)))
		test(t, x.HexString(), s)

		y := LinearRgb(spreadF64(col2linearRgb(col)))
		test(t, y.HexString(), s)
	}
}

func spreadF64(arr [4]float64) (a, b, c, d float64) {
	a = arr[0]
	b = arr[1]
	c = arr[2]
	d = arr[3]
	return
}

// --- Helper functions

func test(t *testing.T, a, b any) {
	if a != b {
		t.Helper()
		t.Errorf("left: %v, right: %v", a, b)
	}
}

func testTrue(t *testing.T, b bool) {
	if !b {
		t.Helper()
		t.Errorf("it false")
	}
}

func testSlice[S ~[]E, E comparable](t *testing.T, a, b S) {
	if len(a) != len(b) {
		t.Helper()
		t.Errorf("different length -> left: %v, right: %v", len(a), len(b))
		return
	}
	for i, val := range a {
		if val != b[i] {
			t.Helper()
			t.Errorf("diff at index: %v, left: %v, right: %v", i, val, b[i])
			return
		}
	}
}

func colors2hex(colors []Color) []string {
	hexColors := make([]string, len(colors))
	for i, c := range colors {
		hexColors[i] = c.HexString()
	}
	return hexColors
}

func isZeroGradient(grad Gradient) bool {
	for _, col := range grad.Colors(13) {
		if col.HexString() != "#00000000" {
			return false
		}
	}
	return true
}
