package utils

import (
	"encoding/base64"
	"fmt"
	"msbase/pkg/conf"

	"github.com/logrusorgru/aurora"
	"gopkg.in/gographics/imagick.v2/imagick"
)

// ConvertFile converts extension
/* @Intention: we needed this to rebase pdf files from a provider */
func ConvertFile(inputStr interface{}, watermark conf.Watermark) string {

	message := inputStr.(string)

	mw := imagick.NewMagickWand()
	if err := mw.SetResolution(200, 200); err != nil {
		fmt.Println("Cannot size image ", err)
	}

	defer mw.Destroy()

	base64Text := make([]byte, base64.StdEncoding.DecodedLen(len(message)))
	n, _ := base64.StdEncoding.Decode(base64Text, []byte(message))
	if err := mw.ReadImageBlob(base64Text[:n]); err != nil {
		fmt.Println(aurora.Yellow(err))
	}

	mw.SetFormat("jpg")
	mw.SetImageMatte(false)
	dw := imagick.NewDrawingWand()
	pw := imagick.NewPixelWand()
	defer dw.Destroy()
	defer pw.Destroy()

	pw.SetColor(watermark.Color)
	dw.SetFontSize(watermark.Size)
	dw.SetFillColor(pw)

	mw.CropImage(mw.GetImageWidth(), 2*mw.GetImageHeight()/3, 0, 0)
	mw.AnnotateImage(dw, 20, float64(mw.GetImageHeight()-20), 0, watermark.Text)
	mw.SetCompressionQuality(80)

	convertedImage := mw.GetImageBlob()

	mw.WriteImage("file.jpg")
	return base64.StdEncoding.EncodeToString(convertedImage)
}
