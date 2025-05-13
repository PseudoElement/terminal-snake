package menu_elements

import (
	"fmt"

	game_abstr "github.com/pseudoelement/terminal-snake/src/game/abstracts"
	"github.com/pseudoelement/terminal-snake/src/game/services/store"
	consts "github.com/pseudoelement/terminal-snake/src/shared/constants"
)

type TextDifficulty struct {
	*game_abstr.BaseViewElement
	store *store.Store
}

func NewTextDifficulty(store *store.Store) *TextDifficulty {
	return &TextDifficulty{
		BaseViewElement: game_abstr.NewBaseViewElement(textDifficulty),
		store:           store,
	}
}

func (this *TextDifficulty) View() string {
	diffLevel := this.store.Get(consts.DIFFICULTY).(game_abstr.IDiffLevel)
	text := fmt.Sprintf("Difficulty: %s", diffLevel.Name())
	return this.TeaElement().Render(text)
}

var _ game_abstr.IViewElement = (*TextDifficulty)(nil)
