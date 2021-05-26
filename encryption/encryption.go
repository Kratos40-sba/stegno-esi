package encryption

import (
	"stegno/utils"
)

type meth string

const (
	None meth = "no-method"
	Xor  meth = "XOR"
	_AES meth = "AES"
)

type Config struct {
	Cle     string
	Methode meth
}

func xorEncDec(s, c string) string {
	return utils.Encdec(s, c)
}
func (con Config) EncDec(s string) string {
	switch con.Methode {
	case Xor:
		return xorEncDec(s, con.Cle)
	case None:
		return s

	}
	return s
}
