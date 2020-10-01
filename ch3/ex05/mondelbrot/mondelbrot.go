package mondelbrot

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

const (
	baseRed   = 127
	baseGreen = 127
	baseBlue  = 127
)

var palette = []color.Color{
	color.RGBA{255, 10, 10, 255}, // red
	color.RGBA{10, 255, 10, 255}, // green
	color.RGBA{10, 10, 255, 255}, // blue
}

func Mondelbrot() {
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
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(os.Stdout, img)
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200

	var v complex128
	for n := 0; n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return palette[n%len(palette)]
		}
	}
	return color.Black
}
