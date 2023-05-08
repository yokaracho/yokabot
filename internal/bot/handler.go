package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func (b *Bot) handleMessage(message *tgbotapi.Message) {

	msg := tgbotapi.NewMessage(message.Chat.ID, "Welcome to my bot! Please select an option:")
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Button️ 1 💎️", "1"),
			tgbotapi.NewInlineKeyboardButtonData("Button 2 💎️", "2"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Button 3 💎 (3-1, 3-2)", "3"),
			tgbotapi.NewInlineKeyboardButtonData("Button 4 💎 (4-1, 4-2)", "4"),
		),
	)
	if _, err := b.Bot.Send(msg); err != nil {
		log.Printf("Error sending message to chat - %v", err)
		return
	}
}

func (b *Bot) handleCall(updates tgbotapi.Update) {

	button := updates.CallbackQuery.Data

	// Handle first layer button presses
	switch button {
	case "1":
		// Respond to button 1 press
		handleButton1(updates, b.Bot)
	case "2":
		// Respond to button 2 press
		handleButton2(updates, b.Bot)

	case "3":
		// Show second layer of buttons
		handleButton3(updates, b.Bot)

	case "4":
		// Show second layer of buttons
		handleButton4(updates, b.Bot)

	case "back":
		// Show first layer of buttons again
		handleButtonBack(updates, b.Bot)

	default:
		// Handle second layer button presses
		if button == "3-1" {
			// Respond to button 3-1 press
			msg := tgbotapi.NewMessage(updates.CallbackQuery.Message.Chat.ID, "You pressed button 3-1!")
			_, err := b.Bot.Send(msg)
			if err != nil {
				log.Printf("Error sending message to chat - %v", err)
				return
			}
		} else if button == "3-2" {
			// Respond to button 3-2 press
			msg := tgbotapi.NewMessage(updates.CallbackQuery.Message.Chat.ID, "You pressed button 3-2!")
			_, err := b.Bot.Send(msg)
			if err != nil {
				log.Printf("Error sending message to chat - %v", err)
				return
			}
		} else if button == "4-1" {
			// Respond to button 4-1 press
			msg := tgbotapi.NewMessage(updates.CallbackQuery.Message.Chat.ID, "You pressed button 4-1!")
			_, err := b.Bot.Send(msg)
			if err != nil {
				log.Printf("Error sending message to chat - %v", err)
				return
			}
		} else if button == "4-2" {
			// Respond to button 4-2 press
			msg := tgbotapi.NewMessage(updates.CallbackQuery.Message.Chat.ID, "You pressed button 4-2!")
			_, err := b.Bot.Send(msg)
			if err != nil {
				log.Printf("Error sending message to chat - %v", err)
				return
			}
		}
	}

}

func handleButton1(update tgbotapi.Update, bot *tgbotapi.BotAPI) {

	// Send a message to the chat
	msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "You pressed button 1!")
	_, err := bot.Send(msg)
	if err != nil {
		log.Printf("Error sending message to chat - %v", err)
		return
	}
}

func handleButton2(update tgbotapi.Update, bot *tgbotapi.BotAPI) {

	// Send a message to the chat
	msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "You pressed button 2!")
	_, err := bot.Send(msg)
	if err != nil {
		log.Printf("Error sending message to chat - %v", err)
		return
	}
}

func handleButton3(update tgbotapi.Update, bot *tgbotapi.BotAPI) {

	// Send a message to the chat
	msg := tgbotapi.NewEditMessageReplyMarkup(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID, tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Button 3-1 💎", "3-1"),
			tgbotapi.NewInlineKeyboardButtonData("Button 3-2 💎", "3-2"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Back ⬅️ ", "back"),
		),
	))
	_, err := bot.Send(msg)
	if err != nil {
		log.Printf("Error sending message to chat - %v", err)
		return
	}
}

func handleButton4(update tgbotapi.Update, bot *tgbotapi.BotAPI) {

	// Send a message to the chat
	msg := tgbotapi.NewEditMessageReplyMarkup(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID, tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Button 4-1 💎", "4-1"),
			tgbotapi.NewInlineKeyboardButtonData("Button 4-2 💎", "4-2"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Back ⬅️", "back"),
		),
	))
	_, err := bot.Send(msg)
	if err != nil {
		log.Printf("Error sending message to chat - %v", err)
		return
	}
}

func handleButtonBack(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	msg := tgbotapi.NewEditMessageReplyMarkup(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID, tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Button 1 💎", "1"),
			tgbotapi.NewInlineKeyboardButtonData("Button 2 💎", "2"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Button 3 💎 (3-1, 3-2)", "3"),
			tgbotapi.NewInlineKeyboardButtonData("Button 4 💎 (4-1, 4-2)", "4"),
		),
	))
	_, err := bot.Send(msg)
	if err != nil {
		log.Printf("Error sending message to chat - %v", err)
		return
	}
}
