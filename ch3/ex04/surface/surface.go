package surface

import (
	"fmt"
	"io"
	"math"
)

const (
	width, height = 600, 320            // canva size in pixels
	cells         = 100                 // number of prid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30degree)
	color         = "grey"
)

type SurfaceGraphOptions struct {
	width  int
	height int
	color  string
}

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30d), cos(30d)

func NewSurfaceGraphOptions(w, h int, c string) SurfaceGraphOptions {
	sgo := SurfaceGraphOptions{width, height, color}
	if w != 0 {
		sgo.width = w
	}
	if h != 0 {
		sgo.height = h
	}
	if c != "" {
		sgo.color = c
	}

	return sgo
}

func Surface(out io.Writer, sgo SurfaceGraphOptions) {
	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", sgo.width, sgo.height)
	c := &Corner{sgo.width, sgo.height}
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := c.corner(i+1, j)
			bx, by := c.corner(i, j)
			cx, cy := c.corner(i, j+1)
			dx, dy := c.corner(i+1, j+1)
			fmt.Fprintf(out, "<polygon points='%g,%g %g,%g %g,%g %g,%g' style='stroke: %s;'/>\n", ax, ay, bx, by, cx, cy, dx, dy, sgo.color)
		}
	}
	fmt.Fprintf(out, "</svg>")
}

type Corner struct {
	width  int
	height int
}

func (c *Corner) corner(i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx, sy).
	sx := float64(c.width)/2 + (x-y)*cos30*xyscale
	sy := float64(c.height)/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy

}

func corner(i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx, sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0, 0)
	return math.Sin(r) / r
}
