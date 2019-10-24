package main

import (
	"log"
	"os"

	bubnobot "bubnobot_heroku"
)

func main() {
	var (
		port  = os.Getenv("PORT")
		token = os.Getenv("TOKEN")
	)

	bubnobot.APIToken = token
	bubnobot.Port = port

	if port == "" {
		bubnobot.Port = "8000"
	}

	bot, err := bubnobot.NewBubnoBot()
	if err != nil {
		log.Fatal(err)
	}

	bot.GreetingMessage = "Всем привет, я только что проснулась! 👵"
	bot.GoodbyeMessage = "Бабушка пошла спать... 😴"

	bot.AddCommand("pashok", bubnobot.PashokCommand{})
	bot.AddCommand("ask", bubnobot.AskCommand{})

	if err = bot.Start(); err != nil {
		log.Fatal(err)
	}
}
