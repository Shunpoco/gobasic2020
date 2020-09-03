package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

var black = color.RGBA{0, 0, 0, 255}
var red = color.RGBA{255, 0, 0, 255}
var green = color.RGBA{0, 255, 0, 255}
var blue = color.RGBA{0, 0, 255, 255}
var palette = []color.Color{black, red, green, blue}

const (
	blackIndex = 0
	redIndex   = 1
	greenIndex = 2
	blueIndex  = 3
)

type LisajousOptions struct {
	cycles  float64
	res     float64
	size    int
	nframes int
	delay   int
}

func main() {
	http.HandleFunc("/", handler)
	if err := http.ListenAndServe(fmt.Sprintf("localhost:%d", 8888), nil); err != nil {
		panic(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	lo := &LisajousOptions{
		fetchProperValue(r.Form, "cycles", 5),
		fetchProperValue(r.Form, "res", 0.001),
		int(fetchProperValue(r.Form, "size", 100)),
		int(fetchProperValue(r.Form, "nframes", 64)),
		int(fetchProperValue(r.Form, "delay", 8)),
	}
	lisajous(w, *lo)
}

func fetchProperValue(values url.Values, key string, defaultValue float64) (result float64) {
	result = defaultValue
	if val := values.Get(key); val != "" {
		value, err := strconv.ParseFloat(val, 64)
		if err != nil {
			return
		}
		result = value
	}
	return
}

func lisajous(out io.Writer, options LisajousOptions) {
	var (
		cycles  = options.cycles
		res     = options.res
		size    = options.size
		nframes = options.nframes
		delay   = options.delay
	)
	rand.Seed(time.Now().UnixNano())
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5), uint8(rand.Int()%3+1))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
