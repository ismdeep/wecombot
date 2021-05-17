package wecombot

func renderMarkdown(str string) *MarkdownMessage {
	postData := &MarkdownMessage{
		MsgType: "markdown",
		Markdown: struct {
			Content string `json:"content"`
		}{
			Content: str,
		},
	}
	return postData
}

// SendMarkdown Send markdown message
func (receiver *Bot) SendMarkdown(str string, options ...Option) error {
	postData := renderMarkdown(str)
	return receiver.send(postData)
}
