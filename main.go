package main

import (
	"fmt"
	"math/rand"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	moves = struct {
		rock     string
		paper    string
		scissors string
	}{
		rock:     "Камень",
		paper:    "Бумага",
		scissors: "Ножницы",
	}
	movesArray = [...]string{moves.rock, moves.paper, moves.scissors}

	commands = struct {
		start string
		love  string
	}{
		start: "start",
		love:  "love",
	}

	keyboard = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(moves.rock),
			tgbotapi.NewKeyboardButton(moves.paper),
			tgbotapi.NewKeyboardButton(moves.scissors)),
	)
)

func commandActions(u tgbotapi.Update, msg *tgbotapi.MessageConfig) {
	switch u.Message.Command() {
	case commands.start:
		msg.Text = "Давай поиграем в игру Камень, ножницы, бумага"
		msg.ReplyMarkup = keyboard
	case commands.love:
		msg.Text = "Ксюшенька моя, люблю тебя))))"
	default:
		msg.Text = "Неизвестная команда"
	}
}

func makeMove() string {
	randIndex := rand.Intn(len(movesArray))
	return movesArray[randIndex]
}

func isValidMove(userMove string) bool {
	for _, move := range movesArray {
		if move == userMove {
			return true
		}
	}
	return false
}

func findWinner(botMove, userMove string) string {
	if userMove == moves.rock && botMove == moves.scissors ||
		userMove == moves.scissors && botMove == moves.paper ||
		userMove == moves.paper && botMove == moves.rock {
		return "Поздравляю! Ты выиграл(а)!"
	} else if botMove == moves.rock && userMove == moves.scissors ||
		botMove == moves.scissors && userMove == moves.paper ||
		botMove == moves.paper && userMove == moves.rock {
		return "Я выиграл! :)"
	}
	return "Ого, у нас ничья"

}

func main() {
	bot, err := tgbotapi.NewBotAPI(ApiKey)
	if err != nil {
		panic(err)
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
			botMove := makeMove()
			userMove := update.Message.Text
			if !isValidMove(userMove) {
				msg.Text = "Неверный ход, должно быть одно из слов: Камень, Ножницы, Бумага"
			} else {
				msg.Text = fmt.Sprintf("Мой ход: %s.\nТвой ход: %s.\n%s\n", botMove, userMove, findWinner(botMove, userMove))
			}
		}

		if _, err := bot.Send(msg); err != nil {
			panic(err)
		}
	}
}
