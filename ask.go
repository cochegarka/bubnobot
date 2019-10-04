package bubnobot

import (
	"math/rand"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// AskCommand ...
type AskCommand struct {
}

// Call ...
func (a AskCommand) call(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	var response string
	if !strings.HasSuffix(message.Text, "?") {
		response = "Я не поняла ваш вопрос ❓👵❔"
	} else {
		if rand.Intn(2) == 0 {
			response = "Да!"
		} else {
			response = "А давайте посмотрим..."
		}
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, response)
	msg.ReplyToMessageID = message.MessageID
	bot.Send(msg)
}
