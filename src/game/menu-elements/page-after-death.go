package menu_elements

import (
	"github.com/charmbracelet/lipgloss"
	game_abstr "github.com/pseudoelement/terminal-snake/src/game/abstracts"
	dead_text "github.com/pseudoelement/terminal-snake/src/game/game-elements/dead-text"
	"github.com/pseudoelement/terminal-snake/src/game/game-elements/score"
	"github.com/pseudoelement/terminal-snake/src/game/services/store"
	consts "github.com/pseudoelement/terminal-snake/src/shared/constants"
)

type AfterDeathPage struct {
	*game_abstr.Page
	store *store.Store
}

func NewAfterDeathPage(store *store.Store) *AfterDeathPage {
	selectableElems := []game_abstr.ISelectableElement{
		NewBackToMenuBtn(),
	}
	selectableElems[0].Select()

	return &AfterDeathPage{
		Page:  game_abstr.NewPage(store, selectableElems),
		store: store,
	}
}

func (this *AfterDeathPage) View() string {
	scoreEl := score.NewScore(this.store)
	scoreEl.UpdateTeaElement(scoreEl.TeaElement().
		UnsetBorderBottom().
		UnsetBorderLeft().
		UnsetBorderRight().
		UnsetBorderTop().
		MarginBottom(1).
		Width(20),
	)

	firstRow := lipgloss.JoinVertical(
		lipgloss.Left,
		dead_text.NewDeadText().View(), scoreEl.View(),
	)
	content := append([]string{firstRow}, this.SelectableElemsToViews()...)

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

func (this *AfterDeathPage) OnInit() {}

var _ game_abstr.IPage = (*AfterDeathPage)(nil)
