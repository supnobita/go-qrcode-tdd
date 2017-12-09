<<<<<<< HEAD
package main

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
	"io/ioutil"
)

func main() {
	fmt.Println("Hello QR code ")
	qrcode := GenerateQRCode("555-2368")
	ioutil.WriteFile("qrcode.png", qrcode, 0644)
}

func GenerateQRCode(code string) []byte {
	img := image.NewNRGBA(image.Rect(0, 0, 21, 21))
	buf := new(bytes.Buffer)
	_ = png.Encode(buf, img)

	return buf.Bytes()
}
=======
//Test
>>>>>>> 090acc5be0a17df2e32af4b3a3604bd38500d516
