package game

import (
	"fmt"
	"math/rand"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Типы

type (
	scoreStruct struct {
		bot  int
		user int
	}
	movesStruct struct {
		Rock     string
		Paper    string
		Scissors string
	}
	winnerStruct struct {
		user   string
		bot    string
		nobody string
	}
)

// Методы scoreStruct

func (s *scoreStruct) updateUserScore() {
	s.user++
}
func (s *scoreStruct) updateBotScore() {
	s.bot++
}

func (s *scoreStruct) resetScore() {
	s.bot = 0
	s.user = 0
}

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

// Функции

func getMove() string {
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

func getWinner(botMove, userMove string) string {
	if userMove == Moves.Rock && botMove == Moves.Scissors ||
		userMove == Moves.Scissors && botMove == Moves.Paper ||
		userMove == Moves.Paper && botMove == Moves.Rock {
		return winner.user
	} else if botMove == Moves.Rock && userMove == Moves.Scissors ||
		botMove == Moves.Scissors && userMove == Moves.Paper ||
		botMove == Moves.Paper && userMove == Moves.Rock {
		return winner.bot
	}
	return winner.nobody
}

func getEndRoundMessage(botMove, userMove, w string, id int64) string {
	score := scoreMap[id]
	var msg string
	switch w {
	case winner.bot:
		msg = "Я выиграл)"
	case winner.user:
		msg = "Поздравляю! Ты победил(а)!"
	case winner.nobody:
		msg = "Ого! У нас ничья!"
	}

	return fmt.Sprintf(
		"Мой ход: %s.\nТвой ход: %s.\n\n%s\n\nСчет: %d:%d",
		botMove,
		userMove,
		msg,
		score.bot,
		score.user)
}

// Процедуры

func ResetScore(id int64) {
	if score, ok := scoreMap[id]; ok {
		score.resetScore()
		scoreMap[id] = score
	}
}

func updateScore(id int64, w string) {
	_, ok := scoreMap[id]
	if !ok {
		scoreMap[id] = scoreStruct{user: 0, bot: 0}
	}
	score := scoreMap[id]
	switch w {
	case winner.bot:
		score.updateBotScore()
	case winner.user:
		score.updateUserScore()
	}
	scoreMap[id] = score
}

// Главная фукция игры

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
