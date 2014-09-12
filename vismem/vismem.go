// vismem visualizes memory locations
// +build !appengine

package main

import (
	"crypto/rand"
	"io"
	"os"

	"github.com/ajstarks/svgo"
)

var canvas = svg.New(os.Stdout)

func main() {
	width := 512
	height := 512
	n := 1024
	rowsize := 32
	diameter := 16
	var value int
	var r io.Reader

	if len(os.Args) > 1 {
		f, err := os.Open(os.Args[1])
		if err != nil {
			panic(err)
		}
		defer f.Close()
		r = f
	} else {
		r = rand.Reader
	}

	mem := make([]byte, n)
	if _, err := io.ReadFull(r, mem); err != nil {
		panic(err)
	}

	canvas.Start(width, height)
	canvas.Title("Visualize Files")
	canvas.Rect(0, 0, width, height, "fill:white")
	dx := diameter / 2
	dy := diameter / 2
	canvas.Gstyle("fill-opacity:1.0")
	for i := 0; i < n; i++ {
		value = int(mem[i])
		if i%rowsize == 0 && i != 0 {
			dx = diameter / 2
			dy += diameter
		}
		canvas.Circle(dx, dy, diameter/2, canvas.RGB(value, value, value))
		dx += diameter
	}
	canvas.Gend()
	canvas.End()
}
