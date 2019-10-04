package main

import (
	"log"
	"net/http"
	"net/url"

	"bitbucket.org/MegaDeathLightsaber/bubnobot"
)

const proxyStr = "yourproxy"

func main() {
	// Bypass the blocking of Telegram
	proxyURL, err := url.Parse(proxyStr)
	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyURL)}}

	bot, err := bubnobot.NewBubnoBotWithClient(client)
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
