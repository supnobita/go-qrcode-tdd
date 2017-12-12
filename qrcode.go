package main

import (
	"fmt"
	"image"
	"image/png"
	"io"
	"os"
)

//main function
func main() {
	fmt.Println("Hello QR code ")
	file, _ := os.Create("qrcode.png")
	defer file.Close()

	GenerateQRCode(file, "555-2368")
	//ioutil.WriteFile("qrcode.png", qrcode, 0644)
}

//function gen QR Code to byte array
func GenerateQRCode(w io.Writer, code string) error {
	img := image.NewNRGBA(image.Rect(0, 0, 21, 21))
	//buf := new(bytes.Buffer)
	png.Encode(w, img)
	return nil
}
