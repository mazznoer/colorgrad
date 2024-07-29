package colorgrad

import (
	"testing"
)

func TestLinspace(t *testing.T) {
	test(t, len(linspace(0, 1, 0)), 0)
	testTrue(t, linspace(0, 1, 1)[0] == 0.0)
	testSlice(t, linspace(0, 1, 2), []float64{0, 1})
	testSlice(t, linspace(0, 1, 3), []float64{0, 0.5, 1})
	testSlice(t, linspace(0, 100, 3), []float64{0, 50, 100})
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
