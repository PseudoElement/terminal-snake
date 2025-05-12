package score

import (
	"fmt"

	game_abstr "github.com/pseudoelement/terminal-snake/src/game/abstracts"
	"github.com/pseudoelement/terminal-snake/src/game/services/store"
	consts "github.com/pseudoelement/terminal-snake/src/shared/constants"
)

type Score struct {
	*game_abstr.BaseViewElement
	store *store.Store
}

func NewScore(store *store.Store) *Score {
	return &Score{
		BaseViewElement: game_abstr.NewBaseViewElement(scoreTab),
		store:           store,
	}
}

func (this *Score) View() string {
	score := this.store.Get(consts.SCORE).(int)
	text := fmt.Sprintf("Score: %d", score)
	return this.TeaElement().Render(text)
}

var _ game_abstr.IViewElement = (*Score)(nil)
