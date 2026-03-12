package main

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	game "github.com/CondoR-R/rock_paper_scissors_tg_bot.git/game"
)

var (
	commands = struct {
		start string
	}{
		start: "start",
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
			botMove := game.MakeMove()
			userMove := update.Message.Text
			if !game.IsValidMove(userMove) {
				msg.Text = "Неверный ход, должно быть одно из слов: Камень, Ножницы, Бумага"
			} else {
				msg.Text = fmt.Sprintf(
					"Мой ход: %s.\nТвой ход: %s.\n%s\n",
					botMove,
					userMove,
					game.FindWinner(botMove, userMove))
			}
		}

		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}
