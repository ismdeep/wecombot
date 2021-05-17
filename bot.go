package wecombot

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Bot WeCom Bot
type Bot struct {
	Key string
}

// New New WeCom Bot
func New(key string) *Bot {
	return &Bot{Key: key}
}

func (receiver *Bot) send(postData interface{}) error {
	httpClient := &http.Client{}
	data, err := json.Marshal(postData)
	if err != nil {
		return err
	}

	resp, err := httpClient.Post(fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=%v", receiver.Key), "application/json", bytes.NewReader(data))
	if err != nil {
		return err
	}

	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	respData := &RespData{}
	if err := json.Unmarshal(all, respData); err != nil {
		return err
	}

	if respData.ErrCode != 0 {
		return errors.New(respData.Msg)
	}

	return nil
}

// Option Option Struct
type Option struct {
	MentionedList       []string
	MentionedMobileList []string
	Articles            []*Article
}

// WithMentionedList With Mentioned List
func WithMentionedList(mentionedList []string) Option {
	return Option{
		MentionedList: mentionedList,
	}
}

// WithMentionAll With Mention All
func WithMentionAll() Option {
	return Option{
		MentionedList: []string{"@all"},
	}
}

// WithMentionedMobileList WithMentioned Mobile List
func WithMentionedMobileList(mentionedMobileList []string) Option {
	return Option{
		MentionedMobileList: mentionedMobileList,
	}
}

// WithMentionAllMobile With Mention All Mobile
func WithMentionAllMobile() Option {
	return Option{
		MentionedMobileList: []string{"@all"},
	}
}

// WithAddArticle With Add Single Article
func WithAddArticle(article *Article) Option {
	return Option{
		Articles: []*Article{article},
	}
}

// WithAddArticles With Article List
func WithAddArticles(articles []*Article) Option {
	return Option{
		Articles: articles,
	}
}
