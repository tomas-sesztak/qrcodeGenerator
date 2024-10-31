package main

import (
	"github.com/yeqown/go-qrcode/v2"
	"github.com/yeqown/go-qrcode/writer/standard"
)

func writeQRCode(filename string, qrc qrcode.QRCode, options ...standard.ImageOption) {
	img, err := standard.New(filename, options...)
	if err != nil {
		panic(err)
	}
	if err = qrc.Save(img); err != nil {
		panic(err)
	}
}
