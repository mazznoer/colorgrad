package colorgrad

import (
    "image/color"
    "testing"
)

func TestBasic(t *testing.T) {
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

    grad, _ = NewGradient().
        Colors(
            color.RGBA{255,0,0,255},
            color.RGBA{255,255,0,255},
            color.RGBA{0,0,255,255},
        ).
        Build()
    colors = grad.Colors(3)

    if len(colors) != 3 {
        t.Errorf("%v", len(colors))
    }

    if colors[0].Hex() != "#ff0000" {
        t.Errorf("%v", colors[0].Hex())
    }

    if colors[1].Hex() != "#ffff00" {
        t.Errorf("%v", colors[0].Hex())
    }

    if colors[2].Hex() != "#0000ff" {
        t.Errorf("%v", colors[1].Hex())
    }
}

func TestDomain(t *testing.T) {
    grad, _ := NewGradient().
        HexColors("#000000", "#ff0000", "#ffffff").
        Domain(0, 0.75, 1).
        Build()
    
    testStr(t, grad.At(0).Hex(), "#000000")
    testStr(t, grad.At(0.75).Hex(), "#ff0000")
    testStr(t, grad.At(1).Hex(), "#ffffff")
    
    grad, _ = NewGradient().
        HexColors("#000000", "#ff0000", "#ffffff").
        Domain(15, 25, 60).
        Build()
    
    testStr(t, grad.At(15).Hex(), "#000000")
    testStr(t, grad.At(25).Hex(), "#ff0000")
    testStr(t, grad.At(60).Hex(), "#ffffff")
}

func testStr(t *testing.T, result, expected string) {
    if a != b {
        t.Errorf("Expected %v, get %v", expected, result)
    }
}
