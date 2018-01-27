// Generates GIF animations of random Lissajous figures

package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

var palette = []color.Color{color.RGBA{0x33, 0x33, 0x33, 0xff}, color.RGBA{0x00, 0xcc, 0x66, 0xff}, color.RGBA{0xff, 0xbf, 0x00, 0xff}}

const (
	grayIndex  = 0
	greenIndex = 1
	goldIndex  = 2
)

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0

	for i := 0; i < nframes; i++ {
		var crayon uint8
		if i%2 == 0 {
			crayon = goldIndex
		} else {
			crayon = greenIndex
		}
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			// randomIndex := rand.Int() % len(palette)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), crayon)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	gif.EncodeAll(out, &anim)
}
