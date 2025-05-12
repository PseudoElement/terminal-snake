package menu_elements

import (
	"fmt"

	game_abstr "github.com/pseudoelement/terminal-snake/src/game/abstracts"
)

type TextRules struct {
	*game_abstr.BaseViewElement
}

func NewTextRules() *TextRules {
	return &TextRules{
		BaseViewElement: game_abstr.NewBaseViewElement(textRules),
	}
}

func (this *TextRules) View() string {
	text := fmt.Sprintf(`
You need to live as long as possible and grab food(red squares).
Every grabbed food addes 1 point to your score.
On easy level of difficulty you'll die when you break down the border of game arena.
On medium and hard levels - you're not allowed cross the border and SNAKE'S BODY, otherwise you will lose.
Controls: 
* keyW - move up
* keyS - move down
* keyA - move left
* keyD - move right
* Escape - quit game and back to main menu

Enjoy :)
`)
	return this.TeaElement().Render(text)
}

var _ game_abstr.IViewElement = (*TextRules)(nil)
