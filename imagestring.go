package imgstr

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/png"
	"strings"
)

// ImageString returns frame serialized image in “IMAGE:<base64-encoded-png>” format.
//
// The usefulness of this serialized format is, if you just output that on
// the Go Playground — https://play.golang.org/ — then it will display it
// as an image.
func ImageString(img image.Image) string {
	var buffer strings.Builder

	buffer.WriteString("IMAGE:")

	{
		var pngBuffer bytes.Buffer

		err := png.Encode(&pngBuffer, img)
		if nil != err {
			return fmt.Sprintf("ERROR:%s", err)
		}

		encoded := base64.StdEncoding.EncodeToString(pngBuffer.Bytes())

		buffer.WriteString(encoded)
	}

	return buffer.String()
}
