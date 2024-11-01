package services

import (
	"github.com/yeqown/go-qrcode/writer/standard"
	"os"
	"github.com/tomas-sesztak/qrcodeGenerator/models"
)

func SetQRCodeOptions(qrcInfo models.QrCodeInfo) []standard.ImageOption {

	imageOptions := []standard.ImageOption{
		standard.WithBuiltinImageEncoder(standard.PNG_FORMAT),
		standard.WithQRWidth(qrcInfo.Width),
		standard.WithLogoSizeMultiplier(int(qrcInfo.LogoSizeMultiplier)),
		standard.WithBorderWidth(qrcInfo.BorderWidth),
	}
	if qrcInfo.LogoPath != "" {
		if _, err := os.Stat(qrcInfo.LogoPath); err != nil {
			panic(err)
		}
		imageOptions = append(
			imageOptions,
			standard.WithLogoImageFilePNG(qrcInfo.LogoPath),
		)
	}
	if qrcInfo.TransparentBg {
		imageOptions = append(
			imageOptions,
			standard.WithBgTransparent(),
		)
	}

	return imageOptions
}
