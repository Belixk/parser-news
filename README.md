# News Parser Bot

A Telegram bot that parses news from TASS and Washington Post RSS feeds.

## Features

- 📰 News parsing from multiple sources
- ⚡ Smart caching system (5-minute cache)
- 🎛️ Interactive menu with buttons
- 🔍 Filter by news source
- 🚀 Fast response times

## Sources

- TASS (Russian news)
- Washington Post (International news)

## Usage

1. Start bot with `/start`
2. Choose news source from menu
3. Get latest news instantly!

## Setup

1. Clone repository
2. Set `TELEGRAM_BOT_TOKEN` in `.env`
3. Run `go run cmd/parser/main.go`
