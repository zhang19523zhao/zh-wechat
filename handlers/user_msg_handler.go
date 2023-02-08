package handlers

import (
	"github.com/eatmoreapple/openwechat"
	"github.com/zhang19523zhao/zh-wechat/config"
	"github.com/zhang19523zhao/zh-wechat/gtp"
	"log"
	"strings"
)

var _ MessageHandlerInterface = (*UserMessageHandler)(nil)

// UserMessageHandler 私聊消息处理
type UserMessageHandler struct {
}

// handle 处理消息
func (g *UserMessageHandler) handle(msg *openwechat.Message) error {
	if msg.IsText() {
		return g.ReplyText(msg)
	}
	return nil
}

// NewUserMessageHandler 创建私聊处理器
func NewUserMessageHandler() MessageHandlerInterface {
	return &UserMessageHandler{}
}

// ReplyText 发送文本消息到群
func (g *UserMessageHandler) ReplyText(msg *openwechat.Message) error {
	// 接收私聊消息
	sender, err := msg.Sender()
	log.Printf("Received User %v Text Msg : %v", sender.NickName, msg.Content)

	requestTextx := strings.TrimSpace(msg.Content)
	requestTextx = strings.Trim(msg.Content, "\n")

	//if requestTextx[:2] != "zh" {
	//	return nil
	//}

	keyword := config.LoadConfig().KeyWord
	if !strings.HasPrefix(requestTextx, keyword) {
		return nil
	}

	// 不处理自己发给自己的消息
	if sender.Self.NickName == sender.NickName {
		return nil
	}

	// 向GPT发起请求
	requestText := strings.TrimSpace(msg.Content)
	requestText = strings.Trim(msg.Content, "\n")
	reply, err := gtp.Completions(requestText)
	if err != nil {
		log.Printf("gtp request error: %v \n", err)
		msg.ReplyText("机器人神了，我一会发现了就去修。")
		return err
	}
	if reply == "" {
		return nil
	}

	// 回复用户
	reply = strings.TrimSpace(reply)
	reply = strings.Trim(reply, "\n")
	_, err = msg.ReplyText(reply)
	if err != nil {
		log.Printf("response user error: %v \n", err)
	}
	return err
}