package colorgrad

import (
	"image/color"
	"testing"

	"github.com/lucasb-eyer/go-colorful"
)

func TestScheme(t *testing.T) {
	testScheme(t, Scheme.Accent)
	testScheme(t, Scheme.Category10)
	testScheme(t, Scheme.Dark2)
	testScheme(t, Scheme.Paired)
	testScheme(t, Scheme.Pastel1)
	testScheme(t, Scheme.Pastel2)
	testScheme(t, Scheme.Set1)
	testScheme(t, Scheme.Set2)
	testScheme(t, Scheme.Set3)
}

func testScheme(t *testing.T, colors []color.Color) {
	grad, _ := NewGradient().
		Colors(colors...).
		Build()

	if grad == nil {
		t.Error("grad is nil")
	}

	c1, _ := colorful.MakeColor(colors[0])
	a1 := c1.Hex()
	a2 := grad.At(0).Hex()

	if a1 != a2 {
		t.Errorf("Expected %v, got %v", a1, a2)
	}

	c2, _ := colorful.MakeColor(colors[len(colors)-1])
	b1 := c2.Hex()
	b2 := grad.At(1).Hex()

	if b1 != b2 {
		t.Errorf("Expected %v, got %v", b1, b2)
	}
}
