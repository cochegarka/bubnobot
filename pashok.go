package bubnobot

import (
	"math/rand"
	"sync"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var (
	pashokStickersInstance *pashokStickers
	pashokOnce             sync.Once
)

// PashokCommand ...
type PashokCommand struct {
}

// Call ...
func (p PashokCommand) call(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	sticker := tgbotapi.NewStickerShare(message.Chat.ID, getPashockStickers().getRandomStickerID())
	bot.Send(sticker)
}

// Contains data that require PashokCommand
type pashokStickers struct {
	stickerIDs []string
	count      int
}

//getPashockStickers initializes data once and refers it as global instance of struct
func getPashockStickers() *pashokStickers {
	pashokOnce.Do(func() {
		pashokStickersInstance = &pashokStickers{}
		pashokStickersInstance.stickerIDs = []string{
			"CAADAgADAQADp0NtG7p38-IhV1ypFgQ",
			"CAADAgADAgADp0NtGxaujZJCnJMIFgQ",
			"CAADAgADAwADp0NtGw6unuDhSreLFgQ",
			"CAADAgADBAADp0NtGxAGyknDgoQnFgQ",
			"CAADAgADBQADp0NtG0SWMxudGtsIFgQ",
			"CAADAgADBgADp0NtG1lfMZDdiQqEFgQ",
			"CAADAgADBwADp0NtG_xivxgs-xlWFgQ",
		}
		pashokStickersInstance.count = len(pashokStickersInstance.stickerIDs)
	})
	return pashokStickersInstance
}

// getRandomSticker returns random sticker ID
func (ps *pashokStickers) getRandomStickerID() string {
	return ps.stickerIDs[rand.Intn(ps.count)]
}
