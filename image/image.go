package image

import (
	"bytes"
	"encoding/binary"
	"errors"
	stdimage "image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

const (
	icoHeaderSize = 6
	icoEntrySize  = 16
)

var (
	ErrImageNotValid = errors.New("image is not valid")
)

func IsImage(value []byte) error {
	if !hasSupportedImageSignature(value) {
		return ErrImageNotValid
	}

	if hasICOSignature(value) {
		if !isICOValid(value) {
			return ErrImageNotValid
		}

		return nil
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

func hasICOSignature(value []byte) bool {
	return len(value) >= icoHeaderSize &&
		value[0] == 0x00 &&
		value[1] == 0x00 &&
		value[2] == 0x01 &&
		value[3] == 0x00
}

func isICOValid(value []byte) bool {
	count := int(binary.LittleEndian.Uint16(value[4:6]))
	if count == 0 || count > (len(value)-icoHeaderSize)/icoEntrySize {
		return false
	}

	directorySize := icoHeaderSize + count*icoEntrySize
	for i := 0; i < count; i++ {
		entryStart := icoHeaderSize + i*icoEntrySize
		entry := value[entryStart : entryStart+icoEntrySize]
		if !isICOEntryValid(value, entry, directorySize) {
			return false
		}
	}

	return true
}

func isICOEntryValid(value []byte, entry []byte, directorySize int) bool {
	if entry[3] != 0 {
		return false
	}

	size := binary.LittleEndian.Uint32(entry[8:12])
	offset := binary.LittleEndian.Uint32(entry[12:16])
	if size == 0 || offset < uint32(directorySize) || offset > uint32(len(value)) || size > uint32(len(value))-offset {
		return false
	}

	return isICOImageDataValid(value[offset : offset+size])
}

func isICOImageDataValid(value []byte) bool {
	if hasPNGSignature(value) {
		_, _, err := stdimage.DecodeConfig(bytes.NewReader(value))
		return err == nil
	}

	return isDIBHeaderValid(value)
}

func isDIBHeaderValid(value []byte) bool {
	if len(value) < 40 {
		return false
	}

	headerSize := binary.LittleEndian.Uint32(value[0:4])
	if headerSize < 40 || headerSize > uint32(len(value)) {
		return false
	}

	width := binary.LittleEndian.Uint32(value[4:8])
	height := binary.LittleEndian.Uint32(value[8:12])
	planes := binary.LittleEndian.Uint16(value[12:14])
	bitCount := binary.LittleEndian.Uint16(value[14:16])
	compression := binary.LittleEndian.Uint32(value[16:20])

	return width > 0 &&
		height > 0 &&
		planes == 1 &&
		isDIBBitCountValid(bitCount) &&
		compression <= 6
}

func isDIBBitCountValid(value uint16) bool {
	switch value {
	case 1, 4, 8, 16, 24, 32:
		return true
	default:
		return false
	}
}
