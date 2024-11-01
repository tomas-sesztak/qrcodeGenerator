package services

import (
	"github.com/yeqown/go-qrcode/v2"
)

func GenerateQRCode(link string) qrcode.QRCode {
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
