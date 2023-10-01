package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type UsersMessage struct {
	NextMessages []string
	ChatId       int64
	UserMessages []string
	Command      string
}

func addNextMessages(message string, user *UsersMessage, chatId int64) {
	msg := tgbotapi.NewMessage(chatId, message)
	user.NextMessages = append(user.NextMessages, msg.Text)

}

func printNewMessage(nextMessages []string) (string, []string) {
	var msg string
	msg, nextMessages = popElemFront(nextMessages)
	return msg, nextMessages
}
