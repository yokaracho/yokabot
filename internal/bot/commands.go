package bot

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

// NewBot create new instance Bot.
func NewBot() *Bot {
	return &Bot{
		Token: "TOKEN",
	}
}

// start bot
func (b *Bot) Start() error {
	// create new exp bot
	bot, err := tgbotapi.NewBotAPI(b.Token)
	if err != nil {
		return fmt.Errorf("failed to create new bot instance: %v", err)
	}
	b.Bot = bot

	b.Bot.Debug = true

	// Start listening for updates with a long poll
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	// Processing incoming updates
	for update := range updates {
		if update.Message != nil {
			if update.Message.IsCommand() {
				switch update.Message.Command() {
				case "start":
					b.handleMessage(update.Message)
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
					_, err := bot.Send(msg)
					if err != nil {
						log.Println(err)
					}
				case "about":
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Welcome to Telegram Bot! The bot has four buttons on the first level and additional buttons on the second level. ")
					_, err := bot.Send(msg)
					if err != nil {
						log.Println(err)
					}
				case "help":
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "If you have any questions or suggestions, you can contact me https://t.me/trusvladimir ")
					_, err := bot.Send(msg)
					if err != nil {
						log.Println(err)
					}
				default:
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Sorry, I don't understand this command.")
					_, err := bot.Send(msg)
					if err != nil {
						log.Println(err)
					}
				}
			} else {
				// non-command messages
			}
		} else if update.CallbackQuery != nil {
			b.handleCall(update)
		}
	}

	return nil
}
