package game

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
