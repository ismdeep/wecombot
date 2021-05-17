package wecombot

// TextMessage Text Message
type TextMessage struct {
	MsgType string `json:"msgtype"`
	Text    struct {
		Content             string   `json:"content"`
		MentionedList       []string `json:"mentioned_list"`
		MentionedMobileList []string `json:"mentioned_mobile_list"`
	} `json:"text"`
}

// MarkdownMessage Markdown Message
type MarkdownMessage struct {
	MsgType  string `json:"msgtype"`
	Markdown struct {
		Content string `json:"content"`
	} `json:"markdown"`
}

// ImageMessage Image Message
type ImageMessage struct {
	MsgType string `json:"msgtype"`
	Image   struct {
		Base64 string `json:"base64"`
		MD5    string `json:"md5"`
	} `json:"image"`
}

// Article Article
type Article struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	PicURL      string `json:"picurl"`
}

// ImageTextMessage Image and Text Message
type ImageTextMessage struct {
	MsgType string `json:"msgtype"`
	News    struct {
		Articles []*Article `json:"articles"`
	} `json:"news"`
}

// RespData Response Data
type RespData struct {
	ErrCode int    `json:"errcode"`
	Msg     string `json:"errmsg"`
}
