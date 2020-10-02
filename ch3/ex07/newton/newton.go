package newton

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"math/cmplx"
	"os"
)

var palette = []color.Color{
	color.RGBA{255, 10, 10, 255}, // red
	color.RGBA{10, 255, 10, 255}, // green
	color.RGBA{10, 10, 255, 255}, // blue
}

func Fractal() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, newton(z))
		}
	}
	png.Encode(os.Stdout, img)
}

// f(z) = z^4 - 1
// f'(z) = 4*z^3
// z_n+1 = z_n - f(z_n) / f'(z_n)
// = z - (z^4 - 1) / (4*z^3)
func newton(z complex128) color.Color {
	const iterations = 200
	const contrast = 3

	for n := uint8(0); n < iterations; n++ {
		z = z - (z*z*z*z-1)/(4*z*z*z)
		zeroApproximate := 1e-5
		if cmplx.Abs(z*z*z*z-1) < zeroApproximate {
			if math.Abs(real(z)-1) < zeroApproximate && math.Abs(imag(z)) < zeroApproximate {
				return color.RGBA{255 - contrast*n, 0, 0, 255}
			}
			if math.Abs(real(z)+1) < zeroApproximate && math.Abs(imag(z)) < zeroApproximate {
				return color.RGBA{0, 255 - contrast*n, 0, 255}
			}
			if math.Abs(real(z)) < zeroApproximate && math.Abs(imag(z)-1) < zeroApproximate {
				return color.RGBA{0, 0, 255 - contrast*n, 255}
			}
			return color.RGBA{0, 255 - contrast*n, 255 - contrast*n, 255}
		}
	}
	return color.Black
}
