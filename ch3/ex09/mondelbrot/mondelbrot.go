package mondelbrot

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"math/cmplx"
)

type MondelbrotOptions struct {
	width, height, scale int
}

func NewMondelbrotOptions(w, h, s int) MondelbrotOptions {
	mo := MondelbrotOptions{1024, 1024, 2}
	if w != 0 {
		mo.width = w
	}
	if h != 0 {
		mo.height = h
	}
	if s != 0 {
		mo.scale = s
	}

	return mo
}

func Mondelbrot(out io.Writer, mo MondelbrotOptions) {
	var (
		xmin, ymin, xmax, ymax = -mo.scale, -mo.scale, mo.scale, mo.scale
		width, height          = mo.width, mo.height
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/float64(height)*float64(ymax-ymin) + float64(ymin)
		for px := 0; px < width; px++ {
			x := float64(px)/float64(width)*float64(xmax-xmin) + float64(xmin)
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(out, img)
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}
