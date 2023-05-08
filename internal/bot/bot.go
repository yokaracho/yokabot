package bot

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type Bot struct {
	Token string
	Bot   *tgbotapi.BotAPI
}
