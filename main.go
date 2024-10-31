package main

func main() {
	qrcInfo := parseFlags()
	qrc := generateQRCode(qrcInfo.url)
	options := setQRCodeOptions(qrcInfo)
	writeQRCode(qrcInfo.file, qrc, options...)
}
