package model

import "github.com/juju/errors"

const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

var ErrCharIndexOutOfRange = errors.New("char index is out of range")

func getCharByIndex(i int64) (string, error) {
	if i < 1 || i > 26 {
		return "", ErrCharIndexOutOfRange
	}

	//a := math.Mod(float64(i), float64(26))
	//i = int64(a)
	//if i == 0 {
	//	i = 26
	//}

	return alphabet[i-1 : i], nil
}
