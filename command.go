package bubnobot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// Command ...
type Command interface {
	call(*tgbotapi.BotAPI, *tgbotapi.Message)
}

// DefaultCommand ...
type DefaultCommand struct{}

// Call ...
func (d DefaultCommand) call(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Бабушка тебя не понимает 👵")
	msg.ReplyToMessageID = message.MessageID
	bot.Send(msg)
}
