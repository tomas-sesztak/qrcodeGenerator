package main

import (
	"github.com/yeqown/go-qrcode/writer/standard"
	"os"
)

func setQRCodeOptions(qrcInfo qrCodeInfo) []standard.ImageOption {

	imageOptions := []standard.ImageOption{
		standard.WithBuiltinImageEncoder(standard.PNG_FORMAT),
		standard.WithQRWidth(qrcInfo.width),
		standard.WithLogoSizeMultiplier(int(qrcInfo.logoSizeMultiplier)),
		standard.WithBorderWidth(qrcInfo.borderWidth),
	}
	if qrcInfo.logoPath != "" {
		if _, err := os.Stat(qrcInfo.logoPath); err != nil {
			panic(err)
		}
		imageOptions = append(
			imageOptions,
			standard.WithLogoImageFilePNG(qrcInfo.logoPath),
		)
	}
	if qrcInfo.transparentBg {
		imageOptions = append(
			imageOptions,
			standard.WithBgTransparent(),
		)
	}

	return imageOptions
}
