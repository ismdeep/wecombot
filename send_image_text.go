package wecombot

func renderImageText(options ...Option) *ImageTextMessage {
	postData := &ImageTextMessage{
		MsgType: "news",
		News: struct {
			Articles []*Article `json:"articles"`
		}{},
	}

	for _, option := range options {
		for _, article := range option.Articles {
			postData.News.Articles = append(postData.News.Articles, article)
		}
	}

	return postData
}

// SendImageText Send Image with Text
func (receiver *Bot) SendImageText(options ...Option) error {
	postData := renderImageText(options...)
	return receiver.send(postData)
}
