package colorgrad_test

import (
	"fmt"

	"github.com/mazznoer/colorgrad"
)

func Example_presetGradient() {
	grad := colorgrad.Rainbow()
	dmin, dmax := grad.Domain()

	fmt.Println(dmin, dmax)
	fmt.Println(grad.At(0).HexString())
	// Output:
	// 0 1
	// #6e40aa
}

func Example_customGradient() {
	grad, err := colorgrad.NewGradient().
		HtmlColors("red", "#FFD700", "lime").
		Domain(0, 0.35, 1).
		Mode(colorgrad.BlendOklab).
		Build()

	if err != nil {
		panic(err)
	}

	fmt.Println(grad.At(0).HexString())
	fmt.Println(grad.At(1).HexString())
	// Output:
	// #ff0000
	// #00ff00
}
