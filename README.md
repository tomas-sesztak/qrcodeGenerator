# QR Code generator
## Purpose
This is a generic QR code generator I have written in GO. The purpose of this project is learning GO and generating QR codes for personal use.
## Usage
```bash
go run . [options]
```
### Required options
```
-file string
	Path to final QRCode file(will be .png)
-url string
	Url QRCode should point to
```
### Optional options
```
-borderWidth int
	QR Border Width, default 0
-logoPath string
	Path to logo to be inserted
-logoSizeMultiplier uint
	Logo size multiplier, default 5 (default 5)
-transparentBg
	Bg transparency, default false
-width uint
	QR size [0-255] (default 128)
```
## Example
```bash
go run . -logoPath mylogo.png -url mysite.com -file final_qr.png
```
