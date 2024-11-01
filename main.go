package main

import (
	"github.com/tomas-sesztak/qrcodeGenerator/services"
	"github.com/tomas-sesztak/qrcodeGenerator/utils"
)

func main() {
	qrcInfo := utils.ParseFlags()
	qrc := services.GenerateQRCode(qrcInfo.Url)
	options := services.SetQRCodeOptions(qrcInfo)
	services.WriteQRCode(qrcInfo.File, qrc, options...)
}
