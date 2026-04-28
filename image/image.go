package image

import (
	"bytes"
	"errors"
	stdimage "image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

var (
	ErrImageNotValid = errors.New("image is not valid")
)

func IsImage(value []byte) error {
	if len(value) == 0 {
		return ErrImageNotValid
	}

	if _, _, err := stdimage.Decode(bytes.NewReader(value)); err != nil {
		return ErrImageNotValid
	}

	return nil
}
