package bot

import (
	"log"
	"time"

	tele "gopkg.in/telebot.v3"
)

type Bot struct {
	Bot_token string // Telegram Bot API token
}

func (t *Bot) Start() {
	b, err := tele.NewBot(tele.Settings{ // Initialize Telegram bot with settings
		Token:  t.Bot_token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Minute},
	})
	if err != nil {
		log.Fatal(err)
	}

	menu := CreateMenu()

	// Register handlers for menu buttons
	b.Handle("🌎 All News", HandleAll)
	b.Handle("🇷🇺 Russia News", HandleTASS)
	b.Handle("🇺🇸 USA News", HandleWashington)

	// Start command handler - displays welcome message and menu
	b.Handle("/start", func(c tele.Context) error {
		return c.Send("Welcome! I'm a news bot that parses articles from various websites. Please choose one of the options below to continue.", menu)
	})

	b.Start() // Start bot polling
}
