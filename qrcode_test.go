package main

import (
	"bytes"
	"image/png"
	"testing"
)

func TestGenerateQRCodeReturnValue(t *testing.T) {
	result := GenerateQRCode("555-555")
	if result == nil {
		t.Errorf("Generate QRCode return Nil")
	}
	if len(result) == 0 {
		t.Errorf("Generate QRCode return 0 byte")
	}
}

func TestFunctionGenerateValidPNG(t *testing.T) {
	result := GenerateQRCode("555-555")
	buffer := bytes.NewBuffer(result)
	_, err := png.Decode(buffer)
	if err != nil {
		t.Error("GenerateQRCode generate invalid PNG")
	}
}
