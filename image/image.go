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
	if !hasSupportedImageSignature(value) {
		return ErrImageNotValid
	}

	if _, _, err := stdimage.DecodeConfig(bytes.NewReader(value)); err != nil {
		return ErrImageNotValid
	}

	return nil
}

func hasSupportedImageSignature(value []byte) bool {
	return hasPNGSignature(value) || hasJPEGSignature(value) || hasGIFSignature(value)
}

func hasPNGSignature(value []byte) bool {
	return len(value) >= 8 &&
		value[0] == 0x89 &&
		value[1] == 'P' &&
		value[2] == 'N' &&
		value[3] == 'G' &&
		value[4] == 0x0D &&
		value[5] == 0x0A &&
		value[6] == 0x1A &&
		value[7] == 0x0A
}

func hasJPEGSignature(value []byte) bool {
	return len(value) >= 3 &&
		value[0] == 0xFF &&
		value[1] == 0xD8 &&
		value[2] == 0xFF
}

func hasGIFSignature(value []byte) bool {
	return len(value) >= 6 &&
		value[0] == 'G' &&
		value[1] == 'I' &&
		value[2] == 'F' &&
		value[3] == '8' &&
		(value[4] == '7' || value[4] == '9') &&
		value[5] == 'a'
}
