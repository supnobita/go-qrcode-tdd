package main

import (
	"bytes"
	"image/png"
	"testing"
)

func TestGenerateQRCodeReturnValue(t *testing.T) {
	buffer := new(bytes.Buffer)
	GenerateQRCode(buffer, "555-5555")
	if buffer.Len() == 0 {
		t.Errorf("No QRCode generated")
	}
}

func TestFunctionGenerateValidPNG(t *testing.T) {
	buffer := new(bytes.Buffer)
	GenerateQRCode(buffer, "555-5555")
	//buffer := bytes.NewBuffer(result)
	_, err := png.Decode(buffer)
	if err != nil {
		t.Errorf("GenerateQRCode generate invalid PNG: %s", err)
	}
}
