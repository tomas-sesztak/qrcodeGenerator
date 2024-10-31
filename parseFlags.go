package main

import (
	"flag"
	"strings"
)

type qrCodeInfo struct {
	borderWidth        int
	file               string
	logoPath           string
	logoSizeMultiplier uint8
	transparentBg      bool
	url                string
	width              uint8
}

func parseFlags() qrCodeInfo {
	borderWidth := flag.Int("borderWidth", 0, "QR Border Width, default 0")
	file := flag.String("file", "", "Path to final QRCode file(will be .png)")
	logoPath := flag.String("logoPath", "", "Path to logo to be inserted")
	logoSizeMultiplier := flag.Uint("logoSizeMultiplier", 5, "Logo size multiplier, default 5")
	transparentBg := flag.Bool("transparentBg", false, "Bg transparency, default false")
	url := flag.String("url", "", "Url QRCode should point to")
	width := flag.Uint("width", 128, "QR size [0-255]")

	flag.Parse()

	if *width > 255 {
		panic("qrWidth is of type uint8, accepted values [0-225]\n")
	}
	if *logoSizeMultiplier > 255 {
		panic("logoSizeMultiplier is of type uint8, accepted values [0-225]\n")
	}

	if !strings.HasSuffix(*file, ".png") {
		*file += ".png"
	}

	return qrCodeInfo{
		borderWidth:        *borderWidth,
		file:               *file,
		logoPath:           *logoPath,
		logoSizeMultiplier: uint8(*logoSizeMultiplier),
		transparentBg:      *transparentBg,
		url:                *url,
		width:              uint8(*width),
	}
}
