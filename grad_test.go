package colorgrad

import (
    "testing"
)

func TestGrad(t *testing.T) {
    grad, _ := NewGradient().Build()
    colors := grad.Colors(2)

    if len(colors) != 2 {
        t.Errorf("%v", len(colors))
    }

    if colors[0].Hex() != "#000000" {
        t.Errorf("%v", colors[0].Hex())
    }

    if colors[1].Hex() != "#ffffff" {
        t.Errorf("%v", colors[1].Hex())
    }
}
