package wecombot

func renderText(str string, options ...Option) *TextMessage {
	postData := &TextMessage{}
	postData.MsgType = "text"
	postData.Text.Content = str
	postData.Text.MentionedList = []string{}
	postData.Text.MentionedMobileList = []string{}

	for _, opt := range options {
		for _, mention := range opt.MentionedList {
			postData.Text.MentionedList = append(postData.Text.MentionedList, mention)
		}
		for _, mention := range opt.MentionedMobileList {
			postData.Text.MentionedMobileList = append(postData.Text.MentionedMobileList, mention)
		}
	}

	return postData
}

// SendText Send text message
func (receiver *Bot) SendText(str string, options ...Option) error {
	postData := renderText(str, options...)
	return receiver.send(postData)
}
