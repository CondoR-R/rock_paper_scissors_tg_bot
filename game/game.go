package game

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Переменные

var (
	Moves movesStruct = movesStruct{
		Rock:     "Камень",
		Paper:    "Бумага",
		Scissors: "Ножницы",
	}
	movesArray = [...]string{Moves.Rock, Moves.Paper, Moves.Scissors}

	scoreMap = map[int64]scoreStruct{}

	winner = winnerStruct{
		user:   "user",
		bot:    "bot",
		nobody: "nobody",
	}
)

// Функции и процедуры, доступные вне пакета

// Сброс счета игры для пользователя с идентефикатором чата id.
// Принимает id чата
func ResetScore(id int64) {
	if score, ok := scoreMap[id]; ok {
		score.resetScore()
		scoreMap[id] = score
	}
}

// Главная фукция игры
// Принимает объект сообщения пользователя
func Game(message *tgbotapi.Message) string {
	userMove := message.Text
	userId := message.Chat.ID
	if !isValidMove(userMove) {
		return "Неверный ход, должно быть одно из слов: Камень, Ножницы, Бумага"
	}
	botMove := getMove()
	w := getWinner(botMove, userMove)
	updateScore(userId, w)

	return getEndRoundMessage(botMove, userMove, w, userId)
}
