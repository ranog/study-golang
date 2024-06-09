package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

// Lissajous gera GIF animado de figuras de Lissajous aleatórias.

var palette = []color.Color{color.Black, color.RGBA{R: 0, G: 255, B: 0, A: 255}}

const blackIndex = 1 // próxima cor da paleta

func main() {
	rand.NewSource(time.Now().UTC().UnixNano())
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // número de revoluções completas do oscilador x
		res     = 0.001 // resolução angular
		size    = 100   // tamanho da imagem [-size..+size]
		nframes = 64    // número de quadros da animação
		delay   = 8     // tempo entre quadros em unidades de 10ms
	)
	freq := rand.Float64() * 3.0 // frequência relativa do oscilador y
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // diferença de fase
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTA: ignorando erros de codificação
}
