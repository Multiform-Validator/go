package image_test

import (
	"bytes"
	"errors"
	stdimage "image"
	"image/color"
	"image/gif"
	"image/jpeg"
	"image/png"
	"testing"

	mvimage "github.com/Multiform-Validator/go/image"
)

func TestIsImage(t *testing.T) {
	tests := []struct {
		name    string
		value   []byte
		wantErr error
	}{
		{"valid jpeg", mustEncodeJPEG(t), nil},
		{"valid png", mustEncodePNG(t), nil},
		{"valid gif", mustEncodeGIF(t), nil},
		{"invalid empty value", nil, mvimage.ErrImageNotValid},
		{"invalid short bytes", []byte{0xFF, 0xD8}, mvimage.ErrImageNotValid},
		{"invalid text bytes", []byte("hello"), mvimage.ErrImageNotValid},
		{"invalid pdf", []byte("%PDF-1.7"), mvimage.ErrImageNotValid},
		{"invalid png signature only", []byte{0x89, 'P', 'N', 'G', 0x0D, 0x0A, 0x1A, 0x0A}, mvimage.ErrImageNotValid},
		{"invalid jpeg signature only", []byte{0xFF, 0xD8, 0xFF, 0xE0}, mvimage.ErrImageNotValid},
		{"invalid gif signature only", []byte("GIF89a"), mvimage.ErrImageNotValid},
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

func mustEncodePNG(t *testing.T) []byte {
	t.Helper()

	var buffer bytes.Buffer
	if err := png.Encode(&buffer, testImage()); err != nil {
		t.Fatalf("png.Encode() error = %v", err)
	}

	return buffer.Bytes()
}

func mustEncodeJPEG(t *testing.T) []byte {
	t.Helper()

	var buffer bytes.Buffer
	if err := jpeg.Encode(&buffer, testImage(), nil); err != nil {
		t.Fatalf("jpeg.Encode() error = %v", err)
	}

	return buffer.Bytes()
}

func mustEncodeGIF(t *testing.T) []byte {
	t.Helper()

	var buffer bytes.Buffer
	if err := gif.Encode(&buffer, testImage(), nil); err != nil {
		t.Fatalf("gif.Encode() error = %v", err)
	}

	return buffer.Bytes()
}

func testImage() stdimage.Image {
	img := stdimage.NewRGBA(stdimage.Rect(0, 0, 1, 1))
	img.Set(0, 0, color.RGBA{R: 255, A: 255})
	return img
}
