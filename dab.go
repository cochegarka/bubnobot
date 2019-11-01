package bubnobot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// DabCommand ...
type DabCommand struct {
}

// Call ...
func (d DabCommand) call(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	response := "даб даб даб я пожилой енотик"

	msg := tgbotapi.NewMessage(message.Chat.ID, response)
	bot.Send(msg)
}
