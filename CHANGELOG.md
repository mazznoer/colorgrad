# Changelog

## v0.12.0

- New `Gradient.ColorsIter()`
- New `Gradient.Linearize()`
- Expose all gradient types:
    - `LinearGradient`
    - `BasisGradient`
    - `SmoothstepGradient`
    - etc.

## v0.11.1

- Fix bug in Lab conversion

## v0.11.0

- New `GradientBuilder.Reset()`
- New `BlendLab`
- `Gradient`'s fields is now public
- `GradientCore` is now public
- New `InterpolationSmoothstep`

## v0.10.2

- Fix bug in `GradientBuilder.Css()`

## v0.10.1

- New `GradientBuilder.Css()`
- Using `uint32` to store preset colors

## v0.10.0

- Added support for transparency
- New `GoColor()`
- Refactor lots of code
- Remove color schemes
