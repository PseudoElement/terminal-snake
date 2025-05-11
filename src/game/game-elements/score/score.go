package score

import (
	"fmt"

	game_abstr "github.com/pseudoelement/terminal-snake/src/game/abstracts"
)

type Score struct {
	*game_abstr.BaseViewElement
	score int
}

func NewScore() *Score {
	return &Score{
		BaseViewElement: game_abstr.NewBaseViewElement(scoreTab),
		score:           0,
	}
}

func (this *Score) Increment() {
	this.score++
}

func (this *Score) Decrement() {
	this.score--
}

func (this *Score) View() string {
	text := fmt.Sprintf("Score: %d", this.score)
	return this.TeaElement().Render(text)
}

var _ game_abstr.IViewElement = (*Score)(nil)
