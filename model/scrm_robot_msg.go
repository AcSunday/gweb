package model

import (
	"go_project_demo/component"
	"net/url"
)

const (
	scrmRobotMessageApi = `https://qyapi.weixin.qq.com/cgi-bin/webhook/send`
)

type ScrmRobotMessage struct {
	MsgType  string                           `json:"msgtype"`
	Text     *ScrmRobotMessageTextContent     `json:"text,omitempty"`
	Markdown *ScrmRobotMessageMarkdownContent `json:"markdown,omitempty"`
}

type ScrmRobotMessageTextContent struct {
	Content             string   `json:"content"`
	MentionedList       []string `json:"mentioned_list,omitempty"`
	MentionedMobileList []string `json:"mentioned_mobile_list,omitempty"`
}

type ScrmRobotMessageMarkdownContent struct {
	Content string `json:"content"`
}

func SendScrmMessage(key string, msg *ScrmRobotMessage) error {
	query := url.Values{
		"key": []string{key},
	}
	return component.HttpClient.Req("POST", scrmRobotMessageApi, query, msg, nil)
}
