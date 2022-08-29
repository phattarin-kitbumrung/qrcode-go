package main

import (
	"image/png"
	"os"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

func main() {
	qrCode, _ := qr.Encode("https://www.codesass.com/courses/golang", qr.M, qr.Auto)
	qrCode, _ = barcode.Scale(qrCode, 200, 200)
	file, _ := os.Create("qr.png")
	defer file.Close()
	png.Encode(file, qrCode)
}