package main

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
	"io/ioutil"
)

//main function
func main() {
	fmt.Println("Hello QR code ")
	qrcode := GenerateQRCode("555-2368")
	ioutil.WriteFile("qrcode.png", qrcode, 0644)
}

//function gen QR Code to byte array
func GenerateQRCode(code string) []byte {
	img := image.NewNRGBA(image.Rect(0, 0, 21, 21))
	buf := new(bytes.Buffer)
	_ = png.Encode(buf, img)

	return buf.Bytes()
}
