package game

import "math/rand"

var (
	Moves = struct {
		Rock     string
		Paper    string
		Scissors string
	}{
		Rock:     "Камень",
		Paper:    "Бумага",
		Scissors: "Ножницы",
	}
	movesArray = [...]string{Moves.Rock, Moves.Paper, Moves.Scissors}
)

func MakeMove() string {
	randIndex := rand.Intn(len(movesArray))
	return movesArray[randIndex]
}

func IsValidMove(userMove string) bool {
	for _, move := range movesArray {
		if move == userMove {
			return true
		}
	}
	return false
}

func FindWinner(botMove, userMove string) string {
	if userMove == Moves.Rock && botMove == Moves.Scissors ||
		userMove == Moves.Scissors && botMove == Moves.Paper ||
		userMove == Moves.Paper && botMove == Moves.Rock {
		return "Поздравляю! Ты выиграл(а)!"
	} else if botMove == Moves.Rock && userMove == Moves.Scissors ||
		botMove == Moves.Scissors && userMove == Moves.Paper ||
		botMove == Moves.Paper && userMove == Moves.Rock {
		return "Я выиграл! :)"
	}
	return "Ого, у нас ничья"
}
