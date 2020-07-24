package colorgrad

import (
    "testing"
)

func TestGrad(t *testing.T) {
    grad := NewGradient().Build()
    colors := grad.Colors(2)

    if len(colors) != 3 {
        t.Errorf("%v", len(colors))
    }

    if colors[0].Hex() != "#000001" {
        t.Errorf("%v", colors[0].Hex())
    }

    if colors[1].Hex() != "#ffffff" {
        t.Errorf("%v", colors[1].Hex())
    }
}
