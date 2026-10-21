package integration_test

import (
	"image/color"
	"math"
	"testing"

	"github.com/mazznoer/colorgrad"
)

// --- Helper functions

var epsilon float64 = 1e-10

func test(t *testing.T, a, b interface{}, desc ...string) {
	if a != b {
		t.Helper()
		t.Errorf("left: %v, right: %v", a, b)

		if len(desc) > 0 {
			t.Errorf("%s", desc[0])
		}
	}
}

func testF(t *testing.T, a, b float64, desc ...string) {
	if math.Abs(a-b) > epsilon {
		t.Helper()
		t.Errorf("left: %v, right: %v", a, b)

		if len(desc) > 0 {
			t.Errorf("%s", desc[0])
		}
	}
}

func testTrue(t *testing.T, b bool, desc ...string) {
	if !b {
		t.Helper()
		t.Errorf("it false")

		if len(desc) > 0 {
			t.Errorf("%s", desc[0])
		}
	}
}

func testColor(t *testing.T, a, b colorgrad.Color) {
	x := arr8(a.RGBA255())
	y := arr8(b.RGBA255())
	if x == y {
		return
	}
	t.Helper()
	t.Errorf("left: %v, right: %v", x, y)
}

func testGoColor(t *testing.T, a, b color.Color) {
	x := arr32(a.RGBA())
	y := arr32(b.RGBA())
	if x == y {
		return
	}
	t.Helper()
	t.Errorf("left: %v, right: %v", x, y)
}

func arr8(r, g, b, a uint8) [4]uint8 {
	return [4]uint8{r, g, b, a}
}

func arr32(r, g, b, a uint32) [4]uint32 {
	return [4]uint32{r, g, b, a}
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

func testSliceF(t *testing.T, a, b []float64) {
	if len(a) != len(b) {
		t.Helper()
		t.Errorf("different length -> left: %v, right: %v", len(a), len(b))
		return
	}
	for i, val := range a {
		if math.Abs(val-b[i]) > epsilon {
			t.Helper()
			t.Errorf("diff at index: %v, left: %v, right: %v", i, val, b[i])
			return
		}
	}
}

func colors2hex(colors []colorgrad.Color) []string {
	hexColors := make([]string, len(colors))
	for i, c := range colors {
		hexColors[i] = c.HexString()
	}
	return hexColors
}
