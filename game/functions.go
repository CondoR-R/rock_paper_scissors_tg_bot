package game

import (
	"fmt"
	"math/rand"
)

// Функции

// Генерирует ход бота.
// Возвращает ход.
func getMove() string {
	randIndex := rand.Intn(len(movesArray))
	return movesArray[randIndex]
}

// Проверяет корректен ли ход пользователя.
// Принимает строку ходом.
// Возвращает true/false в зависимости от корректности хода
func isValidMove(userMove string) bool {
	for _, move := range movesArray {
		if move == userMove {
			return true
		}
	}
	return false
}

// Определяет победиля раунда.
// Принимает ходы бота и пользователя.
// Возвращает победиля (пользователь, бот или ничья).
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

// Генерирует сообщение для конца раунда.
// Принимает ходы бота и пользователя, победиля и id чата.
// Возвращает сообщение для пользователя.
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
		score.user,
		score.bot)
}

// Процедуры

// Обновляет счет игры.
// Принимает id чата и победиля раунда.
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
