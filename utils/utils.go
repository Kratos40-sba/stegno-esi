package utils

import (
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"os"
)

// XOR function
func Encdec(in, cle string) (o string) {
	l := len(cle)
	for i := range in {
		o += string(in[i] ^ cle[i%l])
	}
	return
}

// ByToBin
func ByToBin(b []byte) (bin string) {
	for _, bb := range b {
		bin = fmt.Sprintf("%s%.8b", bin, bb)
	}
	return
}

// ImToRgba
func ImToRgba(im image.Image) *image.RGBA {
	bs := im.Bounds()
	n := image.NewRGBA(image.Rect(0, 0, bs.Dx(), bs.Dy()))
	draw.Draw(n, n.Bounds(), im, bs.Min, draw.Src)
	return n
}
func OpenPng(f string) (image.Image, error) {
	ff, err := os.Open(f)
	if err != nil {
		return nil, err
	}
	defer func(ff *os.File) {
		err := ff.Close()
		if err != nil {

		}
	}(ff)
	im, err := png.Decode(ff)
	if err != nil {
		return nil, err
	}
	return im, nil
}
