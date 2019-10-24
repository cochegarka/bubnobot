package bubnobot

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var (
	// APIToken represents token to access HTTP Bot API
	APIToken = ""
	// Port ...
	Port = ""
)

const ourChatID = -1001261850458

// BubnoBot is love, BubnoBot is life.
type BubnoBot struct {
	bot      *tgbotapi.BotAPI
	client   *http.Client
	commands map[string]Command

	GreetingMessage string
	GoodbyeMessage  string
	Debug           bool

	defaultCommand DefaultCommand
}

// NewBubnoBot creates
func NewBubnoBot() (*BubnoBot, error) {
	return NewBubnoBotWithClient(http.DefaultClient)
}

// NewBubnoBotWithClient ...
func NewBubnoBotWithClient(client *http.Client) (*BubnoBot, error) {
	bot, err := tgbotapi.NewBotAPIWithClient(APIToken, client)
	if err != nil {
		return &BubnoBot{}, err
	}
	return &BubnoBot{
		bot:      bot,
		client:   client,
		Debug:    true,
		commands: make(map[string]Command),
	}, nil
}

// Start ...
func (b *BubnoBot) Start() error {
	log.Printf("le %s has arrived", b.bot.Self.UserName)

	rand.Seed(time.Now().UnixNano())

	b.bot.Debug = b.Debug
	b.bot.Send(tgbotapi.NewMessage(ourChatID, b.GreetingMessage))

	_, err := b.bot.SetWebhook(tgbotapi.NewWebhook("https://bubnobot.herokuapp.com/" + b.bot.Token))
	if err != nil {
		return err
	}
	info, err := b.bot.GetWebhookInfo()
	if err != nil {
		return err
	}
	if info.LastErrorDate != 0 {
		log.Printf("Telegram callback failed: %s", info.LastErrorMessage)
	}

	updates := b.bot.ListenForWebhook("/" + b.bot.Token)
	go http.ListenAndServe(":"+Port, nil)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		if !b.recognizeChat(update.Message) {
			continue
		}

		if update.Message.IsCommand() {
			command, hasCommand := b.commands[update.Message.Command()]
			if !hasCommand {
				command = b.defaultCommand
			}
			command.call(b.bot, update.Message)
		}
	}

	return nil
}

// Kicks away random gappers
func (b *BubnoBot) recognizeChat(message *tgbotapi.Message) bool {
	if message.Chat.ID != ourChatID {
		b.bot.Send(tgbotapi.NewMessage(message.Chat.ID, "Приходите со своей подгруппой! ✨❓👵❔✨"))
		return false
	}
	return true
}

// AddCommand ...
func (b *BubnoBot) AddCommand(shortcut string, command Command) {
	b.commands[shortcut] = command
}
