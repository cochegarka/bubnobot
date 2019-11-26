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
	mark := rand.Intn(maxMark)
	declension := "баллов"

	if mark < 10 || mark > 20 {
		if mark%10 >= 2 && mark%10 <= 4 {
			declension = "балла"
		} else if mark%10 == 1 {
			declension = "балл"
		}
	}

	response := fmt.Sprintf("%d %s", mark, declension)

	msg := tgbotapi.NewMessage(message.Chat.ID, response)
	bot.Send(msg)
}
