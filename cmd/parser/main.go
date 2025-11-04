package main

import (
	"fmt"
	"os"

	"github.com/Belixk/parser-news/internal/bot"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		fmt.Println(".env file not found")
	}

	// Check if Telegram bot token is set
	if os.Getenv("TELEGRAM_BOT_TOKEN") == "" {
		fmt.Println("Error: TELEGRAM_BOT_TOKEN environment variable is not set")
		return
	}

	fmt.Println("Bot token successfully loaded!")
	token := os.Getenv("TELEGRAM_BOT_TOKEN")

	// Initialize and start the bot
	myBot := &bot.Bot{Bot_token: token}
	myBot.Start()
}
