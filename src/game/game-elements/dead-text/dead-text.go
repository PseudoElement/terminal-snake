package dead_text

import game_abstr "github.com/pseudoelement/terminal-snake/src/game/abstracts"

type DeadText struct {
	*game_abstr.BaseViewElement
}

func NewDeadText() *DeadText {
	return &DeadText{
		BaseViewElement: game_abstr.NewBaseViewElement(deadText),
	}
}

func (this *DeadText) View() string {
	return this.TeaElement().Render("YOU DIED.")
}

var _ game_abstr.IViewElement = (*DeadText)(nil)
