package exitinfo

import (
	game_abstr "github.com/pseudoelement/terminal-snake/src/game/abstracts"
)

type ExitInfo struct {
	*game_abstr.BaseViewElement
}

func NewExitInfo() *ExitInfo {
	return &ExitInfo{
		BaseViewElement: game_abstr.NewBaseViewElement(exitInfo),
	}
}

func (this *ExitInfo) View() string {
	return this.TeaElement().Render("Press \"Escape\" to quit.")
}

var _ game_abstr.IViewElement = (*ExitInfo)(nil)
