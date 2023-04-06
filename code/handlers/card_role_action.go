package handlers

import (
	"context"

	larkcard "github.com/larksuite/oapi-sdk-go/v3/card"

	"github.com/narasux/chatgpt-bot/initialization"
	"github.com/narasux/chatgpt-bot/services"
	"github.com/narasux/chatgpt-bot/services/openai"
)

func NewRoleTagCardHandler(cardMsg CardMsg, m MessageHandler) CardHandlerFunc {
	return func(ctx context.Context, cardAction *larkcard.CardAction) (interface{}, error) {
		if cardMsg.Kind == RoleTagsChooseKind {
			newCard, err, done := CommonProcessRoleTag(cardMsg, cardAction,
				m.sessionCache)
			if done {
				return newCard, err
			}
			return nil, nil
		}
		return nil, ErrNextHandler
	}
}

func NewRoleCardHandler(cardMsg CardMsg, m MessageHandler) CardHandlerFunc {
	return func(ctx context.Context, cardAction *larkcard.CardAction) (interface{}, error) {
		if cardMsg.Kind == RoleChooseKind {
			newCard, err, done := CommonProcessRole(cardMsg, cardAction,
				m.sessionCache)
			if done {
				return newCard, err
			}
			return nil, nil
		}
		return nil, ErrNextHandler
	}
}

func CommonProcessRoleTag(
	msg CardMsg, cardAction *larkcard.CardAction, cache services.SessionServiceCacheInterface,
) (interface{}, error, bool) {
	option := cardAction.Action.Option
	roles := initialization.GetTitleListByTag(option)
	SendRoleListCard(context.Background(), &msg.SessionId,
		&msg.MsgId, option, *roles)
	return nil, nil, true
}

func CommonProcessRole(
	msg CardMsg, cardAction *larkcard.CardAction, cache services.SessionServiceCacheInterface,
) (interface{}, error, bool) {
	option := cardAction.Action.Option
	contentByTitle, err := initialization.GetFirstRoleContentByTitle(option)
	if err != nil {
		return nil, err, true
	}
	cache.Clear(msg.SessionId)
	systemMsg := append([]openai.Messages{}, openai.Messages{
		Role: "system", Content: contentByTitle,
	})
	cache.SetMsg(msg.SessionId, systemMsg)
	sendSystemInstructionCard(context.Background(), &msg.SessionId, &msg.MsgId, contentByTitle)
	return nil, nil, true
}
