package utils

import (
	"flag"
	"strings"

	"github.com/tomas-sesztak/go-utils/argparse"

	"github.com/tomas-sesztak/qrcodeGenerator/models"
)

func ParseFlags() models.QrCodeInfo {
	borderWidth := flag.Int("borderWidth", 0, "QR Border Width, default 0")
	file := flag.String("file", "", "Path to final QRCode file(will be .png)")
	logoPath := flag.String("logoPath", "", "Path to logo to be inserted")
	logoSizeMultiplier := flag.Uint("logoSizeMultiplier", 5, "Logo size multiplier, default 5")
	transparentBg := flag.Bool("transparentBg", false, "Bg transparency, default false")
	url := flag.String("url", "", "Url QRCode should point to")
	width := flag.Uint("width", 128, "QR size [0-255]")

	flag.Parse()
	argparse.ParseArgs(false, true)

	if *width > 255 {
		panic("qrWidth is of type uint8, accepted values [0-225]\n")
	}
	if *logoSizeMultiplier > 255 {
		panic("logoSizeMultiplier is of type uint8, accepted values [0-225]\n")
	}

	if !strings.HasSuffix(*file, ".png") {
		*file += ".png"
	}

	return models.QrCodeInfo{
		BorderWidth:        *borderWidth,
		File:               *file,
		LogoPath:           *logoPath,
		LogoSizeMultiplier: uint8(*logoSizeMultiplier),
		TransparentBg:      *transparentBg,
		Url:                *url,
		Width:              uint8(*width),
	}
}
