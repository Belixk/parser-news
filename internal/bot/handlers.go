package bot

import (
	"fmt"
	"log"

	"github.com/Belixk/parser-news/internal/parser"
	tele "gopkg.in/telebot.v3"
)

// CreateMenu generates the reply keyboard with news source options
func CreateMenu() *tele.ReplyMarkup {
	menu := &tele.ReplyMarkup{ResizeKeyboard: true}

	// Create menu buttons for different news sources
	btnAll := menu.Text("🌎 All News")
	btnTASS := menu.Text("🇷🇺 Russia News")
	btnWashington := menu.Text("🇺🇸 USA News")

	// Arrange buttons vertically in separate rows
	menu.Reply(menu.Row(btnAll), menu.Row(btnTASS), menu.Row(btnWashington))
	return menu
}

// HandleAll processes request for all news from all sources
func HandleAll(c tele.Context) error {
	newsAll, err := parser.ParseAll()
	if err != nil {
		log.Printf("Parsing error: %v", err)
		return c.Send("Error loading news. Please try again later.")
	}
	return c.Send(FormatText(newsAll), tele.ModeHTML)
}

// HandleTASS processes request for TASS (Russian) news only
func HandleTASS(c tele.Context) error {
	newsTass, err := parser.ParseBySource("TASS")
	if err != nil {
		log.Printf("TASS parsing error: %v", err)
		return c.Send("Error loading TASS news. Please try again later.")
	}
	return c.Send(FormatText(newsTass), tele.ModeHTML)
}

// HandleWashington processes request for Washington Post (USA) news only
func HandleWashington(c tele.Context) error {
	newsWashington, err := parser.ParseBySource("Washington")
	if err != nil {
		log.Printf("Washington Post parsing error: %v", err)
		return c.Send("Error loading Washington Post news. Please try again later.")
	}
	return c.Send(FormatText(newsWashington), tele.ModeHTML)
}

// FormatText converts news slice into formatted string for Telegram
func FormatText(news []parser.News) string {
	var result string
	for _, item := range news {
		result += fmt.Sprintf("📰 %s\n🔗 %s\n\n", item.Title, item.Link)
	}
	return result
}
