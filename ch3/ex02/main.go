package main

import (
	"ch3/ex02/surface"
	"flag"
)

var surfaceGraphName = flag.String("s", "default", "a name of surface function: default, eggbox, saddle and mogul")

func main() {
	flag.Parse()
	surface.Surface(*surfaceGraphName)
}
