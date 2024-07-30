package colorgrad

import (
	"testing"
)

func TestUtils(t *testing.T) {
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

	type data struct {
		str string
		num float64
	}

	validData := []data{
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
