package main

import (
	"log"
	"stegno/encryption"
	"stegno/lsb"
	"stegno/utils"
)

const (
	KEY     = "123456789"
	MESSAGE = "esi-sba 2021"
	OUTFILE = "steg.png"
	INFILE  = "f.png"
)

func main() {
	im, err := utils.OpenPng(INFILE)
	if err != nil {
		log.Fatalln(err)
	}
	// encode message and image
	err = lsb.StegnoEncod(MESSAGE, OUTFILE, im, &encryption.Config{
		Methode: encryption.Xor,
		Cle:     KEY,
	})
	if err != nil {
		log.Fatalln(err)
	}
	imm, err := utils.OpenPng(OUTFILE)
	if err != nil {
		log.Fatalln(err)
	}
	// decode to message
	message, err := lsb.StegnoDecod(imm, &encryption.Config{
		Methode: encryption.Xor,
		Cle:     KEY,
	})
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf(" Message => %s", message)
}
