package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	game "github.com/CondoR-R/rock_paper_scissors_tg_bot.git/game"
)

var (
	commands = struct {
		start string
		reset string
	}{
		start: "start",
		reset: "reset",
	}

	keyboard = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(game.Moves.Rock),
			tgbotapi.NewKeyboardButton(game.Moves.Paper),
			tgbotapi.NewKeyboardButton(game.Moves.Scissors)),
	)
)

func commandActions(u tgbotapi.Update, msg *tgbotapi.MessageConfig) {
	switch u.Message.Command() {
	case commands.start:
		msg.Text = "Давай поиграем в игру Камень, ножницы, бумага"
		msg.ReplyMarkup = keyboard
	case commands.reset:
		msg.Text = "Счет сброшен"
		game.ResetScore(u.Message.Chat.ID)
	default:
		msg.Text = "Неизвестная команда"
	}
}

func main() {
	bot, err := tgbotapi.NewBotAPI(ApiKey)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30

	updates := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		if update.Message.IsCommand() {
			commandActions(update, &msg)
		} else {
			msg.Text = game.Game(update.Message)
		}

		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}
