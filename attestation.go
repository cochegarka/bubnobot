package bubnobot

import (
	"fmt"
	"math/rand"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const maxMark = 101

// AttestationCommand ...
type AttestationCommand struct {
}

// Call ...
func (d AttestationCommand) call(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	response := fmt.Sprintf("%d баллов", rand.Intn(maxMark))

	msg := tgbotapi.NewMessage(message.Chat.ID, response)
	bot.Send(msg)
}
