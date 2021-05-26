package lsb

import (
	"errors"
	"fmt"
	"image"
	"image/png"
	"os"
	"stegno/encryption"
	"stegno/utils"
	"strconv"
)

var (
	longMessageErr = errors.New("Message trop long")
	noMessageErr   = errors.New("Le message est introuvable")
)

// ajouter un bit a l"octet
func addBite(n, c uint8) (res uint8) {
	i, _ := strconv.ParseInt(string(c), 10, 64)
	res = n&^1 | uint8(i)
	return
}

// encode
func StegnoEncod(m, d string, im image.Image, config *encryption.Config) error {
	m = config.EncDec(m)
	bm := utils.ByToBin([]byte(fmt.Sprintf("%d%s", len(m), m)))
	l := len(bm)
	index := 0
	w, h := im.Bounds().Dx(), im.Bounds().Dy()
	if l > (w * h * 3) {
		return longMessageErr
	}
	rgbIm := utils.ImToRgba(im)
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			col := rgbIm.RGBAAt(x, y)
			if index+1 <= l {
				col.R = addBite(col.R, bm[index])
			}
			if index+2 <= l {
				col.G = addBite(col.G, bm[index+1])

			}
			if index+3 <= l {
				col.B = addBite(col.B, bm[index+2])
			}
			rgbIm.SetRGBA(x, y, col)
			index += 3

		}
	}
	stegoImage, err := os.Create(d)
	if err != nil {
		return err
	}
	defer stegoImage.Close()
	err = png.Encode(stegoImage, rgbIm)
	if err != nil {
		return err
	}
	return nil

}

//
func StegnoDecod(im image.Image, conf *encryption.Config) (message string, err error) {

	w, h := im.Bounds().Dx(), im.Bounds().Dy()
	rgbObj := utils.ImToRgba(im)
	var (
		msgL     int = 0
		msgSl    string
		charBuff int   = 0
		n        uint8 = 0
	)
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			rgbaColorPos := rgbObj.RGBAAt(x, y)
			for _, c := range []uint8{rgbaColorPos.R, rgbaColorPos.G, rgbaColorPos.B} {
				lsb := c & 1
				charBuff += int(lsb << (7 - n))
				if n == 8 {
					b := byte(charBuff)
					if b > 47 && b < 58 && len(message) == 0 {
						msgSl += string(b)
					} else {
						msgL, _ = strconv.Atoi(msgSl)
						if len(message) == msgL {
							goto fin
						}
						message += string(b)

					}
					n, charBuff = 0, 0
				}
				n++

			}
		}
	}
fin:
	if msgL > 0 {
		message = conf.EncDec(message)
	} else {
		message = " "
		err = noMessageErr
	}
	return
}
