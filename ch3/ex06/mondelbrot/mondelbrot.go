package mondelbrot

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func Mondelbrot() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y0 := float64(py)/height*(ymax-ymin) + ymin
		y1 := (float64(py)+1)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x0 := float64(px)/width*(xmax-xmin) + xmin
			x1 := (float64(px)+1)/width*(xmax-xmin) + xmin
			z00 := complex(x0, y0)
			z01 := complex(x0, y1)
			z10 := complex(x1, y0)
			z11 := complex(x1, y1)
			// Image point (px, py) represents complex value z.
			z := averageColor([]color.Color{
				mandelbrot(z00),
				mandelbrot(z01),
				mandelbrot(z10),
				mandelbrot(z11),
			})
			img.Set(px, py, z)
		}
	}
	png.Encode(os.Stdout, img)
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

func averageColor(colors []color.Color) color.Color {
	n := len(colors)
	var r, g, b, a int
	for _, c := range colors {
		dr, dg, db, da := c.RGBA()
		r += int(dr)
		g += int(dg)
		b += int(db)
		a += int(da)
	}
	return color.RGBA{uint8(r / n), uint8(g / n), uint8(b / n), uint8(a / n)}
}
