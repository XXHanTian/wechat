package mass2group

import (
	"github.com/chanxuehong/wechat/mp/core"
)

const (
	MsgTypeText   core.MsgType = "text"
	MsgTypeImage  core.MsgType = "image"
	MsgTypeVoice  core.MsgType = "voice"
	MsgTypeVideo  core.MsgType = "mpvideo"
	MsgTypeNews   core.MsgType = "mpnews"
	MsgTypeWxCard core.MsgType = "wxcard"
)

type MsgHeader struct {
	Filter struct {
		IsToAll bool `json:"is_to_all"`
		TagId   int  `json:"tag_id"`
	} `json:"filter"`
	MsgType     core.MsgType `json:"msgtype"`
	ClientMsgId string       `json:"clientmsgid"`
}

type Text struct {
	MsgHeader
	Text struct {
		Content string `json:"content"`
	} `json:"text"`
}

// 图文消息
type News struct {
	MsgHeader
	News struct {
		MediaId string `json:"media_id"`
	} `json:"mpnews"`
}

// 新建图文消息
//  NOTE: 对于临时素材, mediaId 应该通过 media.UploadNews 得到
func NewNews(tagId int, mediaId string, clientMsgId string) *News {
	var msg News
	msg.MsgType = MsgTypeNews
	msg.Filter.IsToAll = false
	msg.Filter.TagId = tagId
	msg.News.MediaId = mediaId
	msg.ClientMsgId = clientMsgId
	return &msg
}
