package menu_elements

import (
	game_abstr "github.com/pseudoelement/terminal-snake/src/game/abstracts"
	"github.com/pseudoelement/terminal-snake/src/game/services/store"
)

type PlayBtn struct {
	*game_abstr.SelectableElement
}

func NewPlayBtn() *PlayBtn {
	teaElement := bluredBtn
	btn := &PlayBtn{
		SelectableElement: game_abstr.NewSelectableElement(teaElement),
	}

	return btn
}

func (this *PlayBtn) Action(store *store.Store) {}

func (this *PlayBtn) NextPage(store *store.Store) game_abstr.IPage {
	return NewGamePage(store)
}

func (this *PlayBtn) View() string {
	return this.TeaElement().Render("Play")
}

func (this *PlayBtn) Select() {
	selectedTeaElem := selectedBtn
	this.SetSelected(true)
	this.UpdateTeaElement(selectedTeaElem)
}

func (this *PlayBtn) Blur() {
	bluredTeaElem := bluredBtn
	this.SetSelected(false)
	this.UpdateTeaElement(bluredTeaElem)
}

var _ game_abstr.IRedirectableElement = (*PlayBtn)(nil)
