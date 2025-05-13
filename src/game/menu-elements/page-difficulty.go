package menu_elements

import (
	"github.com/charmbracelet/lipgloss"
	game_abstr "github.com/pseudoelement/terminal-snake/src/game/abstracts"
	"github.com/pseudoelement/terminal-snake/src/game/services/store"
	consts "github.com/pseudoelement/terminal-snake/src/shared/constants"
)

type DifficultyPage struct {
	*game_abstr.Page
	store          *store.Store
	textDifficulty game_abstr.IViewElement
}

func NewDifficultyPage(store *store.Store) *DifficultyPage {
	selectableElems := []game_abstr.ISelectableElement{
		NewEasyBtn(),
		NewMediumBtn(),
		NewHardBtn(),
	}
	selectableElems[0].Select()

	return &DifficultyPage{
		Page:           game_abstr.NewPage(store, selectableElems),
		store:          store,
		textDifficulty: NewTextDifficulty(store),
	}
}

func (this *DifficultyPage) View() string {
	content := append([]string{this.textDifficulty.View()}, this.SelectableElemsToViews()...)

	joinVertical := lipgloss.JoinVertical(
		lipgloss.Center,
		content...,
	)

	w := this.Store().Get(consts.WIDTH).(int)
	h := this.Store().Get(consts.HEIGHT).(int)

	flex := lipgloss.Place(
		w, h/2,
		lipgloss.Center, lipgloss.Bottom,
		joinVertical,
	)

	return flex
}

func (this *DifficultyPage) OnInit() {}

var _ game_abstr.IPage = (*DifficultyPage)(nil)
