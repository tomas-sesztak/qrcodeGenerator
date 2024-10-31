package main

import (
	"github.com/yeqown/go-qrcode/v2"
)

func generateQRCode(link string) qrcode.QRCode {
	qrc, err := qrcode.NewWith(
		link,
		qrcode.WithEncodingMode(qrcode.EncModeByte),
		qrcode.WithErrorCorrectionLevel(qrcode.ErrorCorrectionQuart),
	)
	if err != nil {
		panic(err)
	}
	return *qrc
}
