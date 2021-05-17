package wecombot

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
)

func renderImage(data []byte) interface{} {

	md5Ctx := md5.New()
	md5Ctx.Write(data)
	cipherStr := md5Ctx.Sum(nil)

	return &ImageMessage{
		MsgType: "image",
		Image: struct {
			Base64 string `json:"base64"`
			MD5    string `json:"md5"`
		}{
			Base64: base64.StdEncoding.EncodeToString(data),
			MD5:    hex.EncodeToString(cipherStr),
		},
	}
}

// SendImage Send Image
func (receiver *Bot) SendImage(data []byte) error {
	postData := renderImage(data)
	return receiver.send(postData)
}
