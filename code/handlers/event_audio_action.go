package handlers

import (
	"context"
	"fmt"
	"os"

	larkim "github.com/larksuite/oapi-sdk-go/v3/service/im/v1"

	"github.com/narasux/chatgpt-bot/initialization"
	"github.com/narasux/chatgpt-bot/utils/audio"
)

type AudioAction struct {
}

func (*AudioAction) Execute(a *ActionInfo) bool {
	// 只有私聊才解析语音,其他不解析
	if a.info.handlerType != UserHandler {
		return true
	}

	// 判断是否是语音
	if a.info.msgType == "audio" {
		fileKey := a.info.fileKey
		msgId := a.info.msgId
		req := larkim.NewGetMessageResourceReqBuilder().MessageId(
			*msgId).FileKey(fileKey).Type("file").Build()
		resp, err := initialization.GetLarkClient().Im.MessageResource.Get(context.Background(), req)
		if err != nil {
			fmt.Println(err)
			return true
		}
		f := fmt.Sprintf("%s.ogg", fileKey)
		resp.WriteFile(f)
		defer os.Remove(f)

		output := fmt.Sprintf("%s.mp3", fileKey)
		// 等待转换完成
		audio.OggToWavByPath(f, output)
		defer os.Remove(output)

		text, err := a.handler.gpt.AudioToText(output)
		if err != nil {
			fmt.Println(err)

			sendMsg(*a.ctx, fmt.Sprintf("🤖️：语音转换失败，请稍后再试～\n错误信息: %v", err), a.info.msgId)
			return false
		}

		replyMsg(*a.ctx, fmt.Sprintf("🤖️：%s", text), a.info.msgId)
		a.info.qParsed = text
		return true
	}

	return true
}
