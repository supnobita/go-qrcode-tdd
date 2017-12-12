package main

import (
	"bytes"
	"errors"
	"image/png"
	"testing"
)

type ErrorWriter struct{}

//define writer interface
func (e *ErrorWriter) Write(b []byte) (int, error) {
	return 0, errors.New("Expected error")
}

func TestGenerateQRCodeReturnValue(t *testing.T) {
	buffer := new(bytes.Buffer)
	GenerateQRCode(buffer, "555-5555")
	if buffer.Len() == 0 {
		t.Errorf("No QRCode generated")
	}
}

func TestGenerateQRCodeReturnError(t *testing.T) {
	buffer := new(ErrorWriter)
	err := GenerateQRCode(buffer, "555-5555")
	if err == nil || err.Error() != "Expected error" {
		t.Errorf("Error not propagated correctly, got %v", err)
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
