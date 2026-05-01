package image_test

import (
	"encoding/binary"
	"errors"
	"os"
	"path/filepath"
	"testing"

	mvimage "github.com/Multiform-Validator/go/image"
)

func TestIsImage(t *testing.T) {
	tests := []struct {
		name    string
		value   []byte
		wantErr error
	}{
		{"valid jpeg", mustReadMock(t, "valid.jpg"), nil},
		{"valid png", mustReadMock(t, "valid.png"), nil},
		{"valid gif", mustReadMock(t, "valid.gif"), nil},
		{"valid ico", mustReadMock(t, "valid.ico"), nil},
		{"valid ico with png payload", icoWithImageData(t, mustReadMock(t, "valid.png")), nil},
		{"unsupported valid avif", mustReadMock(t, "valid.avif"), mvimage.ErrImageNotValid},
		{"unsupported valid psd", mustReadMock(t, "valid.psd"), mvimage.ErrImageNotValid},
		{"unsupported valid svg", mustReadMock(t, "valid.svg"), mvimage.ErrImageNotValid},
		{"unsupported valid icns", mustReadMock(t, "valid.icns"), mvimage.ErrImageNotValid},
		{"invalid png file", mustReadMock(t, "invalid.png"), mvimage.ErrImageNotValid},
		{"invalid png header file", mustReadMock(t, "invalid-valid-header.png"), mvimage.ErrImageNotValid},
		{"invalid empty value", nil, mvimage.ErrImageNotValid},
		{"invalid empty slice", []byte{}, mvimage.ErrImageNotValid},
		{"invalid short bytes", []byte{0xFF, 0xD8}, mvimage.ErrImageNotValid},
		{"invalid text bytes", []byte("hello"), mvimage.ErrImageNotValid},
		{"invalid bmp bytes", []byte("BMfake-bitmap"), mvimage.ErrImageNotValid},
		{"invalid pdf", []byte("%PDF-1.7"), mvimage.ErrImageNotValid},
		{"invalid png signature only", []byte{0x89, 'P', 'N', 'G', 0x0D, 0x0A, 0x1A, 0x0A}, mvimage.ErrImageNotValid},
		{"invalid png with corrupt payload", append([]byte{0x89, 'P', 'N', 'G', 0x0D, 0x0A, 0x1A, 0x0A}, []byte("corrupt")...), mvimage.ErrImageNotValid},
		{"invalid jpeg signature only", []byte{0xFF, 0xD8, 0xFF, 0xE0}, mvimage.ErrImageNotValid},
		{"invalid gif signature only", []byte("GIF89a"), mvimage.ErrImageNotValid},
		{"invalid gif87a signature only", []byte("GIF87a"), mvimage.ErrImageNotValid},
		{"invalid ico empty directory", []byte{0x00, 0x00, 0x01, 0x00, 0x00, 0x00}, mvimage.ErrImageNotValid},
		{"invalid ico missing entry", []byte{0x00, 0x00, 0x01, 0x00, 0x01, 0x00}, mvimage.ErrImageNotValid},
		{"invalid ico reserved entry byte", invalidICOWithEntry(t, 1, 22, []byte{40, 0, 0, 0}), mvimage.ErrImageNotValid},
		{"invalid ico empty image data", invalidICOWithEntry(t, 0, 22, nil), mvimage.ErrImageNotValid},
		{"invalid ico short dib", icoWithImageData(t, []byte{40}), mvimage.ErrImageNotValid},
		{"invalid ico dib header size", icoWithImageData(t, dibHeader(t, 41, 1, 1, 1, 32, 0)), mvimage.ErrImageNotValid},
		{"invalid ico dib bit count", icoWithImageData(t, dibHeader(t, 40, 1, 1, 1, 2, 0)), mvimage.ErrImageNotValid},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := mvimage.IsImage(tt.value)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("IsImage() error = %v, want %v", err, tt.wantErr)
			}
		})
	}
}

func mustReadMock(t *testing.T, name string) []byte {
	t.Helper()

	value, err := os.ReadFile(filepath.Join("mocks", name))
	if err != nil {
		t.Fatalf("os.ReadFile() error = %v", err)
	}

	return value
}

func invalidICOWithEntry(t *testing.T, reserved byte, offset byte, imageData []byte) []byte {
	t.Helper()

	value := []byte{
		0x00, 0x00, 0x01, 0x00, 0x01, 0x00,
		0x10, 0x10, 0x00, reserved,
		0x01, 0x00, 0x20, 0x00,
		byte(len(imageData)), 0x00, 0x00, 0x00,
		offset, 0x00, 0x00, 0x00,
	}

	return append(value, imageData...)
}

func icoWithImageData(t *testing.T, imageData []byte) []byte {
	t.Helper()

	value := []byte{
		0x00, 0x00, 0x01, 0x00, 0x01, 0x00,
		0x10, 0x10, 0x00, 0x00,
		0x01, 0x00, 0x20, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x16, 0x00, 0x00, 0x00,
	}
	binary.LittleEndian.PutUint32(value[14:18], uint32(len(imageData)))

	return append(value, imageData...)
}

func dibHeader(t *testing.T, headerSize, width, height uint32, planes, bitCount uint16, compression uint32) []byte {
	t.Helper()

	value := make([]byte, 40)
	binary.LittleEndian.PutUint32(value[0:4], headerSize)
	binary.LittleEndian.PutUint32(value[4:8], width)
	binary.LittleEndian.PutUint32(value[8:12], height)
	binary.LittleEndian.PutUint16(value[12:14], planes)
	binary.LittleEndian.PutUint16(value[14:16], bitCount)
	binary.LittleEndian.PutUint32(value[16:20], compression)

	return value
}
