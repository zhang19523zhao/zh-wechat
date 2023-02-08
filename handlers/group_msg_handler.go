package handlers

import (
	"github.com/eatmoreapple/openwechat"
	"github.com/zhang19523zhao/zh-wechat/config"
	"github.com/zhang19523zhao/zh-wechat/gtp"
	"log"
	"strings"
)

var _ MessageHandlerInterface = (*GroupMessageHandler)(nil)

// GroupMessageHandler 群消息处理
type GroupMessageHandler struct {
}

// handle 处理消息
func (g *GroupMessageHandler) handle(msg *openwechat.Message) error {
	if msg.IsText() {
		return g.ReplyText(msg)
	}
	return nil
}

// NewGroupMessageHandler 创建群消息处理器
func NewGroupMessageHandler() MessageHandlerInterface {
	return &GroupMessageHandler{}
}

// ReplyText 发送文本消息到群
func (g *GroupMessageHandler) ReplyText(msg *openwechat.Message) error {
	// 接收群消息
	sender, err := msg.Sender()
	group := openwechat.Group{sender}
	log.Printf("Received Group %v Text Msg : %v", group.NickName, msg.Content)

	if !msg.IsAt() {
		return nil
	}

	replaceTextx := "@" + sender.Self.NickName
	requestTextx := strings.TrimSpace(strings.ReplaceAll(msg.Content, replaceTextx, ""))

	// 判断是否有关键字
	keyword := config.LoadConfig().KeyWord
	if !strings.HasPrefix(requestTextx, keyword) {
		return nil
	}

	// 替换掉@文本，然后向GPT发起请求
	replaceText := "@" + sender.Self.NickName
	requestText := strings.TrimSpace(strings.ReplaceAll(msg.Content, replaceText, ""))
	requestText = strings.TrimPrefix(requestText, "zh")

	reply := ""
	if requestText != "" {
		reply, err = gtp.Completions(requestText)
		if err != nil {
			log.Printf("gtp request error: %v \n", err)
			msg.ReplyText("机器人神了，我一会发现了就去修。")
			return err
		}
		if reply == "" {
			return nil
		}

		// 获取@我的用户
		groupSender, err := msg.SenderInGroup()
		if err != nil {
			log.Printf("get sender in group error :%v \n", err)
			return err
		}

		// 回复@我的用户
		reply = strings.TrimSpace(reply)
		reply = strings.Trim(reply, "\n")
		atText := "@" + groupSender.NickName
		replyText := atText + reply
		_, err = msg.ReplyText(replyText)
		if err != nil {
			log.Printf("response group error: %v \n", err)
		}
	}

	return err
}
